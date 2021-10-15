# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

################################################################################
# Build Go server binary

WORKDIR /build/bin

# Fetch dependencies
COPY src/bin/go.mod .
COPY src/bin/go.sum .
RUN go mod download

COPY src/bin/*.go .
RUN go build -o ./../../out/bin/go .

################################################################################
# Copy over resources

COPY out/rsc/ /out/rsc/

################################################################################
# Run built binary

WORKDIR /
EXPOSE 80
EXPOSE 443
EXPOSE 8080
CMD ["./out/bin/go"]