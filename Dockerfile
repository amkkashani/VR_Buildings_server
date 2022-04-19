FROM golang:1.16


ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir /app
COPY .  /app

WORKDIR /app

RUN go build -o  main .

#ENTRYPOINT [ "/bin/bash", "-c"]

CMD ["/app/main"]