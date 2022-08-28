FROM golang:1.19-alpine

RUN addgroup -S extensions && adduser -S extensions -G extensions
USER extensions

WORKDIR /src/app
COPY . .

RUN go get
RUN go install
