package main

import (
	"fmt"

	"github.com/sourcegraph/sourcegraph/lib/codeintel/tools/lsif-flat/proto"
)

func main() {
	for _, val := range compile().Values {
		fmt.Println(val)
	}
}

func compile() proto.LsifValues {
	vals := []*proto.LsifValue{}

	s := "miso cat miso "
	word := ""
	doc := proto.Document{}
	defs := map[string]int{}
	for i, char := range s {
		role := proto.MonikerOccurrence_ROLE_REFERENCE
		if char == ' ' {
			if _, ok := defs[word]; !ok {
				role = proto.MonikerOccurrence_ROLE_DEFINITION
				defs[word] = i
				vals = append(vals, &proto.LsifValue{Value: &proto.LsifValue_Moniker{Moniker: &proto.Moniker{
					Id:            word,
					MarkdownHover: []string{fmt.Sprintf("Good %s! 🥰", word)},
				}}})
			}
			// TODO proto stuff
			doc.Occurrences = append(doc.Occurrences, &proto.MonikerOccurrence{
				MonikerId: word,
				Role:      role,
				Range: &proto.Range{
					Start: &proto.Position{Line: 0, Character: int32(i - len(word))},
					End:   &proto.Position{Line: 0, Character: int32(i)},
				},
			})
			word = ""
		} else {
			word = word + string(char)
		}
	}
	vals = append(vals, &proto.LsifValue{Value: &proto.LsifValue_Document{Document: &doc}})
	return proto.LsifValues{Values: vals}
}
