
#!/bin/bash
set -ex

dir=$1
mkdir -p $dir && cd $dir


go mod init github.com/HuanLiu-hotstar/$dir

go get github.com/99designs/gqlgen


go run github.com/99designs/gqlgen init
