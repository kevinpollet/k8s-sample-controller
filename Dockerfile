FROM golang:1.17-alpine3.15 as builder
WORKDIR /k8s-sample-controller

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o ./bin/controller ./main.go

FROM alpine:3.15
COPY --from=builder /k8s-sample-controller/bin/controller /usr/local/bin
ENTRYPOINT ["controller"]
