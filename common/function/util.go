package function

import (
	"bytes"
	"cmp"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"path/filepath"
	"strings"
)

func GetRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz123456789"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GetMd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

type pathInfoOut struct {
	DirName   string
	BaseName  string
	Extension string
	Filename  string
}

func GetPathInfo(path string) *pathInfoOut {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)

	dirname, basename := filepath.Split(path)
	basename = basename[:len(basename)-len(ext)]
	result := &pathInfoOut{}
	result.DirName = dirname
	result.BaseName = basename
	result.Extension = ext
	result.Filename = filename
	return result
}

func CheckFileAllowUpload(filename string) bool {
	allowFileExt := []string{
		".zip", ".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".cvs",
		".jpg", ".png", ".jpeg", ".gif",
	}
	for _, s := range allowFileExt {
		if strings.HasSuffix(filename, s) {
			return true
		}
	}
	return false
}

func EncodeURIComponent(s string, excluded ...[]byte) string {
	var b bytes.Buffer
	written := 0
	for i, n := 0, len(s); i < n; i++ {
		c := s[i]
		switch c {
		case '-', '_', '.', '!', '~', '*', '\'', '(', ')':
			continue
		default:
			// Unreserved according to RFC 3986 sec 2.3
			if 'a' <= c && c <= 'z' {
				continue
			}
			if 'A' <= c && c <= 'Z' {
				continue
			}
			if '0' <= c && c <= '9' {
				continue
			}
			if len(excluded) > 0 {
				conti := false
				for _, ch := range excluded[0] {
					if ch == c {
						conti = true
						break
					}
				}
				if conti {
					continue
				}
			}
		}
		b.WriteString(s[written:i])
		fmt.Fprintf(&b, "%%%02X", c)
		written = i + 1
	}
	if written == 0 {
		return s
	}
	b.WriteString(s[written:])
	return b.String()
}

func Base64Encode(obj interface{}) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(obj)
	if err != nil {
		return ""
	}
	encoder.Close()
	return buf.String()
}

func Base64Decode(obj interface{}, enc string) error {
	return json.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(enc))).Decode(obj)
}

func IsEmptyArray[T interface{}](v []T) bool {
	if v == nil {
		return true
	}
	if len(v) == 0 {
		return true
	}
	return false
}

func InArray[T cmp.Ordered](v []T, item T) bool {
	if v == nil {
		return false
	}
	for _, t := range v {
		if t == item {
			return true
		}
	}
	return false
}
