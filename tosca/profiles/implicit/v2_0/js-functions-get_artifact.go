// This file was auto-generated from a YAML file

package v2_0

func init() {
	Profile["/tosca/implicit/2.0/js/functions/get_artifact.js"] = `

// [TOSCA-Simple-Profile-YAML-v1.3] @ 4.8.1
// [TOSCA-Simple-Profile-YAML-v1.2] @ 4.8.1
// [TOSCA-Simple-Profile-YAML-v1.1] @ 4.8.1
// [TOSCA-Simple-Profile-YAML-v1.0] @ 4.8.1

const tosca = require('tosca.lib.utils');

exports.evaluate = function(entity, artifactName, location, remove) {
	if (arguments.length < 2)
		throw 'must have at least 2 arguments';
	let nodeTemplate = tosca.getModelableEntity.call(this, entity).properties;
	if (!nodeTemplate.artifacts || !(artifactName in nodeTemplate.artifacts))
		throw puccini.sprintf('artifact %q not found in %q', artifactName, nodeTemplate.name);
	let artifact = nodeTemplate.artifacts[artifactName];
	if (artifact.$artifact === undefined)
		return artifact.sourcePath;
	return artifact.$artifact;
};
`
}
