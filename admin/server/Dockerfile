FROM golang:1.23-alpine AS builder

ENV GO111MODULE on

# Build project
WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 \
  GOOS=linux \
  go build \
  -a \
  -mod vendor \
  -installsuffix cgo \
  -o app ./cmd/server

FROM alpine:latest

EXPOSE 6001

RUN addgroup -S 997 && adduser -S -g 997 997
USER 997

WORKDIR /app/
COPY --from=builder /go/src/app .
CMD ["./app"]
