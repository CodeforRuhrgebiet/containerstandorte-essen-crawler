FROM golang:1.6
MAINTAINER Martin Schurig <martin@schurig.pw>

RUN mkdir /usr/app
WORKDIR /usr/app

RUN go get github.com/nicostuhlfauth/geoosm
RUN go get github.com/yhat/scrape
RUN go get golang.org/x/net/html
RUN go get golang.org/x/net/html/atom

ADD . /usr/app

ENTRYPOINT ["go", "run"]
CMD ["main.go"]
