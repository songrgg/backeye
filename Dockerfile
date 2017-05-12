FROM songrgg/alpine-debug

MAINTAINER songrgg <songrgg0.0@gmail.com>

WORKDIR /usr/src/app

ENV CONFIGOR_ENV=test
ADD server .
ADD conf/ conf
ADD public/ public
EXPOSE 9876
ENTRYPOINT ["/server"]
