/*
Copyright 2019 The KubeCarrier Authors.

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

package webhooks

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-logr/logr"
	adminv1beta1 "k8s.io/api/admission/v1beta1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"k8c.io/utils/pkg/util"

	corev1alpha1 "k8c.io/kubecarrier/pkg/apis/core/v1alpha1"
)

// CustomResourceDiscoveryWebhookHandler handles mutating/validating of CustomResourceDiscoveries.
type CustomResourceDiscoveryWebhookHandler struct {
	decoder *admission.Decoder
	Log     logr.Logger
	client.Client
	Scheme *runtime.Scheme
}

var _ admission.Handler = (*CustomResourceDiscoveryWebhookHandler)(nil)

// +kubebuilder:webhook:path=/validate-kubecarrier-io-v1alpha1-customresourcediscovery,mutating=false,failurePolicy=fail,groups=kubecarrier.io,resources=customresourcediscoveries,verbs=update;delete,versions=v1alpha1,name=vcustomresourcediscovery.kubecarrier.io

// Handle is the function to handle update/delete requests of CustomResourceDiscovery objects.
func (r *CustomResourceDiscoveryWebhookHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	switch req.Operation {
	case adminv1beta1.Update:
		obj := &corev1alpha1.CustomResourceDiscovery{}
		if err := r.decoder.Decode(req, obj); err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}

		oldObj := &corev1alpha1.CustomResourceDiscovery{}
		if err := r.decoder.DecodeRaw(req.OldObject, oldObj); err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}
		if err := r.validateUpdate(oldObj, obj); err != nil {
			return admission.Denied(err.Error())
		}
	case adminv1beta1.Delete:
		oldObj := &corev1alpha1.CustomResourceDiscovery{}
		if err := r.decoder.DecodeRaw(req.OldObject, oldObj); err != nil {
			return admission.Errored(http.StatusBadRequest, err)
		}
		if err := r.validateDelete(ctx, oldObj); err != nil {
			return admission.Denied(err.Error())
		}
	}
	return admission.Allowed("allowed to commit the request")

}

// CustomResourceDiscoveryWebhookHandler implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (r *CustomResourceDiscoveryWebhookHandler) InjectDecoder(d *admission.Decoder) error {
	r.decoder = d
	return nil
}

func (r *CustomResourceDiscoveryWebhookHandler) validateUpdate(oldObj, newObj *corev1alpha1.CustomResourceDiscovery) error {
	r.Log.Info("validate update", "name", newObj.Name)
	if newObj.Spec.ServiceCluster.Name != oldObj.Spec.ServiceCluster.Name ||
		newObj.Spec.CRD.Name != oldObj.Spec.CRD.Name {
		return fmt.Errorf("the Spec (ServiceCluster, CRD) of CustomResourceDiscovery is immutable")
	}
	return nil
}

func (r *CustomResourceDiscoveryWebhookHandler) validateDelete(ctx context.Context, obj *corev1alpha1.CustomResourceDiscovery) error {
	r.Log.Info("validate delete", "name", obj.Name)
	if obj.Status.ManagementClusterCRD == nil {
		return nil
	}

	crd := &apiextensionsv1.CustomResourceDefinition{}
	if err := r.Get(ctx, types.NamespacedName{
		Name: obj.Status.ManagementClusterCRD.Name,
	}, crd); err != nil {
		return err
	}
	u := &unstructured.UnstructuredList{}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   crd.Spec.Group,
		Version: crd.Spec.Versions[0].Name,
		Kind:    crd.Spec.Names.ListKind,
	})
	if err := r.List(ctx, u); err != nil {
		return err
	}
	if len(u.Items) == 0 {
		return nil
	}

	errorMsg := new(strings.Builder)
	errorMsg.WriteString("management cluster CRD instances are still present in the management cluster\n")
	for _, it := range u.Items {
		errorMsg.WriteString(util.MustLogLine(&it, r.Scheme) + " still present\n")
	}
	return fmt.Errorf("%s", errorMsg)
}
