---
title: "twic cert ls"
description: "The cert ls command description and usage"
keywords: [ "cert", "ls" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_cert"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/cert_ls.md"
---

```markdown
List Docker certificates

Usage:
  twic cert ls [flags]

Aliases:
  ls, list
```

## Examples

```bash
$ twic cert ls
Name                Type                CN                  TSA URL                   Expire
tsa1                client              user1               https://tsa1.example.com  2018-02-02
```

## Related information

* [cert_add](cert_add.md)
* [cert_rm](cert_rm.md)
