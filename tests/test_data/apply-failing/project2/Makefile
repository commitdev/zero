REQUIRED_BINS := ls nonexisting-binary

current_dir:
	@echo "baz: ${baz}" > project.out

summary:

check:
		$(foreach bin, $(REQUIRED_BINS),\
			$(if $(shell command -v $(bin) 2> /dev/null),$(info Found `$(bin)`),$(error Please install `$(bin)`)))
