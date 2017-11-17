SOURCE = ./swagger.yaml
LIB    = ./.lib
JSON   = models/models_easyjson.go

build: $(JSON)

$(JSON): replace
	easyjson --output_filename $@ -pkg -all models

$(LIB): $(SOURCE)
	swagger generate client --skip-validation -f $<
	touch $@

clean::
	rm -rf client models $(LIB)

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
	sed -i '' 's/*os.File/runtime.NamedReadCloser/g' ./client/updates/set_webhook_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/attachments/send_video_note_parameters.go
	sed -i '' 's/\&o.VideoNote/o.VideoNote/g' ./client/attachments/send_video_note_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/chats/set_chat_photo_parameters.go
	sed -i '' 's/\&o.Photo/o.Photo/g' ./client/chats/set_chat_photo_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/stickers/upload_sticker_file_parameters.go
	sed -i '' 's/\&o.PngSticker/o.PngSticker/g' ./client/stickers/upload_sticker_file_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/stickers/create_new_sticker_set_parameters.go
	sed -i '' 's/\&o.PngSticker/o.PngSticker/g' ./client/stickers/create_new_sticker_set_parameters.go
	sed -i '' 's/os.File/runtime.NamedReadCloser/g' ./client/stickers/add_sticker_to_set_parameters.go
	sed -i '' 's/\&o.PngSticker/o.PngSticker/g' ./client/stickers/add_sticker_to_set_parameters.go
	goimports -w ./client
