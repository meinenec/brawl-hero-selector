FROM alpine as certs

RUN apk add --update --no-cache ca-certificates

FROM golang:latest as build

COPY . /heroes-bot

WORKDIR /heroes-bot

RUN CGO_ENABLED=0 go build -o release/heroes-bot

FROM scratch 

WORKDIR /heroes-data

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /heroes-bot/release/heroes-bot /bin/heroes-bot
COPY --from=build /heroes-bot/heroes.json /heroes-data/heroes.json

ENTRYPOINT [ "/bin/heroes-bot" ]