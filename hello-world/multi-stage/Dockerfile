FROM alpine:3.23 AS build
RUN apk update && apk add go
WORKDIR /app
COPY main.go .
RUN go build main.go


FROM alpine:3.23 AS run
WORKDIR /app
COPY --from=build /app/main .
ENTRYPOINT ["./main"]