// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: config/config.go

package config

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

const (
	ffjtConfigbase = iota
	ffjtConfignosuchkey

	ffjtConfigSecond

	ffjtConfigMinute

	ffjtConfigHour

	ffjtConfigDay

	ffjtConfigMonth

	ffjtConfigYear

	ffjtConfigLimitBy

	ffjtConfigHeaderName

	ffjtConfigPath

	ffjtConfigPolicy

	ffjtConfigFaultTolerant

	ffjtConfigHideClientHeaders
)

var ffjKeyConfigSecond = []byte("second")

var ffjKeyConfigMinute = []byte("minute")

var ffjKeyConfigHour = []byte("hour")

var ffjKeyConfigDay = []byte("day")

var ffjKeyConfigMonth = []byte("month")

var ffjKeyConfigYear = []byte("year")

var ffjKeyConfigLimitBy = []byte("limit_by")

var ffjKeyConfigHeaderName = []byte("header_name")

var ffjKeyConfigPath = []byte("path")

var ffjKeyConfigPolicy = []byte("policy")

var ffjKeyConfigFaultTolerant = []byte("fault_tolerant")

var ffjKeyConfigHideClientHeaders = []byte("hide_client_headers")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Config) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Config) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtConfigbase
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
				currentKey = ffjtConfignosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffjKeyConfigDay, kn) {
						currentKey = ffjtConfigDay
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyConfigFaultTolerant, kn) {
						currentKey = ffjtConfigFaultTolerant
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffjKeyConfigHour, kn) {
						currentKey = ffjtConfigHour
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyConfigHeaderName, kn) {
						currentKey = ffjtConfigHeaderName
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyConfigHideClientHeaders, kn) {
						currentKey = ffjtConfigHideClientHeaders
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyConfigLimitBy, kn) {
						currentKey = ffjtConfigLimitBy
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyConfigMinute, kn) {
						currentKey = ffjtConfigMinute
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyConfigMonth, kn) {
						currentKey = ffjtConfigMonth
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyConfigPath, kn) {
						currentKey = ffjtConfigPath
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyConfigPolicy, kn) {
						currentKey = ffjtConfigPolicy
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyConfigSecond, kn) {
						currentKey = ffjtConfigSecond
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'y':

					if bytes.Equal(ffjKeyConfigYear, kn) {
						currentKey = ffjtConfigYear
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyConfigHideClientHeaders, kn) {
					currentKey = ffjtConfigHideClientHeaders
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyConfigFaultTolerant, kn) {
					currentKey = ffjtConfigFaultTolerant
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyConfigPolicy, kn) {
					currentKey = ffjtConfigPolicy
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyConfigPath, kn) {
					currentKey = ffjtConfigPath
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyConfigHeaderName, kn) {
					currentKey = ffjtConfigHeaderName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyConfigLimitBy, kn) {
					currentKey = ffjtConfigLimitBy
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyConfigYear, kn) {
					currentKey = ffjtConfigYear
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyConfigMonth, kn) {
					currentKey = ffjtConfigMonth
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyConfigDay, kn) {
					currentKey = ffjtConfigDay
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyConfigHour, kn) {
					currentKey = ffjtConfigHour
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyConfigMinute, kn) {
					currentKey = ffjtConfigMinute
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyConfigSecond, kn) {
					currentKey = ffjtConfigSecond
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtConfignosuchkey
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

				case ffjtConfigSecond:
					goto handle_Second

				case ffjtConfigMinute:
					goto handle_Minute

				case ffjtConfigHour:
					goto handle_Hour

				case ffjtConfigDay:
					goto handle_Day

				case ffjtConfigMonth:
					goto handle_Month

				case ffjtConfigYear:
					goto handle_Year

				case ffjtConfigLimitBy:
					goto handle_LimitBy

				case ffjtConfigHeaderName:
					goto handle_HeaderName

				case ffjtConfigPath:
					goto handle_Path

				case ffjtConfigPolicy:
					goto handle_Policy

				case ffjtConfigFaultTolerant:
					goto handle_FaultTolerant

				case ffjtConfigHideClientHeaders:
					goto handle_HideClientHeaders

				case ffjtConfignosuchkey:
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

handle_Second:

	/* handler: j.Second type=int64 kind=int64 quoted=false*/

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

			j.Second = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Minute:

	/* handler: j.Minute type=int64 kind=int64 quoted=false*/

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

			j.Minute = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Hour:

	/* handler: j.Hour type=int64 kind=int64 quoted=false*/

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

			j.Hour = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Day:

	/* handler: j.Day type=int64 kind=int64 quoted=false*/

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

			j.Day = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Month:

	/* handler: j.Month type=int64 kind=int64 quoted=false*/

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

			j.Month = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Year:

	/* handler: j.Year type=int64 kind=int64 quoted=false*/

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

			j.Year = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LimitBy:

	/* handler: j.LimitBy type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.LimitBy = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HeaderName:

	/* handler: j.HeaderName type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.HeaderName = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Path:

	/* handler: j.Path type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Path = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Policy:

	/* handler: j.Policy type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Policy = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FaultTolerant:

	/* handler: j.FaultTolerant type=bool kind=bool quoted=false*/

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

				j.FaultTolerant = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.FaultTolerant = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_HideClientHeaders:

	/* handler: j.HideClientHeaders type=bool kind=bool quoted=false*/

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

				j.HideClientHeaders = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.HideClientHeaders = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

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
