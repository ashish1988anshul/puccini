package cloudify_v1_3

import (
	"github.com/tliron/kutil/ard"
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/normal"
	profile "github.com/tliron/puccini/tosca/profiles/cloudify/v5_0_5"
	"github.com/tliron/yamlkeys"
)

//
// Built-in functions
//
// [https://docs.cloudify.co/5.0.5/developer/blueprints/spec-intrinsic-functions/]
//

const functionPathPrefix = "/cloudify/5.0.5/js/functions/"

var FunctionScriptlets = map[string]string{
	tosca.FunctionScriptletPrefix + "concat":         profile.Profile[functionPathPrefix+"get_secret.js"],
	tosca.FunctionScriptletPrefix + "get_attribute":  profile.Profile[functionPathPrefix+"get_attribute.js"],
	tosca.FunctionScriptletPrefix + "get_capability": profile.Profile[functionPathPrefix+"get_capability.js"],
	tosca.FunctionScriptletPrefix + "get_input":      profile.Profile[functionPathPrefix+"get_input.js"],
	tosca.FunctionScriptletPrefix + "get_property":   profile.Profile[functionPathPrefix+"get_property.js"],
	tosca.FunctionScriptletPrefix + "get_secret":     profile.Profile[functionPathPrefix+"get_secret.js"],
}

func ParseFunctionCalls(context *tosca.Context) bool {
	if _, ok := context.Data.(*tosca.FunctionCall); ok {
		// It's already a function call
		return true
	}

	map_, ok := context.Data.(ard.Map)
	if !ok || len(map_) != 1 {
		return false
	}

	for key, data := range map_ {
		scriptletName := tosca.FunctionScriptletPrefix + yamlkeys.KeyString(key)
		_, ok := context.ScriptletNamespace.Lookup(scriptletName)
		if !ok {
			// Not a function call, despite having the right data structure
			return false
		}

		// Some functions accept a list of arguments, some just one argument
		originalArguments, ok := data.(ard.List)
		if !ok {
			originalArguments = ard.List{data}
		}

		// Arguments may be function calls
		arguments := make(ard.List, len(originalArguments))
		for index, argument := range originalArguments {
			argumentContext := context.Clone(argument)
			ParseFunctionCalls(argumentContext)
			arguments[index] = argumentContext.Data
		}

		context.Data = context.NewFunctionCall(scriptletName, arguments)

		// We have only one key
		return true
	}

	return false
}

func ToFunctionCalls(context *tosca.Context) {
	if !ParseFunctionCalls(context) {
		if list, ok := context.Data.(ard.List); ok {
			for index, value := range list {
				childContext := context.ListChild(index, value)
				ToFunctionCalls(childContext)
				list[index] = childContext.Data
			}
		} else if map_, ok := context.Data.(ard.Map); ok {
			for key, value := range map_ {
				childContext := context.MapChild(key, value)
				ToFunctionCalls(childContext)
				map_[key] = childContext.Data
			}
		}
	}
}

func NormalizeFunctionCallArguments(functionCall *tosca.FunctionCall, context *tosca.Context) {
	for index, argument := range functionCall.Arguments {
		// Because the same constraint instance may be shared among more than one value, this
		// func might be called more than once on the same arguments, so we must make sure not
		// to normalize more than once
		if _, ok := argument.(normal.Constrainable); !ok {
			if value, ok := argument.(*Value); ok {
				functionCall.Arguments[index] = value.Normalize()
			} else {
				// Note: this literal value will not have a $type field
				functionCall.Arguments[index] = NewValue(context.ListChild(index, argument)).Normalize()
			}
		}
	}
}
