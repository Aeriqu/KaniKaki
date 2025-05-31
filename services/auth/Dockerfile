FROM golang:1.24.3-alpine3.21 AS build

WORKDIR /app

COPY go.mod  go.sum ./
RUN go mod download && go mod verify

COPY ./common ./common
COPY ./services/auth ./services/auth
RUN go build -o /usr/bin/auth ./services/auth


FROM alpine:3.21

COPY --from=build /usr/bin/auth /usr/bin/

EXPOSE 8080
ENTRYPOINT ["/usr/bin/auth"]