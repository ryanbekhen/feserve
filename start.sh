#!/bin/bash

go build -o tmp/feserve . && sudo setcap 'cap_net_bind_service=+ep' ./tmp/feserve && ./tmp/feserve