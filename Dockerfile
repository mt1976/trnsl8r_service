FROM golang:1.24.0 AS builder

# Set destination for COPY
WORKDIR /build



# Copy the source code. Note the slash at the end, as explained in
COPY main.go ./
COPY res ./res
COPY app ./app/
COPY data ./data/
COPY startupPayload ./startupPayload
COPY README.md ./

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o trnsl8r_service

FROM alpine:latest

# Copy the data/infra artifacts. Note the slash at the end, as explained in
COPY res ./res
COPY data ./data/
COPY startupPayload ./startupPayload
COPY README.md ./
# Copy the the application executable
COPY --from=builder /build/trnsl8r_service /trnsl8r_service

EXPOSE 5050
# Run
CMD ["./trnsl8r_service"]