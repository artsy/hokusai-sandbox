#!/bin/sh

CMD=$1
echo "$CMD"
source /secrets/secrets
"$CMD"

