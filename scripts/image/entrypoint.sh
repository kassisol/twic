#!/bin/bash

TWIC="/usr/local/bin/twic"
DOCKER_TLS_DIR="/etc/docker/tls"

function report_error() {
	local msg=$1

	echo $msg
	exit 1
}

function report_unset_var() {
	local var=$1

	report_error "var \"${var}\" is not set"
}

if [ `ls -1 ${DOCKER_TLS_DIR}/ | wc -l` -gt 0 ]; then
	exit 0
fi

if [ ! -d $DOCKER_TLS_DIR ]; then
	report_error "Directory ${DOCKER_TLS_DIR} is not mounted"
fi

if [ -n "$METADATA_URL" ]; then
	METADATA_V1="${METADATA_URL}/metadata/v1"

	TWIC_TSA_URL=`curl -s ${METADATA_V1}/keys/tsa-url`
	TWIC_TOKEN=`curl -s ${METADATA_V1}/keys/twic-token`
	TWIC_CN=`curl -s ${METADATA_V1}/fqdn`

	TWIC_ALT_NAMES=""
	for t in `curl -s ${METADATA_V1}/interfaces/`; do
		for index in `curl -s ${METADATA_V1}/interfaces/$t`; do
			ip=`curl -s ${METADATA_V1}/interfaces/${t}${index}ipv4/address`
			if [ -z "$TWIC_ALT_NAMES" ]; then
				TWIC_ALT_NAMES="${ip}"
			else
				TWIC_ALT_NAMES="${TWIC_ALT_NAMES},${ip}"
			fi
		done
	done
fi

if [ -z "$TWIC_TSA_URL" ]; then
	report_unset_var TWIC_TSA_URL
fi

if [ -z "$TWIC_TOKEN" ]; then
	report_unset_var TWIC_TOKEN
fi

if [ -z "$TWIC_CN" ]; then
	report_unset_var TWIC_CN
fi

if [ -z "$TWIC_ALT_NAMES" ]; then
	report_unset_var TWIC_ALT_NAMES
fi
TWIC_ALT_NAMES="${TWIC_ALT_NAMES},127.0.0.1"

$TWIC engine create \
	--tsa-url $TWIC_TSA_URL \
	--token $TWIC_TOKEN \
	--common-name $TWIC_CN \
	--alt-names $TWIC_ALT_NAMES
