#ifndef GOOP_GUROBI_HPP_
#define GOOP_GUROBI_HPP_

#include "gurobi_c++.h"
#include "base_solver.hpp"

class GurobiSolver : public Solver
{
public:
    GurobiSolver();
    ~GurobiSolver();
    void setMIPGapTol(double gap);
    void setConcurrentMIP(int numMips);
    void setThreads(int numThreads);
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
    int numVars;
    GRBEnv env;
    GRBModel model;
    GRBVar *vars;
};

#endif
