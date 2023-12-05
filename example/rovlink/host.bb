package api.rovlink.host;

enum RovCamaraType[1] {
    NO_CAM = 0;
    FRONT_CAM = 1;
    BACK_CAM = 2;
}

struct RovHostCamaraData[6] {
    RovCamaraType<6> camaraType;
}
