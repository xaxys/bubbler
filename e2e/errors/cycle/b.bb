// ERROR CASE: Circular import (partner to a.bb)

package cycleb;

import "a.bb";

struct NodeB { uint8 b_val; }
