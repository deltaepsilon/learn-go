FROM ubuntu:latest

# Upgrade the OS
RUN apt-get update && apt-get -y upgrade

# Install base dependencies
RUN apt install -y vim curl git-core zsh build-essential libssl-dev inotify-tools

# Install Go
WORKDIR /root
RUN curl https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz --output go.tar.gz 
RUN tar -zxvf go.tar.gz && mv go /usr/local

# Install Oh My ZSH
WORKDIR /root
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)" || true && rm .zshrc

# Dotfiles
RUN git clone https://github.com/deltaepsilon/dotfiles.git ~/dotfiles
RUN git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim
WORKDIR /root/dotfiles
RUN ./setup.sh
WORKDIR /root
RUN sed -i 's/\/Users\/quiver/\/root/g' .zshrc
COPY .zshrc .append-to-zshrc
RUN sed -i 's/\r//' .append-to-zshrc
RUN cat .append-to-zshrc >> .zshrc

ARG GOPATH=/root/go
ENV GOPATH=${GOPATH}
WORKDIR ${GOPATH}
RUN /usr/local/go/bin/go get -u \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/ramya-rao-a/go-outline \
  github.com/rogpeppe/godef \
  github.com/sqs/goreturns \
  github.com/stamblerre/gocode \
  github.com/uudashr/gopkgs/cmd/gopkgs \
  golang.org/x/lint/golint \
  golang.org/x/tools/cmd/goimports \
  golang.org/x/tools/cmd/gorename \
  golang.org/x/tools/cmd/guru \
  golang.org/x/tools/gopls \
  && echo "Done!"

WORKDIR /app
RUN /usr/local/go/bin/go get -u \
  github.com/onsi/ginkgo/ginkgo 

CMD tail -f /dev/null