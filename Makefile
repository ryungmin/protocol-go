.PHONY: dist clean

all:
ifeq ($(OS),Windows_NT)
	@echo "use 'build.bat'"
else #ifeq ($(OS),Windows_NT)
ifeq ($(shell uname -s),Linux)
	@make -f Makefile.linux
endif #ifeq ($(shell uname -s),Linux)
ifeq ($(shell uname -s),Darwin)
	@make -f Makefile.darwin
endif #ifeq ($(shell uname -s),Darwin)
endif

dist:
ifeq ($(OS),Windows_NT)
	@echo "use 'build.bat dist'"
else #ifeq ($(OS),Windows_NT)
ifeq ($(shell uname -s),Linux)
	@make -f Makefile.linux dist	
endif #ifeq ($(shell uname -s),Linux)
ifeq ($(shell uname -s),Darwin)
	@make -f Makefile.darwin dist
endif #ifeq ($(shell uname -s),Darwin)
endif

clean:
ifeq ($(OS),Windows_NT)
	@echo "use 'build.bat clean'"
else #ifeq ($(OS),Windows_NT)
ifeq ($(shell uname -s),Linux)
	@make -f Makefile.linux clean	
endif #ifeq ($(shell uname -s),Linux)
ifeq ($(shell uname -s),Darwin)
	@make -f Makefile.darwin clean
endif #ifeq ($(shell uname -s),Darwin)
endif
