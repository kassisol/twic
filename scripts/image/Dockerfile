FROM alpine

ARG version

LABEL io.harbormaster.image.maintainer="hbm@kassisol.com"
LABEL io.harbormaster.image.version=$version
LABEL io.harbormaster.image.description="TWIC is an application for managing certificates to connect to the Docker daemon using TLS"

COPY build/twic /usr/local/bin/twic
COPY entrypoint.sh /entrypoint.sh

RUN apk --update --no-cache add \
	bash \
	curl

ENTRYPOINT ["/entrypoint.sh"]
