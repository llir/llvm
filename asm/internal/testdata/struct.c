typedef struct bar {
	int x;
	int y;
} bar;

typedef struct foo {
	bar baz;
	int z;
} foo;

typedef struct qux {
	struct {
		int i;
		int j;
	} foob;
	int m;
} qux;

int main(void) {
	foo a;
	qux b;
	a.baz.x = 1;
	a.baz.y = 2;
	a.z = 3;
	b.foob.i = 4;
	b.foob.j = 5;
	b.m = 6;
	return 42;
}
