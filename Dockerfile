FROM golang:latest as builder

WORKDIR /go/src/app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go mod download

COPY . /go/src/app
RUN go build -ldflags="-s -w" -o server


FROM alpine
ENV DISCORD_TOKEN=""
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app/server server
USER 2000
ENTRYPOINT ./server serve --token=$DISCORD_TOKEN