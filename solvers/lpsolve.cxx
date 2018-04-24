
#include "lp_lib.h"
#include "lpsolve.hpp"
#include <iostream>

using namespace std;

LPSolveSolver::LPSolveSolver()
{
}

LPSolveSolver::~LPSolveSolver()
{
    if (lp != NULL)
    {
        delete_lp(lp);
    }
}

void LPSolveSolver::showLog(bool shouldShow)
{
    if (shouldShow)
    {
        set_verbose(lp, FULL);
    }
    else
    {
        set_verbose(lp, NEUTRAL);
    }
}

void LPSolveSolver::setTimeLimit(double timeLimit)
{
    set_timeout(lp, timeLimit);
}

void LPSolveSolver::addVars(int count, double *lb, double *ub, char *types)
{
    lp = make_lp(0, count);
    set_verbose(lp, NEUTRAL);
    set_add_rowmode(lp, TRUE);
    numVars = count;

    for (size_t i = 0; i < count; i++)
    {
        set_bounds(lp, i + 1, lb[i], ub[i]);
        switch (types[i])
        {
            case 'I':
                set_int(lp, i + 1, TRUE);
                break;
            case 'B':
                set_binary(lp, i + 1, TRUE);
                break;
        }
    }

}

void LPSolveSolver::addConstr(
    int lhs_count, double *lhs_coeffs, uint64 *lhs_var_ids,
    double lhs_constant,
    int rhs_count, double *rhs_coeffs, uint64 *rhs_var_ids,
    double rhs_constant, char sense)
{
    int var_count = lhs_count + rhs_count;
    REAL sparse_row[var_count];
    int colno[var_count];

    for (size_t i = 0; i < lhs_count; i++)
    {
        sparse_row[i] = lhs_coeffs[i];
        colno[i] = (int) lhs_var_ids[i] + 1;
    }

    for (size_t i = 0; i < rhs_count; i++)
    {
        sparse_row[lhs_count + i] = -rhs_coeffs[i];
        colno[lhs_count + i] = (int) rhs_var_ids[i] + 1;
    }

    int constr_type;

    switch (sense)
    {
        case '=':
            constr_type = EQ;
            break;
        case '<':
            constr_type = LE;
            break;
        case '>':
            constr_type = GE;
            break;
    }

    REAL constant = rhs_constant - lhs_constant;
    add_constraintex(lp, var_count, sparse_row, colno, constr_type, constant);
}

void LPSolveSolver::setObjective(
    int count, double *coeffs, uint64 *var_ids, double constant, int sense)
{
    REAL sparse_row[count];
    int colno[count];

    for (size_t i = 0; i < count; i++)
    {
        sparse_row[i] = coeffs[i];
        colno[i] = (int) var_ids[i] + 1;
    }

    set_obj_fnex(lp, count, sparse_row, colno);

    switch (sense)
    {
        case 1:
            set_minim(lp);
            break;
        case -1:
            set_maxim(lp);
            break;
    }
}

MIPSolution LPSolveSolver::optimize()
{
    MIPSolution sol;
    set_add_rowmode(lp, false);
    int res = solve(lp);
    sol.optimal = res == OPTIMAL;
    sol.gap = get_mip_gap(lp, TRUE);
    sol.errorCode = res;
    sol.errorMessage = "No error messages provided for LPSolve";

    sol.values.resize(numVars);
    REAL vars[numVars];
    get_variables(lp, vars);

    for (size_t i = 0; i < numVars; i++)
    {
        sol.values.at(i) = (double) vars[i];
    }

    return sol;
}
