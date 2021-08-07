## apigeecli targetservers update

Update a Target Server

### Synopsis

Update a Target Server

```
apigeecli targetservers update [flags]
```

### Options

```
  -c, --clientAuth        Enable mTLS for the target server
  -d, --desc string       Description for the Target Server
  -b, --enable            Enabling/disabling a TargetServer (default true)
  -g, --grpc              Enable target server for gRPC
  -h, --help              help for update
  -s, --host string       Host name of the target
  -i, --ignoreErr         Ignore TLS validation errors for the target server
      --keyAlias string   Key alias for the target server
      --keyStore string   Key store for the target server
  -n, --name string       Name of the targetserver
  -p, --port int          port number (default -1)
      --tls               Enable TLS for the target server
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --disable-check    Disable check for newer versions
  -e, --env string       Apigee environment name
  -o, --org string       Apigee organization name
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli targetservers](apigeecli_targetservers.md)	 - Manage Target Servers

###### Auto generated by spf13/cobra on 4-Aug-2021