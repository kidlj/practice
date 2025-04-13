#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void character_copy() {
    int c;
    while ((c = getchar()) != EOF) {
        putchar(c);
    }
}

int character_count() {
    int nc = 0;
    while (getchar() != EOF) {
        nc++;
    }

    return nc;
}

int line_count() {
    int c, nl = 0;
    while ((c = getchar()) != EOF) {
        if (c == '\n') {
            nl++;
        }
    }

    return nl;
}

int blank_count() {
    int c, bc = 0;
    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\t' || c == '\n') {
            bc++;
        }
    }

    return bc;
}

void shrink_space() {
    int c, need_shrink = 0;
    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\t' || c == '\n') {
            if (need_shrink) {
                continue;
            }
            need_shrink = 1;
        } else {
            need_shrink = 0;
        }

        putchar(c);
    }
}

int word_count() {
    int c, wc = 0, in_word = 0;
    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\t' || c == '\n') {
            in_word = 0;
            continue;
        }
        if (!in_word) {
            in_word = 1;
            wc++;
        }
    }

    return wc;
}

void word_print() {
    int c, in_word = 0;
    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\t' || c == '\n') {
            if (in_word) {
                in_word = 0;
                putchar('\n');
            }
            continue;
        }

        in_word = 1;
        putchar(c);
    }
}

void test_character_copy() {
    const char *test_string = "Hello world.\nThis is another line.\n";
    freopen("test_input.txt", "w+", stdin);
    fprintf(stdin, "%s", test_string);
    rewind(stdin);

    freopen("test_output.txt", "w+", stdout);
    character_copy();
    rewind(stdout);

    char buffer[256];
    size_t bytes_read = fread(buffer, 1, sizeof(buffer) - 1, stdout);
    buffer[bytes_read] = '\0';
    assert(strcmp(buffer, test_string) == 0);

    freopen("/dev/tty", "r", stdin);
    freopen("/dev/tty", "w", stdout);

    remove("test_input.txt");
    remove("test_output.txt");
}

void test_character_count() {
    freopen("test_input.txt", "w+", stdin);
    fprintf(stdin, "Hello world.");
    rewind(stdin);
    assert(character_count() == 12);

    freopen("/dev/tty", "r", stdin);

    remove("test_input.txt");
}

void test_blank_count() {
    freopen("test_input.txt", "w+", stdin);
    fprintf(stdin, "Hello  world.\nAnother\tline.\n");
    rewind(stdin);
    assert(blank_count() == 5);

    freopen("/dev/tty", "r", stdin);

    remove("test_input.txt");
}

void test_trim_space() {
    char *test_input = "Hello  world. This is  \t a line.\n";
    char *test_output = "Hello world. This is a line.\n";
    freopen("test_input.txt", "w+", stdin);
    fprintf(stdin, "%s", test_input);
    rewind(stdin);

    freopen("test_output.txt", "w+", stdout);
    shrink_space();
    rewind(stdout);

    char buffer[256];
    size_t bytes_read = fread(buffer, 1, sizeof(buffer) - 1, stdout);
    buffer[bytes_read] = '\0';
    assert(strcmp(buffer, test_output) == 0);

    freopen("/dev/tty", "r", stdin);
    freopen("/dev/tty", "w", stdout);

    remove("test_input.txt");
    remove("test_output.txt");
}

void test_line_count() {
    FILE *fp = fopen("test_input.txt", "w");
    if (!fp) {
        perror("failed to create test input file");
        exit(EXIT_FAILURE);
    }
    fprintf(fp, "Hello world.\nThis is another line.\nYet another line.\n");
    fclose(fp);

    freopen("test_input.txt", "r", stdin);
    assert(line_count() == 3);

    freopen("/dev/tty", "r", stdin);

    remove("test_input.txt");
}

void test_word_count() {
    freopen("test_input.txt", "w+", stdin);
    fprintf(stdin, "Hello, world.  This is a line.\n");
    rewind(stdin);

    assert(word_count() == 6);

    freopen("/dev/tty", "r", stdin);

    remove("test_input.txt");
}

void test_word_print() {
    char *test_input = "  Hello  world. This is  \t a line.\n";
    char *test_output = "Hello\nworld.\nThis\nis\na\nline.\n";
    freopen("test_input.txt", "w+", stdin);
    fprintf(stdin, "%s", test_input);
    rewind(stdin);

    freopen("test_output.txt", "w+", stdout);
    word_print();
    rewind(stdout);

    char buffer[256];
    size_t bytes_read = fread(buffer, 1, sizeof(buffer) - 1, stdout);
    buffer[bytes_read] = '\0';
    assert(strcmp(buffer, test_output) == 0);

    freopen("/dev/tty", "r", stdin);
    freopen("/dev/tty", "w", stdout);

    remove("test_input.txt");
    remove("test_output.txt");
}

int main() {
    test_character_count();
    test_blank_count();
    test_line_count();
    test_character_copy();
    test_trim_space();
    test_word_count();
    test_word_print();

    printf("All tests passed.\n");
    return 0;
}