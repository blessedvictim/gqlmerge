package command

import (
	"fmt"
)

func Usage() string {
	flags := fmt.Sprintf(`
Flags:

	%s
`, flagIdentMsg)

	return helpMsg + flags
}

const helpMsg = `👋 'gqlmerge' is the tool to merge & stitch GraphQL files and generate a GraphQL schema
Author : Woonki Moon <woonki.moon@gmail.com>

Usage:	gqlmerge [FLAG ...] [PATH ...] [OUTPUT]

e.g.

	gqlmerge ./schema schema.graphql

Options:

	-v	: check the version
	-h	: help
`

const flagIdentMsg = `-ident	: defines the padding in the generated GraphQL scheme.

	It follows the next pattern: ident={n}{i},
		* n - amount of idents
		* i - ident ("t" for tabs and "s" for spaces)
	If "n" is not stated 1 will be used, 
	so "--ident=1t" is equal to "--ident=t"`
