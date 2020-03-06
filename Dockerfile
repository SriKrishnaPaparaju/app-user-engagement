FROM registry.svc.ci.openshift.org/openshift/release:golang-1.12 AS buildsource
WORKDIR /go/src/github.com/redhat-developer/app-user-engagement
COPY . .
RUN go mod tidy
RUN go mod vendor
# Now build the source code
RUN make build && make install

FROM docker.io/library/centos:7
RUN INSTALL_PKGS=" \
      bind-utils bsdtar findutils git hostname lsof socat \
      sysvinit-tools tar tree util-linux wget which \
      " && \
    yum install -y --setopt=skip_missing_names_on_install=False ${INSTALL_PKGS} && \
    yum clean all

COPY imagecontent/policy.json /etc/containers/
COPY imagecontent/registries.conf /etc/containers/
COPY imagecontent/storage.conf /etc/containers/
RUN mkdir -p /var/cache/blobs \
    /var/lib/shared/overlay-images \
    /var/lib/shared/overlay-layers && \
    touch /var/lib/shared/overlay-images/images.lock \
    /var/lib/shared/overlay-layers/layers.lock

COPY --from=buildsource /go/src/github.com//redhat-developer/app-user-engagement/userengagement /usr/bin
EXPOSE 8080
CMD ["/usr/bin/userengagement"]
