package api.rovlink.excontrol;

struct RovExcontrolClampData[6] {
    int16 clamp[2];
    int16 wrist[2] {
        get wrist_percent(float64): value / 3000.0 * 100.0;
        get wrist_percent_int(int32): (int32)(value / 3000.0 * 100.0);
        set wrist_percent(float64): value > -0.5 && value < 0.5 ? 0 : value;
    };
    void [2];
}

struct RovExcontrolArmAData[6] {
    uint16 a1[2];
    uint16 a2[2];
    uint16 a3[2];
}

struct RovExcontrolArmBData[6] {
    uint16 b1[2];
    uint16 b2[2];
    uint16 b3[2];
}

struct RovExcontrolArmCData[6] {
    uint16 c1[2];
    uint16 c2[2];
    uint16 c3[2];
}