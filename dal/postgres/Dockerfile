
FROM postgres:alpine
#COPY ./init.sql /docker-entrypoint-initdb.d
ENV POSTGRES_USER bf
ENV POSTGRES_PASSWORD bf123
ENV POSTGRES_DB bfdb
RUN addgroup -S postgres && adduser -S -G bf bf123
COPY . .

EXPOSE 5432

#RUN #chmod 755 createuser.sh && /bin/bash -c /bin/createuser.sh
ENTRYPOINT ["createuser.sh"]
CMD ["postgres"]