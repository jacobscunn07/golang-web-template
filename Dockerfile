FROM golang:1.14

RUN WORKDIRS="/app" && \
        echo "Creating working directories: $WORKDIRS..." && \
        mkdir -p $WORKDIRS && \
        ls -laR $WORKDIRS

WORKDIR /app

COPY . .

RUN go build ./cmd/web

ENTRYPOINT /app/web
