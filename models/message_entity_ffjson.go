// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: models/message_entity.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *MessageEntity) MarshalJSON() ([]byte, error) {
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
func (mj *MessageEntity) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Length != 0 {
		buf.WriteString(`"length":`)
		fflib.FormatBits2(buf, uint64(mj.Length), 10, mj.Length < 0)
		buf.WriteByte(',')
	}
	if mj.Offset != 0 {
		buf.WriteString(`"offset":`)
		fflib.FormatBits2(buf, uint64(mj.Offset), 10, mj.Offset < 0)
		buf.WriteByte(',')
	}
	if len(mj.Type) != 0 {
		buf.WriteString(`"type":`)
		fflib.WriteJsonString(buf, string(mj.Type))
		buf.WriteByte(',')
	}
	if len(mj.URL) != 0 {
		buf.WriteString(`"url":`)
		fflib.WriteJsonString(buf, string(mj.URL))
		buf.WriteByte(',')
	}
	if mj.User != nil {
		if true {
			/* Struct fall back. type=models.User kind=struct */
			buf.WriteString(`"user":`)
			err = buf.Encode(mj.User)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_MessageEntitybase = iota
	ffj_t_MessageEntityno_such_key

	ffj_t_MessageEntity_Length

	ffj_t_MessageEntity_Offset

	ffj_t_MessageEntity_Type

	ffj_t_MessageEntity_URL

	ffj_t_MessageEntity_User
)

var ffj_key_MessageEntity_Length = []byte("length")

var ffj_key_MessageEntity_Offset = []byte("offset")

var ffj_key_MessageEntity_Type = []byte("type")

var ffj_key_MessageEntity_URL = []byte("url")

var ffj_key_MessageEntity_User = []byte("user")

func (uj *MessageEntity) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *MessageEntity) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_MessageEntitybase
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
				currentKey = ffj_t_MessageEntityno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'l':

					if bytes.Equal(ffj_key_MessageEntity_Length, kn) {
						currentKey = ffj_t_MessageEntity_Length
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffj_key_MessageEntity_Offset, kn) {
						currentKey = ffj_t_MessageEntity_Offset
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_MessageEntity_Type, kn) {
						currentKey = ffj_t_MessageEntity_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffj_key_MessageEntity_URL, kn) {
						currentKey = ffj_t_MessageEntity_URL
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_MessageEntity_User, kn) {
						currentKey = ffj_t_MessageEntity_User
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_MessageEntity_User, kn) {
					currentKey = ffj_t_MessageEntity_User
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_MessageEntity_URL, kn) {
					currentKey = ffj_t_MessageEntity_URL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_MessageEntity_Type, kn) {
					currentKey = ffj_t_MessageEntity_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_MessageEntity_Offset, kn) {
					currentKey = ffj_t_MessageEntity_Offset
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_MessageEntity_Length, kn) {
					currentKey = ffj_t_MessageEntity_Length
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_MessageEntityno_such_key
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

				case ffj_t_MessageEntity_Length:
					goto handle_Length

				case ffj_t_MessageEntity_Offset:
					goto handle_Offset

				case ffj_t_MessageEntity_Type:
					goto handle_Type

				case ffj_t_MessageEntity_URL:
					goto handle_URL

				case ffj_t_MessageEntity_User:
					goto handle_User

				case ffj_t_MessageEntityno_such_key:
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

handle_Length:

	/* handler: uj.Length type=int64 kind=int64 quoted=false*/

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

			uj.Length = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Offset:

	/* handler: uj.Offset type=int64 kind=int64 quoted=false*/

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

			uj.Offset = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Type:

	/* handler: uj.Type type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Type = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_URL:

	/* handler: uj.URL type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.URL = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_User:

	/* handler: uj.User type=models.User kind=struct quoted=false*/

	{
		/* Falling back. type=models.User kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.User)
		if err != nil {
			return fs.WrapErr(err)
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