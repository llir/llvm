//Potentially ambiguous if. Note that indentation may be misleading.

int main(void){
  int x;
  int y;

  if(x) 
    if (y) x = 4711;
    else x=42;
}
