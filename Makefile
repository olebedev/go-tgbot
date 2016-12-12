SOURCE    = ./swagger.yaml
LIB       = ./.lib
MODELS  = $(patsubst %.go, %_ffjson.go, $(shell find models -name "*.go" -type f -not -name "*ffjson.go" -not -name "*_expose.go" -not -name "ffjson-*.go"))

build: $(LIB)

$(LIB): $(SOURCE)
	swagger generate client -f $<
	touch $@

models: clean-ffjson $(MODELS)

clean::
	rm -rf client models $(LIB)

models/%_ffjson.go: models/%.go
	ffjson -w=$@ $<

clean-ffjson::
	@find $(@D) -type f -name "*_ffjson_expose.go" -exec rm {} +
	@find $(@D) -type d -name "ffjson-*" -exec rm -r {} +
