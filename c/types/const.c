#include <stdio.h>

void pointer_to_variable(void)
{
    int i = 10;
    int j = 20;
    int *ptr = &i;

    printf("*ptr = %d\n", *ptr);

    ptr = &j;
    printf("*ptr = %d\n", *ptr);

    *ptr = 30;
    printf("*ptr = %d\n", *ptr);
}

void pointer_to_const(void)
{
    int i = 10;
    int j = 20;
    int const *ptr = &i;

    printf("*ptr = %d\n", *ptr);

    // up qualification
    ptr = &j;
    printf("*ptr = %d\n", *ptr);

    // error: expression must be a modifiable lvalue
    // *ptr = 30;
    // printf("*ptr = %d\n", *ptr);
}

void pointer_to_const_down_qualification(void)
{
    int i = 10;
    int const j = 20;
    int *ptr = &i;

    printf("*ptr = %d\n", *ptr);

    // down qualification
    // warning: assigning to 'int *' from 'const int *' discards qualifiers
    // not allowed in C++
    // ptr = &j;
    // printf("*ptr = %d\n", *ptr);
}

void constant_pointer_to_variable(void)
{
    int i = 10;
    int j = 20;
    int *const ptr = &i;

    printf("*ptr = %d\n", *ptr);

    *ptr = 100;
    printf("*ptr = %d\n", *ptr);

    // error: expression must be a modifiable lvalue
    // ptr = &j;
    // printf("*ptr = %d\n", *ptr);
}

void constant_pointer_to_constant(void)
{
    int i = 10;
    int j = 20;
    int const *const ptr = &i; // or
    // const int *const ptr = &i;

    printf("*ptr = %d\n", *ptr);

    // error
    // ptr = &j;

    // error
    // *ptr = 100;
}

int main()
{
    pointer_to_variable();
    pointer_to_const();
    pointer_to_const_down_qualification();
    constant_pointer_to_variable();
    constant_pointer_to_constant();

    return 0;
}
