#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    DO \$\$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'admin') THEN
            CREATE ROLE admin WITH LOGIN PASSWORD 'dbpassword';
        ELSE
            RAISE NOTICE 'Role admin already exists. Skipping creation.';
        END IF;

   END
    \$\$;
EOSQL


