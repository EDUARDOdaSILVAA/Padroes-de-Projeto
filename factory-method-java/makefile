
SRC_DIR = src/main/java
OUT_DIR = bin
MAIN_CLASS = com.Main

all: run

compile:
	@echo "A compilar o projeto Java..."
	@mkdir -p $(OUT_DIR)
	@find $(SRC_DIR) -name "*.java" | xargs javac -d $(OUT_DIR)

run: compile
	@echo "A executar o Factory Method..."
	@java -cp $(OUT_DIR) $(MAIN_CLASS)

clean:
	@echo "A limpar ficheiros binários..."
	@rm -rf $(OUT_DIR)