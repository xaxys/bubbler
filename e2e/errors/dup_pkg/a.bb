// ERROR CASE: Duplicate package name
// Both a.bb and b.bb declare the same package name.
// When entry.bb imports both, the compiler must reject this.

package shared_pkg;

struct TypeA { uint8 a; }
