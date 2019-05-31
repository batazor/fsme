FROM node:12-lapine AS build

WORKDIR /src
COPY ./ ./
RUN npm i
RUN npm run build

FROM nginx:1.16-alpine

WORKDIR /usr/share/nginx/html
COPY --from=build /src/build ./
