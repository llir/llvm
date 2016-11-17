void putint(int x);

/* Test the implementation of global and local arrays. Note that
   output will appear as "123456".*/

int a[10];

int main(void) {
  int b[10];
  a[7] = 123;
  b[5] = 456;
  putint(a[7]);
  putint(b[5]);
}
