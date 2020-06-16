package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"io/ioutil"

	// "bufio"
	// "encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Users struct{
	Browsers []string `json:"browsers"`
	company  string
	country  string
	Email    string `json:"email"`
	job      string
	Name     string `json:"name"`
	phone    int
}

func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	filecont, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	seenBrowsers := make(map[string]interface{})
	uniqueBrowsers := 0
	email := ""
	foundUsers := ""

	dataLine := strings.Split(string(filecont), "\n")
	users := make([]Users, 0)
	var user Users
	for _, data := range dataLine {
		err := easyjson.Unmarshal([]byte(data), &user)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
		user = Users{}
	}
	/* r := reflect.TypeOf(u)
	// t := reflect.ValueOf(u)
	for i := 0; i < r.NumField(); i++{
		fmt.Printf("%s = %v\n", r.Field(i).Tag.Get("json"))
	} */

	for i, u := range users {

		isAndroid := false
		isMSIE := false

		for _, browser := range u.Browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true
				seenBrowsers[browser] = true

				//notSeenBefore := true
				if _, saw := seenBrowsers[browser]; !saw {
					_ = false
					uniqueBrowsers++
				}
			}
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				seenBrowsers[browser] = true

				//notSeenBefore := true
				if _, saw := seenBrowsers[browser]; !saw {
					_ = false
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email = strings.ReplaceAll(u.Email, "@", " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, u.Name, email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}

		/*if isAndroid && isMSIE {
			// log.Println("Android and MSIE user:", user["name"], user["email"])
			email = strings.ReplaceAll(u.Email, "@", " [at] ")
			foundUsers = fmt.Sprintf("[%d] %s <%s>", i, u.Name, email)
			fmt.Fprintln(out, foundUsers)
		}
	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}*/

 var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB4ad3f7dDecodeCourseraGoHw3Struct(in *jlexer.Lexer, out *Users) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "email":
			out.Email = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browsers = nil
			} else {
				in.Delim('[')
				if out.Browsers == nil {
					if !in.IsDelim(']') {
						out.Browsers = make([]string, 0, 4)
					} else {
						out.Browsers = []string{}
					}
				} else {
					out.Browsers = (out.Browsers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browsers = append(out.Browsers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}

func easyjsonB4ad3f7dEncodeCourseraGoHw3Struct(out *jwriter.Writer, in Users) {
	out.RawByte('{')
	// first := true
	_ = true
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix)
		if in.Browsers == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browsers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Users) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB4ad3f7dEncodeCourseraGoHw3Struct(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Users) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB4ad3f7dEncodeCourseraGoHw3Struct(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Users) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB4ad3f7dDecodeCourseraGoHw3Struct(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Users) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB4ad3f7dDecodeCourseraGoHw3Struct(l, v)
}
