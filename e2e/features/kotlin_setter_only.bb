package kotlinnegative;

struct SetterOnly {
    uint16 raw {
        set calibrated(float64): (uint16)(value);
    };
}
