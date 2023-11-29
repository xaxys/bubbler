struct RovModeModeAData[6] {
    bool side_push[1];
    bool tilt[1];
    bool roll[1];
    bool auxiliary[1];
    bool rescue[1];
    bool module[1];
}

struct RovModeModeBData[6] {
    bool keep_heading[1];
    bool keep_depth[1];
    bool auto_stabilize[1];
    bool auto_schedule[1];
    bool auto_execute[1];
    bool auto_avoiding_obstacles[1]; 
}