PATTERNS := factory abstract_factory builder singleton prototype adapter bridge composite decorator facade flyweight proxy
pattern = $(word 1, $@)


.PHONY: $(PATTERNS)
$(PATTERNS):
	go run main.go -pattern $(pattern)
