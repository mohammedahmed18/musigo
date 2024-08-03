# Use the official golang image as the base image
FROM golang:1.22.4

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Install Redis
RUN apt-get update && apt-get install -y redis-server


# Install SSH server
RUN apt-get update && \
 apt-get install -y openssh-server

 # Allow SSH access
RUN mkdir /var/run/sshd

# Set the working directory
WORKDIR /workspace

# Install necessary Go tools
RUN go install golang.org/x/tools/gopls@latest

# Copy the project into the container
COPY . /workspace
