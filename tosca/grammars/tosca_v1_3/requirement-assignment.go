package tosca_v1_3

import (
	"github.com/tliron/kutil/ard"
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/grammars/tosca_v2_0"
)

//
// RequirementAssignment
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.8.2
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.8.2
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.7.2
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.7.2
//

// tosca.Reader signature
func ReadRequirementAssignment(context *tosca.Context) tosca.EntityPtr {
	context.SetReadTag("Count", "")
	context.SetReadTag("Directives", "")
	context.SetReadTag("Optional", "")
	//context.SetReadTag("Allocation", "")

	self := tosca_v2_0.NewRequirementAssignment(context)

	if context.Is(ard.TypeMap) {
		// Long notation
		context.ValidateUnsupportedFields(append(context.ReadFields(self), "occurrences"))

		if occurrences := ard.NewNode(self.Context.Data).Get("occurrences"); occurrences.Value != nil {
			occurrences_ := tosca_v2_0.ReadRange(context.FieldChild("occurrences", occurrences.Value)).(*tosca_v2_0.Range)
			lower := int64(occurrences_.Lower)
			self.Count = &lower
			// TODO: have no idea what to do with max bound in "occurrences" keyname
		}
	} else if context.ValidateType(ard.TypeMap, ard.TypeString) {
		// Short notation
		self.TargetNodeTemplateNameOrTypeName = context.FieldChild("node", context.Data).ReadString()
	}

	return self
}
