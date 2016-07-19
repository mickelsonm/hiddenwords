package hiddenwords

import (
	"errors"
	"sort"
)

//TextPiece is a piece of the text
type TextPiece struct {
	Text  string
	Count int
}

//TextPieces is a slice of TextPiece
type TextPieces []TextPiece

//Sets up sorting for text pieces
//Default sorting will be based on the count and in descending order
func (s TextPieces) Len() int           { return len(s) }
func (s TextPieces) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TextPieces) Less(i, j int) bool { return s[i].Count > s[j].Count }

//GetCharLookup - gets the character lookup for provided charset
func GetCharLookup(charsText string) (map[string]TextPiece, error) {
	if len(charsText) < 3 {
		return nil, errors.New("Must have at least 3 characters")
	}
	pieceStore := make(map[string]TextPiece)
	for _, c := range charsText {
		char := string(c)
		pieceStore[char] = TextPiece{
			Text: char,
		}
	}
	return pieceStore, nil
}

//ProcessText - analyzes and sorts the text constructing the pieces
func ProcessText(searchText string, lookup map[string]TextPiece) (TextPieces, error) {
	if len(searchText) < 50 {
		return nil, errors.New("Must have at least 50 characters")
	}
	for _, t := range searchText {
		char := string(t)
		if piece, found := lookup[char]; found {
			piece.Count++
			lookup[char] = piece
		}
	}

	var pieces TextPieces
	for _, p := range lookup {
		pieces = append(pieces, p)
	}
	sort.Sort(pieces)

	return pieces, nil
}

//GetAnswer - gets the current answer for the pieces
func (s TextPieces) GetAnswer() string {
	answer := ""
	for _, p := range s {
		//drops all pieces after and including _
		if p.Text == "_" {
			break
		}
		answer += p.Text
	}
	return answer
}
