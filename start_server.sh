#!/bin/sh

# in prod, this file should contain app secret KEY=VALUE configs
# if file exists, populate env with the vars
SECRETS_FILE=/secrets/secrets
if [ -f "$SECRETS_FILE" ]
then
  export $(cat "$SECRETS_FILE" | xargs)
fi

# start app as you normally would
./server
