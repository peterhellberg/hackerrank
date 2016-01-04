#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname "$0"); pwd)

function runChallenge {
  spy -c c --dir "$1" ./$1/run.sh $SCRIPT_DIR/$1 $2
}

if which spy >/dev/null; then
  if [[ -z "$1" ]]; then
    echo "Usage: ./spy.sh <challenge> [variant]"
    exit
  fi

  if [[ -z "$2" ]]; then
    SCRIPT_NAME=$1
  else
    SCRIPT_NAME="$1_$2"
  fi

  runChallenge $1 $SCRIPT_NAME
else
  echo "You need to install spy: go get -v github.com/jpillora/spy"
fi
