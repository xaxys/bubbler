// Import test: this file references types from types.bb.
// Tests that imported types are correctly resolved across files
// and that custom package options are applied.

package protocol;

option go_package       = "github.com/example/protocol";
option java_package     = "com.example.protocol";
option csharp_namespace = "Example.Protocol";

import "types.bb";

struct Packet {
    uint8     seq;
    Direction dir;
    Vec2      pos;
    uint16    checksum;
}

struct Batch {
    uint8    count;
    Packet<4> packets;
}
