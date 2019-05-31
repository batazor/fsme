FROM node:12-alpine AS build

WORKDIR /src
COPY ./ ./
RUN npm i
RUN npm run build

FROM nginx:1.16-alpine

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

WORKDIR /usr/share/nginx/html
COPY ./ops/root /
COPY --from=build /src/build ./

CMD /bin/sh -c "envsubst '\$PORT' < /etc/nginx/conf.d/default.conf > /etc/nginx/conf.d/default.conf" && nginx -g 'daemon off;'
