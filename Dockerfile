FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/freddysilber/nfl-loser-pool-api/
WORKDIR /go/src/github.com/freddysilber/nfl-loser-pool-api
RUN go mod download
COPY . /go/src/github.com/freddysilber/nfl-loser-pool-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/nfl-loser-pool-api github.com/freddysilber/nfl-loser-pool-api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/freddysilber/nfl-loser-pool-api/build/nfl-loser-pool-api /usr/bin/nfl-loser-pool-api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/nfl-loser-pool-api"]