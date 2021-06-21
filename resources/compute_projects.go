package resources

import (
	"context"
	"fmt"
	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	compute "google.golang.org/api/compute/v1"
)

func ComputeProjects() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_projects",
		Description:  "Represents a Project resource  A project is used to organize resources in a Google Cloud Platform environment For more information, read about the  Resource Hierarchy (== resource_for {$api_version}",
		Resolver:     fetchComputeProjects,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name:        "common_instance_metadata_fingerprint",
				Description: "Specifies a fingerprint for this request, which is essentially a hash of the metadata's contents and used for optimistic locking The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet  To see the latest fingerprint, make a get() request to retrieve the resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CommonInstanceMetadata.Fingerprint"),
			},
			{
				Name:        "common_instance_metadata_items",
				Description: "Array of key/value pairs The total size of all keys and values must be less than 512 KB",
				Type:        schema.TypeJSON,
				Resolver:    resolveComputeProjectCommonInstanceMetadataItems,
			},
			{
				Name:        "common_instance_metadata_kind",
				Description: "Type of the resource Always compute#metadata for metadata",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CommonInstanceMetadata.Kind"),
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "default_network_tier",
				Description: "This signifies the default network tier used for configuring resources of the project and can only take the following values: PREMIUM, STANDARD Initially the default network tier is PREMIUM",
				Type:        schema.TypeString,
			},
			{
				Name:        "default_service_account",
				Description: "Default service account used by VMs running in this project",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional textual description of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled_features",
				Description: "Restricted features enabled for use on this project",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#project for projects",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The project ID For example: my-example-project Use the project ID to make requests to Compute Engine",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "usage_export_location_bucket_name",
				Description: "The name of an existing bucket in Cloud Storage where the usage report object is stored The Google Service Account is granted write access to this bucket This can either be the bucket name by itself, such as example-bucket, or the bucket name with gs:// or https://storagegoogleapiscom/ in front of it, such as gs://example-bucket",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UsageExportLocation.BucketName"),
			},
			{
				Name:        "usage_export_location_report_name_prefix",
				Description: "An optional prefix for the name of the usage report object stored in bucketName If not supplied, defaults to usage The report is stored as a CSV file named report_name_prefix_gce_YYYYMMDDcsv where YYYYMMDD is the day of the usage according to Pacific Time If you supply a prefix, it should conform to Cloud Storage object naming conventions",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UsageExportLocation.ReportNamePrefix"),
			},
			{
				Name:        "xpn_project_status",
				Description: "The role this project has in a shared VPC configuration Currently, only projects with the host role, which is specified by the value HOST, are differentiated",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "gcp_compute_project_quotas",
				Description: "A quotas entry",
				Resolver:    fetchComputeProjectQuotas,
				Columns: []schema.Column{
					{
						Name:        "project_id",
						Description: "Unique ID of gcp_compute_projects table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "limit",
						Description: "Quota limit for this metric",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "metric",
						Description: "Name of the quota metric",
						Type:        schema.TypeString,
					},
					{
						Name:        "owner",
						Description: "Owning resource This is the resource on which this quota is applied",
						Type:        schema.TypeString,
					},
					{
						Name:        "usage",
						Description: "Current usage of this metric",
						Type:        schema.TypeFloat,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	call := c.Services.Compute.Projects.
		Get(c.ProjectId).
		Context(ctx)
	output, err := call.Do()
	if err != nil {
		return err
	}
	res <- output
	return nil
}
func resolveComputeProjectCommonInstanceMetadataItems(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(*compute.Project)
	if !ok {
		return fmt.Errorf("expected *compute.Project but got %T", p)
	}
	m := make(map[string]interface{})
	for _, i := range p.CommonInstanceMetadata.Items {
		m[i.Key] = i.Value
	}
	return resource.Set(c.Name, m)
}
func fetchComputeProjectQuotas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(*compute.Project)
	if !ok {
		return fmt.Errorf("expected *compute.Project but got %T", p)
	}
	res <- p.Quotas
	return nil
}
