#!/usr/bin/env bash

curl -L http://guvm.mvd.ru/upload/expired-passports/list_of_expired_passports.csv.bz2 | bzip2 -d > /tmp/list_of_expired_passports.csv
PID=$(pidof -s invalid-passports-go)
if [ ! -z "${PID}" ]; then
  kill -SIGUSR1 "${PID}"
  echo Sent signal to "${PID}"
else
  echo "Not found running process 'invalid-passports-go'"
fi
