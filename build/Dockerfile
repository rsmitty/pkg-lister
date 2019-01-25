FROM ubuntu:latest

ADD ./bin/pkg-lister-linux pkg-lister

RUN chmod +x pkg-lister

ENTRYPOINT ["./pkg-lister"]