/* Fibbonacci, in the simple and naive form */
/* Prints fibbonacci numbers for n=1..12 */

void putint(int i);
void putstring(char s[]);

int fib(int n) {
  if (n == 0) return 1;
  if (n == 1) return 1;
  return fib(n-1) + fib(n-2);
}

int main (void) {
  int i;
  char space[2];
  char cr[2];
  space[0] = ' ';  space[1] = 0;
  cr[0]    = '\n'; cr[1]    = 0;


  i = 0;

  while (i<=12) {
    putint(i);
    putstring(space);
    putint(fib(i));
    putstring(cr);
    i = i + 1;
  }
}
