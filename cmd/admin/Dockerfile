# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="E-ecommerce"

# Set the Current Working Directory inside the container
RUN mkdir -p /server-ecommmerce
WORKDIR /server-ecommmerce

CMD ls

COPY ./ go.mod go.sum ./

# Install dependencies
RUN go mod download


# Copy data to working dir
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o ./cmd/admin/main ./cmd/admin/main.go

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add tzdata zip ca-certificates

WORKDIR /server-ecommmerce

# Copy the Pre-built binary file from the previous stage
COPY --from=builder  /server-ecommmerce .

# Command to run the executable
CMD ["./cmd/admin/main"]