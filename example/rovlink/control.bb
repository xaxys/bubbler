package api.rovlink.control;

struct RovControlPropellerAData[6] {
    uint16 a1[2];
    uint16 a2[2];
    uint16 a3[2];
}

struct RovControlPropellerBData[6] {
    uint16 b1[2];
    uint16 b2[2];
    uint16 b3[2];
}

struct RovControlPropellerCData[6] {
    uint16 c1[2];
    uint16 c2[2];
    uint16 c3[2];
}

struct RovControlPropellerDData[6] {
    uint16 d1[2];
    uint16 d2[2];
    uint16 d3[2];
}

struct RovControlLightAData[6] {
    uint16 l1[2] {
        get l1_percent(float64): (value / 20000.0) * 100.0;
        get l1_percent_int(int32): (int32)((value / 20000.0) * 100.0 + 0.5);
        set l1_percent_int(float64): value < 0.5 ? 0 : value;
    };
    uint16 l2[2] {
        get l2_percent(float64): (value / 20000.0) * 100.0;
        get l2_percent_int(int32): (int32)((value / 20000.0) * 100.0 + 0.5);
        set l2_percent(float64): value < 0.5 ? 0 : value;
    };
    uint16 l3[2] {
        get l3_percent(float64): (value / 20000.0) * 100.0;
        get l3_percent_int(int32): (int32)((value / 20000.0) * 100.0 + 0.5);
        set l3_percent(float64): value < 0.5 ? 0 : value;
    };
}

struct RovControlLightBData[6] {
    uint16 l4[2] {
        get l4_percent(float64): (value / 20000.0) * 100.0;
        get l4_percent_int(int32): (int32)((value / 20000.0) * 100.0 + 0.5);
        set l4_percent(float64): value < 0.5 ? 0 : value;
    };
    uint16 l5[2] {
        get l5_percent(float64): (value / 20000.0) * 100.0;
        get l5_percent_int(int32): (int32)((value / 20000.0) * 100.0 + 0.5);
        set l5_percent(float64): value < 0.5 ? 0 : value;
    };
    uint16 l6[2] {
        get l6_percent(float64): (value / 20000.0) * 100.0;
        get l6_percent_int(int32): (int32)((value / 20000.0) * 100.0 + 0.5);
        set l6_percent(float64): value < 0.5 ? 0 : value;
    };
}

struct RovControlPtzData[6] {
    uint16 th1[2] {
        get th1_angle(float64): ((int32)value - 1500.0) / 1000.0 * 166.0;
        get th1_angle_int(int32): (int32)(((int32)value - 1500.0) / 1000.0 * 166.0 + 0.5);
    };
    uint16 th2[2] {
        get th2_angle(float64): ((int32)value - 1500.0) / 1000.0 * 166.0;
        get th2_angle_int(int32): (int32)(((int32)value - 1500.0) / 1000.0 * 166.0 + 0.5);
    };
    uint16 th3[2] {
        get th3_angle(float64): ((int32)value - 1500.0) / 1000.0 * 166.0;
        get th3_angle_int(int32): (int32)(((int32)value - 1500.0) / 1000.0 * 166.0 + 0.5);
    };
}

struct RovControlServoAData[6] {
    uint16 pwm1[2];
    uint16 pwm2[2];
    uint16 pwm3[2];
}

struct RovControlServoBData[6] {
    uint16 pwm4[2];
    uint16 pwm5[2];
    uint16 pwm6[2];
}

struct RovControlServoCData[6] {
    uint16 pwm7[2];
    uint16 pwm8[2];
    uint16 pwm9[2];
}

struct RovControlPostureData[6] {
    uint16 forward_backward[2] {
        get forward_backward_percent(float64): ((int32)value - 1530.0) / 10.0 / 0.45;
        get forward_backward_percent_int(int32): (int32)((((int32)value - 1530.0) / 10.0 / 0.45) + 0.5);
        set forward_backward_percent(float64): (uint16)(value * 10 * 0.45 + 1530.0);
    };
    uint16 left_right[2] {
        get left_right_percent(float64): ((int32)value - 1530.0) / 10.0 / 0.45;
        get left_right_percent_int(int32): (int32)((((int32)value - 1530.0) / 10.0 / 0.45) + 0.5);
        set left_right_percent(float64): (uint16)(value * 10 * 0.45 + 1530.0);
    };
    uint16 up_down[2] {
        get up_down_percent(float64): ((int32)value - 1530.0) / 10.0 / 0.45;
        get up_down_percent_int(int32): (int32)((((int32)value - 1530.0) / 10.0 / 0.45) + 0.5);
        set up_down_percent(float64): (uint16)(value * 10 * 0.45 + 1530.0);
    };
}
