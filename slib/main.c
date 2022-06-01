// from sys/stat
extern int chmod(const char *pathname, unsigned int mode);

int entry_point(void) {
        return !! chmod("asagi", 0700);
}

char *desc[] = {0};
struct bash_builtin {
        char *name;
        int (*fn)(void);
        int on;
        char **long_doc;
        char *short_doc;
        char *other;
} executable_struct = { "executable", entry_point, 1, desc, "chmod 0700 asagi", 0 };
