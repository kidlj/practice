#include "test.h"
#include <stdio.h>

typedef struct s t;

int main(void)
{
    t v = {1};
    printf("%d\n", v.i);
}
