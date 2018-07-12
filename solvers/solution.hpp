#ifndef GOOP_SOLUTION_HPP
#define GOOP_SOLUTION_HPP

#include <vector>

using namespace std;

struct MIPSolution
{
    vector<double> values;
    double obj;
    double gap;
    bool optimal;
    int errorCode;
    string errorMessage;

    double getValue(int i)
    {
        return values.at(i);
    }

    ~MIPSolution()
    {
    }
};

#endif
