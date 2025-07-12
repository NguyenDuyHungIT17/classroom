FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux


WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY etc /app/etc
COPY templates /app/templates
RUN go build -ldflags="-s -w" -o /app/backend main.go

FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /app

COPY --from=builder /app/backend /app/backend
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/templates /app/templates

EXPOSE 8888

CMD ["./backend"]