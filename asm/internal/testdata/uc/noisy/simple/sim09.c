/* Passing arrays as arguments of functions poses some particular
   difficulties. This program should output

   12345678
   ABCD

*/

void putint(int i);
void putstring(char s[]);

void f(int a[]) {
  putint(a[3]);
}

void g(int b[]) {
  f(b);
}

void fc(char a[]) {
  putstring(a);
}

void gc(char b[]) {
  fc(b);
}

int x[10];
char xc[10];

int main (void) {
  int y[10];
  char yc[10];

  x[3] = 12;
  y[3] = 34;

  f(x);
  f(y);
  
  x[3] = 56;
  y[3] = 78;

  g(x);
  g(y);

  xc[0] = '\n';
  xc[1] = 'A';
  xc[2] = 0;

  yc[0] = 'B';
  yc[1] = 0;

  fc(xc);
  fc(yc);

  xc[0] = 'C';
  xc[1] = 0;

  yc[0] = 'D';
  yc[1] = '\n';
  yc[2] = 0;
 
  gc(xc);
  gc(yc);
}
