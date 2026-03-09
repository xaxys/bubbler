// ERROR CASE: Entry point that imports two files with the same package name.

package entry;

import "a.bb";
import "b.bb";

struct Root { uint8 root; }
