
var tosca = {};

tosca.isTosca = function(o, kind) {
	if (o.metadata === undefined)
		return false;
	o = o.metadata['puccini-tosca'];
	if (o === undefined)
		return false;
	if (o.version !== '1.0')
		return false;
	if (kind !== undefined)
		return kind === o.kind;
	return true;
};

tosca.isNodeTemplate = function(vertex, typeName) {
	if (tosca.isTosca(vertex, 'NodeTemplate')) {
		if (typeName !== undefined)
			return typeName in vertex.properties.types;
		return true;
	}
	return false;
};

tosca.getNodeTemplate = function(entity) {
	var vertex;
	switch (entity) {
	case 'SELF':
		vertex = site;
		break;
	case 'SOURCE':
		vertex = source;
		break;
	case 'TARGET':
		vertex = target;
		break;
	case 'HOST':
		vertex = tosca.getHost(site);
		break;
	default:
		for (var vertexId in clout.vertexes) {
			var vertex = clout.vertexes[vertexId];
			if (tosca.isNodeTemplate(vertex) && (vertex.properties.name === entity))
				return vertex.properties;
		}
		throw puccini.sprintf('node template "%s" not found', entity);
	}
	if (!tosca.isNodeTemplate(vertex))
		throw puccini.sprintf('node template "%s" not found', entity);
	return vertex.properties;
};

tosca.getHost = function(vertex) {
	for (var e = 0; e < vertex.edgesOut.length; e++) {
		var edge = vertex.edgesOut[e];
		if (tosca.isTosca(edge, 'Relationship')) {
			for (var typeName in edge.properties.types) {
				var type = edge.properties.types[typeName];
				if (type.metadata.role === 'host')
					return edge.target;
			}
		}
	}
	return {};
};

tosca.getComparable = function(v) {
	if ((v === undefined) || (v === null))
		return null;
	var c = v.$number;
	if (c !== undefined)
		return c;
	c = v.$string;
	if (c !== undefined)
		return c;
	return v;
};

tosca.compare = function(v1, v2) {
	var c = v1.$comparer;
	if (c === undefined)
		c = v2.$comparer;
	if (c !== undefined)
		return clout.call(c, 'compare', [v1, v2]);
	v1 = tosca.getComparable(v1);
	v2 = tosca.getComparable(v2);
	if (v1 == v2)
		return 0;
	else if (v1 < v2)
		return -1;
	else
		return 1;
}