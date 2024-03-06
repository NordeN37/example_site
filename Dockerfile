FROM golang:1.21.1-alpine as builder

# ARG GO_MAIN_PATH
ARG VERSION=0.1
WORKDIR /src
COPY . .

RUN apk --no-cache update && apk --no-cache add git gcc libc-dev

RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -mod=vendor -a -installsuffix cgo -o app -ldflags "-X 'main.Version=${VERSION}'" ./main.go
FROM alpine

WORKDIR /root/
COPY --from=builder /src .
EXPOSE 80

CMD ["./app"]
