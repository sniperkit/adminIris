package util

import (
	"html/template"
	"time"
)

// TemplateFuncCSSAndJavascript ...
func TemplateFuncCSSAndJavascript(tags []string, tagType string) template.HTML {
	t := time.Now()
	ts := t.Format("20060102150405")
	var resultStr string
	for _, value := range tags {
		if tagType == "css" {
			resultStr += `<link rel="stylesheet" type="text/css" href="/resources/css/` + value + `?t=` + ts + `" />`
		} else if tagType == "js" {
			resultStr += `<script type="text/javascript" src="/resources/js/` + value + `?t=` + ts + `"></script>`
		}
	}
	return template.HTML(resultStr)
}
