# truenas-cli

Some simple things to get your NAS automated

```shell
./bin/truenas-cli certificate add certificate cert.crt cert.key
./bin/truenas-cli system general set-ui-cert "$(./bin/truenas-cli system general get-ui-certs | grep certificate | awk '{print $1}' | tr -d \" | tr -d :)"
```
