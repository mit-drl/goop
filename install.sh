#!/usr/bin/env bash

set -e

WORK=$(pwd)

go generate

if [[ $(uname) == 'Linux' ]]; then
    sudo apt install -y libpcre3 libpcre3-dev autotools-dev byacc \
        flex cmake build-essential autoconf
elif [[ $(uname) == 'Darwin' ]]; then
    brew install pcre autoconf
fi

mkdir -p .third_party
scripts/install_swig.sh

# Install govendor
cd $WORK
go get -u github.com/kardianos/govendor

# Install Go packages
govendor sync
