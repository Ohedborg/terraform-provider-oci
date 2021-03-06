// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// API for the Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service
// build on Hadoop, Spark and Data Science distribution, which can be fully integrated with existing enterprise
// data in Oracle Database and Oracle Applications..
//

package bds

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AddCloudSqlDetails The information about added Cloud SQL capability
type AddCloudSqlDetails struct {

	// Shape of the node
	Shape *string `mandatory:"true" json:"shape"`

	// Base-64 encoded password for Cloudera Manager admin user
	ClusterAdminPassword *string `mandatory:"true" json:"clusterAdminPassword"`

	// The size of block volume in GB that needs to be attached to a given node.
	// All the necessary details needed for attachment are managed by service itself.
	BlockVolumeSizeInGBs *int64 `mandatory:"false" json:"blockVolumeSizeInGBs"`
}

func (m AddCloudSqlDetails) String() string {
	return common.PointerString(m)
}
