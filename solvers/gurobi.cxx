
#include <iostream>
#include "gurobi_c++.h"
#include "gurobi.hpp"

using namespace std;

GurobiSolver::GurobiSolver() : env(GRBEnv()), model(env)
{
}

GurobiSolver::~GurobiSolver()
{
    delete[] vars;
}

void GurobiSolver::setMIPGapTol(double gap)
{
    model.set(GRB_DoubleParam_MIPGap, gap);
}

void GurobiSolver::showLog(bool shouldShow)
{
    if (shouldShow)
    {
        model.getEnv().set(GRB_IntParam_OutputFlag, 1);
    }
    else
    {
        model.getEnv().set(GRB_IntParam_OutputFlag, 0);
    }

}

void GurobiSolver::setTimeLimit(double timeLimit)
{
    model.getEnv().set(GRB_DoubleParam_TimeLimit, timeLimit);
}

void GurobiSolver::addVars(int count, double *lb, double *ub, char *types)
{
    vars = model.addVars(lb, ub, NULL, types, NULL, count);
    numVars = count;
    model.update();
}

void GurobiSolver::addConstr(
        int lhs_count, double *lhs_coeffs, uint64 *lhs_var_ids,
        double lhs_constant,
        int rhs_count, double *rhs_coeffs, uint64 *rhs_var_ids,
        double rhs_constant, char sense)
{
    GRBLinExpr lhs_expr = lhs_constant;
    GRBLinExpr rhs_expr = rhs_constant;
    GRBVar lhs_vars[lhs_count];
    GRBVar rhs_vars[rhs_count];

    for (int i = 0; i < lhs_count; i++)
    {
        lhs_vars[i] = vars[lhs_var_ids[i]];
    }

    for (int i = 0; i < rhs_count; i++)
    {
        rhs_vars[i] = vars[rhs_var_ids[i]];
    }

    lhs_expr.addTerms(lhs_coeffs, lhs_vars, lhs_count);
    rhs_expr.addTerms(rhs_coeffs, rhs_vars, rhs_count);
    model.addConstr(lhs_expr, sense, rhs_expr);

}

void GurobiSolver::setObjective(int count, double *coeffs, uint64 *var_ids,
        double constant, int sense)
{
    GRBLinExpr expr = constant;
    GRBVar *vs = new GRBVar[count];

    for (int i = 0; i < count; i++)
    {
        vs[i] = vars[var_ids[i]];
    }

    expr.addTerms(coeffs, vs, count);
    model.setObjective(expr, sense);
    delete[] vs;
}

MIPSolution GurobiSolver::optimize()
{
    MIPSolution sol;

    try
    {
        model.update();
        model.optimize();
        sol.values.resize(numVars);

        for (int i = 0; i < numVars; i++)
        {
            sol.values.at(i) = vars[i].get(GRB_DoubleAttr_X);
        }

        sol.obj = model.get(GRB_DoubleAttr_ObjVal);
        sol.gap = model.get(GRB_DoubleAttr_MIPGap);
        sol.optimal = model.get(GRB_IntAttr_Status) == GRB_OPTIMAL;
        sol.errorCode = 0;
        sol.errorMessage = "No error";
    }
    catch (GRBException e)
    {
        sol.errorCode = e.getErrorCode();
        sol.errorMessage = e.getMessage();
        cerr << "Code: " << e.getErrorCode() << endl;
        cerr << "Message: " << e.getMessage() << endl;
    }

    return sol;
}
