/*
Should produce this::

#########################################
############                 ############
#########                       #########
######                             ######
####                                 ####
###                                   ###
##                                     ##
#                                       #
#                                       #
#                                       #
#                                       #
#                                       #
##                                     ##
###                                   ###
####                                 ####
######                             ######
#########                       #########
############                 ############
#########################################

The output is adjusted so that the image will look circular if the
relation between the height and the width of a character is
approximately 22/10.

*/

void putstring(char s[]);

void drawpos(char c) {
  char s[2];

  if (c) s[0] = '#';
  else s[0] = ' ';

  s[1] = 0;
  putstring(s);
}
 

void nl(void) {
  char cr[2];
  cr[0]    = '\n'; 
  cr[1]    = 0;
  putstring(cr);
}
  
int main(void) {
  int i;
  int j;
  
  i = -9;

  while (i<= 9) {
    j = -20;
    nl();
    while (j<=20) {
      drawpos(i*i*22*22/(10*10)+j*j>380);
      j=j+1;
    }
    i=i+1;
  }
  nl(); nl();
}

    

  
