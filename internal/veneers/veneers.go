package veneers

import (
	"github.com/grafana/cog/internal/veneers/builder"
	"github.com/grafana/cog/internal/veneers/option"
)

func Common() *Rewriter {
	return NewRewrite(
		[]builder.RewriteRule{
			// We don't want these builders at all
			builder.Omit(builder.ByName("GridPos")),
			builder.Omit(builder.ByName("DataSourceRef")),
			builder.Omit(builder.ByName("LibraryPanelRef")),

			// No need for builders for structs generated from a disjunction
			builder.Omit(builder.StructGeneratedFromDisjunction()),

			// rearrange things a bit
			builder.MergeInto(
				builder.ByName("Panel"),
				"FieldConfig",
				"fieldConfig.defaults",

				[]string{
					// don't copy these over as they clash with a similarly named options from Panel
					"description", "links",

					// TODO: check if these are actually relevant
					"displayNameFromDS", "filterable", "path", "writeable",
				},
			),

			// remove builders that were previously merged into something else
			builder.Omit(builder.ByName("FieldConfig")),
			builder.Omit(builder.ByName("FieldConfigSource")),
		},

		[]option.RewriteRule{
			/********************************************
			 * Dashboards
			 ********************************************/

			// Let's make the dashboard constructor more friendly
			option.PromoteToConstructor(
				option.ByName("Dashboard", "title"),
			),

			// `Tooltip` looks better than `GraphTooltip`
			option.Rename(
				option.ByName("Dashboard", "graphTooltip"),
				"tooltip",
			),

			// `panels` refers to RowPanel only for now
			option.Rename(
				option.ByName("Dashboard", "panels"),
				"rows",
			),
			/*
				// TODO: finish implementing this rule
				option.ArrayToAppend(
					option.ByName("Dashboard", "rows"),
				),
			*/

			// Editable() + Readonly() instead of Editable(val bool)
			option.UnfoldBoolean(
				option.ByName("Dashboard", "editable"),
				option.BooleanUnfold{OptionTrue: "editable", OptionFalse: "readonly"},
			),

			// Refresh(string) instead of Refresh(struct StringOrBool)
			option.StructFieldsAsArguments(
				option.ByName("Dashboard", "refresh"),
				"ValString",
			),

			// We don't want these options at all
			option.Omit(option.ByName("Dashboard", "schemaVersion")),

			/********************************************
			 * Panels
			 ********************************************/

			option.Omit(option.ByName("Panel", "id")),            // generated by the backend
			option.Omit(option.ByName("Panel", "fieldConfig")),   // merged with another builder
			option.Omit(option.ByName("Panel", "options")),       // comes from a panel plugin
			option.Omit(option.ByName("Panel", "custom")),        // comes from a panel plugin
			option.Omit(option.ByName("Panel", "pluginVersion")), // TODO: check if it's relevant or not
			option.Omit(option.ByName("Panel", "repeatPanelId")), // TODO: check if it's relevant or not

			/********************************************
			 * Rows
			 ********************************************/

			// Let's make the row constructor more friendly
			option.PromoteToConstructor(
				option.ByName("RowPanel", "title"),
			),
		},
	)
}
