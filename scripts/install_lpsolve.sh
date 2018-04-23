#!/usr/bin/env bash

set -e

LP_SOLVE_TAR_NAME=lp_solve_5.5.2.0_dev_ux64.tar.gz
LP_SOLVE_URL=http://sourceforge.net/projects/lpsolve/files/lpsolve/5.5.2.0/lp_solve_5.5.2.0_dev_ux64.tar.gz

mkdir -p .third_party/lpsolve
cd .third_party
wget $LP_SOLVE_URL
tar -xzf $LP_SOLVE_TAR_NAME -C lpsolve
cd -
