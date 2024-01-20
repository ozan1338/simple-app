FROM golang:1.20-alpine3.16 AS builder

WORKDIR /app

COPY . .

RUN apk --no-cache add bash \
	curl \
	git \
	gcc \
	g++ \
	tzdata

RUN go build -o favourite-book-app

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/favourite-book-app .

EXPOSE 8080

CMD ["./favourite-book-app"]