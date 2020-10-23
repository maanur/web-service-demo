FROM golang:1.15.3-alpine as builder
WORKDIR /opt/app
COPY . .
RUN go build .

FROM golang:1.15.3-alpine
COPY --from=builder /opt/app/web-service-demo /usr/local/bin/web-service-demo
ENTRYPOINT [ "web-service-demo" ]
