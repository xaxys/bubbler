// ERROR CASE: Circular import A → B → A
// Compiler must detect the cycle and exit with a non-zero code.

package cyclea;

import "b.bb";

struct NodeA { uint8 a_val; }
