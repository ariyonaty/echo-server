# syntax=docker/dockerfile:1
ARG BASE_REGISTRY=cgr.dev
ARG BASE_IMAGE=chainguard/go
ARG BASE_TAG=latest

FROM ${BASE_REGISTRY}/${BASE_IMAGE}:${BASE_TAG} AS builder

COPY . /app

RUN cd /app && \
    CGO_ENABLED=0 GOOS=linux \
        go build -a -tags netgo -ldflags '-w' .

FROM scratch
COPY --from=builder /app/echo-server /usr/bin/
CMD ["/usr/bin/echo-server"]
