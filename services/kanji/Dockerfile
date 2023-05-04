FROM golang:1.20.3-alpine3.17 AS build

WORKDIR /usr/src/app

COPY go.mod  go.sum ./
RUN go mod download && go mod verify

COPY ./common ./common
COPY ./services/wanikani/proto ./services/wanikani/proto
COPY ./services/kanji ./services/kanji
RUN go build -o /usr/bin/kanji ./services/kanji


FROM alpine:3.17

COPY --from=build /usr/bin/kanji /usr/bin/

EXPOSE 8080
ENTRYPOINT ["/usr/bin/kanji"]