#!/usr/bin/env sh
export $(cat /tmp/env/file | xargs)
env
#rm /tmp/env/file
./server
