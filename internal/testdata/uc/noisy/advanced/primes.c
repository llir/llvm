//Computes all primes less than 1000

void putint(int i);
void putstring(char s[]);

char notprime[1001];
int max;
	
int main(void) {
  int i;
  int j;
  char cr[2];
  char space[2];

  cr[0] = '\n';
  cr[1] = 0;
  
  space[0] = ' ';
  space[1] = 0;

  max = 1001;
  i = 2;
  while (i<max) {
    if (!notprime[i]) {
      j = i + i;
      while (j<max) {
	notprime[j] = 1;
	j = j + i;
      }
    }
    i = i + 1;
  }
  i = 2;
  while (i + 10 < max) {
    j = i;
    putstring(cr);
    while (j < i+10) {
      if (!notprime[j]) {
	putint(j);
	putstring(space);
      }
      j = j+1;
    }
    i = i + 10;
  }
}
