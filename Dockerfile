#  FOR DEBUGGING
FROM golang:1.17 as debug

RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /go/src/go-ambassador
COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . /go/src/go-ambassador/

COPY ./dlv.sh /
RUN chmod +x /dlv.sh

CMD ["/dlv.sh"]


# FOR DEVELOPMENT
FROM golang:1.17 as dev

WORKDIR /go/src/go-ambassador

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]