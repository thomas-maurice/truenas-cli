#!/bin/bash

echo "cleaning up old ezcrypt config"
rm -rf ezcrypt.yml

echo "generating a new pki..."
ezcrypt pki init --key-type rsa --rsa-bits 4096 -d --name testing > /dev/null
echo "generating a new nas cert..."
ezcrypt pki cert new truenas --pki testing --server --cn truenas-test.local -k rsa -b 4096 > /dev/null
certificate=$(ezcrypt pki cert ls -ojson | jq '.data[] | select(.commonName == "truenas-test.local") | .id' -r)
ezcrypt pki cert get-cert ${certificate} > cert.crt
ezcrypt pki cert chain ${certificate} >> cert.crt
ezcrypt pki cert get-pk ${certificate} > cert.key
echo "written cert.crt and cert.key"
echo
echo "reconfiguring truenas UI"
./bin/truenas-cli certificate add certificate cert.crt cert.key
./bin/truenas-cli system general set-ui-cert "$(./bin/truenas-cli system general get-ui-certs | grep certificate | awk '{ print \$1 }' | tr -d \" | tr -d :)"
./bin/truenas-cli raw get system/general/ui_restart
