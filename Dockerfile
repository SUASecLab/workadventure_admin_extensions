FROM golang:1.22-alpine as golang-builder

RUN addgroup -S extensions && adduser -S extensions -G extensions

WORKDIR /src/app
COPY --chown=extensions:extensions . .

RUN go get
RUN go build -o extensions-bin

FROM scratch
COPY --from=golang-builder /src/app/extensions-bin /extensions
COPY --from=golang-builder /etc/passwd /etc/passwd

USER extensions
CMD [ "/extensions" ]
