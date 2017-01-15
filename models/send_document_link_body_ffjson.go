// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/send_document_link_body.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *SendDocumentLinkBody) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *SendDocumentLinkBody) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if len(mj.Caption) != 0 {
		buf.WriteString(`"caption":`)
		fflib.WriteJsonString(buf, string(mj.Caption))
		buf.WriteByte(',')
	}
	buf.WriteString(`"chat_id":`)
	/* Interface types must use runtime reflection. type=interface {} kind=interface */
	err = buf.Encode(mj.ChatID)
	if err != nil {
		return err
	}
	buf.WriteByte(',')
	if mj.DisableNotification != false {
		if mj.DisableNotification {
			buf.WriteString(`"disable_notification":true`)
		} else {
			buf.WriteString(`"disable_notification":false`)
		}
		buf.WriteByte(',')
	}
	if mj.Document != nil {
		buf.WriteString(`"document":`)
		fflib.WriteJsonString(buf, string(*mj.Document))
	} else {
		buf.WriteString(`"document":null`)
	}
	buf.WriteByte(',')
	if mj.ReplyMarkup != nil {
		buf.WriteString(`"reply_markup":`)
		/* Interface types must use runtime reflection. type=interface {} kind=interface */
		err = buf.Encode(mj.ReplyMarkup)
		if err != nil {
			return err
		}
		buf.WriteByte(',')
	}
	if mj.ReplyToMessageID != 0 {
		buf.WriteString(`"reply_to_message_id":`)
		fflib.FormatBits2(buf, uint64(mj.ReplyToMessageID), 10, mj.ReplyToMessageID < 0)
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_SendDocumentLinkBodybase = iota
	ffj_t_SendDocumentLinkBodyno_such_key

	ffj_t_SendDocumentLinkBody_Caption

	ffj_t_SendDocumentLinkBody_ChatID

	ffj_t_SendDocumentLinkBody_DisableNotification

	ffj_t_SendDocumentLinkBody_Document

	ffj_t_SendDocumentLinkBody_ReplyMarkup

	ffj_t_SendDocumentLinkBody_ReplyToMessageID
)

var ffj_key_SendDocumentLinkBody_Caption = []byte("caption")

var ffj_key_SendDocumentLinkBody_ChatID = []byte("chat_id")

var ffj_key_SendDocumentLinkBody_DisableNotification = []byte("disable_notification")

var ffj_key_SendDocumentLinkBody_Document = []byte("document")

var ffj_key_SendDocumentLinkBody_ReplyMarkup = []byte("reply_markup")

var ffj_key_SendDocumentLinkBody_ReplyToMessageID = []byte("reply_to_message_id")

func (uj *SendDocumentLinkBody) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *SendDocumentLinkBody) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_SendDocumentLinkBodybase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_SendDocumentLinkBodyno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'c':

					if bytes.Equal(ffj_key_SendDocumentLinkBody_Caption, kn) {
						currentKey = ffj_t_SendDocumentLinkBody_Caption
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_SendDocumentLinkBody_ChatID, kn) {
						currentKey = ffj_t_SendDocumentLinkBody_ChatID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_SendDocumentLinkBody_DisableNotification, kn) {
						currentKey = ffj_t_SendDocumentLinkBody_DisableNotification
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_SendDocumentLinkBody_Document, kn) {
						currentKey = ffj_t_SendDocumentLinkBody_Document
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffj_key_SendDocumentLinkBody_ReplyMarkup, kn) {
						currentKey = ffj_t_SendDocumentLinkBody_ReplyMarkup
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_SendDocumentLinkBody_ReplyToMessageID, kn) {
						currentKey = ffj_t_SendDocumentLinkBody_ReplyToMessageID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_SendDocumentLinkBody_ReplyToMessageID, kn) {
					currentKey = ffj_t_SendDocumentLinkBody_ReplyToMessageID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_SendDocumentLinkBody_ReplyMarkup, kn) {
					currentKey = ffj_t_SendDocumentLinkBody_ReplyMarkup
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SendDocumentLinkBody_Document, kn) {
					currentKey = ffj_t_SendDocumentLinkBody_Document
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_SendDocumentLinkBody_DisableNotification, kn) {
					currentKey = ffj_t_SendDocumentLinkBody_DisableNotification
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffj_key_SendDocumentLinkBody_ChatID, kn) {
					currentKey = ffj_t_SendDocumentLinkBody_ChatID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_SendDocumentLinkBody_Caption, kn) {
					currentKey = ffj_t_SendDocumentLinkBody_Caption
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_SendDocumentLinkBodyno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_SendDocumentLinkBody_Caption:
					goto handle_Caption

				case ffj_t_SendDocumentLinkBody_ChatID:
					goto handle_ChatID

				case ffj_t_SendDocumentLinkBody_DisableNotification:
					goto handle_DisableNotification

				case ffj_t_SendDocumentLinkBody_Document:
					goto handle_Document

				case ffj_t_SendDocumentLinkBody_ReplyMarkup:
					goto handle_ReplyMarkup

				case ffj_t_SendDocumentLinkBody_ReplyToMessageID:
					goto handle_ReplyToMessageID

				case ffj_t_SendDocumentLinkBodyno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Caption:

	/* handler: uj.Caption type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Caption = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ChatID:

	/* handler: uj.ChatID type=interface {} kind=interface quoted=false*/

	{
		/* Falling back. type=interface {} kind=interface */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.ChatID)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DisableNotification:

	/* handler: uj.DisableNotification type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.DisableNotification = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.DisableNotification = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Document:

	/* handler: uj.Document type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

			uj.Document = nil

		} else {

			var tval string
			outBuf := fs.Output.Bytes()

			tval = string(string(outBuf))
			uj.Document = &tval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReplyMarkup:

	/* handler: uj.ReplyMarkup type=interface {} kind=interface quoted=false*/

	{
		/* Falling back. type=interface {} kind=interface */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.ReplyMarkup)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ReplyToMessageID:

	/* handler: uj.ReplyToMessageID type=int64 kind=int64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.ReplyToMessageID = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
