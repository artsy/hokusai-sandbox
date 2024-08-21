#!/bin/sh

CMD="$@"

SECRETS_FILE_DIR=/secrets

if [ -d "$SECRETS_FILE_DIR" ]
then
  echo "Sourcing secrets file..."
  source "$SECRETS_FILE_DIR/secrets"
fi

echo "Running command: $CMD"
$CMD
