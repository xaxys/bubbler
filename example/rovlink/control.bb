package api.control;

struct RovControlPropellerAData[6] {
    uint16 a1[2] [order = "big"];
    uint16 a2[2] [order = "big"];
    uint16 a3[2] [order = "big"];
}

struct RovControlLightAData[6] {
    uint16 l1[2] [order = "big"];
    uint16 l2[2] [order = "big"];
    uint16 l3[2] [order = "big"];
}

struct RovControlPtzData[6] {
    uint16 th1[2] [order = "big"];
    uint16 th2[2] [order = "big"];
    uint16 th3[2] [order = "big"];
}

struct RovControlServoAData[6] {
    uint16 pwm1[2] [order = "big"];
    uint16 pwm2[2] [order = "big"];
    uint16 pwm3[2] [order = "big"];
}

struct RovControlPostureData[6] {
    uint16 straight[2] [order = "big"];
    uint16 rotate[2] [order = "big"];
    uint16 vertical[2] [order = "big"];
}

struct RovControlMovementData[6] {
    uint16 speed_x[2] [order = "big"];
    uint16 speed_y[2] [order = "big"];
    uint16 speed_z[2] [order = "big"];
}

struct RovControlHeadingData[6] {
    uint16 pitch[2] [order = "big"];
    uint16 roll[2] [order = "big"];
    uint16 yaw[2] [order = "big"];
}
