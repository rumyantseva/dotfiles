# echo "source ~/my-dot-configs/.go" >> ~/.zshrc 

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

export CDPATH=$CDPATH:$GOPATH/src:$GOPATH/src/github.com
