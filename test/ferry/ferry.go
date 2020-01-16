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

package ferry

import (
	"context"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	catalogv1alpha1 "github.com/kubermatic/kubecarrier/pkg/apis/catalog/v1alpha1"
	corev1alpha1 "github.com/kubermatic/kubecarrier/pkg/apis/core/v1alpha1"
	operatorv1alpha1 "github.com/kubermatic/kubecarrier/pkg/apis/operator/v1alpha1"
	"github.com/kubermatic/kubecarrier/pkg/testutil"
	"github.com/kubermatic/kubecarrier/test/framework"
)

var _ suite.SetupAllSuite = (*FerrySuite)(nil)

// FerrySuite checks the Ferry related operations creations and deletions.
type FerrySuite struct {
	suite.Suite
	*framework.Framework

	masterClient  client.Client
	serviceClient client.Client
}

func (s *FerrySuite) SetupSuite() {
	var err error
	s.serviceClient, err = s.ServiceClient()
	s.Require().NoError(err, "creating service client")
	s.masterClient, err = s.MasterClient()
	s.Require().NoError(err, "creating master client")
}

func (s *FerrySuite) TestCreationAndDeletion() {
	t := s.T()
	ctx := context.Background()

	serviceKubeconfig, err := ioutil.ReadFile(s.Framework.Config().ServiceInternalKubeconfigPath)
	require.NoError(t, err, "cannot read service internal kubeconfig")

	provider := &catalogv1alpha1.Provider{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "steel-inquisitor",
			Namespace: "kubecarrier-system",
		},
	}
	require.NoError(t, s.masterClient.Create(ctx, provider))
	require.NoError(t, testutil.WaitUntilReady(s.masterClient, provider))

	sec := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "eu-west-1",
			Namespace: provider.Status.NamespaceName,
		},
		Data: map[string][]byte{
			"kubeconfig": serviceKubeconfig,
		},
	}
	scr := &operatorv1alpha1.ServiceClusterRegistration{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "eu-west-1",
			Namespace: provider.Status.NamespaceName,
		},
		Spec: operatorv1alpha1.ServiceClusterRegistrationSpec{
			KubeconfigSecret: operatorv1alpha1.ObjectReference{
				Name: "eu-west-1",
			},
		},
	}
	require.NoError(t, client.IgnoreNotFound(s.masterClient.Delete(ctx, sec.DeepCopy())))
	require.NoError(t, client.IgnoreNotFound(s.masterClient.Delete(ctx, scr.DeepCopy())))

	require.NoError(t, s.masterClient.Create(ctx, sec))
	require.NoError(t, s.masterClient.Create(ctx, scr))

	require.NoError(t, testutil.WaitUntilReady(
		s.masterClient,
		scr,
	), "serviceClusterRegistration object not ready within time limit")

	serviceCluster := &corev1alpha1.ServiceCluster{}
	serviceCluster.SetName(scr.GetName())
	serviceCluster.SetNamespace(provider.Status.NamespaceName)
	require.NoError(t, testutil.WaitUntilReady(s.masterClient, serviceCluster))

	assert.NoError(t, testutil.WaitUntilReady(s.masterClient, scr), "scr object not ready within the time limit")

	require.NoError(t, s.masterClient.Delete(ctx, scr))
	assert.NoError(t, testutil.WaitUntilNotFound(s.masterClient, scr))
	assert.NoError(t, s.masterClient.Delete(ctx, sec))
}