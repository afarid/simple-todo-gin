FROM golang:1.17-alpine as builder
RUN mkdir /build
ADD .  /build/
WORKDIR /build
RUN apk add -U --no-cache ca-certificates
RUN apk add  --no-cache git
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

# generate clean, final image for end users
FROM scratch as prod

COPY --from=builder /build/main /app/main
COPY --from=builder /build/.env /app/.env
COPY --from=builder /build/db/migrations /app/db/migrations
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app

ENTRYPOINT [ "./main" ]