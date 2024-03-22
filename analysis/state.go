package analysis

import (
	"educationalsp/lsp"
	"strings"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri string, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri string, text string) {
	s.Documents[uri] = text
}
func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	document := s.Documents[uri]
	// get the word  at the given position
	// Get the line from the docuemnt at position.Line
	// Get the word from the line at position.Character

	word := ""
	// Split document into lines
	lines := strings.Split(document, "\n")
	// Get the line at position.Line
	line := lines[position.Line]
	// Get the word at position.Character
	wordPosition := position.Character
	for i := wordPosition; i >= 0; i-- {
		if line[i] == ' ' {
			break
		}
		wordPosition = i
	}
	for i := wordPosition; i < len(line); i++ {
		if line[i] == ' ' {
			break
		}
		word += string(line[i])
	}

	// Convert the word to lower case and remove all non-alphanumeric characters
	word = strings.ToLower(word)
	word = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1
	}, word)

	// Check if the word is "hotdog"
	responseText := "Not hotdog"

	if word == "hotdog" {
		responseText = "Hotdog"
	}

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: responseText,
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}
