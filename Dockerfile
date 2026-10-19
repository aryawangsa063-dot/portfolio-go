FROM golang:1.25-alpine AS builder

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 go build -o /out/server .

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=builder /out/server /server
COPY --from=builder /src/templates /templates
COPY --from=builder /src/static /static

EXPOSE 8000

CMD ["/server"]
