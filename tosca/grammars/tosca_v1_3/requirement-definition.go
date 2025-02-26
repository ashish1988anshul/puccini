package tosca_v1_3

import (
	"github.com/tliron/kutil/ard"
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/grammars/tosca_v2_0"
)

//
// RequirementDefinition
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.7.3
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.7.3
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.6.3
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.6.2
//

// tosca.Reader signature
func ReadRequirementDefinition(context *tosca.Context) tosca.EntityPtr {
	context.SetReadTag("CountRange", "occurrences,RangeEntity")

	self := tosca_v2_0.ReadRequirementDefinition(context).(*tosca_v2_0.RequirementDefinition)
	self.DefaultCountRange = ard.List{1, 1}
	return self
}
