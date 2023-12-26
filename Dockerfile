# BUILD
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy

WORKDIR /app/cmd/
RUN go build -o telebot

# RUN IMAGE
FROM alpine
WORKDIR /app
COPY --from=build /app/cmd/telebot .

CMD ["./telebot"]