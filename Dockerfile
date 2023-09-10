FROM golang:1.21.1-alpine3.18 AS build-stage

ENV GOOS=linux

WORKDIR /go/src/test_quadro

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /test_quadro  /go/src/test_quadro/app/.

FROM build-stage AS test-stage
RUN go test -v ./...

FROM scratch
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-stage main /bin/main

ENTRYPOINT [ "/test_quadro" ]
