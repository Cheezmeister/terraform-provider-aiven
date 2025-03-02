package pg

import (
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourcePG() *schema.Resource {
	return &schema.Resource{
		ReadContext: schemautil.DatasourceServiceRead,
		Description: "The PG data source provides information about the existing Aiven PostgreSQL service.",
		Schema:      schemautil.ResourceSchemaAsDatasourceSchema(aivenPGSchema(), "project", "service_name"),
	}
}
