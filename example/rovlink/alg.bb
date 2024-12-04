package api.alg;

struct RovAlgPidKpData[6] {
    float32 kp[4] [order = "big"];
    uint16 loop_id[2] [order = "big"];
}

struct RovAlgPidKiData[6] {
    float32 ki[4] [order = "big"];
    uint16 loop_id[2] [order = "big"];
}

struct RovAlgPidKdData[6] {
    float32 kd[4] [order = "big"];
    uint16 loop_id[2] [order = "big"];
}
