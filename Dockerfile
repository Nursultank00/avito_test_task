FROM golang:1.19.3

RUN go version
ENV GOPATH=/

COPY ./ ./

# psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# wait-for-postgres.sh
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o api ./cmd/main.go

CMD ["./api"]