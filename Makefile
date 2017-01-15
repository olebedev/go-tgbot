SOURCE = ./swagger.yaml
LIB    = ./.lib
MODELS = $(patsubst %.go, %_ffjson.go, $(shell find models -name "*.go" -type f -not -name "*ffjson.go" -not -name "*_expose.go" -not -name "ffjson-*.go"))

build: replace

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

replace: $(LIB)
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/attachments/send_document_parameters.go
	sed -i '' 's/\&o.Document/o.Document/g' ./client/attachments/send_document_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/attachments/send_audio_parameters.go
	sed -i '' 's/\&o.Audio/o.Audio/g' ./client/attachments/send_audio_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/attachments/send_photo_parameters.go
	sed -i '' 's/\&o.Photo/o.Photo/g' ./client/attachments/send_photo_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/attachments/send_sticker_parameters.go
	sed -i '' 's/\&o.Sticker/o.Sticker/g' ./client/attachments/send_sticker_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/attachments/send_video_parameters.go
	sed -i '' 's/\&o.Video/o.Video/g' ./client/attachments/send_video_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/attachments/send_voice_parameters.go
	sed -i '' 's/\&o.Voice/o.Voice/g' ./client/attachments/send_voice_parameters.go
	goimports -w ./client/attachments
