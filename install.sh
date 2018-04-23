#!/usr/bin/env bash

set -e

WORK=$(pwd)

if [[ $(uname) == 'Linux' ]]; then
    sudo apt install -y libpcre3 libpcre3-dev autotools-dev byacc \
        flex cmake build-essential autoconf
elif [[ $(uname) == 'Darwin' ]]; then
    brew install pcre autoconf
fi

mkdir -p .third_party
scripts/install_swig.sh
scripts/install_lpsolve.sh

# Install govendor
cd $WORK
go get -u github.com/kardianos/govendor

# Install Go packages
govendor sync

ln -sf $GUROBI_HOME $WORK/.third_party/gurobi
