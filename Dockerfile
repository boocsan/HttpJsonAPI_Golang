FROM golang:latest

LABEL MAINTAINER iedred7584

WORKDIR /

ADD main.go /

EXPOSE 8888
CMD ["go", "run", "main.go"]
