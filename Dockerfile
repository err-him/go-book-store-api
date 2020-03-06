FROM mysql:8.0.19

COPY db.sql /docker-entrypoint-initdb.d

# Start from golang base image
FROM golang:alpine as builder

ENV GO111MODULE=auto

# Add Maintainer info
LABEL maintainer="Himanshu Gupta <hayhimanshu009@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN apk add --no-cache bash
# Set the current working directory inside the container
RUN mkdir /app

WORKDIR /app

RUN echo "$PWD"

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

COPY . .

RUN echo "$PWD"

# Build the Go app
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .


# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/


RUN echo "$PWD"

COPY --from=builder /app/main .
COPY --from=builder /app/db.sql .
COPY --from=builder /app/wait-for ./wait-for
COPY --from=builder /app/config ./config


# Expose port 8080 to the outside world
EXPOSE 9002

#Command to run the executable
CMD ["./main"]
