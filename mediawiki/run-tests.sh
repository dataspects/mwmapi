#!/bin/bash

export MWAPI=https://dserver/w/api.php
export MWROOT=/home/lex/mediawiki-manager/mediawiki_root/w
export MWCONTAINER=mediawiki_canasta

go test -v -coverprofile=cp.out
go tool cover -func=cp.out
go tool cover -html=cp.out -o coverage.html