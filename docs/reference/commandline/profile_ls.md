---
title: "twic profile ls"
description: "The profile ls command description and usage"
keywords: [ "profile", "ls" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_profile"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/profile_ls.md"
---

```markdown
List Docker profiles

Usage:
  twic profile ls [flags]

Aliases:
  ls, list
```

## Examples

```bash
$ twic profile ls
NAME                CERTIFICATE NAME    DOCKER HOST
docker1             tsa1                tcp://docker1.example.com:2376
```

## Related information

* [profile_add](profile_add.md)
* [profile_env](profile_env.md)
* [profile_rm](profile_rm.md)
* [profile_status](profile_status.md)
