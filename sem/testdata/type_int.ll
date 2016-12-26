; Integer types.
%0 = type i0       ; error: invalid integer type bit width; expected > 0, got 0
%1 = type i1       ; valid
%2 = type i8388607 ; valid
%3 = type i8388608 ; error: invalid integer type bit width; expected < 2^24, got 8388608
%4 = type i32      ; valid
