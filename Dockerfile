FROM golang:latest

RUN mkdir /build 
WORKDIR /build

RUN export GO111MODULE=on

RUN go install github.com/cindy1408/gym/src@latest

RUN cd /build && git clone https://github.com/cindy1408/gym.git

RUN cd /build/gym/src && go build 

ENV PORT=8080

EXPOSE 8080