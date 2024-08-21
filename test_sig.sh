#!/bin/sh

trap "echo eating sigterm; exit" SIGTERM

while true
do
  sleep 5
done
