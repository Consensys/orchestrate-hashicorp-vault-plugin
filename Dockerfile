FROM golang:1.14-alpine as builder
RUN apk add --update alpine-sdk
RUN apk update && apk add git openssh gcc musl-dev linux-headers

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY  / .
RUN mkdir -p /build/bin \
    && CGO_ENABLED=1 GOOS=linux go build -a -v -i -o /build/bin/orchestrate-hashicorp-vault-plugin . \
    && sha256sum -b /build/bin/orchestrate-hashicorp-vault-plugin > /build/bin/SHA256PLUGIN

FROM vault:latest
ARG always_upgrade
RUN echo ${always_upgrade} > /dev/null && apk update && apk upgrade
RUN apk add bash openssl jq

USER vault
WORKDIR /app
RUN mkdir -p /home/vault/plugins

COPY --from=builder /build/bin/orchestrate-hashicorp-vault-plugin /home/vault/plugins/orchestrate-hashicorp-vault-plugin
COPY --from=builder /build/bin/SHA256PLUGIN /home/vault/plugins/SHA256PLUGIN
RUN ls -la /home/vault/plugins
HEALTHCHECK CMD nc -zv 127.0.0.1 9200 || exit 1