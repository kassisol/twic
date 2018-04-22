---
title: "twic profile rm"
description: "The profile rm command description and usage"
keywords: [ "profile", "rm" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_profile"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/profile_rm.md"
---

```markdown
Remove Docker profile

Usage:
  twic profile rm [name] [flags]

Aliases:
  rm, remove
```

## Examples

```bash
$ twic profile ls
NAME                CERTIFICATE NAME    DOCKER HOST
docker1             tsa1                tcp://docker1.example.com:2376
docker2             tsa1                tcp://docker2.example.com:2376
$ twic profile rm docker2
$ twic profile ls
NAME                CERTIFICATE NAME    DOCKER HOST
docker1             tsa1                tcp://docker1.example.com:2376
```

## Related information

* [profile_add](profile_add.md)
* [profile_env](profile_env.md)
* [profile_ls](profile_ls.md)
* [profile_status](profile_status.md)
