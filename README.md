# Goop [![Go Report Card](https://goreportcard.com/badge/github.com/mit-drl/goop)](https://goreportcard.com/report/github.com/mit-drl/goop) [![Build Status](https://travis-ci.org/mit-drl/goop.svg?branch=master)](https://travis-ci.org/mit-drl/goop) [![Go Doc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=round-square)](https://godoc.org/github.com/mit-drl/goop)

General Linear Optimization in Go. `goop` provides general interface for solving
mixed integer linear optimization problems using a variety of back-end solvers.

# Quickstart

We are going to start with a simple example showing how `goop` can be used to
solve integer linear programs. The example below seeks to maximize the following
MIP:

```
maximize    x +   y + 2 z
subject to  x + 2 y + 3 z <= 4
            x +   y       >= 1
x, y, z binary
```

This is is the same example implemented [here](http://www.gurobi.com/documentation/7.5/examples/mip1_py.html). Below
we have implemented the model using `goop` and have optimized the model using
the supported Gurobi solver.

```go
package main

import (
    "fmt"
    "github.com/mit-drl/goop"
    "github.com/mit-drl/goop/solvers"
)

func main() {
    // Instantiate a new model
    m := goop.NewModel()

    // Add your variables to the model
    x := m.AddBinaryVar()
    y := m.AddBinaryVar()
    z := m.AddBinaryVar()

    // Add your constraints
    m.AddConstr(goop.Sum(x, y.Mult(2), z.Mult(3)).LessEq(goop.K(4)))
    m.AddConstr(goop.Sum(x, y).GreaterEq(goop.One))

    // Set a linear objective using your variables
    obj := goop.Sum(x, y, z.Mult(2))
    m.SetObjective(obj, goop.SenseMaximize)

    // Optimize the variables according to the model
    sol, err := m.Optimize(solvers.Gurobi)

    // Check if there is an error from the solver. No error should be returned
    // for this model
    if err != nil {
    	panic("Should not have an error")
    }

    // Print out the solution
    fmt.Println("x =", sol.Value(x))
    fmt.Println("y =", sol.Value(y))
    fmt.Println("z =", sol.Value(z))

    // Output:
    // x = 1
    // y = 0
    // z = 1
}
```

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
