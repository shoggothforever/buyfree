#!/bin/bash
set -e
#adduser bf
#bf123
#su - bf
#psql -d bfdb
if [ "$1" = 'postgres' ]; then
    chown -R postgres "$PGDATA"
    if [ -z "$(ls -A "$PGDATA")" ]; then
        gosu bf initdb
    fi
    exec gosu bf "$@"
fi
exec "$@"

