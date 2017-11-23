@nest = global double fmul (double 4.33, double fadd (double 2.1, double 1.2))

define double @nestedinstruction() {
entry:
	%0 = load double, double* @nest
	ret double %0
}

