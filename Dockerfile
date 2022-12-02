# Add Maintainer info
# LABEL maintainer="Utkarsh Chourasia<utkarshchourasia.in>"

# Postgres Database setup
FROM postgres:14-alpine AS database

ENV POSTGRES_PASSWORD=secret

ENV POSTGRES_PORT=5432

EXPOSE 5432

COPY models/init.sql /docker-entrypoint-initdb.d/
