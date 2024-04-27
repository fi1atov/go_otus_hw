FROM golang:1.22-alpine3.19 as builder
WORKDIR /usr/local/go/src/
COPY hw15_go_sql /usr/local/go/src/

RUN go clean --modcache && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o service-one ./main.go

FROM scratch
WORKDIR /app
COPY --from=builder /usr/local/go/src/service-one /app/

EXPOSE 8080
CMD ["/app/service-one"]
