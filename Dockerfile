# Postgres Database setup
FROM postgres:14-alpine AS database

COPY models/init.sql /docker-entrypoint-initdb.d/
