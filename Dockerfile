FROM golang:1.19-alpine

# set up a workdir
WORKDIR /app

COPY pkg /app/pkg
COPY cmd/main.go /app
COPY go.mod /app

RUN go build -o /mock

CMD [ "/mock" ]