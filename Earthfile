VERSION 0.8

tailwindcss:
  FROM alpine/curl
  RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v4.0.0-beta.8/tailwindcss-linux-x64 && \
      chmod +x tailwindcss-linux-x64 && \
      mv tailwindcss-linux-x64 tailwindcss
  SAVE ARTIFACT ./tailwindcss

deps:
  FROM golang:1.24
  WORKDIR /src
  COPY go.mod go.sum ./
  RUN go mod download

build:
  FROM +deps
  COPY +tailwindcss/tailwindcss /usr/local/bin/tailwindcss
  COPY *.go ./
  COPY --dir internal/ static/ ./
  RUN --mount=type=cache,id=go-build-cache,target=/root/.cache/go-build \
      GOOS=linux CGO_ENABLED=0 go build -o http-tool main.go
  RUN tailwindcss -i ./static/base.css -o ./static/style.css
  SAVE ARTIFACT ./http-tool AS LOCAL ./bin/http-tool
  SAVE ARTIFACT ./static
  SAVE ARTIFACT ./internal

run:
  FROM alpine:latest
  LABEL org.opencontainers.image.source = "https://github.com/hmaier-dev/http-tool"
  ARG tag
  WORKDIR /root
  COPY +build/static /root/static/
  COPY +build/internal /root/internal/
  COPY +build/http-tool .
  EXPOSE 8080
  ENTRYPOINT ["./http-tool"]
  SAVE IMAGE --push ghcr.io/hmaier-dev/http-tool:$tag

