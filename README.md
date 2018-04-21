# Goop [![Go Report Card](https://goreportcard.com/badge/github.com/mit-drl/goop)](https://goreportcard.com/report/github.com/mit-drl/goop) [![Build Status](https://travis-ci.org/mit-drl/goop.svg?branch=master)](https://travis-ci.org/mit-drl/goop) [![Go Doc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=round-square)](https://godoc.org/github.com/mit-drl/goop)
General Linear Optimization in Go

# Installation

1. First get the code
```
mkdir -p $GOPATH/github.com/mit-drl && cd $GOPATH/github.com/mit-drl
git clone https://github.com/mit-drl/goop && cd goop
```

2. Next build install the dependencies
```
./install.sh
```

3. Follow the [instructions](#Solver Notes) for your solver of choice. Note,
currently only Gurobi is supported

4. Finally build the library
```
go build
```

5. (Optional) Test our installation
```
govendor test -v +local
```

# Solver Notes

Currently we only support Gurobi. Since Gurobi is proprietary, you need to
complete the following steps in order for the project to build

- You must have [Gurobi](http://www.gurobi.com/downloads/download-center)
installed and have a valid license.
- The `GUROBI_HOME` environment variable must be set to the home directory
of your Gurobi installation

# Testing
