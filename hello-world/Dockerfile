FROM alpine:3.23
RUN apk update && apk add go
WORKDIR /app
COPY main.go .
RUN go build main.go
ENTRYPOINT ["./main"]