package utils

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"
)

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReplacePlaceholders(template string, values map[string]string) string {
	for key, value := range values {
		template = strings.ReplaceAll(template, key, value)
	}
	return template
}

func ParsePromptTemplate(templateName string, values map[string]any) (string, error) {
	tmplContent, tmplContentErr := ReadFile(templateName)
	if tmplContentErr != nil {
		return "", tmplContentErr
	}
	templ, templErr := template.New("interviewPrompt").Parse(tmplContent)
	if templErr != nil {
		return "", templErr
	}
	var promptBuffer bytes.Buffer
	if err := templ.Execute(&promptBuffer, values); err != nil {
		return "", err
	}
	return promptBuffer.String(), nil
}

func IncludeString(list []string, item string) bool {
	include := false
	for _, listItem := range list {
		if listItem == item {
			include = true
			break
		}
	}
	return include
}
