---
title: "twic cert rm"
description: "The cert rm command description and usage"
keywords: [ "cert", "rm" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_cert"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/cert_rm.md"
---

```markdown
Remove Docker certificate

Usage:
  twic cert rm [name] [flags]

Aliases:
  rm, remove

Flags:
  -p, --password string   Password
  -u, --username string   Username
```

## Examples

```bash
$ twic cert ls
Name                Type                CN                  TSA URL                   Expire
tsa1                client              user1               https://tsa1.example.com  2018-02-02
tsa2                client              user1               https://tsa2.example.com  2018-02-03
$ twic cert rm tsa2
$ twic cert ls
Name                Type                CN                  TSA URL                   Expire
tsa1                client              user1               https://tsa1.example.com  2018-02-02
```

## Related information

* [cert_add](cert_add.md)
* [cert_ls](cert_ls.md)
