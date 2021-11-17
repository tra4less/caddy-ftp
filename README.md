# caddy-ftp

access ftp through caddy

# usage

```shell
{
  auto_https off
  http_port 80
  https_port 443
}

:80 {
    route /abc/* {
        uri replace /abc/ /
        ftp {
            addr test.rebex.net:21
            user demo
            pass password
            # dial_timeout 3s
            # disable_epsv
            # disable_utf8
        }
    }
    # curl localhost:80/abc/pub/example/imap-console-client.png
}
```
