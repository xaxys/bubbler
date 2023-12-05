package api.rovlink.excomponent;

struct RovExcomponentRelayData[6] {
    bool lazer[1];
    bool clamp[1];
    bool arm[1];
    bool sonar[1];
    bool propeller[1];
    bool ptz[1];
}

struct RovExcomponentRescueData[6] {
    uint8 stretch[1];
    uint8 left_servo[1];
    uint8 right_servo[1];
    uint8 clamp[1];
    uint8 turn_servo[1];
    uint8 solution[1];
}
