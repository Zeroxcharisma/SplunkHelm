// Copyright (c) 2018-2022 Splunk Inc. All rights reserved.

//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package enterprise

import (
	"context"
	enterpriseApi "github.com/splunk/splunk-operator/api/v4"

	enterpriseApiV3 "github.com/splunk/splunk-operator/api/v3"
	splcommon "github.com/splunk/splunk-operator/pkg/splunk/common"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// K8EventPublisher structure used to publish k8s event
type K8EventPublisher struct {
	client   splcommon.ControllerClient
	instance interface{}
}

// private function to get new k8s event publisher
func newK8EventPublisher(client splcommon.ControllerClient, instance interface{}) (*K8EventPublisher, error) {
	eventPublisher := &K8EventPublisher{
		client:   client,
		instance: instance,
	}

	return eventPublisher, nil
}

// publishEvents adds events to k8s
func (k *K8EventPublisher) publishEvent(ctx context.Context, eventType, reason, message string) {

	var event corev1.Event

	// in the case of testing, client is not passed
	if k.client == nil {
		return
	}

	// based on the custom resource instance type find name, type and create new event
	switch k.instance.(type) {
	case *enterpriseApi.Standalone:
		cr, ok := k.instance.(*enterpriseApi.Standalone)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	case *enterpriseApiV3.LicenseMaster:
		cr, ok := k.instance.(*enterpriseApiV3.LicenseMaster)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	case *enterpriseApi.LicenseManager:
		cr, ok := k.instance.(*enterpriseApi.LicenseManager)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	case *enterpriseApi.IndexerCluster:
		cr, ok := k.instance.(*enterpriseApi.IndexerCluster)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	case *enterpriseApi.ClusterManager:
		cr, ok := k.instance.(*enterpriseApi.ClusterManager)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	case *enterpriseApiV3.ClusterMaster:
		cr, ok := k.instance.(*enterpriseApiV3.ClusterMaster)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	case *enterpriseApi.MonitoringConsole:
		cr, ok := k.instance.(*enterpriseApi.MonitoringConsole)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	case *enterpriseApi.SearchHeadCluster:
		cr, ok := k.instance.(*enterpriseApi.SearchHeadCluster)
		if ok {
			return
		}
		event = cr.NewEvent(eventType, reason, message)
	default:
		return
	}

	reqLogger := log.FromContext(ctx)
	scopedLog := reqLogger.WithName("PublishEvent")
	scopedLog.Info("publishing event", "reason", event.Reason, "message", event.Message)

	err := k.client.Create(ctx, &event)
	if err != nil {
		scopedLog.Error(err, "failed to record event, ignoring",
			"reason", event.Reason, "message", event.Message, "error", err)
	}
}

// Normal publish normal events to k8s
func (k *K8EventPublisher) Normal(ctx context.Context, reason, message string) {
	k.publishEvent(ctx, corev1.EventTypeNormal, reason, message)
}

// Warning publish warning events to k8s
func (k *K8EventPublisher) Warning(ctx context.Context, reason, message string) {
	k.publishEvent(ctx, corev1.EventTypeWarning, reason, message)
}
