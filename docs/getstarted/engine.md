---
title: "Engine"
description: "Engine"
keywords: [ "getting started", "TWIC", "engine" ]
date: "2017-02-02"
url: "/docs/getstarted/twic/engine/"
menu:
  docs:
    parent: "getstarted_twic"
    weight: -85
github_edit: "https://github.com/kassisol/twic/edit/master/docs/getstarted/engine.md"
toc: true
---

## Docker Engine

Either the "admin" user or a member of the admin group (auth_group_admin) set during TSA configuration can manage the certificates for Docker engine.

### Create certificates
#### Standalone installation

Run the command bellow from the Docker host to generate the engine certificates.

> Make sure that the directory `/etc/docker/tls` is created.

```bash
# twic engine create
Common Name (CN) : HOST
Alt Names : HOST,CNAME,IPADDRESS,127.0.0.1
TSA URL : https://tsa1.example.com
Username : admin
Password : *******
```

> HOST and CNAME must be a FQDN!

#### RancherOS installation

You can create the twic token that will be needed to generate the Docker engine certificates. A token will be displayed on the screen. Copy that token, save it for next step.

> The token generated is valid only for 24 hours after creation.

```bash
$ twic access
TSA URL: https://tsa1.example.com
User: admin
Password: *******
```

Generate an engine certificate using this command (make sure you are using the latest version of kassisol/twic-engine:x.x.x):

```bash
$ sudo system-docker run -t --rm -v /etc/docker:/etc/docker -e "TWIC_TSA_URL=https://tsa1.example.com" -e "TWIC_TOKEN=" -e "TWIC_CN=docker1.example.com" -e TWIC_ALT_NAMES="192.168.10.100" kassisol/twic-engine:0.1.6
```

Restart docker:

```bash
$ sudo system-docker restart docker
```

### Recreate certificates

First revoke the certificate in TSA.

```bash
# tsa cert ls (look for serial number of the engine cert)
# tsa cert revoke <serial number>
```

On the Docker host, run the following commands as root to recreate the certificates.

```bash
$ docker node update --availability drain <node name>
$ sudo ros service stop docker
$ sudo system-docker rm -v docker
$ sudo rm -f /etc/docker/tls/*
```

Run the same command to generate the certificate on a RancherOS host. To finish, start docker service and activate the node back to the swarm cluster.

```bash
$ sudo ros service up docker
$ docker node update --availability active <node name>
```
