// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// File Storage Service API
//
// APIs for OCI file storage service.
//

package ffsw

import (
    "github.com/oracle/oci-go-sdk/common"
)


    
 // MountTargetSummary Summary information for a MountTarget.
type MountTargetSummary struct {
    
 // The OCID of the compartment that contains the mount target.
    CompartmentId *string `mandatory:"true" json:"compartmentId"`
    
 // A user-friendly name. Does not have to be unique, and it is changeable.
 // Avoid entering confidential information.
 // Example: `My mount target`
    DisplayName *string `mandatory:"true" json:"displayName"`
    
 // The OCID of the mount target.
    Id *string `mandatory:"true" json:"id"`
    
 // The current state of the mount target.
    LifecycleState MountTargetSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
    
 // The OCIDs of the private IPs associated with this mount target.
    PrivateIpIds []string `mandatory:"true" json:"privateIpIds"`
    
 // The OCID of the subnet the mount target is in.
    SubnetId *string `mandatory:"true" json:"subnetId"`
    
 // The date and time the mount target was created, in the format defined by RFC3339.
 // Example: `2016-08-25T21:10:29.600Z`
    TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
    
 // The Availability Domain the mount target is in. May be unset.
 // Example: `Uocm:PHX-AD-1`
    AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
    
 // The OCID of the associated export set. Controls what file
 // systems will be exported via NFS on this mount target.
    ExportSetId *string `mandatory:"false" json:"exportSetId"`
}

func (m MountTargetSummary) String() string {
    return common.PointerString(m)
}


// MountTargetSummaryLifecycleStateEnum Enum with underlying type: string
type MountTargetSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for MountTargetSummaryLifecycleState
const (
    MountTargetSummaryLifecycleStateCreating MountTargetSummaryLifecycleStateEnum = "CREATING"
    MountTargetSummaryLifecycleStateActive MountTargetSummaryLifecycleStateEnum = "ACTIVE"
    MountTargetSummaryLifecycleStateDeleting MountTargetSummaryLifecycleStateEnum = "DELETING"
    MountTargetSummaryLifecycleStateDeleted MountTargetSummaryLifecycleStateEnum = "DELETED"
    MountTargetSummaryLifecycleStateFailed MountTargetSummaryLifecycleStateEnum = "FAILED"
    MountTargetSummaryLifecycleStateUnknown MountTargetSummaryLifecycleStateEnum = "UNKNOWN"
)

var mappingMountTargetSummaryLifecycleState = map[string]MountTargetSummaryLifecycleStateEnum { 
    "CREATING": MountTargetSummaryLifecycleStateCreating,
    "ACTIVE": MountTargetSummaryLifecycleStateActive,
    "DELETING": MountTargetSummaryLifecycleStateDeleting,
    "DELETED": MountTargetSummaryLifecycleStateDeleted,
    "FAILED": MountTargetSummaryLifecycleStateFailed,
    "UNKNOWN": MountTargetSummaryLifecycleStateUnknown,
}

// GetMountTargetSummaryLifecycleStateEnumValues Enumerates the set of values for MountTargetSummaryLifecycleState
func GetMountTargetSummaryLifecycleStateEnumValues() []MountTargetSummaryLifecycleStateEnum {
   values := make([]MountTargetSummaryLifecycleStateEnum, 0)
   for _, v := range mappingMountTargetSummaryLifecycleState {
      if v != MountTargetSummaryLifecycleStateUnknown {
         values = append(values, v)
      }
   }
   return values
}



