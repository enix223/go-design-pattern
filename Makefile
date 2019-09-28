PATTERNS := factory abstract_factory builder singleton prototype
pattern = $(word 1, $@)


.PHONY: $(PATTERNS)
$(PATTERNS):
	go run main.go -pattern $(pattern)