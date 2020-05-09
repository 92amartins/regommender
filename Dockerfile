FROM golang:1.14

WORKDIR /go/src/github.com/regommender
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build

CMD ["go", "run", "main.go"]