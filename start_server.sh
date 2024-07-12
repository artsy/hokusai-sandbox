#!/usr/bin/env sh
export $(cat /tmp/env/file | xargs)
env
./server
