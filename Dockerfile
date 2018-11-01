FROM golang:1.10 AS builder

# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/vincedgy/databases
ADD src/github.com/vincedgy/databases ./
ADD src/github.com/vincedgy/databases/* ./

# Run the dep for dependencies resolution
RUN dep ensure --vendor-only

# Run the executable
ADD bin ./

# Compile statically
RUN CGO_ENABLED=0 GOOS=linux nice go build -a -installsuffix nocgo -o /app .

# The real RUNNING container with our go program is here
FROM scratch
COPY --from=builder /app ./
ENTRYPOINT ["./app"]