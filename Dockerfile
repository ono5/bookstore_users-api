FROM golang:latest

RUN go get github.com/cespare/reflex

WORKDIR /api
COPY ./api .
ENTRYPOINT ["reflex", "-c", "reflex.conf"]
