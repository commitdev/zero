# Terraform init x509 issue 

```
x509: certificate signed by unknown authority.
```

[Cacert](https://curl.haxx.se/ca/cacert.pem) from https://curl.haxx.se/docs/caextract.html

[Curl PEM issue solution](https://github.com/hashicorp/terraform/issues/10779#issuecomment-304664405)
[Go x509 issue solution](https://stackoverflow.com/a/29295887/2990066)
[Go x509 envar setup](https://golang.org/src/crypto/x509/root_unix.go)
