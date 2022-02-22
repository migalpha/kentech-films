# build stage
FROM golang:1.16 as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go build -ldflags "-X main.version=$(git describe --tags HEAD || git log --format="%h" -1)" ./cmd/kentech-films

# final stage
FROM golang:1.16

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.10.0/migrate.linux-amd64.tar.gz | tar xvz && \
  mv migrate.linux-amd64 /usr/bin/migrate

COPY --from=builder /app /app
WORKDIR /app

EXPOSE 8000

RUN groupadd docker
RUN usermod -a -G docker root
RUN newgrp - docker