#include "SensorAdapter.h"
#include <algorithm>

SensorAdapter::SensorAdapter(LegacySensor *sensor) : legacySensor(sensor) {}

std::string SensorAdapter::GetData() const {
    std::string data = legacySensor->GetReversedData();
    std::reverse(data.begin(), data.end());
    return "Adapter traduziu: " + data;
}