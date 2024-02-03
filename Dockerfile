FROM golang:alpine as builder

COPY ./ /build/

WORKDIR /build

ENV GO111MODULE=on CGO_ENABLED=0 GOPROXY="https://goproxy.cn,https://proxy.golang.com.cn,direct"

RUN apk update && apk add --no-cache upx git

RUN go mod tidy

RUN go build -a \
  -trimpath \
  -ldflags="-s -w -X 'main.Version=$(git describe --tags --always || git rev-parse --short HEAD)'" \
  -o drone-git-sync .

RUN upx -9 /build/drone-git-sync

FROM alpine:3

MAINTAINER 陌竹 <mozhu233@outlook.com>

LABEL org.label-schema.name="drone-git-sync" \
  org.label-schema.vendor="陌竹" \
  org.label-schema.schema-version="1.0" \
  org.label-schema.vcs-url="https://github.com/virzz/drone-git-sync" \
  org.opencontainers.image.source="https://github.com/virzz/drone-git-sync" \
  org.opencontainers.image.description="Drone Git Sync" \
  org.opencontainers.image.licenses="MIT"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories  && \
  apk update && \
  apk add --no-cache ca-certificates git git-lfs openssh curl perl && \
  rm -rf /var/cache/apk/*

COPY --from=builder --chmod=755 /build/drone-git-sync /bin/

ENTRYPOINT ["/bin/drone-git-sync"]
