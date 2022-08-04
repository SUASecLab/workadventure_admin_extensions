FROM golang:1.19-alpine

WORKDIR /src/app
COPY . .

RUN go get
RUN go install
