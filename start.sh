#!/bin/bash
pkill notebook
rm -rf ./notebook
go build -o notebook main.go
nohup ./notebook &
