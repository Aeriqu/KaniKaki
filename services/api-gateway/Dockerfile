FROM golang:1.24.3-alpine3.21 AS build

WORKDIR /app

COPY go.mod  go.sum ./
RUN go mod download && go mod verify

COPY ./common ./common
COPY ./services/api-gateway ./services/api-gateway
COPY ./services/auth/proto ./services/auth/proto
COPY ./services/kanji/proto ./services/kanji/proto
RUN go build -o /usr/bin/api-gateway ./services/api-gateway


FROM alpine:3.21

ARG GRAPHQL_SCHEMA_ROOT=/etc/api-gateway/graphql
COPY ./services/api-gateway/schema/graphql /etc/api-gateway/graphql/
COPY --from=build /usr/bin/api-gateway /usr/bin/

EXPOSE 8080
ENTRYPOINT ["/usr/bin/api-gateway"]