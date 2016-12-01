#include <stdio.h>
#include <stdint.h>

int main(void) {
	uint64_t a = 5;
	uint32_t b = a; // Trunc

	uint32_t c = 3;
	uint64_t d = c; // ZExt

	int32_t e = 3;
	int64_t f = e; // SExt

	double g = 5.0;
	float h = g; // FPTrunc

	float i = 3.0;
	double j = i; // FPExt

	double k = 5.0;
	uint64_t l = k; // FPToUI

	double m = 3.0;
	int64_t n = m; // FPToSI

	uint64_t o = 5;
	double p = o; // UIToFP

	int64_t q = 3;
	double r = q; // SIToFP

	void *s = NULL;
	uint64_t t = (uint64_t) s; // PtrToInt

	uint64_t u = 0;
	void *v = (void *) u; // IntToPtr

	return 0;
}
