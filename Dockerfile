FROM golang:1.15.3-alpine3.12 as build

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/web/main.go

FROM golang:1.15.3-alpine3.12

WORKDIR /dist

RUN mkdir configs
RUN mkdir migrations

COPY --from=build /build/main .
COPY --from=build /build/internal/data/migrations /dist/migrations
COPY --from=build /build/configs /dist/configs

CMD ["/dist/main"]
