void putint(int x);

/* Test the implementation of global and local variables. Note that
   output will appear as "7654321".*/

int foo;
int main(void) {
  int bar;
  foo = 76;
  bar = 54321;
  putint(foo);
  putint(bar);
}
