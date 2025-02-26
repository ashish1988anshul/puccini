package grammars

import (
	"time"

	"github.com/tliron/kutil/ard"
	urlpkg "github.com/tliron/kutil/url"
	"github.com/tliron/puccini/tosca"
)

func Detect(context *tosca.Context) bool {
	if context.Grammar == nil {
		var errorContext *tosca.Context
		if context.Grammar, errorContext = GetGrammar(context); errorContext != nil {
			errorContext.ReportKeynameUnsupportedValue()
		}
	}
	return context.Grammar != nil
}

func GetGrammar(context *tosca.Context) (*tosca.Grammar, *tosca.Context) {
	if versionContext, version := DetectVersion(context); version != nil {
		if grammars, ok := Grammars[versionContext.Name]; ok {
			if grammar, ok := grammars[*version]; ok {
				return grammar, nil
			} else {
				return nil, versionContext
			}
		} else {
			return nil, versionContext
		}
	}
	return nil, nil
}

func CompatibleGrammars(context1 *tosca.Context, context2 *tosca.Context) bool {
	grammar1, _ := GetGrammar(context1)
	grammar2, _ := GetGrammar(context2)
	return grammar1 == grammar2
}

func DetectVersion(context *tosca.Context) (*tosca.Context, *string) {
	var versionContext *tosca.Context
	var ok bool

	for keyword := range Grammars {
		if versionContext, ok = context.GetFieldChild(keyword); ok {
			if keyword == "heat_template_version" {
				// Hack to allow HOT to use YAML !!timestamp values

				if versionContext.Is(ard.TypeString) {
					return versionContext, versionContext.ReadString()
				}

				switch data := versionContext.Data.(type) {
				case time.Time:
					versionContext.Data = data.Format("2006-01-02")
					return versionContext, versionContext.ReadString()
				}

				versionContext.ReportValueWrongType(ard.TypeString, ard.TypeTimestamp)
			} else {
				if versionContext.ValidateType(ard.TypeString) {
					return versionContext, versionContext.ReadString()
				}
			}
		}
	}

	return nil, nil
}

func GetImplicitImportSpec(context *tosca.Context) (*tosca.ImportSpec, bool) {
	if versionContext, version := DetectVersion(context); version != nil {
		if paths, ok := ImplicitProfilePaths[versionContext.Name]; ok {
			if path, ok := paths[*version]; ok {
				if url, err := urlpkg.NewValidInternalURL(path, nil); err == nil {
					return &tosca.ImportSpec{url, nil, true}, true
				} else {
					context.ReportError(err)
				}
			}
		}
	}

	return nil, false
}
