/* 
   Factorial function, interactive.
*/


void putint(int i);
int getint(void);

int fac(int n) {
  if (n == 0) return 1;
  else return n*fac(n-1);
}

int a;
int main(void) {
  a = getint();
  putint( fac(a));
}
