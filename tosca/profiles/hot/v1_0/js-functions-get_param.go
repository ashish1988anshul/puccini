// This file was auto-generated from a YAML file

package v1_0

func init() {
	Profile["/hot/1.0/js/functions/get_param.js"] = `

// [https://docs.openstack.org/heat/wallaby/template_guide/hot_spec.html#get_param]

const tosca = require('tosca.lib.utils');

exports.evaluate = function(input) {
	if (arguments.length !== 1)
		throw 'must have 1 argument';
	if (!tosca.isTosca(clout))
		throw 'Clout is not TOSCA';
	let inputs = clout.properties.tosca.inputs;
	if (!(input in inputs))
		throw puccini.sprintf('parameter %q not found', input);
	let r = inputs[input];
	r = clout.coerce(r);
	return r;
};
`
}
