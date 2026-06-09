#ifndef SENSORADAPTER_H
#define SENSORADAPTER_H
#include "target.h"
#include "LegacySensor.h"

class SensorAdapter : public Target {
private:
    LegacySensor *legacySensor;
public:
    SensorAdapter(LegacySensor *sensor);
    std::string GetData() const override;
};
#endif