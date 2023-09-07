package sql

import (
	"github.com/cockroachdb/cockroachdb-parser/pkg/sql/parser"
	"github.com/cockroachdb/cockroachdb-parser/pkg/sql/sem/tree"
	"regexp"
	"strings"
	"unicode"
)

type SQLFormat struct {
}

var ignoreComments = regexp.MustCompile(`^--.*\s*`)

func (obj *SQLFormat) Format(input string) (string, error) {
	cfg := tree.DefaultPrettyCfg()
	cfg.UseTabs = true
	cfg.LineWidth = 60
	cfg.TabWidth = 4
	cfg.Simplify = true
	cfg.Align = tree.PrettyNoAlign
	cfg.Case = strings.ToUpper
	cfg.JSONFmt = true
	cfg.Align = tree.PrettyAlignAndDeindent

	stmts := []string{input}
	var prettied strings.Builder
	for _, stmt := range stmts {
		for len(stmt) > 0 {
			stmt = strings.TrimSpace(stmt)
			hasContent := false
			// Trim comments, preserving whitespace after them.
			for {
				found := ignoreComments.FindString(stmt)
				if found == "" {
					break
				}
				// Remove trailing whitespace but keep up to 2 newlines.
				prettied.WriteString(strings.TrimRightFunc(found, unicode.IsSpace))
				newlines := strings.Count(found, "\n")
				if newlines > 2 {
					newlines = 2
				}
				prettied.WriteString(strings.Repeat("\n", newlines))
				stmt = stmt[len(found):]
				hasContent = true
			}
			// Split by semicolons
			next := stmt
			if pos, _ := parser.SplitFirstStatement(stmt); pos > 0 {
				next = stmt[:pos]
				stmt = stmt[pos:]
			} else {
				stmt = ""
			}
			// This should only return 0 or 1 responses.
			allParsed, err := parser.Parse(next)
			if err != nil {
				return "", err
			}
			for _, parsed := range allParsed {
				prettied.WriteString(cfg.Pretty(parsed.AST))
				prettied.WriteString(";\n")
				hasContent = true
			}
			if hasContent {
				prettied.WriteString("\n")
			}
		}
	}

	return strings.TrimRightFunc(prettied.String(), unicode.IsSpace), nil
}
