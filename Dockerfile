#  FOR DEBUGGING
FROM golang:1.17 as debug

RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . /app

COPY ./dlv.sh /
RUN chmod +x /dlv.sh

CMD ["/dlv.sh"]


# FOR DEVELOPMENT
FROM golang:1.17 as dev

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . /app

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY ./dev.sh /
RUN chmod +x /dev.sh

CMD ["/dev.sh"]