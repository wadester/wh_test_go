#!/usr/bin/make
# Module:  Makefile
# Purpose:  build the GLIB examples
# Author:   Wade Hampton
# Date:     9/25/2015
#

BINS=hello go_func go_func1

all:  $(OBJS)
	$(MAKE) $(BINS)


hello:  hello.go
	go build hello.go

go_func:  go_func.go
	go build go_func.go

go_func1:  go_func1.go
	go build go_func.go 

clean::
	rm -f $(BINS)

