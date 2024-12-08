FROM node:19.9.0-alpine AS build

ARG REACT_APP_API_URL
ENV REACT_APP_API_URL=$REACT_APP_API_URL

WORKDIR /src
COPY ./ ./
RUN \
  echo "REACT_APP_API_URL=$REACT_APP_API_URL" >> .env && \
  npm i && \
  npm run build

FROM nginx:1.23-alpine

RUN apk add --update bash curl openssh python

WORKDIR /usr/share/nginx/html
COPY ./ops/root /
COPY --from=build /src/build ./
