package api.rovlink.sensor;

struct RovSensorCabTempHumidPressData[6] {
    uint16 temperature[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    uint16 humidity[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    uint16 pressure[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
}

struct RovSensorWaterTempDepthPressData[6] {
    uint16 temperature[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    uint16 depth[2] {
        get trustable(bool): (value / 100.0) < 1.0 ? false : true;
        get confidence(int32): (value / 100.0) < 1.0 ? 0 : 100;
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    uint16 pressure[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
}

struct RovSensorAccelerationData[6] {
    uint16 x[2] {
        get (float64): value / 32768.0 * 16.0 * 9.8;
        set (float64): (uint16)(value / 16.0 / 9.8 * 32768.0);
    };
    uint16 y[2] {
        get (float64): value / 32768.0 * 16.0 * 9.8;
        set (float64): (uint16)(value / 16.0 / 9.8 * 32768.0);
    };
    uint16 z[2] {
        get (float64): value / 32768.0 * 16.0 * 9.8;
        set (float64): (uint16)(value / 16.0 / 9.8 * 32768.0);
    };
}

struct RovSensorAngularVelocityData[6] {
    uint16 x[2] {
        get (float64): value / 32768.0 * 2000.0;
        set (float64): (uint16)(value / 2000.0 * 32768.0);
    };
    uint16 y[2] {
        get (float64): value / 32768.0 * 2000.0;
        set (float64): (uint16)(value / 2000.0 * 32768.0);
    };
    uint16 z[2] {
        get (float64): value / 32768.0 * 2000.0;
        set (float64): (uint16)(value / 2000.0 * 32768.0);
    };
}

struct RovSensorEulerAngleData[6] {
    int16 pitch[2] {
        get (float64): (float64)value;
        set (float64): (int16)value;
    };
    int16 roll[2] {
        get (float64): (float64)value;
        set (float64): (int16)value;
    };
    uint16 yaw[2] {
        get (float64): (float64)((value + 180) % 360);
        set (float64): (uint16)(((int16)value + 180) % 360);
    };
}

struct RovSensorMagneticFieldData[6] {
    uint16 x[2] {
        get (float64): value / 32768.0 * 4912.0;
        set (float64): (uint16)(value / 4912.0 * 32768.0);
    };
    uint16 y[2] {
        get (float64): value / 32768.0 * 4912.0;
        set (float64): (uint16)(value / 4912.0 * 32768.0);
    };
    uint16 z[2] {
        get (float64): value / 32768.0 * 4912.0;
        set (float64): (uint16)(value / 4912.0 * 32768.0);
    };
}

struct RovSensorHeightSonarData[6] {
    uint16 height[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    uint16 confidence[2] {
        get trustable(bool): (value / 100.0) <= 95 ? false : true;
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    void [2];
}

struct RovSensorDistanceSonarData[6] {
    uint16 distance[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    uint16 confidence[2] {
        get (float64): value / 100.0;
        set (float64): (uint16)(value * 100.0);
    };
    void [2];
}
