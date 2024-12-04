package api.types;

enum RovDeviceType[1] {
    NONE = 0x00,
    POWER_CAB = 0x01,
    CONTROL_CAB = 0x02,
    COMM_CAB = 0x03,
    MAIN_CAB = 0x04,
    HOST = 0x05,
}

enum RovlinkFrameType[1] {
    // SENSOR
    SENSOR_NONE = 0x00,
    SENSOR_CAB_TEMP_HUMID_PRESS = 0x01,
    SENSOR_WATER_TEMP_DEPTH_PRESS = 0x02,
    SENSOR_ACCELERATION = 0x03,
    SENSOR_ANGULAR_VELOCITY = 0x04,
    SENSOR_EULER_ANGLE = 0x05,
    SENSOR_MAGNETIC_FIELD = 0x06,
    SENSOR_HEIGHT_SONAR = 0x07,
    SENSOR_DISTANCE_SONAR = 0x08,

    // CORE
    EXDATA_LEAKAGE = 0x10,
    EXDATA_KEEP_ALIVE = 0x11,

    // CONTROL
    CONTROL_PROPELLER_A = 0x21,
    CONTROL_LIGHT_A = 0x25,
    CONTROL_PTZ = 0x27,
    CONTROL_SERVO_A = 0x28,
    CONTROL_POSTURE = 0x2B,
    CONTROL_MOVEMENT = 0x2C,
    CONTROL_HEADING = 0x2D,

    // ALGORITHM
    ALGORITHM_PID_KP = 0x41,
    ALGORITHM_PID_KI = 0x42,
    ALGORITHM_PID_KD = 0x43,

    // MODE
    MODE_MODE_A = 0x71,

    // EXCONTROL
    EXCONTROL_CLAMP = 0x81,
    EXCONTROL_MANIPULATOR_A = 0x82,

    // BATTERY
    BETTERY_VOLTAGE = 0xA0,
    BATTERY_CURRENT = 0xA3,
    BATTERY_SOC = 0xA4,

    // MONITOR
    MONITOR_PROPELLER_A = 0xB0,
    MONITOR_LIGHT_A = 0xB5,
    MONITOR_PTZ = 0xB7,
    MONITOR_SERVO_A = 0xB5,
}

struct RovlinkFrame[8] {
    RovlinkFrameType opcode[1];
    bool sensor[#1];
    bool valid[#1];
    bool dlc[#2];
    RovDeviceType device[#4];
    uint8<6> payload[6];
}

struct RovlinkStdFrame[10] {
    uint8 magic[1] = 0xFD;
    RovlinkFrame;
    uint8 crc[1];
}
