#include <stdio.h>

int fac(int n) {
  int i;
  int p;
  if (n<0) return 0;
  i = 0;
  p = 1;
  while (i<n) {
    i = i + 1;
    p = p * i;
  }
  return p;
}

int main(void){
  fac(5);
}
