#!/usr/bin/env ash

cd /go/src/surls && \
./bin/surls_linux_amd64 --run-mode container &

tail -f /tmp/block;