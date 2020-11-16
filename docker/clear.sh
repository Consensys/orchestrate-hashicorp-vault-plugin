#!/bin/bash

PWD="$(pwd)"
OPERATOR_JSON="$PWD/docker/config/operator.json"
OPERATOR_DATA="$PWD/docker/config/data"

echo "Clearing previous state..."
echo "rm -fr $OPERATOR_DATA"
rm -fr "$OPERATOR_DATA"
echo "rm $OPERATOR_JSON"
rm "$OPERATOR_JSON"
