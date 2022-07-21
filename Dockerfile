# create postgresql dockerfile

FROM postgres:9.6

ENV POSTGRES_USER nonadmin
ENV POSTGRES_PASSWORD postingdata
ENV POSTGRES_DB cshare

RUN postgresql-setup initdb

CMD ["postgresql-setup", "start"]

EXPOSE 5432

COPY . /docker-entrypoint-initdb.d

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

## create postgresql docker image
#docker build -t postgresql .
#
## run postgresql docker image
#docker run -d -p 5432:5432 --name postgresql postgresql
#
## connect to postgresql docker image
#docker exec -it postgresql psql -U postgres -d postgres
#
## create postgresql database
#CREATE DATABASE postgres;
#
## create postgresql user
#CREATE USER postgres WITH PASSWORD 'postgres';
#
## grant privileges to postgresql user
#GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;
#
## connect to postgresql docker image
#docker exec -it postgresql psql -U postgres -d postgres
#
## create postgresql database
#CREATE DATABASE postgres;
#
## create postgresql user
#CREATE USER postgres WITH PASSWORD 'postgres';
#
## grant privileges to postgresql user
#GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;
#
## connect to postgresql docker image
#docker exec -it postgresql psql -U postgres -d postgres
#
## create postgresql database
#CREATE DATABASE postgres;
#
## create postgresql user
#CREATE USER postgres WITH PASSWORD 'postgres';
#
## grant privileges to postgresql user
#GRANT ALL PRIVILEGES ON DATABASE postgres TO postgres;
#
## connect to postgresql docker image
#docker exec -it postgresql psql -U postgres -d postgres
#
## create postgresql database
#CREATE DATABASE postgres;