#!/bin/sh
echo resolver $(awk 'BEGIN{ORS=" "} $1=="nameserver" {print $2";"}' /etc/resolv.conf) > /etc/nginx/resolvers.conf

echo "Starting nginx"
exec nginx -g "daemon off;"
