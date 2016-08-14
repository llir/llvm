/* Let's try a bunch of function calls. Should output
   01234567890123456789 */

void putint(int i);

void foo(int v1, int v2, int v3, int v4, int v5) {
  putint( v1);
  putint( v2);
  putint( v3);
  putint( v4);
  putint( v5);
}

int f(int x) {
  return x + 1;
}

char g(char x) {
  if (! x ) return 1; // Equivalent to x == 0
  else return 2*g(x-1);
}

int main(void) {	
  int x;
  foo(0,1,2,3,4);
  x = 5;
  foo(x+0, x+1, x+2, x+x-2, x*2-1);

  foo(0, f(0), f(f(0)), f(f(f(0))), g(2));
  foo(g(2)+g(0), g(2)+g(1), g(0)+g(1)+g(2), g(3), g(4)-7);
}
