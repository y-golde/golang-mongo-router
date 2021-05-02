FROM golang:latest

LABEL maintainer="y-golde <yon.golde@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 1948

RUN go build

CMD [ "./helloworld" ]
