// This file was auto-generated from a YAML file

package v1_1

func init() {
	Profile["/tosca/simple/1.1/nodes.yaml"] = `
tosca_definitions_version: tosca_simple_yaml_1_1

imports:

- relationships.yaml
- interfaces.yaml

node_types:

  tosca.nodes.Root:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.1
    description: >-
      The TOSCA Root Node Type is the default type that all other TOSCA base Node Types derive from.
      This allows for all TOSCA nodes to have a consistent set of features for modeling and
      management (e.g., consistent definitions for requirements, capabilities and lifecycle
      interfaces).
    attributes:
      tosca_id:
        description: >-
          A unique identifier of the realized instance of a Node Template that derives from any
          TOSCA normative type.
        type: string
      tosca_name:
        description: >-
          This attribute reflects the name of the Node Template as defined in the TOSCA service
          template. This name is not unique to the realized instance model of corresponding deployed
          application as each template in the model can result in one or more instances (e.g.,
          scaled) when orchestrated to a provider environment.
        type: string
      state:
        description: >-
          The state of the node instance.
        type: string
        default: initial
    interfaces:
      Standard:
        type: tosca.interfaces.node.lifecycle.Standard
    capabilities:
      feature:
        type: tosca.capabilities.Node
    requirements:
    - dependency:
        capability: tosca.capabilities.Node
        node: tosca.nodes.Root
        relationship: tosca.relationships.DependsOn
        occurrences: [ 0, UNBOUNDED ]

  tosca.nodes.Compute:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.2
    description: >-
      The TOSCA Compute node represents one or more real or virtual processors of software
      applications or services along with other essential local resources. Collectively, the
      resources the compute node represents can logically be viewed as a (real or virtual) "server".
    derived_from: tosca.nodes.Root
    attributes:
      private_address:
        description: >-
          The primary private IP address assigned by the cloud provider that applications may use to
          access the Compute node.
        type: string
      public_address:
        description: >-
          The primary public IP address assigned by the cloud provider that applications may use to
          access the Compute node.
        type: string
      networks:
        description: >-
          The list of logical networks assigned to the compute host instance and information about
          them.
        type: map
        entry_schema:
          type: tosca.datatypes.network.NetworkInfo
      ports:
        description: >-
          The list of logical ports assigned to the compute host instance and information about
          them.
        type: map
        entry_schema:
          type: tosca.datatypes.network.PortInfo
    capabilities:
      host:
        type: tosca.capabilities.Container
        valid_source_types: [ tosca.nodes.SoftwareComponent ]
      endpoint:
        type: tosca.capabilities.Endpoint.Admin
      os:
         type: tosca.capabilities.OperatingSystem
      scalable:
         type: tosca.capabilities.Scalable
      binding:
         type: tosca.capabilities.network.Bindable
    requirements:
    - local_storage:
        capability: tosca.capabilities.Attachment
        node: tosca.nodes.Storage.BlockStorage
        relationship: tosca.relationships.AttachesTo
        occurrences: [ 0, UNBOUNDED ]

  tosca.nodes.SoftwareComponent:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.3
    description: >-
      The TOSCA SoftwareComponent node represents a generic software component that can be managed
      and run by a TOSCA Compute Node Type.
    derived_from: tosca.nodes.Root
    properties:
      component_version:
        description: >-
          The optional software component's version.
        type: version
        required: false
      admin_credential:
        description: >-
          The optional credential that can be used to authenticate to the software component.
        type: tosca.datatypes.Credential
        required: false
    requirements:
    - host:
        capability: tosca.capabilities.Container
        node: tosca.nodes.Compute
        relationship: tosca.relationships.HostedOn

  tosca.nodes.WebServer:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.4
    description: >-
      This TOSCA WebServer Node Type represents an abstract software component or service that is
      capable of hosting and providing management operations for one or more WebApplication nodes.
    derived_from: tosca.nodes.SoftwareComponent
    capabilities:
      data_endpoint:
        type: tosca.capabilities.Endpoint
      admin_endpoint:
        type: tosca.capabilities.Endpoint.Admin
      host:
        type: tosca.capabilities.Container
        valid_source_types: [ tosca.nodes.WebApplication ]

  tosca.nodes.WebApplication:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.5
    description: >-
      The TOSCA WebApplication node represents a software application that can be managed and run by
      a TOSCA WebServer node. Specific types of web applications such as Java, etc. could be derived
      from this type.
    derived_from: tosca.nodes.Root
    properties:
      context_root:
        description: >-
          The web application's context root which designates the application's URL path within the
          web server it is hosted on.
        type: string
        required: false
    capabilities:
      app_endpoint:
        type: tosca.capabilities.Endpoint
    requirements:
    - host:
        capability: tosca.capabilities.Container
        node: tosca.nodes.WebServer
        relationship: tosca.relationships.HostedOn

  tosca.nodes.DBMS:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.6
    description: >-
      The TOSCA DBMS node represents a typical relational, SQL Database Management System software
      component or service.
    derived_from: tosca.nodes.SoftwareComponent
    properties:
      root_password:
        description: >-
          The optional root password for the DBMS server.
        type: string
        required: false
      port:
        description: >-
          The DBMS server's port.
        type: integer
        required: false
    capabilities:
      host:
        type: tosca.capabilities.Container
        valid_source_types: [ tosca.nodes.Database ]

  tosca.nodes.Database:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.7
    description: >-
      The TOSCA Database node represents a logical database that can be managed and hosted by a
      TOSCA DBMS node.
    derived_from: tosca.nodes.Root
    properties:
      name:
        description: >-
          The logical database Name.
        type: string
      port:
        description: >-
          The port the database service will use to listen for incoming data and requests.
        type: integer
        required: false
      user:
        description: >-
          The special user account used for database administration.
        type: string
        required: false
      password:
        description: >-
          The password associated with the user account provided in the 'user' property.
        type: string
        required: false
    capabilities:
      database_endpoint:
        type: tosca.capabilities.Endpoint.Database
    requirements:
    - host:
        capability: tosca.capabilities.Container
        node: tosca.nodes.DBMS
        relationship: tosca.relationships.HostedOn

  tosca.nodes.Storage.ObjectStorage:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.8
    description: >-
      The TOSCA ObjectStorage node represents storage that provides the ability to store data as
      objects (or BLOBs of data) without consideration for the underlying filesystem or devices.
    derived_from: tosca.nodes.Root
    properties:
      name:
        description: >-
          The logical name of the object store (or container).
        type: string
      size:
        description: >-
          The requested initial storage size (default unit is in Gigabytes).
        type: scalar-unit.size
        constraints:
        - greater_or_equal: 0 GB
        required: false
      maxsize:
        description: >-
          The requested maximum storage size (default unit is in Gigabytes).
        type: scalar-unit.size
        constraints:
        - greater_or_equal: 0 GB
        required: false
    capabilities:
      storage_endpoint:
        type: tosca.capabilities.Endpoint

  tosca.nodes.Storage.BlockStorage:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.9
    description: >-
      The TOSCA BlockStorage node currently represents a server-local block storage device (i.e., not
      shared) offering evenly sized blocks of data from which raw storage volumes can be created. 
    derived_from: tosca.nodes.Root
    properties:
      size:
        description: >-
          The requested storage size (default unit is MB).
        type: scalar-unit.size
        constraints:
        - greater_or_equal: 1 MB
      volume_id:
        description: >-
          ID of an existing volume (that is in the accessible scope of the requesting application).
        type: string
        required: false
      snapshot_id:
        description: >-
          Some identifier that represents an existing snapshot that should be used when creating the
          block storage (volume).
        type: string
        required: false
    capabilities:
      attachment:
        type: tosca.capabilities.Attachment

  tosca.nodes.Container.Runtime:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.10
    description: >-
      The TOSCA Container Runtime node represents operating system-level virtualization technology
      used to run multiple application services on a single Compute host.
    derived_from: tosca.nodes.SoftwareComponent
    capabilities:
      host:
        type: tosca.capabilities.Container
      scalable:
        type: tosca.capabilities.Scalable

  tosca.nodes.Container.Application:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.11
    description: >-
      The TOSCA Container Application node represents an application that requires Container-level
      virtualization technology.
    derived_from: tosca.nodes.Root
    requirements:
    - host:
        capability: tosca.capabilities.Container
        node: tosca.nodes.Container.Runtime
        relationship: tosca.relationships.HostedOn
    - storage:
        capability: tosca.capabilities.Storage
    - network:
        capability: tosca.capabilities.Network

  tosca.nodes.LoadBalancer:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 5.9.12
    description: >-
      The TOSCA Load Balancer node represents logical function that be used in conjunction with a
      Floating Address to distribute an application's traffic (load) across a number of instances of
      the application (e.g., for a clustered or scaled application).
    derived_from: tosca.nodes.Root
    properties:
      algorithm:
        description: >-
          No description in spec.
        type: string
        required: false
        status: experimental
    capabilities:
      client:
        description: >-
          The Floating (IP) client's on the public network can connect to.
        type: tosca.capabilities.Endpoint.Public
        occurrences: [ 0, UNBOUNDED ]
    requirements:
    - application:
        capability: tosca.capabilities.Endpoint
        relationship: tosca.relationships.RoutesTo
        occurrences: [ 0, UNBOUNDED ]

  #
  # Network
  #

  tosca.nodes.network.Network:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 8.5.1
    description: >-
      The TOSCA Network node represents a simple, logical network service.
    derived_from: tosca.nodes.Root
    properties:
      ip_version:
        description: >-
          The IP version of the requested network.
        type: integer
        constraints:
        - valid_values: [ 4, 6 ]
        default: 4
        required: false
      cidr:
        description: >-
          The cidr block of the requested network.
        type: string
        required: false
      start_ip:
        description: >-
          The IP address to be used as the 1st one in a pool of addresses derived from the cidr
          block full IP range.
        type: string
        required: false
      end_ip:
        description: >-
          The IP address to be used as the last one in a pool of addresses derived from the cidr
          block full IP range.
        type: string
        required: false
      gateway_ip:
        description: >-
          The gateway IP address.
        type: string
        required: false
      network_name:
        description: >-
          An Identifier that represents an existing Network instance in the underlying cloud
          infrastructure - OR - be used as the name of the new created network.
        type: string
        required: false
      network_id:
        description: >-
          An Identifier that represents an existing Network instance in the underlying cloud
          infrastructure. This property is mutually exclusive with all other properties except
          network_name.
        type: string
        required: false
      segmentation_id:
        description: >-
          A segmentation identifier in the underlying cloud infrastructure (e.g., VLAN id, GRE
          tunnel id). If the segmentation_id is specified, the network_type or physical_network
          properties should be provided as well.
        type: string
        required: false
      network_type:
        description: >-
          Optionally, specifies the nature of the physical network in the underlying cloud
          infrastructure. Examples are flat, vlan, gre or vxlan. For flat and vlan types,
          physical_network should be provided too.
        type: string
        required: false
      physical_network:
        description: >-
          Optionally, identifies the physical network on top of which the network is implemented,
          e.g. physnet1. This property is required if network_type is flat or vlan.
        type: string
        required: false
      dhcp_enabled:
        description: >-
          Indicates the TOSCA container to create a virtual network instance with or without a DHCP
          service.
        type: boolean
        default: true
        required: false
    attributes:
      segmentation_id:
        description: >-
          The actual segmentation_id that is been assigned to the network by the underlying cloud
          infrastructure. 
        type: string
    capabilities:
      link:
        type: tosca.capabilities.network.Linkable

  tosca.nodes.network.Port:
    metadata:
      tosca.normative: 'true'
      specification.citation: '[TOSCA-Simple-Profile-YAML-v1.1]'
      specification.location: 8.5.2
    description: >-
      The TOSCA Port node represents a logical entity that associates between Compute and Network
      normative types.

      The Port node type effectively represents a single virtual NIC on the Compute node instance.
    derived_from: tosca.nodes.Root
    properties:
      ip_address:
        description: >-
          Allow the user to set a fixed IP address. Note that this address is a request to the
          provider which they will attempt to fulfill but may not be able to dependent on the
          network the port is associated with.
        type: string
        required: false
      order:
        description: >-
          The order of the NIC on the compute instance (e.g. eth2). Note: when binding more than one
          port to a single compute (aka multi vNICs) and ordering is desired, it is *mandatory* that
          all ports will be set with an order value and. The order values must represent a positive,
          arithmetic progression that starts with 0 (e.g. 0, 1, 2, ..., n).
        type: integer
        constraints:
        - greater_or_equal: 0
        default: 0
        required: false
      is_default:
        description: >-
          Set is_default=true to apply a default gateway route on the running compute instance to
          the associated network gateway. Only one port that is associated to single compute node
          can set as default=true.
        type: boolean
        default: false
        required: false
      ip_range_start:
        description: >-
          Defines the starting IP of a range to be allocated for the compute instances that are
          associated by this Port. Without setting this property the IP allocation is done from the
          entire CIDR block of the network.
        type: string
        required: false
      ip_range_end:
        description: >-
          Defines the ending IP of a range to be allocated for the compute instances that are
          associated by this Port. Without setting this property the IP allocation is done from the
          entire CIDR block of the network.
        type: string
        required: false
    attributes:
      ip_address:
        description: >-
          The IP address would be assigned to the associated compute instance.
        type: string
    requirements:
    - link:
        capability: tosca.capabilities.network.Linkable
        relationship: tosca.relationships.network.LinksTo
    - binding:
        capability: tosca.capabilities.network.Bindable
        relationship: tosca.relationships.network.BindsTo
`
}
