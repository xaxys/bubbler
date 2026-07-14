package kotlinaccessor;

struct GetterOnly {
    uint16 raw {
        get doubled(uint32): (uint32)(value);
    };
}
