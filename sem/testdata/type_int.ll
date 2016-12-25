%a = type i0       ; error: invalid integer type bit width; expected > 0, got 0
%b = type i1       ; valid
%d = type i8388607 ; valid
%e = type i8388608 ; error: invalid integer type bit width; expected < 2^24, got 8388608
%c = type i32      ; valid
