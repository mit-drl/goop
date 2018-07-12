#ifndef RS_Solver_HPP
#define RS_Solver_HPP

#include <iostream>
#include <vector>
#include <string>
#include "solution.hpp"

using namespace std;

#define uint64 unsigned long long

class Solver
{
    public:
        Solver() {};
        virtual ~Solver() {};
        virtual void addVars(
            int count, double *lb, double *ub, char *types) = 0;
        virtual void addConstr(
            int lhs_count, double *lhs_coeffs,
            uint64 *lhs_vars, double lhs_constant,
            int rhs_count, double *rhs_coeffs,
            uint64 *rhs_vars, double rhs_constant,
            char sense) = 0;
        virtual void setObjective(int count, double *coeffs, uint64 *var_ids,
                double constant, int sense) = 0;
        virtual void showLog(bool shouldShow) = 0;
        virtual void setTimeLimit(double timeLimit) = 0;
        virtual MIPSolution optimize() = 0;
};

#endif
