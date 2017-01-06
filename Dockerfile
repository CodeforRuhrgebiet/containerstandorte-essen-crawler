FROM golang:1.6
MAINTAINER Martin Schurig <martin@schurig.pw>

RUN mkdir /usr/app
WORKDIR /usr/app

RUN go get github.com/kellydunn/golang-geo
RUN go get github.com/yhat/scrape

ADD . /usr/app

ENTRYPOINT ["go", "run"]
CMD ["main.go"]
