// Should first print 42 nine times, then an approximation of PI
// times one million. Finally it should print 110001000.

void putint(int i);
void putstring(char s[]);

int main(void) {
  char cr[1];
  int x; int y;

  cr[0]='\n';
  cr[1]=0;

  putint(42);
  putstring(cr);

  putint(35+7);
  putstring(cr);

  putint(6*7);
  putstring(cr);

  putint(3*4+5*6);
  putstring(cr);

  putint(7*8-3*4-2);
  putstring(cr);

  putint((-6)*(-7));    
  putstring(cr);

  putint( ((9 + 9 + 9) 
	     * ((9 + 9) / 9) 
	     * (9 * 9 - (9 + 9)) 
	     + (9 + 9)*(9+9+9)) 
	    / (9 * 9+9) 
	    - 9 / 9);
  putstring(cr);

  x = y = 6;
  putint(x*(y+1)); 
  putstring(cr);

  x = (y = 5)+3;
  putint(x*y+2); 
  putstring(cr);


  putint((3*113+((2*2)*(2*2)))*1000000/113);
  putstring(cr);

  putint( 1  > 0);
  putint( 1 >= 0);
  putint( 1 == 0);
  putint( 1  < 0);
  putint( 1 <= 0);
  putint( 0 < 1 && 0 < 1);
  putint( 0 < 1 && 1 < 0);
  putint( 1 < 0 && 0 < 1);
  putint( 1 < 0 && 1 < 0);
  putstring(cr);
}
