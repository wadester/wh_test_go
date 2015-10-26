#!/usr/bin/make
# Module:  Makefile
# Purpose:  build the GLIB examples
# Author:   Wade Hampton
# Date:     9/25/2015
#

# object files to build, source is same name + ".go"
BINS=hello exercise-slices go_func go_func1 go_vars \
        go_loop go_ptr go_struct go_map go_funcs go_method \
	go_error go_reader1 go_reader2 go_websvr

# all rule to build all object files listed above
all:  $(OBJS)
	$(MAKE) $(BINS)

# go build rule, single files only
%: %.go
	go build $<

# cleanup rule
clean::
	rm -f $(BINS)

