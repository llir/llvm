int f(int x) {
  int y;
  if (x > 1) {
    y = f(x-1);
    y = y + f(x-1);
    return y;
  }
  return 1;
}
      
int main (void) {
  f(8);
}

