#ifndef GOOP_LPSOLVE_HPP_
#define GOOP_LPSOLVE_HPP_

#include "base_solver.hpp"
#include "lp_lib.h"

class LPSolveSolver : public Solver
{
public:
    LPSolveSolver();
    ~LPSolveSolver();
    void addVars(int count, double *lb, double *ub, char *types);
    void addConstr(
        int lhs_count, double *lhs_coeffs,
        uint64 *lhs_vars, double lhs_constant,
        int rhs_count, double *rhs_coeffs,
        uint64 *rhs_vars, double rhs_constant,
        char sense);
    void setObjective(int count, double *coeffs, uint64 *var_ids,
            double constant, int sense);
    void showLog(bool shouldShow);
    void setTimeLimit(double timeLimit);
    MIPSolution optimize();
private:
    lprec *lp;
    int numVars;
};

#endif
