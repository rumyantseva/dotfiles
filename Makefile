all: test

test:
	for file in $(shell find $(CURDIR)/dotfiles -type f); do \
		f=$$(basename $$file); \
		ln -sfn $$file $(HOME)/$$f; \
	done
