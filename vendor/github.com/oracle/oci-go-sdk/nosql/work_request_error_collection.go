// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// ndcs-control-plane API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestErrorCollection Results of ListWorkRequestErrors
type WorkRequestErrorCollection struct {

	// A page of WorkRequestError objects.
	Items []WorkRequestError `mandatory:"false" json:"items"`
}

func (m WorkRequestErrorCollection) String() string {
	return common.PointerString(m)
}
