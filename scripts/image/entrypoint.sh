#!/bin/bash

TWIC="/usr/local/bin/twic"
DOCKER_TLS_DIR="/etc/docker/tls"

function report_error() {
	local var=$1

	echo "var \"${var}\" is not set"
	exit 1
}

if [ `ls -1 ${DOCKER_TLS_DIR}/ | wc -l` -gt 0 ]; then
	exit 0
fi

if [ ! -d $DOCKER_TLS_DIR ]; then
	mkdir $DOCKER_TLS_DIR
fi

if [ -n $METADATA_URL ]; then
	TWIC_TSA_URL=`curl -s ${METADATA_URL}/key/tsa-url`
	TWIC_TOKEN=`curl -s ${METADATA_URL}/key/twic-token`
	TWIC_CN=`curl -s ${METADATA_URL}/fqdn`

	TWIC_ALT_NAMES=""
	for t in `curl -s ${METADATA_URL}/interfaces/`; do
		for index in `curl -s ${METADATA_URL}/interfaces/$t`; do
			ip=`curl -s ${METADATA_URL}/interfaces/${t}${index}address`
			if [ $TWIC_ALT_NAMES == "" ]; then
				TWIC_ALT_NAMES="${ip}"
			else
				TWIC_ALT_NAMES="${TWIC_ALT_NAMES},${ip}"
			fi
		done
	done
	TWIC_ALT_NAMES="${TWIC_ALT_NAMES},127.0.0.1"
else
	if [ -z $TWIC_TSA_URL ]; then
		report_error TWIC_TSA_URL
	fi

	if [ -z $TWIC_TOKEN ]; then
		report_error TWIC_TOKEN
	fi

	if [ -z $TWIC_CN ]; then
		report_error TWIC_CN
	fi

	if [ -z $TWIC_ALT_NAMES ]; then
		report_error TWIC_ALT_NAMES
	fi
fi

$TWIC engine create \
	--tsa-url $TWIC_TSA_URL \
	--token $TWIC_TOKEN \
	--common-name $TWIC_CN \
	--alt-names $TWIC_ALT_NAMES
