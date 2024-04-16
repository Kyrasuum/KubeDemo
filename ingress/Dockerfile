# arguements
ARG IMAGE_NAME=alpine
ARG TAG=latest
# production stage
FROM ${IMAGE_NAME}:${TAG} as production-stage

RUN apk add nginx openrc gettext lua5.1 lua5.1-dev luarocks nginx-mod-http-lua --no-cache &&\
    addgroup www &&\
    mkdir -p /var/cache/nginx/logs &&\
    luarocks-5.1 install lua-resty-http &&\
    chown -R nginx:www /var/lib/nginx

COPY ./startup.sh /etc/nginx/startup.sh
COPY ./nginx.conf /etc/nginx/nginx.conf.template
RUN chown -R nginx:www /etc/nginx

EXPOSE 80
USER nginx
CMD ["/etc/nginx/startup.sh"]
