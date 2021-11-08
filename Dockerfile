FROM golang:latest as builder

COPY . /root/go/src/app
WORKDIR /root/go/src/app

ENV CGO_ENABLED=0

COPY go.* ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY main.go ./
RUN go build -a -ldflags="-s -w" -o server


FROM alpine
ENV DISCORD_TOKEN=""
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /root/go/src/app/server server
USER 2000
ENTRYPOINT ./server serve --token=$DISCORD_TOKEN