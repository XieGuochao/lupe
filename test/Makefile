# Define the directories to exclude
EXCLUDE_DIRS := bin/

# Get all directories except the excluded ones
# Remove / from the end of the directory name
DIRS := $(filter-out $(EXCLUDE_DIRS), $(wildcard */))
DIRS := $(patsubst %/,%,$(DIRS))

# Define the build target
# Execute build on every directory except EXCLUDE_DIRS

build:
	@for dir in $(DIRS); do \
		cd $$dir; \
		go build -o ../bin/$$dir; \
		cd ..; \
	done

# Define the default target
default: build

clean:
	rm -rf bin/*

# Define the phony targets
.PHONY: build clean default