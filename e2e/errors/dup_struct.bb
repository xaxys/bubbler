// ERROR CASE: Duplicate struct name in the same file.
// Compiler must reject this and exit with a non-zero code.

package errtest;

struct Foo { uint8 x; }
struct Foo { uint8 y; }   // duplicate — should cause a compile error
