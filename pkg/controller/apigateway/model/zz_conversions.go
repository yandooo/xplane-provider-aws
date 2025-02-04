/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package model

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigateway/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetModelInput returns input for read
// operation.
func GenerateGetModelInput(cr *svcapitypes.Model) *svcsdk.GetModelInput {
	res := &svcsdk.GetModelInput{}

	return res
}

// GenerateModel returns the current state in the form of *svcapitypes.Model.
func GenerateModel(resp *svcsdk.Model) *svcapitypes.Model {
	cr := &svcapitypes.Model{}

	if resp.ContentType != nil {
		cr.Spec.ForProvider.ContentType = resp.ContentType
	} else {
		cr.Spec.ForProvider.ContentType = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Id != nil {
		cr.Status.AtProvider.ID = resp.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.Name != nil {
		cr.Spec.ForProvider.Name = resp.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}
	if resp.Schema != nil {
		cr.Spec.ForProvider.Schema = resp.Schema
	} else {
		cr.Spec.ForProvider.Schema = nil
	}

	return cr
}

// GenerateCreateModelInput returns a create input.
func GenerateCreateModelInput(cr *svcapitypes.Model) *svcsdk.CreateModelInput {
	res := &svcsdk.CreateModelInput{}

	if cr.Spec.ForProvider.ContentType != nil {
		res.SetContentType(*cr.Spec.ForProvider.ContentType)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}
	if cr.Spec.ForProvider.Schema != nil {
		res.SetSchema(*cr.Spec.ForProvider.Schema)
	}

	return res
}

// GenerateUpdateModelInput returns an update input.
func GenerateUpdateModelInput(cr *svcapitypes.Model) *svcsdk.UpdateModelInput {
	res := &svcsdk.UpdateModelInput{}

	return res
}

// GenerateDeleteModelInput returns a deletion input.
func GenerateDeleteModelInput(cr *svcapitypes.Model) *svcsdk.DeleteModelInput {
	res := &svcsdk.DeleteModelInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
