#!/bin/sh
set -o errexit

set -o pipefail

set -o nounset

postgres_ready() {
python << END
import sys
import psycopg2
try:
    psycopg2.connect(
        dbname="${DEFAULT_DATABASE_NAME}",
        user="${DEFAULT_DATABASE_USER}",
        password="${DEFAULT_DATABASE_PASSWORD}",
        host="${DEFAULT_DATABASE_HOST}",
        port="${DEFAULT_DATABASE_PORT}",
    )
except psycopg2.OperationalError:
    sys.exit(-1)
sys.exit(0)
END
}

until postgres_ready; do
 >&2 echo "Waiting for PostgreSQL to become available....:-("
 sleep 1
done
>&2 echo "PostgreSQL is ready!!!!...:-)"

python manage.py migrate
python manage.py collectstatic --noinput

exec "$@"
