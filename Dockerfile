FROM golang:1.17 as go-builder

WORKDIR /api-mutant

ENV GOPROXY "https://proxy.golang.org"
ENV GOOS "linux"
ENV CGO_ENABLED "0"

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd cmd
COPY internal internal

WORKDIR /api-mutant/cmd/api

RUN go build -ldflags="-s -w" -o /main main.go

### Build Final Image ###

FROM scratch

COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-builder /main /api-mutant

CMD ["/api-mutant"]
