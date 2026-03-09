// omit_empty option test:
// When omit_empty = "true" and the file defines no local structs/enums,
// the target file should NOT be generated.

package importonly;

option omit_empty = true;

import "imports/types.bb";

// Intentionally empty — no local struct definitions.
