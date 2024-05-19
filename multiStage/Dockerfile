FROM golang:1.17-alpine AS build

WORKDIR /app
COPY . .
RUN go mod init app
RUN go build -o main .


FROM alpine:latest
WORKDIR /app
# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .
COPY --from=build /app/repos/database.json .
EXPOSE 8080
CMD ["./main"]
