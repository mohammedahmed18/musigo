FROM ubuntu:22.04

ENV GO_VERSION=${1.22.4}

# Install necessary dependencies
RUN apt-get update && \
    apt-get install -y \
    wget \
    fuse \
    git \
    gcc \
    fontconfig \
    unzip \
    curl \
    ripgrep


RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"


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
RUN git clone https://github.com/NvChad/starter ~/.config/nvim

COPY nvim-config/lspconfig.lua /root/.config/nvim/lua/configs
COPY nvim-config/null-ls.lua /root/.config/nvim/lua/configs
COPY nvim-config/init-plugins.lua /root/.config/nvim/lua/plugins/init.lua

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
