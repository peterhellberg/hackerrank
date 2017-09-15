#!/usr/bin/env bash

cat $1/input/input00.txt | go run $1/$2.go
cat $1/input/input05.txt | go run $1/$2.go
cat $1/input/input08.txt | go run $1/$2.go
