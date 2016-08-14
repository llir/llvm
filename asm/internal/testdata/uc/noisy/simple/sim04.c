void putstring(char s[]);

/* A simple program that tests output of strings. One string is a
   global array, the other is local. Of course, neither array has a
   length that is a multiple of 4. Outputs

   Hello
   Good bye
*/

char x[7];

int main(void) {
  char y[10];

  x[0] = 'H';
  x[1] = 'e';
  x[2] = 'l';
  x[3] = 'l';
  x[4] = 'o';
  x[5] = '\n';
  x[6] = 0;

  y[0] = 'G';
  y[1] = 'o';
  y[2] = 'o';
  y[3] = 'd';
  y[4] = ' ';
  y[5] = 'b';
  y[6] = 'y';
  y[7] = 'e';
  y[8] = '\n';  
  y[9] = 0;

  putstring(x);   
  putstring(y);
}
