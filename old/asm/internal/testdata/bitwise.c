int main(void) {
	int a = 5;
	int b = 3;
	unsigned int c = 5;
	unsigned int d = 3;

	int e = a << b;            // ShlInst
	unsigned int f = c >> d;   // LShrInst
	int g = a >> b;            // AShrInst
	int h = a & b;             // AndInst
	int i = a | b;             // OrInst
	int j = a ^ b;             // XorInst

	return 0;
}
