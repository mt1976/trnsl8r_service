# syntax=docker/dockerfile:1

FROM golang:1.24.0

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY main.go ./
COPY res ./res
COPY app ./app/
COPY data ./data/
COPY startupPayload ./startupPayload/
COPY README.md ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./trnsl8r_service

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 5050

# Run
CMD ["./trnsl8r_service"]