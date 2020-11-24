FROM golang:latest

RUN mkdir -p /go/src/prom-receiver
ENV GO111MODULE=on
WORKDIR /go/src/prom-receiver
COPY . /go/src/prom-receiver
RUN go install github.com/hellozee/prom-receiver
CMD /go/bin/prom-receiver

EXPOSE 5001
