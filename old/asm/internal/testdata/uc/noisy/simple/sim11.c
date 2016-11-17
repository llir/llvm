/* A variation of sim06.c. Same output */

void putint(int i);
void putstring(char s[]);

void f(int i, char v[]) {
  v[0] = i;
}

int main(void) {
  char t[2];
  int b;
  b = 10;
  t[1] = 0;

  while (b) {
    f(48+b-1, t);
    putstring(t);
    b = b - 1;
  }

  f('\n', t);
  putstring(t);
}









