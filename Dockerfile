FROM golang:1.15-alpine as build
COPY . /app
WORKDIR /app
RUN go build -o gool

FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/gool .
ENTRYPOINT ["./gool"]
