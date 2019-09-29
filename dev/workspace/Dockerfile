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

CMD tail -f /dev/null