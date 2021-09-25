FROM golang:1.16.8 as reflex

ENV CGO_ENABLED=0

RUN set -ex \
 && go get github.com/cespare/reflex

WORKDIR /GoBB

COPY go.mod .
COPY go.sum .
RUN go mod download

CMD reflex -r '\.go$' --start-service=true -- sh -c 'go build -o ./local/build/app github.com/cjmarkham/GoBB/cmd/app && ./local/build/app'
