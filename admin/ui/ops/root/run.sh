#!/bin/bash

envsubst '\$PORT' < /etc/nginx/conf.d/default.conf > /etc/nginx/conf.d/default.conf
envsubst '\$HEROKU_PRIVATE_IP' < /etc/nginx/conf.d/default.conf > /etc/nginx/conf.d/default.conf
nginx -g 'daemon off;'
