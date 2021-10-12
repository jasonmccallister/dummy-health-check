
FROM golang:1.17-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN GOOS=linux go build -a -o healthcheck .

FROM alpine:3
RUN mkdir /app
COPY --from=builder /build/healthcheck /app
ENTRYPOINT [ "/app/healthcheck" ]
EXPOSE 8080
