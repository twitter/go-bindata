# Variables
TESTDATA_DIR := ./testdata

# Targets
.PHONY: all clean

all:
	make -C $(TESTDATA_DIR)

clean:
	-rm $(TESTDATA_DIR)/generated_file
