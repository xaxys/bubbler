// ERROR CASE: Duplicate option declaration in the same file.
// Compiler must reject this and exit with a non-zero code.

package errtest;

option go_package = "errtest";
option go_package = "errtest2";   // duplicate — should cause a compile error

struct Bar { uint8 z; }
