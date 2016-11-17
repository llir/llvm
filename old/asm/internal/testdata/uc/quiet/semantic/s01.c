/* Some strange but legal expressions and statements. 
   For more examples of semantically correct code, see noisy/simple 
*/

int x;
char y;

int main(void) {
  int z;
  char w;

  x = x+y+z+w;
  
  x = z = 42;

  x == z == 42;

  x == (z = 99); // Not legal without the paren

  while (x) x = 0;

  if (123) y = 4; else y = 7;

  w = x > y;

  y = 0 < x < 10;

  42;
}









