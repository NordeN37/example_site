ARG GO_VER=1.21.1
FROM golang:${GO_VER}-alpine as builder

# ARG GO_MAIN_PATH
ARG VERSION=0.1
WORKDIR /app
COPY . /app

RUN go build -o main
# Expose port 8080 to the outside world
FROM alpine

WORKDIR /root/
COPY --from=builder /app .

EXPOSE 8081

CMD ["./main"]