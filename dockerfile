# Use Ubuntu 22.04 as the base image
FROM ubuntu:22.04

# Set environment variables for Go installation
ENV GOLANG_VERSION 1.18
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Update and install necessary packages
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y postgresql postgresql-contrib inotify-tools wget && \
    rm -rf /var/lib/apt/lists/*

# Install Go
RUN wget https://dl.google.com/go/go$GOLANG_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
    rm go$GOLANG_VERSION.linux-amd64.tar.gz

# Set up the working directory
WORKDIR $GOPATH

# Optional: Copy your Go application code and PostgreSQL scripts

# Optional: Set the default command or entrypoint

# Expose the PostgreSQL port
EXPOSE 5432

# Start the PostgreSQL service (or use an entrypoint script)
CMD ["postgres"]
