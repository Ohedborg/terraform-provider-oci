// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateOceInstanceDetails The information about new OceInstance.
type CreateOceInstanceDetails struct {

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OceInstance Name
	Name *string `mandatory:"true" json:"name"`

	// Tenancy Identifier
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// Identity Cloud Service access token identifying a stripe and service administrator user
	IdcsAccessToken *string `mandatory:"true" json:"idcsAccessToken"`

	// Tenancy Name
	TenancyName *string `mandatory:"true" json:"tenancyName"`

	// Object Storage Namespace of Tenancy
	ObjectStorageNamespace *string `mandatory:"true" json:"objectStorageNamespace"`

	// Admin Email for Notification
	AdminEmail *string `mandatory:"true" json:"adminEmail"`

	// OceInstance description
	Description *string `mandatory:"false" json:"description"`

	IdentityStripe *IdentityStripeDetails `mandatory:"false" json:"identityStripe"`

	// Instance type based on its usage
	InstanceUsageType CreateOceInstanceDetailsInstanceUsageTypeEnum `mandatory:"false" json:"instanceUsageType,omitempty"`

	// Web Application Firewall(WAF) primary domain
	WafPrimaryDomain *string `mandatory:"false" json:"wafPrimaryDomain"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOceInstanceDetails) String() string {
	return common.PointerString(m)
}

// CreateOceInstanceDetailsInstanceUsageTypeEnum Enum with underlying type: string
type CreateOceInstanceDetailsInstanceUsageTypeEnum string

// Set of constants representing the allowable values for CreateOceInstanceDetailsInstanceUsageTypeEnum
const (
	CreateOceInstanceDetailsInstanceUsageTypePrimary    CreateOceInstanceDetailsInstanceUsageTypeEnum = "PRIMARY"
	CreateOceInstanceDetailsInstanceUsageTypeNonprimary CreateOceInstanceDetailsInstanceUsageTypeEnum = "NONPRIMARY"
)

var mappingCreateOceInstanceDetailsInstanceUsageType = map[string]CreateOceInstanceDetailsInstanceUsageTypeEnum{
	"PRIMARY":    CreateOceInstanceDetailsInstanceUsageTypePrimary,
	"NONPRIMARY": CreateOceInstanceDetailsInstanceUsageTypeNonprimary,
}

// GetCreateOceInstanceDetailsInstanceUsageTypeEnumValues Enumerates the set of values for CreateOceInstanceDetailsInstanceUsageTypeEnum
func GetCreateOceInstanceDetailsInstanceUsageTypeEnumValues() []CreateOceInstanceDetailsInstanceUsageTypeEnum {
	values := make([]CreateOceInstanceDetailsInstanceUsageTypeEnum, 0)
	for _, v := range mappingCreateOceInstanceDetailsInstanceUsageType {
		values = append(values, v)
	}
	return values
}
