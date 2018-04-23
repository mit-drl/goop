package goop

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/mit-drl/goop/solvers"
)

// Model represents the overall constrained linear optimization model to be
// solved. Model contains all the variables associated with the optimization
// problem, constraints, objective, and parameters. New variables can only be
// created using an instantiated Model.
type Model struct {
	vars      []*Var
	constrs   []*Constr
	obj       *Objective
	showLog   bool
	timeLimit time.Duration
}

// NewModel returns a new model with some default arguments such as not to show
// the log and no time limit.
func NewModel() *Model {
	return &Model{showLog: false}
}

// ShowLog instructs the solver to show the log or not.
func (m *Model) ShowLog(shouldShow bool) {
	m.showLog = shouldShow
}

// SetTimeLimit sets the solver time limit for the model.
func (m *Model) SetTimeLimit(dur time.Duration) {
	m.timeLimit = dur
}

// AddVar adds a variable of a given variable type to the model given the lower
// and upper value limits. This variable is returned.
func (m *Model) AddVar(lower, upper float64, vtype VarType) *Var {
	id := uint64(len(m.vars))
	newVar := &Var{id, lower, upper, vtype}
	m.vars = append(m.vars, newVar)
	return newVar
}

// AddBinaryVar adds a binary variable to the model and returns said variable.
func (m *Model) AddBinaryVar() *Var {
	return m.AddVar(0, 1, Binary)
}

// AddVarVector adds a vector of variables of a given variable type to the
// model. It then returns the resulting slice.
func (m *Model) AddVarVector(
	num int, lower, upper float64, vtype VarType,
) []*Var {
	stID := uint64(len(m.vars))
	vs := make([]*Var, num)
	for i := range vs {
		vs[i] = &Var{stID + uint64(i), lower, upper, vtype}
	}

	m.vars = append(m.vars, vs...)
	return vs
}

// AddBinaryVarVector adds a vector of binary variables to the model and
// returns the slice.
func (m *Model) AddBinaryVarVector(num int) []*Var {
	return m.AddVarVector(num, 0, 1, Binary)
}

// AddVarMatrix adds a matrix of variables of a given type to the model with
// lower and upper value limits and returns the resulting slice.
func (m *Model) AddVarMatrix(
	rows, cols int, lower, upper float64, vtype VarType,
) [][]*Var {
	vs := make([][]*Var, rows)
	for i := range vs {
		vs[i] = m.AddVarVector(cols, lower, upper, vtype)
	}

	return vs
}

// AddBinaryVarMatrix adds a matrix of binary variables to the model and returns
// the resulting slice.
func (m *Model) AddBinaryVarMatrix(rows, cols int) [][]*Var {
	return m.AddVarMatrix(rows, cols, 0, 1, Binary)
}

// AddConstr adds a the given constraint to the model.
func (m *Model) AddConstr(constr *Constr) {
	m.constrs = append(m.constrs, constr)
}

// SetObjective sets the objective of the model given an expression and
// objective sense.
func (m *Model) SetObjective(e Expr, sense ObjSense) {
	m.obj = NewObjective(e, sense)
}

// Optimize optimizes the model using the given solver type and returns the
// solution or an error.
func (m *Model) Optimize(solverType solvers.SolverType) (*Solution, error) {
	lbs := make([]float64, len(m.vars))
	ubs := make([]float64, len(m.vars))
	types := new(bytes.Buffer)
	for i, v := range m.vars {
		lbs[i] = v.Lower()
		ubs[i] = v.Upper()
		types.WriteByte(byte(v.Type()))
	}

	mipModel := solvers.NewSolver(solverType)
	mipModel.ShowLog(m.showLog)

	if m.timeLimit > 0 {
		mipModel.SetTimeLimit(m.timeLimit.Seconds())
	}

	mipModel.AddVars(len(m.vars), &lbs[0], &ubs[0], types.String())

	for _, constr := range m.constrs {
		mipModel.AddConstr(
			constr.lhs.NumVars(),
			getCoeffsPtr(constr.lhs),
			getVarsPtr(constr.lhs),
			constr.lhs.Constant(),
			constr.rhs.NumVars(),
			getCoeffsPtr(constr.rhs),
			getVarsPtr(constr.rhs),
			constr.rhs.Constant(),
			byte(constr.sense),
		)
	}

	if m.obj != nil {
		mipModel.SetObjective(
			m.obj.NumVars(),
			getCoeffsPtr(m.obj),
			getVarsPtr(m.obj),
			m.obj.Constant(),
			int(m.obj.sense),
		)
	}

	mipSol := mipModel.Optimize()
	defer solvers.DeleteSolver(mipModel)

	if mipSol.GetErrorCode() != 0 {
		msg := fmt.Sprintf(
			"[Code = %d] %s",
			mipSol.GetErrorCode(),
			mipSol.GetErrorMessage(),
		)
		return nil, errors.New(msg)
	}

	sol := newSolution(mipSol)
	return sol, nil
}
