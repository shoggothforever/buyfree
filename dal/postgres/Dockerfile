
FROM postgres:alpine
#COPY init.sql /docker-entrypoint-initdb.d
ENV POSTGRES_USER bf
ENV POSTGRES_PASSWORD bf123
ENV POSTGRES_DB bfdb
COPY . .

RUN chmod 755 createuser.sh && /bin/bash -c /bin/createuser.sh