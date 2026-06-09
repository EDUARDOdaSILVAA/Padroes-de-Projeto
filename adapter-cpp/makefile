BINARY_NAME = adapter_app
CXX = g++
CXXFLAGS = -Wall -Wextra -std=c++11 -Iinclude

all: run

build:
	@echo "A compilar o código C++..."
	@$(CXX) $(CXXFLAGS) src/main.cpp src/LegacySensor.cpp src/SensorAdapter.cpp -o $(BINARY_NAME)

run: build
	@echo "A executar o Adapter..."
	@./$(BINARY_NAME)

clean:
	@echo "A limpar o executável..."
	@rm -f $(BINARY_NAME)