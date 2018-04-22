---
title: "twic profile add"
description: "The profile add command description and usage"
keywords: [ "profile", "add" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_profile"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/profile_add.md"
---

```markdown
Add Docker profile

Usage:
  twic profile add [name] [flags]

Flags:
  -c, --cert-name string       Certificate Name
  -a, --docker-host string     Docker Host
  -p, --docker-port string     Docker Port (default "2376")
  -s, --docker-scheme string   Docker Scheme (default "tcp")
```

## Examples

```bash
$ twic profile add docker1
Certificate Name : tsa1
Docker Host : docker1.example.com
$ twic profile ls
NAME                CERTIFICATE NAME    DOCKER HOST
docker1             tsa1                tcp://docker1.example.com:2376
```

## Related information

* [profile_env](profile_env.md)
* [profile_ls](profile_ls.md)
* [profile_rm](profile_rm.md)
* [profile_status](profile_status.md)
