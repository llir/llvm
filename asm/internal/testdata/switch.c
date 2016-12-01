int a;

int main(void) {
	int b = 50;

	switch (a) {
	case 0:
		b = 10;
		break;
	case 1:
		b = 20;
		break;
	case 2:
		b = 30;
		break;
	case 3:
		b = 40;
		break;
	default:
		break;
	}

	return b;
}
