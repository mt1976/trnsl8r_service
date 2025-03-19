# syntax=docker/dockerfile:1

FROM golang:1.24.0 AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./trnsl8r_service.run

# Deploy
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/trnsl8r_service.run .
COPY res ./res/
COPY data ./data/
COPY startupPayload ./startupPayload/
COPY README.md ./

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 5050

# Run
CMD ["./trnsl8r_service.run"]