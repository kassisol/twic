---
title: "twic cert add"
description: "The cert add command description and usage"
keywords: [ "cert", "add" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_cert"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/cert_add.md"
---

```markdown
Add Docker certificate

Usage:
  twic cert add [name] [flags]

Flags:
  -a, --alt-names string     Certificate Alternative Names
  -n, --common-name string   Certificate Common Name
  -p, --password string      Password
  -c, --tsa-url string       TSA URL
  -t, --type string          Certificate type (default "client")
  -u, --username string      Username
```

## Examples

```bash
$ twic cert add tsa1
TSA URL : https://tsa1.example.com
Username : user1
Password: ********
$ twic cert ls
Name                Type                CN                  TSA URL                   Expire
tsa1                client              user1               https://tsa1.example.com  2018-02-02
```

## Related information

* [cert_ls](cert_ls.md)
* [cert_rm](cert_rm.md)
