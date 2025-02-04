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

package integrationresponse

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetIntegrationResponseInput returns input for read
// operation.
func GenerateGetIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.GetIntegrationResponseInput {
	res := &svcsdk.GetIntegrationResponseInput{}

	if cr.Status.AtProvider.IntegrationResponseID != nil {
		res.SetIntegrationResponseId(*cr.Status.AtProvider.IntegrationResponseID)
	}

	return res
}

// GenerateIntegrationResponse returns the current state in the form of *svcapitypes.IntegrationResponse.
func GenerateIntegrationResponse(resp *svcsdk.GetIntegrationResponseOutput) *svcapitypes.IntegrationResponse {
	cr := &svcapitypes.IntegrationResponse{}

	if resp.ContentHandlingStrategy != nil {
		cr.Spec.ForProvider.ContentHandlingStrategy = resp.ContentHandlingStrategy
	} else {
		cr.Spec.ForProvider.ContentHandlingStrategy = nil
	}
	if resp.IntegrationResponseId != nil {
		cr.Status.AtProvider.IntegrationResponseID = resp.IntegrationResponseId
	} else {
		cr.Status.AtProvider.IntegrationResponseID = nil
	}
	if resp.IntegrationResponseKey != nil {
		cr.Spec.ForProvider.IntegrationResponseKey = resp.IntegrationResponseKey
	} else {
		cr.Spec.ForProvider.IntegrationResponseKey = nil
	}
	if resp.ResponseParameters != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range resp.ResponseParameters {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		cr.Spec.ForProvider.ResponseParameters = f3
	} else {
		cr.Spec.ForProvider.ResponseParameters = nil
	}
	if resp.ResponseTemplates != nil {
		f4 := map[string]*string{}
		for f4key, f4valiter := range resp.ResponseTemplates {
			var f4val string
			f4val = *f4valiter
			f4[f4key] = &f4val
		}
		cr.Spec.ForProvider.ResponseTemplates = f4
	} else {
		cr.Spec.ForProvider.ResponseTemplates = nil
	}
	if resp.TemplateSelectionExpression != nil {
		cr.Spec.ForProvider.TemplateSelectionExpression = resp.TemplateSelectionExpression
	} else {
		cr.Spec.ForProvider.TemplateSelectionExpression = nil
	}

	return cr
}

// GenerateCreateIntegrationResponseInput returns a create input.
func GenerateCreateIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.CreateIntegrationResponseInput {
	res := &svcsdk.CreateIntegrationResponseInput{}

	if cr.Spec.ForProvider.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*cr.Spec.ForProvider.ContentHandlingStrategy)
	}
	if cr.Spec.ForProvider.IntegrationResponseKey != nil {
		res.SetIntegrationResponseKey(*cr.Spec.ForProvider.IntegrationResponseKey)
	}
	if cr.Spec.ForProvider.ResponseParameters != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range cr.Spec.ForProvider.ResponseParameters {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetResponseParameters(f2)
	}
	if cr.Spec.ForProvider.ResponseTemplates != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range cr.Spec.ForProvider.ResponseTemplates {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		res.SetResponseTemplates(f3)
	}
	if cr.Spec.ForProvider.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*cr.Spec.ForProvider.TemplateSelectionExpression)
	}

	return res
}

// GenerateUpdateIntegrationResponseInput returns an update input.
func GenerateUpdateIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.UpdateIntegrationResponseInput {
	res := &svcsdk.UpdateIntegrationResponseInput{}

	if cr.Spec.ForProvider.ContentHandlingStrategy != nil {
		res.SetContentHandlingStrategy(*cr.Spec.ForProvider.ContentHandlingStrategy)
	}
	if cr.Status.AtProvider.IntegrationResponseID != nil {
		res.SetIntegrationResponseId(*cr.Status.AtProvider.IntegrationResponseID)
	}
	if cr.Spec.ForProvider.IntegrationResponseKey != nil {
		res.SetIntegrationResponseKey(*cr.Spec.ForProvider.IntegrationResponseKey)
	}
	if cr.Spec.ForProvider.ResponseParameters != nil {
		f5 := map[string]*string{}
		for f5key, f5valiter := range cr.Spec.ForProvider.ResponseParameters {
			var f5val string
			f5val = *f5valiter
			f5[f5key] = &f5val
		}
		res.SetResponseParameters(f5)
	}
	if cr.Spec.ForProvider.ResponseTemplates != nil {
		f6 := map[string]*string{}
		for f6key, f6valiter := range cr.Spec.ForProvider.ResponseTemplates {
			var f6val string
			f6val = *f6valiter
			f6[f6key] = &f6val
		}
		res.SetResponseTemplates(f6)
	}
	if cr.Spec.ForProvider.TemplateSelectionExpression != nil {
		res.SetTemplateSelectionExpression(*cr.Spec.ForProvider.TemplateSelectionExpression)
	}

	return res
}

// GenerateDeleteIntegrationResponseInput returns a deletion input.
func GenerateDeleteIntegrationResponseInput(cr *svcapitypes.IntegrationResponse) *svcsdk.DeleteIntegrationResponseInput {
	res := &svcsdk.DeleteIntegrationResponseInput{}

	if cr.Status.AtProvider.IntegrationResponseID != nil {
		res.SetIntegrationResponseId(*cr.Status.AtProvider.IntegrationResponseID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
