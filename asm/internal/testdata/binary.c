int main(void) {
	int a = 5;
	int b = 3;

	int c = a + b;   // AddInst
	int d = a - b;   // SubInst
	int e = a * b;   // MulInst
	int f = a / b;   // SDivInst
	int g = a % b;   // SRemInst

	unsigned int h = 5;
	unsigned int i = 3;

	unsigned int j = h / i;   // UDivInst
	unsigned int k = h % i;   // URemInst

	float l = 5.0;
	float m = 3.0;

	float n = l + m;   // FAddInst
	float o = l - m;   // FSubInst
	float p = l * m;   // FMulInst
	float q = l / m;   // FDivInst

	return 0;
}
