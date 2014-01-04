#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname "$0"); pwd)

function runChallenge {
  rerun -c -x -b -p "$1/**/*" -- ./$1/run.sh $SCRIPT_DIR/$1 $2
}

if which rerun >/dev/null; then
  if [[ -z "$1" ]]; then
    echo "Usage: ./rerun <challenge> [variant]"
    exit
  fi

  if [[ -z "$2" ]]; then
    SCRIPT_NAME=$1
  else
    SCRIPT_NAME="$1_$2"
  fi

  runChallenge $1 $SCRIPT_NAME
else
  echo "You need to install rerun"
fi
