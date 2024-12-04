package api.battery;

struct RovBatteryVoltageData[6] {
    uint16 voltage[2] [order = "big"];
    uint16 power[2] [order = "big"];
    uint16 discharge[2] [order = "big"];
}

struct RovBatteryCurrentData[6] {
    uint16 ca1[2] [order = "big"];
    uint16 ca2[2] [order = "big"];
    uint16 ca3[2] [order = "big"];
}

struct RovBatterySocData[6] {
    uint16 soc[2] [order = "big"];
    uint16 discharge_time[2] [order = "big"];
    uint16 res_time[2] [order = "big"];
}
