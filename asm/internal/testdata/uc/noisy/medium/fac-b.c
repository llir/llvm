/* 
   Factorial function. Prints factorial of all numbers from 0 to 10.
*/

void putint(int i);
void putstring(char s[]);

int fac(int n) {
  if (n == 0) return 1;
  else return n*fac(n-1);
}

int a;
int main(void) {
  int i;
  char space[2];
  char cr[2];
  space[0] = ' ';  space[1] = 0;
  cr[0]    = '\n'; cr[1]    = 0;

  i = 0;

  while (i<=10) {
    putint(i);
    putstring(space);
    putint(fac(i));
    putstring(cr);
    i = i + 1;
  }
}
