// ERROR CASE: Duplicate package name (partner to a.bb)
// Same package name as a.bb — entry.bb importing both must fail.

package shared_pkg;

struct TypeB { uint8 b; }
