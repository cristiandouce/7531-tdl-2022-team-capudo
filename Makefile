mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(abspath $(patsubst %/,%,$(dir $(mkfile_path))))

run:
	$(MAKE) -C ./proyecto run

run_big: export PATH_RECORRIDOS="$(current_dir)/proyecto/datasets/recorridos_big_350.csv"
run_big: run


