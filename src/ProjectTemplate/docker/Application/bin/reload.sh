#!/bin/bash

RUN_PROCESS_NAME=GoRunning

# Watch all *.go files in the specified directory
# Call the restart function when they are changed
function monitor() {
  inotifywait -q -m -r -e modify -e move -e create -e delete --exclude '(.glide|vendor|_test.go|[^g][^o]$)' $1 |
  while read line; do
    for pid in $(jobs -p); do
      silentKill "$pid"
    done
    restart &
  done
}

function silentKill {
  kill -9 $1 2>/dev/null
  wait $1 2>/dev/null
}

# Terminate and rerun the main Go program
function restart {
  # when multiple changes happen (like git checkout other-branch)
  # don't start expensive go build, but wait for all changes
  sleep .5

  echo ">> Rebuilding..."
  go build -o $BIN_PATH $MAIN_PATH
  echo ">> Reloading..."
  pkill -f $RUN_PROCESS_NAME
  bash -c "exec -a $RUN_PROCESS_NAME $BIN_PATH &"
}

# Make sure all background processes get terminated
function close {
  killall -q -w -9 inotifywait
  exit 0
}

trap close INT
echo ">> Watching started"

MAIN_PATH=$1
BIN_PATH=$2

# Start the main Go program
restart

# Monitor current dir
monitor $PWD

wait
