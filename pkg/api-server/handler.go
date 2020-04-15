/*
Copyright 2020 The KubeCarrier Authors.

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

package apiserver

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/kubermatic/kubecarrier/pkg/apis/apiserver/v1alpha1"
	"github.com/kubermatic/kubecarrier/pkg/internal/version"
)

type kubecarrierHandler struct{}

func (v kubecarrierHandler) VersionSteam(e *empty.Empty, server v1alpha1.Kubecarrier_VersionSteamServer) error {
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		apiVersion, err := v.Version(ctx, &v1alpha1.VersionRequest{})
		if err != nil {
			return err
		}
		if err := server.Send(apiVersion); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func (v kubecarrierHandler) WhoAmI(ctx context.Context, _ *empty.Empty) (*v1alpha1.UserInfo, error) {
	user, present := ExtractUserInfo(ctx)
	if !present {
		return nil, fmt.Errorf("Unauthorized")
	}
	return &v1alpha1.UserInfo{
		User:   user.User.GetName(),
		Groups: user.User.GetGroups(),
	}, nil
}

var _ v1alpha1.KubecarrierServer = (*kubecarrierHandler)(nil)

func (v kubecarrierHandler) Version(context.Context, *v1alpha1.VersionRequest) (*v1alpha1.APIVersion, error) {
	versionInfo := version.Get()
	return &v1alpha1.APIVersion{
		Version: versionInfo.Version,
		Branch:  versionInfo.Branch,
		BuildDate: &timestamp.Timestamp{
			Seconds: versionInfo.BuildDate.Unix(),
			Nanos:   int32(versionInfo.BuildDate.Nanosecond()),
		},
		GoVersion: versionInfo.GoVersion,
		Platform:  versionInfo.Platform,
	}, nil
}
