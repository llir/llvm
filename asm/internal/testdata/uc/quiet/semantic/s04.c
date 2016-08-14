int foo(char x[]) {
  return x[0];
}

int main(void) {
  char a[10];
  foo(a);	
}
