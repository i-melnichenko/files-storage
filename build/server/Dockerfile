FROM golang:1.19 as builder
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o bin/server cmd/server/main.go

FROM scratch
COPY --from=builder /go/src/bin/server /
CMD ["/server"]
EXPOSE 8000
