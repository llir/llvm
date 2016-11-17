/* Let's try all I/O operations. The program waits until you have
   answered the questions. */

void putint(int x);
void putstring(char s[]);
int getint(void);
int getstring(char s[]);


int main(void) {
  char nameq[12];
  char ageq[10];
  char youare[10];
  char cr[2];

  char name[80];
  int age;
  
  nameq[0]='Y';
  nameq[1]='o';
  nameq[2]='u';
  nameq[3]='r';
  nameq[4]=' ';
  nameq[5]='n';
  nameq[6]='a';
  nameq[7]='m';
  nameq[8]='e';
  nameq[9]='?';
  nameq[10]=' ';
  nameq[11]=0;

  ageq[0]='Y';
  ageq[1]='o';
  ageq[2]='u';
  ageq[3]='r';
  ageq[4]=' ';
  ageq[5]='a';
  ageq[6]='g';
  ageq[7]='e';
  ageq[8]=' ';
  ageq[9]=0;

  youare[0]='Y';
  youare[1]='o';
  youare[2]='u';
  youare[3]=' ';
  youare[4]='a';
  youare[5]='r';
  youare[6]='e';
  youare[7]=':';
  youare[8]=' ';
  youare[9]=0;
  
  cr[0] = '\n';
  cr[1] = 0;

  putstring(nameq);
  getstring(name);
  
  putstring(ageq);
  age = getint();
  
  putstring(youare);
  putstring(name);
  putstring(cr);

  putstring(youare);
  putint(age);
  putstring(cr);


}
