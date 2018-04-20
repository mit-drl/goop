package mip

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/mit-drl/goop/mip/solvers"
)

type Model struct {
	vars      []*Var
	constrs   []*Constr
	obj       *Objective
	showLog   bool
	timeLimit time.Duration
}

func NewModel() *Model {
	return &Model{showLog: false}
}

func (m *Model) ShowLog(shouldShow bool) {
	m.showLog = shouldShow
}

func (m *Model) SetTimeLimit(dur time.Duration) {
	m.timeLimit = dur
}

func (m *Model) AddVar(lower, upper float64, vtype VarType) *Var {
	id := uint64(len(m.vars))
	newVar := &Var{id, lower, upper, vtype}
	m.vars = append(m.vars, newVar)
	return newVar
}

func (m *Model) AddBinaryVar() *Var {
	return m.AddVar(0, 1, Binary)
}

func (m *Model) AddVarVector(
	num int, lower, upper float64, vtype VarType,
) []*Var {
	stId := uint64(len(m.vars))
	vs := make([]*Var, num)
	for i := range vs {
		vs[i] = &Var{stId + uint64(i), lower, upper, vtype}
	}

	m.vars = append(m.vars, vs...)
	return vs
}

func (m *Model) AddBinaryVarVector(num int) []*Var {
	return m.AddVarVector(num, 0, 1, Binary)
}

func (m *Model) AddVarMatrix(
	rows, cols int, lower, upper float64, vtype VarType,
) [][]*Var {
	vs := make([][]*Var, rows)
	for i := range vs {
		vs[i] = m.AddVarVector(cols, lower, upper, vtype)
	}

	return vs
}

func (m *Model) AddBinaryVarMatrix(rows, cols int) [][]*Var {
	return m.AddVarMatrix(rows, cols, 0, 1, Binary)
}

func (m *Model) AddConstr(constr *Constr) {
	m.constrs = append(m.constrs, constr)
}

func (m *Model) SetObjective(e Expr, sense ObjSense) {
	m.obj = NewObjective(e, sense)
}

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

	sol := NewSolution(mipSol)
	return sol, nil
}
