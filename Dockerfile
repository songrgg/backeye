FROM songrgg/alpine-debug

MAINTAINER songrgg <songrgg0.0@gmail.com>

ENV CONFIGOR_ENV=test
ADD server /
ADD conf/ /
ENTRYPOINT ["/server"]
