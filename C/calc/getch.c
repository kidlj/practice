#include <stdio.h>

#define BUFFSIZE 100

char buf[BUFFSIZE];
int bufp = 0;

int _getch(void) {
    return (bufp > 0) ? buf[--bufp] : getchar();
}

void _ungetch(int c) {
    if (bufp >= BUFFSIZE) {
        printf("ungetch: too many characters\n");
    } else {
        buf[bufp++] = c;
    }
}