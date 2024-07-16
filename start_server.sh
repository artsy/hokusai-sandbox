#!/usr/bin/env sh
export $(cat /secrets/secrets | xargs)
./server
