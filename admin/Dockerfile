### This Dockerfile create for deploy to heroku #%@!

# UI image =====================================================================
FROM node:19.9.0-alpine AS build_ui

ARG REACT_APP_API_URL
ENV REACT_APP_API_URL=$REACT_APP_API_URL

WORKDIR /src
COPY ./ui ./
RUN \
  echo "REACT_APP_API_URL=$REACT_APP_API_URL" >> .env && \
  npm i && \
  npm run build

# SERVER image =================================================================
FROM golang:1.23-alpine AS build_server

ENV GO111MODULE on

# Build project
WORKDIR /go/src
COPY ./server .
RUN CGO_ENABLED=0 \
  GOOS=linux \
  go build \
  -a \
  -mod vendor \
  -installsuffix cgo \
  -o app ./cmd/server

# Application image ============================================================
FROM nginx:1.23-alpine

ARG PORT=80
ENV PORT=$PORT

RUN apk add --update bash curl openssh python py-pip supervisor && \
	  pip install supervisor-stdout

WORKDIR /usr/share/nginx/html
COPY ./heroku/ops/root /
COPY --from=build_ui /src/build ./
COPY --from=build_server /go/src/app /app

CMD /bin/bash /heroku-exec.sh && /bin/bash -c "envsubst '\$PORT' < /etc/nginx/conf.d/default.conf > /etc/nginx/conf.d/default.conf" && supervisord -c /etc/supervisord.conf
