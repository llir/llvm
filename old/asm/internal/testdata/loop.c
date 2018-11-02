int main(void) {
	int sum = 0;
	for (int i = 0; i < 10; i++) {
		sum += i;
	}
	return sum%256;
}
