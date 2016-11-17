/* Prints the numbers from 1 to 20. Puts an X next to the number if
   even, and a Y if divisble by 3. If the number is *not* divisible by
   3 but greater than 10, put a W next to it. Note that two if
   statements are written without braces.  */

void putint(int i);
void putstring(char s[]);

int main(void) {
  int i;
  char space[2];
  char X[2];  char Y[2];
  char W[2];  char nl[2];

  space[0] = ' ';  space[1] = 0;
  X[0]     = 'X';  X[1]     = 0;
  Y[0]     = 'Y';  Y[1]     = 0;
  W[0]     = 'W';  W[1]     = 0;
  nl[0]    = '\n'; nl[1]    = 0;

  i = 0;
  while (i != 21) { // same as  i<21
    putint(i);
    putstring(space);
    if (i/2*2==i) putstring(X);
    if (i/3*3==i) {
      putstring(Y);
    } 
    else if (i>10) putstring(W);
    putstring(nl);
    i = i + 1;
  }
}
