---
title: "twic profile env"
description: "The profile env command description and usage"
keywords: [ "profile", "env" ]
date: "2017-02-14"
menu:
  docs:
    parent: "twic_cli_profile"
github_edit: "https://github.com/kassisol/twic/edit/master/docs/reference/commandline/profile_env.md"
---

```markdown
Set / Unset Docker environment variables

Usage:
  twic profile env [name] [flags]

Flags:
  -s, --shell string   Force environment to be configured for a specified shell: (tcsh, bash) (default "bash")
  -u, --unset          Unset variables instead of setting them
```

## Examples

```bash
$ twic profile ls
NAME                CERTIFICATE NAME    DOCKER HOST
docker1             tsa1                tcp://docker1.example.com:2376
```

To use with the Bash shell which is default.

```bash
$ twic profile env docker1

export DOCKER_HOST=tcp://docker1.example.com:2376
export DOCKER_TLS_VERIFY=1
export DOCKER_CERT_PATH=/root/.twic/profiles/docker1
$ eval $(twic profile env docker1)
```

Or with TCSH.

```bash
$ twic profile env -s tcsh docker1

set DOCKER_HOST tcp://docker1.example.com:2376
set DOCKER_TLS_VERIFY 1
set DOCKER_CERT_PATH /root/.twic/profiles/docker1
```

## Related information

* [profile_add](profile_add.md)
* [profile_ls](profile_ls.md)
* [profile_rm](profile_rm.md)
* [profile_status](profile_status.md)
