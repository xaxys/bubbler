// Types shared across imports test
// Used by protocol.bb via: import "types.bb"

package types;

option go_package       = "types";
option java_package     = "com.example.types";
option csharp_namespace = "Example.Types";

enum Direction[1] {
    NORTH = 0,
    SOUTH = 1,
    EAST  = 2,
    WEST  = 3,
}

struct Vec2 {
    int16 x;
    int16 y;
}
