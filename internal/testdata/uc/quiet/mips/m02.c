int f(int x) {
  if (x > 0)
    return 2 + f(x-1);
  return 112;
}
      
int main (void) {
  f(8);
}

