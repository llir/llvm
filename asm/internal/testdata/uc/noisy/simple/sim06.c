/* Our first control structure! Should print 9876543210 */

void putint(int i);
void putstring(char s[]);

char t[2];

int main(void) {
  int b;
  b = 10;
  t[1] = 0;

  while (b) {
    t[0] = 48+b-1;
    putstring(t);
    b = b - 1;
  }

  t[0] = '\n';
  putstring(t);
}









