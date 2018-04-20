cd $HOME

wget http://packages.gurobi.com/7.5/gurobi7.5.2_linux64.tar.gz
tar -xzf gurobi7.5.2_linux64.tar.gz

export GUROBI_HOME=$HOME/gurobi752/linux64
export GUROBI_LIB_NAME=gurobi75
cd -
