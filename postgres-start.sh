docker run -it -d --name postgresql --rm \
    -v /root/postgresql-data:/postgressqldata \
    -p 0.0.0.0:5432:5432 \
    -e POSTGRES_USER=root \
    -e POSTGRES_PASSWORD=990130 \
    -e POSTGRES_DB=eventbackend \
    -e PGDATA=/postgressqldata \
    postgres:14