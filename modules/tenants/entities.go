package tenants

import (
	"xhantos/sqlboiler/models"
)

type Tenant struct {
	ID     int    `json:"id,omitempty"`
	Alias  string `json:"alias" validate:"required,min=3,max=12"`
	APIKey string `json:"api_key,omitempty"`
	Status string `json:"status,omitempty"`
}

// mapTenantFromModel converts the SQLBoiler-generated *models.Tenant to the Tenant interface.
func mapTenantFromModel(modelTenant *models.Tenant) *Tenant {
	return &Tenant{
		ID:     modelTenant.ID,
		Alias:  modelTenant.Alias,
		APIKey: modelTenant.APIKey,
		Status: string(modelTenant.Status),
	}
}

// mapTenantToModel converts the Tenant interface to the SQLBoiler-generated *models.Tenant.
func mapTenantToModel(tenant Tenant) *models.Tenant {
	return &models.Tenant{
		ID:     tenant.ID,
		Alias:  tenant.Alias,
		APIKey: tenant.APIKey,
		Status: models.TenantStatus(tenant.Status),
	}
}