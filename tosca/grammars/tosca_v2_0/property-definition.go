package tosca_v2_0

import (
	"github.com/tliron/puccini/tosca"
)

//
// PropertyDefinition
//
// [TOSCA-v2.0] @ ?
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.6.10
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.6.9
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.5.8
// [TOSCA-Simple-Profile-YAML-v1.0] @ 3.5.8
//

type PropertyDefinition struct {
	*AttributeDefinition `name:"property definition"`

	Required          *bool             `read:"required"`
	ConstraintClauses ConstraintClauses `read:"constraints,[]ConstraintClause" traverse:"ignore"`
}

func NewPropertyDefinition(context *tosca.Context) *PropertyDefinition {
	return &PropertyDefinition{AttributeDefinition: NewAttributeDefinition(context)}
}

// tosca.Reader signature
func ReadPropertyDefinition(context *tosca.Context) tosca.EntityPtr {
	self := NewPropertyDefinition(context)
	var ignore []string
	if context.HasQuirk(tosca.QuirkAnnotationsIgnore) {
		ignore = append(ignore, "annotations")
	}
	context.ValidateUnsupportedFields(append(context.ReadFields(self), ignore...))
	return self
}

func (self *PropertyDefinition) Inherit(parentDefinition *PropertyDefinition) {
	logInherit.Debugf("property definition: %s", self.Name)

	if self.Required != nil {
		if parentDefinition.IsRequired() && !*self.Required {
			self.Context.FieldChild("required", *self.Required).ReportRefinement(parentDefinition.IsRequired())
		}
	}

	self.AttributeDefinition.Inherit(parentDefinition.AttributeDefinition)

	if (self.Required == nil) && (parentDefinition.Required != nil) {
		self.Required = parentDefinition.Required
	}
	if parentDefinition.ConstraintClauses != nil {
		self.ConstraintClauses = parentDefinition.ConstraintClauses.Append(self.ConstraintClauses)
	}
}

// tosca.Renderable interface
func (self *PropertyDefinition) Render() {
	self.renderOnce.Do(self.render)
}

func (self *PropertyDefinition) render() {
	logRender.Debugf("property definition: %s", self.Name)

	self.doRender()
	self.ConstraintClauses.Render(self.DataType, self.AttributeDefinition)

	if (self.Default != nil) && (self.DataType != nil) {
		// The "default" value must be a valid value of the type
		self.Default.RenderProperty(self.DataType, self)
	}
}

func (self *PropertyDefinition) IsRequired() bool {
	// defaults to true
	return (self.Required == nil) || *self.Required
}

//
// PropertyDefinitions
//

type PropertyDefinitions map[string]*PropertyDefinition

func (self PropertyDefinitions) Inherit(parentDefinitions PropertyDefinitions) {
	for name, definition := range parentDefinitions {
		if _, ok := self[name]; !ok {
			self[name] = definition
		}
	}

	for name, definition := range self {
		if parentDefinition, ok := parentDefinitions[name]; ok {
			if definition != parentDefinition {
				definition.Inherit(parentDefinition)
			}
		}
	}
}
