package system

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// SetupDiff -
type SetupDiff struct {
	InContainerButNotInInstance []string `json:"inContainerButNotInInstance"`
	InInstanceButNotInContainer []string `json:"inInstanceButNotInContainer"`
}

// GetSetupDiff -
func GetSetupDiff(containerName string, containerWPath string, instanceWPath string, depth int) (SetupDiff, error) {
	mwcwtFN, _ := containerInternalTree(containerName, containerWPath, depth)
	mwmitFN, _ := fileSystemTree(instanceWPath, depth)
	ct := doCompareTrees(mwcwtFN, mwmitFN)
	sud := SetupDiff{
		InContainerButNotInInstance: ct.In0NotIn1,
		InInstanceButNotInContainer: ct.In1NotIn0,
	}
	return sud, nil
}

type compareTrees struct {
	In0NotIn1 []string
	In1NotIn0 []string
}

func doCompareTrees(f0 string, f1 string) compareTrees {
	file0, _ := ioutil.ReadFile(f0)
	file1, _ := ioutil.ReadFile(f1)
	return compareTrees{
		In0NotIn1: scanTree(file0, file1),
		In1NotIn0: scanTree(file1, file0),
	}
}

func scanTree(f0 []byte, f1 []byte) []string {
	sl := []string{}
	scanner := bufio.NewScanner(bytes.NewReader(f0))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		currentLine := scanner.Text()
		// FIXME: consider adding ^ and $
		re := regexp.MustCompile(regexp.QuoteMeta(currentLine))
		if re.Find(f1) == nil {
			sl = append(sl, currentLine)
		}
	}
	return sl
}

func fileSystemTree(path string, depth int) (string, string) {
	cmd := exec.Command("tree", "-L", strconv.Itoa(depth), path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fname := "/tmp/fstree.txt"
	f, err := os.Create(fname)
	f.Write(harmonizeTree(out))
	f.Close()
	return fname, out.String()
}

func containerInternalTree(containerName string, path string, depth int) (string, string) {
	cmd := exec.Command("sudo", "docker", "exec", "-t", containerName, "tree", "-L", strconv.Itoa(depth), path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("LEX2103051159 " + err.Error() + out.String())
	}
	fname := "/tmp/citree.txt"
	f, err := os.Create(fname)
	f.Write(harmonizeTree(out))
	f.Close()
	return fname, out.String()
}

func harmonizeTree(t bytes.Buffer) []byte {
	nt := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(t.String(), "├", "|", -1), "─", "-", -1), "│", "|", -1), "`", "|", -1), "└", "|", -1), " ", "\u00a0", -1)
	return []byte(nt)
}

// jr, _ := json.MarshalIndent(dockerOut, "", "")
// log.Print(string(jr))
