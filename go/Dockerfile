FROM golang:1.17-alpine AS build

WORKDIR /build
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN CGO_ENABLED=0 go build -o apps_discoverer
#GOOS=linux GOARCH=arm64
FROM alpine

WORKDIR /app

COPY --from=build /build/apps_discoverer apps_discoverer

CMD ["/app/apps_discoverer"]
