#!/bin/sh

set -e

FILE="/usr/share/nginx/html/internal/env.js"
echo -n "export const env={" > $FILE
env | grep ^PUBLIC_ | while read -r line; do
  key=${line%%=*}
  value=${line#*=}
  entry="\"${key}\":\"${value}\""
  if [ -z "$first" ]; then
    first=true
    echo -n "$entry" >> $FILE
  else
    echo -n ",$entry" >> $FILE
  fi
done
echo -n "};" >> $FILE
echo "$FILE"

exit 0
