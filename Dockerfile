FROM golang:1.15 AS builder
WORKDIR ./src
RUN pwd
COPY . .
RUN go mod download -x
RUN go build -o api_keypass
RUN pwd

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
WORKDIR /bin/
COPY --from=builder ./go/src/api_keypass .
COPY --from=builder ./go/src/config.yaml .
CMD ["./api_keypass"]