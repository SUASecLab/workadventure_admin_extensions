FROM golang:1.20-alpine

RUN addgroup -S extensions && adduser -S extensions -G extensions
USER extensions

WORKDIR /src/app
COPY --chown=extensions:extensions . .

RUN go get
RUN go install
