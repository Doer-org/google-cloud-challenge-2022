FROM golang:1.19 AS dev

WORKDIR /go/src/github.com/Doer-org/google-cloud-challenge-2022/server/src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/github.com/Doer-org/google-cloud-challenge-2022/server/src/infrastructure/ent

CMD [ "go","generate","./..." ]
