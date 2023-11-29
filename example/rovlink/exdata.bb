struct RovExdataLeakageData[6] {
    bool<48> leakage[6];
}

enum RovExdataKeepAliveStatus[2] {
    HEALTHY = 0;
    WATCHDOG_TIMEOUT = 1;
}

struct RovExdataKeepAliveData[6] {
    uint32 timestamp[4];
    RovExdataKeepAliveStatus status[2];
}
