package util

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"sort"
	"strings"
)

const MODIFIER_PREFIX = "//@ahum:"

// ModifyCodeByMarkers finds special marker comments in Go source code and inserts
// new code on the line just before the marker.
//
// Parameters:
//
//	src: The Go source code as a byte slice.
//	insertions: A map where the key is the marker comment (e.g., "// @ahum: imports")
//	            and the value is the new code to be inserted before it.
//
// Returns:
//
//	The modified source code as a byte slice.
//	An error if parsing fails or if any of the specified markers are not found in the source.
func ModifyCodeByMarkers(src []byte, insertions map[string]string) ([]byte, error) {
	// Validate that all markers are actual comments
	cleanInsertions := make(map[string]string)
	for marker := range insertions {
		cleanInsertions[strings.TrimSpace(marker)] = insertions[marker]
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "source.go", src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse source: %w", err)
	}

	// replacements maps a line's starting offset to the new text for that line.
	replacements := make(map[int]string)
	foundMarkers := make(map[string]bool)

	for _, commentGroup := range file.Comments {
		for _, comment := range commentGroup.List {
			trimmedComment := strings.TrimSpace(comment.Text)
			cleanComment := strings.ReplaceAll(trimmedComment, " ", "")
			cleanComment = strings.ToLower(cleanComment)
			cleanComment = strings.TrimPrefix(cleanComment, MODIFIER_PREFIX)

			if newCode, ok := cleanInsertions[cleanComment]; ok {
				pos := comment.Pos()
				offset := fset.File(pos).Offset(pos)

				// Find the start of the line containing the marker.
				lineStart := offset
				for lineStart > 0 && src[lineStart-1] != '\n' {
					lineStart--
				}

				// Get the original line's indentation.
				indentation := src[lineStart:offset]

				// *** KEY CHANGE IS HERE ***
				// Construct the replacement text by adding the new code, a newline,
				// the proper indentation, and the original marker.
				var replacementTextBuilder strings.Builder
				replacementTextBuilder.WriteString(string(indentation))
				replacementTextBuilder.WriteString(newCode)
				replacementTextBuilder.WriteString("\n")
				replacementTextBuilder.WriteString(string(indentation))
				replacementTextBuilder.WriteString(trimmedComment)

				replacements[lineStart] = replacementTextBuilder.String()
				foundMarkers[cleanComment] = true
			}
		}
	}

	// Verify that all requested markers were found.
	if len(foundMarkers) != len(cleanInsertions) {
		var missing []string
		for marker := range cleanInsertions {
			if !foundMarkers[marker] {
				missing = append(missing, marker)
			}
		}
		return nil, fmt.Errorf("could not find required markers: %s", strings.Join(missing, ", "))
	}

	var offsets []int
	for offset := range replacements {
		offsets = append(offsets, offset)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(offsets)))

	output := src
	for _, offset := range offsets {
		endOfLine := offset
		for endOfLine < len(output) && output[endOfLine] != '\n' {
			endOfLine++
		}

		var buf bytes.Buffer
		buf.Write(output[:offset])
		buf.WriteString(replacements[offset])
		buf.Write(output[endOfLine:])
		output = buf.Bytes()
	}

	return output, nil
}
