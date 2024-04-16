#!/bin/sh
if [ "$NAMESERVER" == "" ]; then
    export DNS_SRV=$(cat /etc/resolv.conf |grep -i '^nameserver'|head -n1|cut -d ' ' -f2)
fi

echo "Nameserver is: $DNS_SRV"

echo "Copying nginx config"
envsubst '$DNS_SRV' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

echo "Using nginx config:"
cat /etc/nginx/nginx.conf

echo "Starting nginx"
exec nginx -g "daemon off;"
