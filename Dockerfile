FROM golang:1.18-alpine

WORKDIR /src/app
COPY . .

RUN go get
RUN go install
