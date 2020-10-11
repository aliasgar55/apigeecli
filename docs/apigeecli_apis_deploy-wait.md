## apigeecli apis deploy-wait

Deploys a revision of an existing API proxy and waits for deployment status

### Synopsis

Deploys a revision of an existing API proxy to an environment and waits for deployment status

```
apigeecli apis deploy-wait [flags]
```

### Options

```
  -e, --env string    Apigee environment name
  -h, --help          help for deploy-wait
  -n, --name string   API proxy name
  -r, --ovr           Forces deployment of the new revision
  -v, --rev int       API Proxy revision (default -1)
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --disable-check    Disable check for newer versions
  -o, --org string       Apigee organization name
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli apis](apigeecli_apis.md)	 - Manage Apigee API proxies in an org

###### Auto generated by spf13/cobra on 11-Oct-2020