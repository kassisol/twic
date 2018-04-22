---
title: "Client"
description: "Client"
keywords: [ "getting started", "TWIC", "client" ]
date: "2017-02-02"
url: "/docs/getstarted/twic/client/"
menu:
  docs:
    parent: "getstarted_twic"
    weight: -85
github_edit: "https://github.com/kassisol/twic/edit/master/docs/getstarted/client.md"
toc: true
---

## Client
### Add a user certificate

```bash
$ twic cert add tsa1
TSA URL : https://tsa1.example.com
Username : user1
Password : ******
$ twic cert ls
NAME                TYPE                CN                  TSA URL                   EXPIRE
tsa1                client              user1               https://tsa1.example.com  2018-02-02
```

> TSA URL is https only

### Create a connection profile

Here you can name your profile with any meaningful name. You use the name of the certificate you created in the previous step. You also need the FQDN of the Docker back-end host you want to connect to:

```bash
$ twic profile add docker1
Certificate Name : tsa1
Docker Host : docker1.example.com
$ twic profile ls
NAME                CERTIFICATE NAME    DOCKER HOST
docker1             tsa1                tcp://docker1.example.com:2376
```

### Use profile
#### Set Env variables
##### Bash

```bash
$ twic profile env docker1
DOCKER_HOST=tcp://docker1.example.com:2376
DOCKER_TLS_VERIFY=1
DOCKER_CERT_PATH=/home/user1/.twic/profile/docker1
$ eval $(twic profile env docker1)
$ twic profile status
DOCKER_HOST=tcp://docker1.example.com:2376
DOCKER_TLS_VERIFY=1
DOCKER_CERT_PATH=/home/user1/.twic/profile/docker1
$ env | grep DOCKER
DOCKER_HOST=tcp://docker1.example.com:2376
DOCKER_TLS_VERIFY=1
DOCKER_CERT_PATH=/home/user1/.twic/profile/docker1
```

##### Tcsh

```bash
$ twic profile env -s tcsh docker1
setenv DOCKER_HOST tcp://docker1.example.com:2376
setenv DOCKER_TLS_VERIFY 1
setenv DOCKER_CERT_PATH /home/user1/.twic/profile/docker1

** Copy / paste the output of the above command **

$ twic profile status
DOCKER_HOST=tcp://docker1.example.com:2376
DOCKER_TLS_VERIFY=1
DOCKER_CERT_PATH=/home/user1/.twic/profile/docker1
```

#### Unset Env variables

```bash
$ eval $(twic profile env -u docker1)
```
