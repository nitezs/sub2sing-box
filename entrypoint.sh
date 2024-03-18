#!/bin/sh
if [ ! -d "/app/templates" ]; then
	mkdir /app/templates
fi
if [ -z "$(ls -A /app/templates)" ]; then
	cp -r /app/templates-origin/* /app/templates
fi
cd /app
/app/sub2sing-box server
