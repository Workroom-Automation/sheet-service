#!/bin/sh

[ -d "/etc/app/" ] && ln -s /etc/app/* .

# and add this at the end
ls -lah
exec "./app" $@