#!/usr/bin/env bash

set -e

WORK=$(pwd)

if [[ $(uname) == 'Linux' ]]; then
    sudo apt install -y libpcre3 libpcre3-dev autotools-dev byacc \
        flex cmake build-essential autoconf
    # Original Method for Installation
    scripts/install_swig.sh
    scripts/install_lpsolve.sh

    mkdir -p .third_party

    ln -sf $GUROBI_HOME $WORK/.third_party/gurobi

elif [[ $(uname) == 'Darwin' ]]; then
    brew install pcre autoconf
    brew install swig 
    brew install lp_solve 

    go run scripts/make_lib.go --go-fname solvers/lib.go --pkg solvers
fi

# Install Go packages
go mod init github.com/mit-drl/goop
