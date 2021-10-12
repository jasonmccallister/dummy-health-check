
FROM golang:1.17-alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN GOOS=linux go build -a -o health-check .

FROM alpine:3
RUN mkdir /app
COPY --from=builder /build/health-check /app
ENTRYPOINT [ "/app/health-check" ]
EXPOSE 8080
