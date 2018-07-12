/* Parse the header file to generate wrappers */
%include <typemaps.i>
%include "std_vector.i"
%include "std_string.i"

namespace std {
    %template(DoubleVector) vector<double>;
}
