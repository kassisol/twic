---
title: "twic profile status"
description: "The profile status command description and usage"
keywords: [ "profile", "status" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_profile"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/profile_status.md"
---

```markdown
Display Docker environment variables if set

Usage:
  twic profile status [flags]
```

## Examples

```bash
$ twic profile status
DOCKER_HOST=tcp://docker1.example.com:2376
DOCKER_TLS_VERIFY=1
DOCKER_CERT_PATH=/root/.twic/profiles/docker1
```

## Related information

* [profile_add](profile_add.md)
* [profile_env](profile_env.md)
* [profile_ls](profile_ls.md)
* [profile_rm](profile_rm.md)
