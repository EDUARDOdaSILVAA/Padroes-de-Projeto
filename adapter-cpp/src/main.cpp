#include <iostream>
#include "target.h"
#include "LegacySensor.h"
#include "SensorAdapter.h"

void ClientCode(const Target *target) {
    std::cout << target->GetData() << std::endl;
}

int main() {
    std::cout << "--- Sistema de Leitura de Sensores ---\n\n";

    Target *modernTarget = new Target();
    std::cout << "1. Testando com o alvo moderno nativo:\n";
    ClientCode(modernTarget);
    std::cout << "\n";

    LegacySensor *oldSensor = new LegacySensor();
    std::cout << "2. Sensor legado (formato incompatível):\n";
    std::cout << oldSensor->GetReversedData() << "\n\n";

    SensorAdapter *adapter = new SensorAdapter(oldSensor);
    std::cout << "3. Usando o Adapter:\n";
    ClientCode(adapter);

    delete modernTarget;
    delete oldSensor;
    delete adapter;

    return 0;
}