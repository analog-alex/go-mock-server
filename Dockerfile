FROM golang:1.16-alpine

# set up a workdir
WORKDIR /app

COPY src /app/src
COPY main.go /app
COPY go.mod /app

RUN go build -o /mock

CMD [ "/mock" ]