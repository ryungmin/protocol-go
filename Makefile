all:
ifeq ($(OS),Windows_NT)
ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
	@make -f Makefile.windows
else # ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
	@make -f Makefile.windows
endif
ifeq ($(PROCESSOR_ARCHITECTURE),x86)
	@make -f Makefile.windows build_win32
endif
endif #ifeq ($(PROCESSOR_ARCHITEW6432),AMD64)
else #ifeq ($(OS),Windows_NT)
ifeq ($(shell uname -s),Linux)
	@make -f Makefile.linux
endif #ifeq ($(shell uname -s),Linux)
ifeq ($(shell uname -s),Darwin)
	@make -f Makefile.darwin
endif #ifeq ($(shell uname -s),Darwin)
endif

clean:
ifeq ($(OS),Windows_NT)
	@make -f Makefile.windows clean
else #ifeq ($(OS),Windows_NT)
ifeq ($(shell uname -s),Linux)
	@make -f Makefile.linux clean	
endif #ifeq ($(shell uname -s),Linux)
ifeq ($(shell uname -s),Darwin)
	@make -f Makefile.darwin clean
endif #ifeq ($(shell uname -s),Darwin)
endif
