package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	tagMinLen   = 1
	tagMaxLen   = 128
	imageMinLen = 1
	imageMaxLen = 255
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func StrNow() string {
	t := time.Now()
	year, month, day := t.Date()
	hour, min, _ := t.Clock()
	return fmt.Sprintf("%d%02d%02d%02d%02d", year, month, day, hour, min)
}

var varRegex = regexp.MustCompile(`\$\{[\w-.]+\}`)

func TemplateStringRender(template string, data map[string]string) string {
	var indexList = varRegex.FindAllStringIndex(template, -1)

	var buffer bytes.Buffer
	var i = 0
	for _, indexRange := range indexList {
		var begin = indexRange[0]
		var end = indexRange[1]
		var key = template[begin+2 : end-1]
		var value = data[key]
		var tmp = template[i:begin]
		buffer.WriteString(replaceDeprecatedFormat(tmp, data))
		buffer.WriteString(value)
		i = end
	}
	buffer.WriteString(replaceDeprecatedFormat(template[i:], data))
	return buffer.String()
}

var varRegex2 = regexp.MustCompile(`\$(?:branch|commit|time)`)

func replaceDeprecatedFormat(template string, data map[string]string) string {
	return varRegex2.ReplaceAllStringFunc(template, func(match string) string {
		return data[match[1:]]
	})
}

func GenImageTag(tagFormat, branch, commit string) (string, error) {

	// fmt.Printf("GenImageName, tagFormat=%s, args=%+v", tagFormat, args)
	fmt.Printf("GenImageName, tagFormat=%s", tagFormat)

	//if err := ValidateImageName(image); err != nil {
	//	return "", err
	//}

	var data = make(map[string]string, 3)
	data["branch"] = branch
	data["commit"] = commit
	data["time"] = StrNow()
	// for key, value := range args {
	// 	data[key] = value
	// }
	var tag = TemplateStringRender(tagFormat, data)

	if err := ValidateTagName(tag); err != nil {
		return "", err
	}

	// var result = fmt.Sprintf("%s:%s", image, tag)
	return tag, nil
}

func ShortCommitSHA(commitSHA string) string {
	if len(commitSHA) > 7 {
		return commitSHA[:7]
	}

	return commitSHA
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func isIllegalLength(s string, min int, max int) bool {
	if min == -1 {
		return (len(s) > max)
	}
	if max == -1 {
		return (len(s) < min)
	}
	return (len(s) < min || len(s) > max)
}

func ValidateImageName(image string) error {
	if isIllegalLength(image, imageMinLen, imageMaxLen) {
		return fmt.Errorf("image is illegal in length. (greater than %d or less than %d)",
			imageMinLen, imageMaxLen)
	}
	validImage := regexp.MustCompile("^[a-z0-9]+(?:[._-][a-z0-9]+)*$")
	legal := validImage.MatchString(image)
	if !legal {
		return fmt.Errorf("image name:%s contains illegal characters!", image)
	}
	return nil
}

var validTag = regexp.MustCompile("^\\w+[\\w.-]*$")

func ValidateTagName(tag string) error {
	if isIllegalLength(tag, tagMinLen, tagMaxLen) {
		return fmt.Errorf("tag name is illegal in length. (greater than %d or less than %d)",
			tagMinLen, tagMaxLen)
	}
	legal := validTag.MatchString(tag)
	if !legal {
		return fmt.Errorf("tag:%s contains illegal characters!", tag)
	}
	return nil
}

func ParseSimpleArg(s string) (string, string) {
	var index = strings.Index(s, "=")
	if index < 0 {
		return s, ""
	}
	return s[:index], s[index+1:]
}
