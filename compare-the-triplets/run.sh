#!/usr/bin/env bash

GOT=`cat $1/input/0.txt | go run $1/$2.go`
WANT=`cat $1/output/0.txt`
echo "want $WANT, got $GOT"
