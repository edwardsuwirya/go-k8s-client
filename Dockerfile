FROM golang:alpine as build-env
RUN apk update && apk add --no-cache git

WORKDIR /src

COPY . .

RUN go mod tidy
RUN go build -o binary

FROM alpine
WORKDIR /app
COPY --from=build-env /src/binary /app
ENTRYPOINT ["./binary"]