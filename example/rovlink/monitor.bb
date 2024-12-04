package api.monitor;

struct RovMonitorPropellerAData[6] {
    uint16 a1[2] [order = "big"];
    uint16 a2[2] [order = "big"];
    uint16 a3[2] [order = "big"];
}

struct RovMonitorLightAData[6] {
    uint16 l1[2] [order = "big"];
    uint16 l2[2] [order = "big"];
    uint16 l3[2] [order = "big"];
}

struct RovMonitorPtzData[6] {
    uint16 th1[2] [order = "big"];
    uint16 th2[2] [order = "big"];
    uint16 th3[2] [order = "big"];
}

struct RovMonitorServoAData[6] {
    uint16 pwm1[2] [order = "big"];
    uint16 pwm2[2] [order = "big"];
    uint16 pwm3[2] [order = "big"];
}