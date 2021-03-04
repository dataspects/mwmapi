VERSION=210303

source testConfig.env

sudo docker run -p 3000:3000 dataspects/mwmapi:${VERSION}