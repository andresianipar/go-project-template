#!/bin/bash

CONTAINER=default

if [ -n "$1" ]; then
  CONTAINER=$1
else
  echo "ERROR: command requires exactly 1 argument: container name"
  exit 1
fi

if [ "$(docker ps -a | grep -c $CONTAINER)" -gt 0 ]; then
  docker ps -a | grep $CONTAINER | awk '{ print $1 }' | xargs docker rm -f
  echo "Removed container: $CONTAINER"
else
  echo "No container to remove: $CONTAINER"
fi

exit
