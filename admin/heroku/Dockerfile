# UI image =====================================================================
FROM node:12-alpine AS build_ui

ARG REACT_APP_API_URL
ENV REACT_APP_API_URL=$REACT_APP_API_URL

WORKDIR /src
COPY ./admin/ui ./
RUN \
  echo "REACT_APP_API_URL=$REACT_APP_API_URL" >> .env && \
  npm i && \
  npm run build

# SERVER image =================================================================
FROM golang:1.12-alpine AS build_server

ENV GO111MODULE on

# Build project
WORKDIR /go/src
COPY ./admin/server .
RUN CGO_ENABLED=0 \
  GOOS=linux \
  go build \
  -a \
  -mod vendor \
  -installsuffix cgo \
  -o app ./cmd/server

# Application image ============================================================
FROM nginx:1.16-alpine

RUN apk add --update bash curl openssh python supervisor && \
	  ln -sf /dev/stdout /var/log/nginx/access.log && ln -sf /dev/stderr /var/log/nginx/error.log

WORKDIR /usr/share/nginx/html
COPY ./admin/heroku/ops/root /
COPY --from=build_ui /src/build ./
COPY --from=build_server /go/src/app /app

CMD /bin/bash /heroku-exec.sh && /bin/bash -c "envsubst '\$PORT' < /etc/nginx/conf.d/default.conf > /etc/nginx/conf.d/default.conf" && supervisord -c /etc/supervisord.conf
