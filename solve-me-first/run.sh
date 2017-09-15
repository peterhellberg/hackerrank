#!/usr/bin/env bash

GOT=`cat $1/input/1.txt | go run $1/$2.go`
WANT=`cat $1/output/1.txt`
echo "want $WANT, got $GOT"

GOT=`cat $1/input/2.txt | go run $1/$2.go`
WANT=`cat $1/output/2.txt`
echo "want $WANT, got $GOT"
