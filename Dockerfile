FROM golang:latest

EXPOSE 32147

WORKDIR /usr/app/

COPY . .

CMD ["make"]

FROM alpine:latest

ENTRYPOINT [ "./server" ]
