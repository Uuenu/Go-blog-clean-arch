# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /cmd/app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /cmd/app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 3000
CMD [ "/cmd/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]