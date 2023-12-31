/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1alpha1 "kubedb.dev/provider-azure/apis/azure/v1alpha1"
	rconfig "kubedb.dev/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this PrivateDNSZone.
func (mg *PrivateDNSZone) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PrivateDNSZoneVirtualNetworkLink.
func (mg *PrivateDNSZoneVirtualNetworkLink) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateDNSZoneName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PrivateDNSZoneNameRef,
		Selector:     mg.Spec.ForProvider.PrivateDNSZoneNameSelector,
		To: reference.To{
			List:    &PrivateDNSZoneList{},
			Managed: &PrivateDNSZone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateDNSZoneName")
	}
	mg.Spec.ForProvider.PrivateDNSZoneName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateDNSZoneNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkIDRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkIDSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkID")
	}
	mg.Spec.ForProvider.VirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RouteTable.
func (mg *RouteTable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroup.
func (mg *SecurityGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkNameRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkNameSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkName")
	}
	mg.Spec.ForProvider.VirtualNetworkName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetNetworkSecurityGroupAssociation.
func (mg *SubnetNetworkSecurityGroupAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkSecurityGroupID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NetworkSecurityGroupIDRef,
		Selector:     mg.Spec.ForProvider.NetworkSecurityGroupIDSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkSecurityGroupID")
	}
	mg.Spec.ForProvider.NetworkSecurityGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkSecurityGroupIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetRouteTableAssociation.
func (mg *SubnetRouteTableAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouteTableID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RouteTableIDRef,
		Selector:     mg.Spec.ForProvider.RouteTableIDSelector,
		To: reference.To{
			List:    &RouteTableList{},
			Managed: &RouteTable{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouteTableID")
	}
	mg.Spec.ForProvider.RouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouteTableIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetwork.
func (mg *VirtualNetwork) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VirtualNetworkPeering.
func (mg *VirtualNetworkPeering) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RemoteVirtualNetworkID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.RemoteVirtualNetworkIDRef,
		Selector:     mg.Spec.ForProvider.RemoteVirtualNetworkIDSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RemoteVirtualNetworkID")
	}
	mg.Spec.ForProvider.RemoteVirtualNetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RemoteVirtualNetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1alpha1.ResourceGroupList{},
			Managed: &v1alpha1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualNetworkName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VirtualNetworkNameRef,
		Selector:     mg.Spec.ForProvider.VirtualNetworkNameSelector,
		To: reference.To{
			List:    &VirtualNetworkList{},
			Managed: &VirtualNetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualNetworkName")
	}
	mg.Spec.ForProvider.VirtualNetworkName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualNetworkNameRef = rsp.ResolvedReference

	return nil
}
