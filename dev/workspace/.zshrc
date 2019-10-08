# Lines to append to .zshrc from deltaepsilon/dotfiles
# Don't leave blank lines... it copies poorly
# Functions

# https://dave.cheney.net/2016/06/21/automatically-run-your-packages-tests-with-inotifywait
watch() { while inotifywait --exclude .swp -e modify -r .; do $@; done; }


# Exports
export GOPATH=/root/go
export PATH=$GOPATH/bin:$PATH
export PATH=/usr/local/go/bin:$PATH
export PATH=/root/go/bin:$PATH

# Aliases
alias ll="ls -al"
alias reload="source ~/.zshrc"
