#!/bin/sh

python wordcounter/manage.py migrate

exec "$@"
