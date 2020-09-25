FROM golang:1.15.2 AS builder
WORKDIR /opt/app
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -o redirect .

FROM alpine
COPY --from=builder /opt/app/redirect /opt/app/redirect
EXPOSE 8080
CMD ["/opt/app/redirect"]
