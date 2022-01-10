FROM golang:1.17.3

WORKDIR /
COPY . .

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /eventbackend

COPY . /eventbackend

RUN go build -o backend .

EXPOSE 8080

ENTRYPOINT ["./backend"]