FROM  golang:1.20 as builder

WORKDIR /tmp/tiny-golang-image

COPY . .

RUN go mod tidy \
    && go get -u -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o main main.go

FROM scratch
COPY --from=builder /tmp/tiny-golang-image /
CMD ["/main"]
