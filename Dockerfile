FROM golang:alpine

RUN apk update && apk add --no-cache git
RUN apk add chromium

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]