FROM golang:1.24.4-alpine3.22 AS build

WORKDIR /usr/src/app

COPY go.mod  go.sum ./
RUN go mod download && go mod verify

COPY ./common ./common
COPY ./services/wanikani ./services/wanikani
RUN go build -o /usr/bin/wanikani ./services/wanikani


FROM alpine:3.22

COPY --from=build /usr/bin/wanikani /usr/bin/

EXPOSE 8080
ENTRYPOINT ["/usr/bin/wanikani"]