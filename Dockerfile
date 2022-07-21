FROM golang:1.18-alpine3.16 as build-stage

COPY ./ /go/src/as-seating-consumer
WORKDIR /go/src/as-seating-consumer

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o as-seating-consumer

FROM alpine:latest as production-stage

RUN apk --no-cache add ca-certificates

COPY --from=build-stage /go/src/as-seating-consumer /as-seating-consumer
WORKDIR /as-seating

ENTRYPOINT ["./as-seating-consumer"]