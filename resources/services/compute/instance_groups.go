package compute

import (
	"context"
	"strings"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func InstanceGroups() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_instance_groups",
		Description: "Represents an Instance Group resource",
		Resolver:    fetchComputeInstanceGroups,
		Multiplex:   client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "creation_timestamp",
				Description: "The creation timestamp for this instance group in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "fingerprint",
				Description: "The fingerprint of the named ports",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "A unique identifier for this instance group, generated by the server",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "kind",
				Description: "The resource type, which is always compute#instanceGroup for instance groups",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the instance group",
				Type:        schema.TypeString,
			},
			{
				Name:        "named_ports",
				Description: "Assigns a name to a port number",
				Type:        schema.TypeJSON,
				Resolver:    resolveInstanceGroupsNamedPorts,
			},
			{
				Name:        "network",
				Description: "The URL of the network to which all instances in the instance group belong",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "The URL of the region where the instance group is located (for regional resources)",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "The URL for this instance group",
				Type:        schema.TypeString,
			},
			{
				Name:        "size",
				Description: "The total number of instances in the instance group",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "subnetwork",
				Description: "The URL of the subnetwork to which all instances in the instance group belong",
				Type:        schema.TypeString,
			},
			{
				Name:        "zone",
				Description: "The URL of the zone where the instance group is located (for zonal resources)",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "gcp_compute_instance_group_instances",
				Resolver: fetchComputeInstanceGroupInstances,
				Columns: []schema.Column{
					{
						Name:        "instance_group_cq_id",
						Description: "Unique CloudQuery ID of gcp_compute_instance_groups table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance",
						Description: "The URL of the instance",
						Type:        schema.TypeString,
					},
					{
						Name:        "named_ports",
						Description: "The named ports that belong to this instance group",
						Type:        schema.TypeJSON,
						Resolver:    resolveInstanceGroupInstancesNamedPorts,
					},
					{
						Name:        "status",
						Description: "\"DEPROVISIONING\" - The Nanny is halted and we are performing tear down tasks like network deprogramming, releasing quota, IP, tearing down disks etc   \"PROVISIONING\" - Resources are being allocated for the instance   \"REPAIRING\" - The instance is in repair   \"RUNNING\" - The instance is running   \"STAGING\" - All required resources have been allocated and the instance is being started   \"STOPPED\" - The instance has stopped successfully   \"STOPPING\" - The instance is currently stopping (either being deleted or killed)   \"SUSPENDED\" - The instance has suspended   \"SUSPENDING\" - The instance is suspending   \"TERMINATED\" - The instance has stopped (either by explicit action or underlying failure)",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchComputeInstanceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.InstanceGroups.AggregatedList(c.ProjectId).PageToken(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return err
		}
		output := list.(*compute.InstanceGroupAggregatedList)

		var instanceGroups []*compute.InstanceGroup
		for _, items := range output.Items {
			instanceGroups = append(instanceGroups, items.InstanceGroups...)
		}
		res <- instanceGroups

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveInstanceGroupsNamedPorts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.InstanceGroup)
	j := map[string]interface{}{}
	for _, v := range r.NamedPorts {
		j[v.Name] = v.Port
	}
	return resource.Set(c.Name, j)
}
func fetchComputeInstanceGroupInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*compute.InstanceGroup)
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		if r.Zone == "" {
			return nil
		}
		zoneParts := strings.Split(r.Zone, "/")
		zone := zoneParts[len(zoneParts)-1]

		call := c.Services.Compute.InstanceGroups.ListInstances(c.ProjectId, zone, r.Name, &compute.InstanceGroupsListInstancesRequest{}).PageToken(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return err
		}
		output := list.(*compute.InstanceGroupsListInstances)

		res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
func resolveInstanceGroupInstancesNamedPorts(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*compute.InstanceWithNamedPorts)
	j := map[string]interface{}{}
	for _, v := range r.NamedPorts {
		j[v.Name] = v.Port
	}
	return resource.Set(c.Name, j)
}
