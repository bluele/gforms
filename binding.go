package gforms

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

func parseReuqestBody(req *http.Request) (*Data, *RawData, error) {
	if isNilValue(req) {
		return nil, nil, errors.New("*http.Request is nil.")
	}
	contentType := req.Header.Get("Content-Type")
	if req.Method == "POST" || req.Method == "PUT" || contentType != "" {
		if strings.Contains(contentType, "json") {
			return bindJson(req)
		} else if strings.Contains(contentType, "multipart/form-data") {
			return bindMultiPartForm(req)
		} else {
			return bindForm(req)
		}
	}
	return nil, nil, nil
}

func bindJson(req *http.Request) (*Data, *RawData, error) {
	var jsonBody map[string]json.RawMessage
	data := Data{}
	rawData := RawData{}
	if req.Body == nil {
		return &data, &rawData, nil
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range jsonBody {
		switch c := v[0]; c {
		case 'n':
			data[k] = nilV()
			rawData[k] = string(v)
		case 't', 'f':
			if c == 't' {
				data[k] = newV(true, reflect.Bool)
			} else {
				data[k] = newV(false, reflect.Bool)
			}
			rawData[k] = string(v)
		case '"':
			s, ok := unquoteBytes(v)
			if ok {
				str := string(s)
				data[k] = newV([]string{str}, reflect.String)
				rawData[k] = str
			}
		default:
			data[k] = newV([]string{string(v)}, reflect.String)
			rawData[k] = string(v)
		}
	}
	return &data, &rawData, nil
}

func bindForm(req *http.Request) (*Data, *RawData, error) {
	req.ParseForm()
	data := Data{}
	rawData := RawData{}
	for name, v := range req.Form {
		if len(v) != 0 {
			data[name] = newV(v, reflect.String)
			rawData[name] = v[0]
		}
	}
	return &data, &rawData, nil
}

func bindMultiPartForm(req *http.Request) (*Data, *RawData, error) {
	req.ParseMultipartForm(32 << 20)
	data := Data{}
	rawData := RawData{}
	for name, v := range req.MultipartForm.Value {
		if len(v) != 0 {
			data[name] = newV(v, reflect.String)
		}
	}
	for name, v := range req.MultipartForm.File {
		if len(v) != 0 {
			data[name] = newV(v, reflect.Array)
		}
	}
	return &data, &rawData, nil
}

func bindValues(uv url.Values) (*Data, *RawData, error) {
	data := Data{}
	rawData := RawData{}
	for name, v := range uv {
		if len(v) != 0 {
			data[name] = newV(v, reflect.String)
			rawData[name] = v[0]
		}
	}
	return &data, &rawData, nil
}

func unquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return
	}
	s = s[1 : len(s)-1]

	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room?  Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == '\\':
			r++
			if r >= len(s) {
				return
			}
			switch s[r] {
			default:
				return
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])
				if rr < 0 {
					return
				}
				r += 6
				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)
						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == '"', c < ' ':
			return

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w], true
}

func getu4(s []byte) rune {
	if len(s) < 6 || s[0] != '\\' || s[1] != 'u' {
		return -1
	}
	r, err := strconv.ParseUint(string(s[2:6]), 16, 64)
	if err != nil {
		return -1
	}
	return rune(r)
}
