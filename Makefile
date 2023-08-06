TESTS := $(wildcard tests/*.sh)

build:
	go build

test: build
	$(MAKE) $(TESTS)
	@printf "\e[32mPass all tests\e[0m\n"

clean:
	go clean
	rm -rf out/

 $(TESTS):
	@echo "Testing" $@
	@./$@
	@printf "\e[32mOK\e[0m\n"

.PHONY: build clean test $(TESTS)
