package js

import (
	"errors"
	"fmt"

	"github.com/tliron/puccini/ard"
	cloutpkg "github.com/tliron/puccini/clout"
)

func GetScriptletsMetadata(clout *cloutpkg.Clout) (ard.StringMap, error) {
	// TODO: check that version=1.0
	if scriptlets, ok := ard.NewNode(clout.Metadata).Get("puccini").Get("scriptlets").StringMap(false); ok {
		return scriptlets, nil
	} else {
		return nil, errors.New("no \"puccini.scriptlets\" metadata in Clout")
	}
}

func GetScriptletsMetadataSection(name string, clout *cloutpkg.Clout) (ard.Value, error) {
	segments, final, err := parseScriptletName(name)
	if err != nil {
		return nil, err
	}

	metadata, err := GetScriptletsMetadata(clout)
	if err != nil {
		return nil, err
	}

	m := metadata
	for _, s := range segments {
		o := m[s]
		var ok bool
		if m, ok = o.(ard.StringMap); !ok {
			return nil, fmt.Errorf("scriptlet metadata not found: %s", name)
		}
	}

	section, ok := m[final]
	if !ok {
		return nil, fmt.Errorf("scriptlet metadata not found: %s", name)
	}

	return section, nil
}
