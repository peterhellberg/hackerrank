#!/usr/bin/env bash

cd $1 && cat input.txt | lein run -m $2
