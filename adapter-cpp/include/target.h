#ifndef TARGET_H
#define TARGET_H
#include <string>

class Target {
public:
    virtual ~Target() = default;
    virtual std::string GetData() const {
        return "Dados padrão do sistema moderno.";
    }
};
#endif