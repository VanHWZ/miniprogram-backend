FROM bitnami/golang:1.17.3

WORKDIR /
COPY . .

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /eventbackend

COPY . /eventbackend

RUN wget -O /etc/apt/sources.list http://mirrors.cloud.tencent.com/repo/debian10_sources.list && \
    apt-get clean all && \
    apt-get update && \
    apt-get install -y vim net-tools
RUN go build -o backend .

EXPOSE 8080

ENTRYPOINT ["bash","-c", "nohup /eventbackend/backend &> /eventbackend/log/out.log"]