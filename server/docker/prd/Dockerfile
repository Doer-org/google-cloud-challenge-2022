## Build
FROM golang:1.19 AS build

WORKDIR /go/src/github.com/Doer-org/google-cloud-challenge-2022/server/src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /server

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/server"]
# TODO: cloud buildのpathを変える