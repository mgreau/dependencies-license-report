FROM alpine:latest

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

ENV PROJECT_NAME licenses-report
ENV PROJECT_PATH /go/src/github.com/mgreau/$PROJECT_NAME

RUN	apk add --no-cache \
	ca-certificates

COPY src /go/src

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		go \
		git \
		gcc \
		libc-dev \
		libgcc

# Build app and configure user
RUN set -o errexit -o nounset \
	&& cd $PROJECT_PATH \
	&& go get github.com/ryanuber/go-license \
	&& go get github.com/mitchellh/colorstring \
	&& go get gopkg.in/yaml.v2 \
	&& go get github.com/Sirupsen/logrus \
	&& go get gopkg.in/urfave/cli.v1 \
	&& go build -o /usr/bin/$PROJECT_NAME \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete." \
	&& echo "Adding go user and group" \
	&& addgroup -S -g 1000 golang \
	&& adduser -D -S -G golang -u 1000 -s /bin/ash golang \
	&& mkdir -p /home/golang/.go \
	&& mkdir -p /home/golang/reports \
	&& chown -R golang:golang /home/golang \
	&& chown golang:golang /usr/bin/$PROJECT_NAME

USER golang

WORKDIR /home/golang

ENTRYPOINT [ "licenses-report" ]