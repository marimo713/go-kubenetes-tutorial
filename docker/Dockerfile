############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git and gcc.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -v all
# Build the binary.
RUN go build -o /go/bin/main app/main.go 
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main
# Run the main binary.
WORKDIR /go/bin
ENTRYPOINT ["./main"]