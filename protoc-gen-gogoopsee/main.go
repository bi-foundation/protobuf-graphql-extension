package main

import (
	_ "github.com/bi-foundation/protobuf-graphql-extension/plugin/bitflags"
	_ "github.com/bi-foundation/protobuf-graphql-extension/plugin/graphql"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	vanity.ForEachFile(files, vanity.TurnOnTestGenAll)
	vanity.ForEachFile(files, vanity.TurnOnEqualAll)
	vanity.ForEachFile(files, vanity.TurnOnPopulateAll)
	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)
	resp := command.Generate(req)
	command.Write(resp)
}
