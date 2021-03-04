#!/bin/bash

go test -v -coverprofile=cp.out
go tool cover -func=cp.out
go tool cover -html=cp.out -o coverage.html