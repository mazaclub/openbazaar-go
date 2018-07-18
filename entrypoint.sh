#!/bin/sh

# ends0/ndots options in resolv.conf may prevent 
# openbazaard from resolving names. docker prevents sed -i from working
# so do this gyration
sed -e 's/ndots:[0-9]//g'  -e 's/edns0//g'  /etc/resolv.conf > /opt/resolv.conf
cat /opt/resolv.conf > /etc/resolv.conf

/opt/openbazaard  "$@"
