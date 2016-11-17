// A simple calculator, based on predictive parsing.
//
// Try:
// 1+2+3+4
// 100-10-1
// (12-4)+(99-11+16)*19 
// (3*113+((2*2)*(2*2)))*1000000/113

/* Important characters:

'0' = 48
'9' = 57
'+' = 43 
'-' = 45 
'*' = 42
'/' = 47
'(' = 40
')' = 41

*/

void getstring(char s[]);
void putstring(char s[]);
void putint(int i);


int p;
char s[80];

int zero;
int nine;
int plus;
int minus;
int times;
int div;
int lpar;
int rpar;

char bad_number[11];
char bad_expression[15];

char cr[2];

char test_data[21];

int isNumber(char c) {
  return  (s[p] >= zero) && (s[p] <= nine);
}

// l == 0 => expr; l == 1 => term; l ==0 => factor
int expr(int l) {
  int a;
  int b;
  if (l==0) {
    a = expr(1);
    while (!(s[p] != plus && s[p]!= minus)) {
      if (s[p] == plus)
	{
	  p = p + 1;
	  b = expr(1);
	  a = a + b;
	}
      else 
	{
	  p = p + 1;
	  b = expr(1);
	  a = a - b;
	}
    }
    return a;
  }
  else if (l==1) {
    a = expr(2);
    while (!(s[p] != times && s[p]!= div)) {
      if (s[p] == times)
	{
	  p = p + 1;
	  b = expr(2);
	  a = a * b;
	}
      else 
	{
	  p = p + 1;
	  b = expr(2);
	  a = a / b;
	}
    }
    return a;
  }
  else if (l == 2) {
    if (s[p] ==lpar) {
      p = p+1;
      a = expr(0);
      if (s[p] !=rpar)  {
	putstring( bad_expression );
	putstring( cr );
      }
      p = p + 1;
      return a;
    } 
    else if (!isNumber(s[p])) {
      putstring( bad_number );
      putstring( cr );
      return 0;
    }
    else {
      a = 0;
      while (isNumber(s[p])) {
	a = a * 10 + (s[p] - zero);
	p = p + 1;
      }
      return a;
    }
  }
}

int main(void) {
  zero = 48;
  nine = 57;
  plus = 43;
  minus = 45;
  times = 42;
  div = 47;
  lpar = 40;
  rpar = 41;

  bad_number[ 0]='B';
  bad_number[ 1]='a';
  bad_number[ 2]='d';
  bad_number[ 3]=' ';
  bad_number[ 4]='n';
  bad_number[ 5]='u';
  bad_number[ 6]='m';
  bad_number[ 7]='b';
  bad_number[ 8]='e';
  bad_number[ 9]='r';
  bad_number[10]=0  ;


  bad_expression[ 0]='B';
  bad_expression[ 1]='a';
  bad_expression[ 2]='d';
  bad_expression[ 3]=' ';
  bad_expression[ 4]='e';
  bad_expression[ 5]='x';
  bad_expression[ 6]='p';
  bad_expression[ 7]='r';
  bad_expression[ 8]='e';
  bad_expression[ 9]='s';
  bad_expression[10]='s';
  bad_expression[11]='i';
  bad_expression[12]='o';
  bad_expression[13]='n';
  bad_expression[14]= 0;

  cr[0] = '\n';
  cr[1] = 0;

  test_data[ 0]='(';
  test_data[ 1]='1';
  test_data[ 2]='2';
  test_data[ 3]='-';
  test_data[ 4]='4';
  test_data[ 5]=')';
  test_data[ 6]='+';
  test_data[ 7]='(';
  test_data[ 8]='9';
  test_data[ 9]='9';
  test_data[10]='-';
  test_data[11]='1';
  test_data[12]='1';
  test_data[13]='+';
  test_data[14]='1';
  test_data[15]='6';
  test_data[16]=')';
  test_data[17]='*';
  test_data[18]='1';
  test_data[19]='9';
  test_data[20]= 0 ;

  getstring(s);

  if (s[0]=='t') { //if input string begins with t, use test string
    p = 0;
    while (test_data[p] != 0) {
      s[p] = test_data[p];
      p = p + 1;
    }
    s[p] = 0;
  }

  p = 0;
  putint(expr(0));
  putstring(cr);
}




