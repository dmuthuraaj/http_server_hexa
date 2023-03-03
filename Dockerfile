# Build Stage
FROM golang:latest as build
WORKDIR /app
COPY . .
RUN go build --ldflags '-linkmode external -extldflags "-static"' -o main ./cmd

# Run stage
FROM alpine:latest
EXPOSE 9100
WORKDIR /app
COPY --from=build /app/main .
CMD ["./main"]