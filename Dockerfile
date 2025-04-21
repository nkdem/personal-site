FROM node:22.14-alpine AS builder

ARG GIT_SHA
RUN if [ -z "${GIT_SHA}" ]; then \
        echo "Error: GIT_SHA is not set."; \
        exit 1; \
    fi
ENV ELEVENTY_GIT_SHA=${GIT_SHA}

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN corepack enable pnpm
RUN pnpm install --frozen-lockfile 

COPY . .

RUN pnpm run build 

# Stage 2: Build Caddy v2.10.0 from source
FROM golang:1.24-alpine AS caddy_builder

RUN apk add --no-cache go git

ENV CADDY_VERSION=v2.10.0

# Clone Caddy repository and checkout the specific version
RUN git clone https://github.com/caddyserver/caddy.git /caddy
WORKDIR /caddy
RUN git checkout ${CADDY_VERSION}

RUN go build -o /usr/local/bin/caddy ./cmd/caddy

# Final Stage: Copy over the caddy binary from caddy_builder, the static files generated from eleventy and run the Caddyfile
FROM alpine:3.21

COPY --from=caddy_builder /usr/local/bin/caddy /usr/local/bin/caddy

WORKDIR /srv

COPY --from=builder /app/dist /srv 

COPY Caddyfile /etc/caddy/Caddyfile

EXPOSE 80
CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile"]