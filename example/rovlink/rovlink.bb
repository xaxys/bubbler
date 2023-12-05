package api.rovlink;

import "sensor.bb";
import "control.bb";

enum RovDeviceType[1] {
    NONE = 0x00,
    POWER_CAB = 0x01,
    CONTROL_CAB = 0x02,
    COMM_CAB = 0x03,
    MAIN_CAB = 0x04,
    HOST = 0x05,
}

enum RovlinkFrameType[1] {
    SENSOR_NONE = 0x00,
    SENSOR_CAB_TEMP_HUMID_PRESS = 0x01,
    SENSOR_WATER_TEMP_DEPTH_PRESS = 0x02,
    SENSOR_ACCELERATION = 0x03,
    SENSOR_ANGULAR_VELOCITY = 0x04,
    SENSOR_EULER_ANGLE = 0x05,
    SENSOR_MAGNETIC_FIELD = 0x06,
    SENSOR_HEIGHT_SONAR = 0x07,
    SENSOR_DISTANCE_SONAR = 0x08,

    // EXDATA
    EXDATA_LEAKAGE = 0x10,
    EXDATA_KEEP_ALIVE = 0x11,
    // EXDATA_POWER_CAB_ID = 0x12,
    // EXDATA_CONTROL_CAB_ID = 0x13,
    // EXDATA_COMM_CAB_ID = 0x14,
    // EXDATA_MAIN_CAB_ID = 0x15,
    // EXDATA_HOST_ID = 0x16,

    // CONTROL
    CONTROL_PROPELLER_A = 0x21,
    CONTROL_PROPELLER_B = 0x22,
    CONTROL_PROPELLER_C = 0x23,
    CONTROL_PROPELLER_D = 0x24,
    CONTROL_LIGHT_A = 0x25,
    CONTROL_LIGHT_B = 0x26,
    CONTROL_PTZ = 0x27,
    CONTROL_SERVO_A = 0x28,
    CONTROL_SERVO_B = 0x29,
    CONTROL_SERVO_C = 0x2A,
    CONTROL_POSTURE = 0x2B,

    // HOST
    HOST_CAMARA = 0x31,
    // HOST_VIRTUAL = 0x3F,

    // ALGORITHM
    ALGORITHM_PID_KP = 0x41,
    ALGORITHM_PID_KI = 0x42,
    ALGORITHM_PID_KD = 0x43,

    // COMPONENT
    COMPONENT_PROPELLER = 0x51,
    COMPONENT_LIGHT = 0x52,
    COMPONENT_PTZ = 0x53,
    COMPONENT_SERVO_A = 0x54,
    COMPONENT_SERVO_B = 0x55,

    // EXCOMPONENT
    EXCOMPONENT_RELAY = 0x61,
    EXCOMPONENT_RESCUE = 0x62,

    // MODE
    MODE_MODE_A = 0x71,
    MODE_MODE_B = 0x72,

    // EXCONTROL
    EXCONTROL_CLAMP = 0x81,
    EXCONTROL_ARM_A = 0x82,
    EXCONTROL_ARM_B = 0x83,
    EXCONTROL_ARM_C = 0x84,

    // BETTERY
    BETTERY_VOLTAGE = 0xA0,
    BETTERY_CURRENT_GAIN = 0xA1,
    BETTERY_CURRENT_BIAS = 0xA2,
    BETTERY_CURRENT = 0xA3,
    BETTERY_REMAIN = 0xA4,

    // CURRENT
    CURRENT_PROPELLER_A = 0xB0,
    CURRENT_PROPELLER_B = 0xB1,
    CURRENT_PROPELLER_C = 0xB2,
    CURRENT_PROPELLER_D = 0xB3,
    CURRENT_PTZ = 0xB4,
    CURRENT_SERVO_A = 0xB5,
    CURRENT_SERVO_B = 0xB6,
}

struct RovlinkFrame[8] {
    RovlinkFrameType opcode[1];
    bool sensor[#1];
    bool valid[#1];
    void [#2];
    RovDeviceType device[#4];
    uint8<6> payload[6];
}

struct RovlinkFullFrame[10] {
    uint8 header[1];
    RovlinkFrame;
    uint8 crc[1];
}
