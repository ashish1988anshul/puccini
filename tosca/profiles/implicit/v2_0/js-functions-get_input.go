// This file was auto-generated from a YAML file

package v2_0

func init() {
	Profile["/tosca/implicit/2.0/js/functions/get_input.js"] = `

// [TOSCA-Simple-Profile-YAML-v1.3] @ 4.4.1
// [TOSCA-Simple-Profile-YAML-v1.2] @ 4.4.1
// [TOSCA-Simple-Profile-YAML-v1.1] @ 4.4.1
// [TOSCA-Simple-Profile-YAML-v1.0] @ 4.4.1

const tosca = require('tosca.lib.utils');

exports.evaluate = function(input) {
	if (arguments.length !== 1)
		throw 'must have 1 argument';
	if (!tosca.isTosca(clout))
		throw 'Clout is not TOSCA';
	let inputs = clout.properties.tosca.inputs;
	if (!(input in inputs))
		throw puccini.sprintf('input %q not found', input);
	let r = inputs[input];
	r = clout.coerce(r);
	return r;
};
`
}
