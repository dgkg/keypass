FROM golang:1.15 AS builder
WORKDIR ./src
RUN pwd
COPY . .
RUN echo "download modules"
RUN go mod download
RUN echo "build binary"
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o api_keypass .
RUN pwd

FROM alpine:3.12
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder ./go/src/api_keypass .
COPY --from=builder ./go/src/config.yaml .
ENV GLIBC_VERSION 2.32-r0
# Download and install glibc
RUN apk add --update curl && \
  curl -Lo /etc/apk/keys/sgerrand.rsa.pub https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub && \
  curl -Lo glibc.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk" && \
  curl -Lo glibc-bin.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-bin-${GLIBC_VERSION}.apk" && \
  apk add glibc-bin.apk glibc.apk && \
  /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc-compat/lib && \
  echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf && \
  apk del curl && \
  rm -rf glibc.apk glibc-bin.apk /var/cache/apk/*
RUN ls -alsh
CMD ["./api_keypass"]