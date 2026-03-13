package newsletter

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"strings"

	"github.com/yoanbernabeu/daybrief/internal/gemini"
)

//go:embed templates/email.html
var emailTemplate string

func paragraphs(s string) template.HTML {
	parts := strings.Split(s, "\n\n")
	var buf strings.Builder
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		buf.WriteString(`<p style="color:#333333;font-size:16px;line-height:1.7;margin:0 0 16px;">`)
		buf.WriteString(template.HTMLEscapeString(p))
		buf.WriteString("</p>\n")
	}
	return template.HTML(buf.String())
}

func RenderHTML(newsletter *gemini.Newsletter) (string, error) {
	funcMap := template.FuncMap{
		"paragraphs": paragraphs,
	}

	tmpl, err := template.New("email").Funcs(funcMap).Parse(emailTemplate)
	if err != nil {
		return "", fmt.Errorf("parsing email template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, newsletter); err != nil {
		return "", fmt.Errorf("executing email template: %w", err)
	}

	return buf.String(), nil
}
