package tosca_v2_0

import (
	"github.com/tliron/kutil/ard"
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/normal"
	profile "github.com/tliron/puccini/tosca/profiles/implicit/v2_0"
	"github.com/tliron/yamlkeys"
)

const constraintPathPrefix = "/tosca/implicit/2.0/js/constraints/"

// Built-in constraint functions
var ConstraintClauseScriptlets = map[string]string{
	tosca.ConstraintScriptletPrefix + "equal":            profile.Profile[constraintPathPrefix+"equal.js"],
	tosca.ConstraintScriptletPrefix + "greater_than":     profile.Profile[constraintPathPrefix+"greater_than.js"],
	tosca.ConstraintScriptletPrefix + "greater_or_equal": profile.Profile[constraintPathPrefix+"greater_or_equal.js"],
	tosca.ConstraintScriptletPrefix + "less_than":        profile.Profile[constraintPathPrefix+"less_than.js"],
	tosca.ConstraintScriptletPrefix + "less_or_equal":    profile.Profile[constraintPathPrefix+"less_or_equal.js"],
	tosca.ConstraintScriptletPrefix + "in_range":         profile.Profile[constraintPathPrefix+"in_range.js"],
	tosca.ConstraintScriptletPrefix + "valid_values":     profile.Profile[constraintPathPrefix+"valid_values.js"],
	tosca.ConstraintScriptletPrefix + "length":           profile.Profile[constraintPathPrefix+"length.js"],
	tosca.ConstraintScriptletPrefix + "min_length":       profile.Profile[constraintPathPrefix+"min_length.js"],
	tosca.ConstraintScriptletPrefix + "max_length":       profile.Profile[constraintPathPrefix+"max_length.js"],
	tosca.ConstraintScriptletPrefix + "pattern":          profile.Profile[constraintPathPrefix+"pattern.js"],
	tosca.ConstraintScriptletPrefix + "schema":           profile.Profile[constraintPathPrefix+"schema.js"], // introduced in TOSCA 1.3
}

var ConstraintClauseNativeArgumentIndexes = map[string][]int{
	tosca.ConstraintScriptletPrefix + "equal":            {0},
	tosca.ConstraintScriptletPrefix + "greater_than":     {0},
	tosca.ConstraintScriptletPrefix + "greater_or_equal": {0},
	tosca.ConstraintScriptletPrefix + "less_than":        {0},
	tosca.ConstraintScriptletPrefix + "less_or_equal":    {0},
	tosca.ConstraintScriptletPrefix + "in_range":         {0, 1},
	tosca.ConstraintScriptletPrefix + "valid_values":     {-1}, // -1 means all
}

//
// ConstraintClause
//
// [TOSCA-v2.0] @ ?
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.6.3
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.6.3
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.5.2
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.5.2
//

type ConstraintClause struct {
	*Entity `name:"constraint clause"`

	Operator              string
	Arguments             ard.List
	NativeArgumentIndexes []int
	DataType              *DataType            `traverse:"ignore" json:"-" yaml:"-"` // TODO: unncessary, this entity should never be traversed
	Definition            *AttributeDefinition `traverse:"ignore" json:"-" yaml:"-"`
}

func NewConstraintClause(context *tosca.Context) *ConstraintClause {
	return &ConstraintClause{Entity: NewEntity(context)}
}

// tosca.Reader signature
func ReadConstraintClause(context *tosca.Context) tosca.EntityPtr {
	self := NewConstraintClause(context)

	if context.ValidateType(ard.TypeMap) {
		map_ := context.Data.(ard.Map)
		if len(map_) != 1 {
			context.ReportValueMalformed("constraint clause", "map size not 1")
			return self
		}

		for key, value := range map_ {
			operator := yamlkeys.KeyString(key)

			scriptletName := tosca.ConstraintScriptletPrefix + operator
			scriptlet, ok := context.ScriptletNamespace.Lookup(scriptletName)
			if !ok {
				context.Clone(operator).ReportValueMalformed("constraint clause", "unsupported operator")
				return self
			}

			self.Operator = operator

			if list, ok := value.(ard.List); ok {
				self.Arguments = list
			} else {
				self.Arguments = ard.List{value}
			}

			self.NativeArgumentIndexes = scriptlet.NativeArgumentIndexes

			// We have only one key
			break
		}
	}

	return self
}

func (self *ConstraintClause) ToFunctionCall(context *tosca.Context, strict bool) *tosca.FunctionCall {
	// Special case: "in_range" for a "range" accepts a range (two integers) rather than two ranges
	var isRangeInRange bool
	if self.Operator == "in_range" {
		if rangeType, ok := self.Context.Namespace.Lookup("range"); ok {
			if isRangeInRange = self.Context.Hierarchy.IsCompatible(rangeType, self.DataType); isRangeInRange {
				range_ := ReadRange(context.Clone(self.Arguments)).(*Range)
				return context.NewFunctionCall(tosca.ConstraintScriptletPrefix+self.Operator, []any{
					ReadValue(context.ListChild(0, range_.Lower)).(*Value),
					ReadValue(context.ListChild(1, range_.Upper)).(*Value),
				})
			}
		}
	}

	arguments := make([]any, len(self.Arguments))
	for index, argument := range self.Arguments {
		if self.IsNativeArgument(index) {
			if _, ok := argument.(*Value); !ok {
				if self.DataType != nil {
					value := ReadValue(self.Context.ListChild(index, argument)).(*Value)
					value.Render(self.DataType, self.Definition, true, false) // bare
					argument = value
				} else if strict {
					panic("no data type for native argument")
				}
			}
		}

		arguments[index] = argument
	}

	return context.NewFunctionCall(tosca.ConstraintScriptletPrefix+self.Operator, arguments)
}

func (self *ConstraintClause) IsNativeArgument(index int) bool {
	for _, i := range self.NativeArgumentIndexes {
		if (i == -1) || (i == index) {
			return true
		}
	}
	return false
}

//
// ConstraintClauses
//

type ConstraintClauses []*ConstraintClause

func (self ConstraintClauses) Append(constraints ConstraintClauses) ConstraintClauses {
	length := len(self)
	if length > 0 {
		r := make(ConstraintClauses, length)
		copy(r, self)
		return append(r, constraints...)
	} else {
		r := make(ConstraintClauses, len(constraints))
		copy(r, constraints)
		return r
	}
}

func (self ConstraintClauses) Render(dataType *DataType, definition *AttributeDefinition) {
	for _, constraint := range self {
		constraint.DataType = dataType
		constraint.Definition = definition
	}
}

/*
func (self ConstraintClauses) Validate(dataType *DataType) {
	for _, constraintClause := range self {
		if (constraintClause.DataType != nil) && (constraintClause.DataType != dataType) {
			panic(fmt.Sprintf("constraint clause for data type %q cannot be used with data type %q", constraintClause.DataType.Name, dataType.Name))
		}
	}
}
*/

func (self ConstraintClauses) Normalize(context *tosca.Context) normal.FunctionCalls {
	var normalFunctionCalls normal.FunctionCalls
	for _, constraintClause := range self {
		functionCall := constraintClause.ToFunctionCall(context, false)
		NormalizeFunctionCallArguments(functionCall, context)
		normalFunctionCalls = append(normalFunctionCalls, normal.NewFunctionCall(functionCall))
	}
	return normalFunctionCalls
}

func (self ConstraintClauses) NormalizeConstrainable(context *tosca.Context, normalConstrainable normal.Constrainable) {
	for _, constraintClause := range self {
		functionCall := constraintClause.ToFunctionCall(context, true)
		NormalizeFunctionCallArguments(functionCall, context)
		normalConstrainable.AddConstraint(functionCall)
	}
}

func (self ConstraintClauses) NormalizeListEntries(context *tosca.Context, normalList *normal.List) {
	for _, constraintClause := range self {
		functionCall := constraintClause.ToFunctionCall(context, true)
		NormalizeFunctionCallArguments(functionCall, context)
		normalList.AddEntryConstraint(functionCall)
	}
}

func (self ConstraintClauses) NormalizeMapKeys(context *tosca.Context, normalMap *normal.Map) {
	for _, constraintClause := range self {
		functionCall := constraintClause.ToFunctionCall(context, true)
		NormalizeFunctionCallArguments(functionCall, context)
		normalMap.AddKeyConstraint(functionCall)
	}
}

func (self ConstraintClauses) NormalizeMapValues(context *tosca.Context, normalMap *normal.Map) {
	for _, constraintClause := range self {
		functionCall := constraintClause.ToFunctionCall(context, true)
		NormalizeFunctionCallArguments(functionCall, context)
		normalMap.AddValueConstraint(functionCall)
	}
}
