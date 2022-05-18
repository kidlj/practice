#include <ctype.h>
#include <stdio.h>

#include "calc.h"

int getop(char s[]) {
    int i, c;
    while ((s[0] = c = _getch()) == ' ' || c == '\t') {
        ;
    }
    s[1] = '\0';
    if (!isdigit(c) && c != '.') {
        return c;
    }
    i = 0;
    if (isdigit(c)) {
        while (isdigit(c = s[++i] = _getch())) {
            ;
        }
    }
    if (c == '.') {
        while (isdigit(c = s[++i] = _getch())) {
            ;
        }
    }
    s[i] = '\0';
    if (c != EOF) {
        _ungetch(c);
    }
    return NUMBER;
}