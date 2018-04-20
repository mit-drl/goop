
import argparse
import os
import platform

parser = argparse.ArgumentParser(
    description="Creates a lib.go file for the SWIG library")
parser.add_argument("--gurobi-home", "-p", dest="gurobi_home", type=str,
                    default=os.environ["GUROBI_HOME"])
parser.add_argument("--go-fname", dest="go_fname", type=str)
parser.add_argument("--pkg", dest="package_name", type=str)
args = parser.parse_args()

platform_str = platform.platform().lower()

if "xenial" in platform_str:
    gurobi_libs = ["gurobi_g++5.2", "gurobi75"]
elif "darwin" in platform_str:
    gurobi_libs = ["gurobi_c++", "gurobi75"]
elif "trusty" in platform_str:
    gurobi_libs = ["gurobi_c++", "gurobi75"]
elif "generic" in platform_str:
    gurobi_libs = ["gurobi_g++5.2", "gurobi75"]
else:
    print("OS not supported!: {}".format(platform_str))
    exit(1)


cxx_flags_fmt = "// #cgo CXXFLAGS: --std=c++11 -I{}/include "\
    + "-I${{SRCDIR}}/../include\n"
ld_flags_fmt = "// #cgo LDFLAGS: -L{}/lib {}\n"
libs_str = " ".join("-l" + lib for lib in gurobi_libs)
go_str = "package {}\n\n".format(args.package_name)
go_str += cxx_flags_fmt.format(args.gurobi_home)
go_str += ld_flags_fmt.format(args.gurobi_home, libs_str)
go_str += "import \"C\""

with open(args.go_fname, "w") as f:
    f.write(go_str)
