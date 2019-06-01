FROM node:12-alpine AS build

ARG REACT_APP_API_URL
ENV REACT_APP_API_URL=$REACT_APP_API_URL

WORKDIR /src
COPY ./ ./
RUN \
  echo "REACT_APP_API_URL=$REACT_APP_API_URL" >> .env && \
  npm i && \
  npm run build

FROM nginx:1.16-alpine

RUN apk add --update bash curl openssh python

WORKDIR /usr/share/nginx/html
COPY ./ops/root /
COPY --from=build /src/build ./

CMD /bin/bash /heroku-exec.sh && /bin/bash -c "envsubst '\$PORT' < /etc/nginx/conf.d/default.conf > /etc/nginx/conf.d/default.conf" && nginx -g 'daemon off;'
