# Date: 
# Dockerfile Reference: https://docs.docker.com/engine/reference/builder/

#################################################################################################################
# Build Docker Image
#################################################################################################################
FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build

# Build the binary
RUN go build -o main .

#################################################################################################################
# Build a small image
# Alpine Docker Official Image
#     A building block for Alpine Linux Docker containers.
#     The image includes source code, libraries, tools, and other core dependencies that your application needs.
#################################################################################################################
FROM alpine

# Add new user
RUN adduser -S -D -H -h /app appuser
USER appuser

# Copy dependencies
COPY . /app
COPY --from=builder /build/main /app/
WORKDIR /app

# Inform Docker that the container listens on the network port "8000" at runtime.
EXPOSE 8000

# Provide default (executabe) for an executing container.
CMD ["./main"]
