#!/usr/bin/env bash

cd .third_party
wget https://github.com/swig/swig/archive/rel-3.0.12.tar.gz
tar -xzf rel-3.0.12.tar.gz
cd swig-rel-3.0.12
./autogen.sh && ./configure && make && sudo make install
