#!/usr/bin/env bash

GOT=`cat $1/input/1.txt | go run $1/$2.go`
WANT=`cat $1/output/1.txt`

echo WANT
echo ----
echo "$WANT"
echo
echo GOT
echo ---
echo "$GOT"
