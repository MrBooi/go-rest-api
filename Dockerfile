# have required goland tools
FROM golang:1.18 AS builder 
# create directory for the project
RUN mkdir /app
ADD . /app
WORKDIR /app

# build binary 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["/app"]
