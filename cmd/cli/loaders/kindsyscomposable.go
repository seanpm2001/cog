package loaders

import (
	"fmt"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue"
	"github.com/grafana/cog/internal/ast"
	"github.com/grafana/cog/internal/simplecue"
)

func kindsysComposableLoader(input CueInput) (ast.Schemas, error) {
	libraries, err := simplecue.ParseImports(input.CueImports)
	if err != nil {
		return nil, err
	}

	schemaRootValue, err := parseCueEntrypoint(input.Entrypoint, libraries, "grafanaplugin")
	if err != nil {
		return nil, err
	}

	variant, err := schemaVariant(schemaRootValue)
	if err != nil {
		return nil, err
	}

	kindIdentifier, err := inferComposableKindIdentifier(schemaRootValue)
	if err != nil {
		return nil, err
	}

	schema, err := simplecue.GenerateAST(schemaFromThemaLineage(schemaRootValue), simplecue.Config{
		Package:              filepath.Base(input.Entrypoint), // TODO: extract from somewhere else?
		ForceVariantEnvelope: variant == ast.SchemaVariantDataQuery,
		SchemaMetadata: ast.SchemaMeta{
			Kind:       ast.SchemaKindComposable,
			Variant:    variant,
			Identifier: kindIdentifier,
		},
		Libraries: libraries,
	})
	if err != nil {
		return nil, err
	}

	return input.filterSchema(schema)
}

func schemaVariant(kindRoot cue.Value) (ast.SchemaVariant, error) {
	schemaInterface, err := kindRoot.LookupPath(cue.ParsePath("schemaInterface")).String()
	if err != nil {
		return "", err
	}

	switch schemaInterface {
	case "PanelCfg":
		return ast.SchemaVariantPanel, nil
	case "DataQuery":
		return ast.SchemaVariantDataQuery, nil
	default:
		return "", fmt.Errorf("unknown schema variant '%s'", schemaInterface)
	}
}

// TODO: the schema should explicitly tell us the "plugin ID"/panel ID/dataquery type/...
func inferComposableKindIdentifier(kindRoot cue.Value) (string, error) {
	schemaInterface, err := kindRoot.LookupPath(cue.ParsePath("schemaInterface")).String()
	if err != nil {
		return "", err
	}

	kindName, err := kindRoot.LookupPath(cue.ParsePath("name")).String()
	if err != nil {
		return "", err
	}

	return strings.ToLower(strings.TrimSuffix(kindName, schemaInterface)), nil
}
