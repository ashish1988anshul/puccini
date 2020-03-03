package hot

import (
	"github.com/tliron/puccini/ard"
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/normal"
	"github.com/tliron/yamlkeys"
)

//
// Value
//

type Value struct {
	*Entity `name:"value"`
	Name    string

	Constraints Constraints
	Description *string
}

func NewValue(context *tosca.Context) *Value {
	return &Value{
		Entity: NewEntity(context),
		Name:   context.Name,
	}
}

// tosca.Reader signature
func ReadValue(context *tosca.Context) interface{} {
	ToFunctionCalls(context)
	return NewValue(context)
}

// tosca.Mappable interface
func (self *Value) GetKey() string {
	return self.Name
}

func (self *Value) Normalize() normal.Constrainable {
	var normalConstrainable normal.Constrainable

	switch data := self.Context.Data.(type) {
	case ard.List:
		normalList := normal.NewList(len(data))
		for index, value := range data {
			normalList.Set(index, NewValue(self.Context.ListChild(index, value)).Normalize())
		}
		normalConstrainable = normalList

	case ard.Map:
		normalMap := normal.NewMap()
		for key, value := range data {
			if _, ok := key.(string); !ok {
				// HOT does not support complex keys
				self.Context.MapChild(key, yamlkeys.KeyData(key)).ReportValueWrongType("string")
			}
			name := yamlkeys.KeyString(key)
			normalMap.Put(name, NewValue(self.Context.MapChild(name, value)).Normalize())
		}
		normalConstrainable = normalMap

	case *tosca.FunctionCall:
		NormalizeFunctionCallArguments(data, self.Context)
		normalConstrainable = normal.NewFunctionCall(data)

	default:
		normalConstrainable = normal.NewValue(data)
	}

	self.Constraints.Normalize(self.Context, normalConstrainable)

	if self.Description != nil {
		normalConstrainable.SetDescription(*self.Description)
	}

	return normalConstrainable
}

//
// Values
//

type Values map[string]*Value

func (self Values) Normalize(normalConstrainables normal.Constrainables) {
	for key, value := range self {
		normalConstrainables[key] = value.Normalize()
	}
}
