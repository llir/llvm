// This program illustrates the bubblesort algorithm by sorting an
// array of char and printing the intermediate states of the array.


void putstring(char s[]);

char eol[1];
int n;


void bubble(char a[]) {
  int i;
  int j;
  char t;

  putstring (a);
  putstring (eol);
  i=n-1;
  while (i>0) {
    j = 0;
    while (j<i) {
      if (a[j] > a[j+1]) { 
	  t = a[j];
	  a[j] = a[j+1];
	  a[j+1] = t;
	}
      j = j + 1;
    }
    putstring (a);
    putstring (eol);
    i = i -1;
  }
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

  // Fill the string with a permutation of the characters a-z
  i = 0;
  q = 11;
  while (i<n) {
    t = q - (q / 26)*26; // q mod 26
    s[i] = 'a'+t;
    i = i + 1;
    q = q + 17;
  }

  bubble(s);
}
