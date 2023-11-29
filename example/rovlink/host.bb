enum RovCamaraType[1] {
    NO_CAM = 0;
    FRONT_CAM = 1;
    BACK_CAM = 2;
}

struct RovHostCamaraData[6] {
    RovCamaraType cam1;
    RovCamaraType cam2;
    RovCamaraType cam3;
    RovCamaraType cam4;
    RovCamaraType cam5;
    RovCamaraType cam6;
}
