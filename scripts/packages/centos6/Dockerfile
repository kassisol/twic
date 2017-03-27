FROM centos:6

COPY entrypoint.sh /entrypoint.sh
COPY build /usr/local/src/twic

ENV DAPPER_SOURCE /tmp
ENV DAPPER_OUTPUT dist
ENV SHELL /bin/bash

WORKDIR ${DAPPER_SOURCE}

ENV RPMBUILD_PATH="/srv/rpmbuild"

RUN build="rpm-build" \
	&& set -x \
	&& yum -y install $build \
	&& yum clean all

RUN mkdir -p ${RPMBUILD_PATH} \
	&& mkdir ${RPMBUILD_PATH}/BUILD \
	&& mkdir ${RPMBUILD_PATH}/RPMS \
	&& mkdir ${RPMBUILD_PATH}/SOURCES \
	&& mkdir ${RPMBUILD_PATH}/SPECS \
	&& mkdir ${RPMBUILD_PATH}/SRPMS \
	&& mkdir ${RPMBUILD_PATH}/tmp \
	&& echo "%_topdir ${RPMBUILD_PATH}" > /root/.rpmmacros \
	&& echo "%_tmppath ${RPMBUILD_PATH}/tmp" >> /root/.rpmmacros

COPY twic.spec ${RPMBUILD_PATH}/SPECS/twic.spec

RUN set -x \
	&& tar cvzf ${RPMBUILD_PATH}/SOURCES/twic.tar.gz -C /usr/local/src twic	 \
	&& rm -rf /usr/local/src/twic

ENTRYPOINT ["/entrypoint.sh"]
