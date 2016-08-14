// This program illustrates the quick sort algorithm by sorting an
// array of char and printing the intermediate results.
//
// Adapted from N.Wirth: Algorithms + Data Structures = Programs


void putstring(char s[]);

char eol[2];
int n;


void sort(char a[], int l, int r) {
  int i;
  int j;
  char x;
  char w;


  i = l;
  j = r;
  x = a[(l+r) / 2];
  
  while ( i<= j) {
    while (a[i] < x) i = i + 1;
    while (x < a[j]) j = j - 1;
    if (i<= j) {
      w = a[i];
      a[i] = a[j];
      a[j] = w;
      i = i + 1;
      j = j - 1;
    }
  }

  putstring (a);
  putstring (eol);
  if (l < j) sort(a, l,j);
  if (i < r) sort(a, i, r);

}

int main(void)
{
  char s[27];
  int i;
  char t;
  int q;

  eol[0] = '\n';
  eol[1] = 0;

  n = 26;

  s[n] = 0;

  i = 0;

  // Fill the string with random-looking data
  q = 11;
  while (i<n) {
    t = q - (q / 26)*26;
    s[i] = 'a'+t;
    i = i + 1;
    q = q + 17;
  }


  putstring (s); // print it ...
  putstring (eol);
  sort(s, 0, n-1); // sort it ...
  putstring(s);  // and print again
  putstring (eol);

}
