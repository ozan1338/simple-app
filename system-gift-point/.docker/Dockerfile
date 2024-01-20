FROM golang:1.20-alpine3.16 AS builder

WORKDIR /app

COPY . .

RUN apk --no-cache add bash \
	curl \
	git \
	gcc \
	g++ \
	tzdata

RUN go build -o system-gift-point-app

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/system-gift-point-app .

EXPOSE 8080

CMD ["./system-gift-point-app"]