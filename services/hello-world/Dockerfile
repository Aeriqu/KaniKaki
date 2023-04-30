FROM golang:1.20.3-alpine3.17 AS build

WORKDIR /app

COPY go.mod  go.sum ./
RUN go mod download && go mod verify

COPY ./services/hello-world ./services/hello-world
RUN go build -o /usr/bin/hello-world ./services/hello-world


FROM alpine:3.17

COPY --from=build /usr/bin/hello-world /usr/bin/

EXPOSE 8080
ENTRYPOINT ["/usr/bin/hello-world"]