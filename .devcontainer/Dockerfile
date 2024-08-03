# Use the official golang image as the base image
FROM golang:1.22.4


# Install necessary dependencies
RUN apt-get update && \
    apt-get install -y \
    wget \
    fuse \
    fontconfig \
    unzip \
    curl \
    ripgrep

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# install yarn
Run npm i -g yarn

# install redis
RUN apt-get update && apt-get install -y redis-server 

#install nvim
RUN curl -LO https://github.com/neovim/neovim/releases/download/v0.10.1/nvim-linux64.tar.gz
RUN rm -rf /opt/nvim
RUN tar -C /opt -xzf nvim-linux64.tar.gz
ENV PATH="/opt/nvim-linux64/bin:${PATH}"
Run chmod +x /opt/nvim-linux64/bin/nvim 

# copy nvim configuration
RUN mkdir -p ~/.config/nvim
ADD nvim-config /root/.config/nvim

RUN wget https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/FiraCode.zip -O /tmp/FiraCode.zip && \
unzip /tmp/FiraCode.zip -d /usr/share/fonts && \
fc-cache -fv && \
rm /tmp/FiraCode.zip

# Set up environment variables
ENV PATH="/usr/local/bin:$PATH"
# Set the working directory
WORKDIR /workspace

# Install necessary Go tools
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/air-verse/air@latest


# Copy the project into the container
COPY . /workspace
