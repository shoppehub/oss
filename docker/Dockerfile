FROM alpine:3.13



ADD cmd /usr/local/cmd
ADD config /usr/local/config


WORKDIR /usr/local/


ENV  GIN_MODE=release

CMD  "./cmd"
