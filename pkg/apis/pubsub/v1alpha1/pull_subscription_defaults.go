/*
Copyright 2019 Google LLC

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

package v1alpha1

import (
	"context"
	"time"
)

var (
	defaultAckDeadline       = 30 * time.Second
	defaultRetentionDuration = 7 * 24 * time.Hour
)

func (s *PullSubscription) SetDefaults(ctx context.Context) {
	s.Spec.SetDefaults(ctx)
}

func (ss *PullSubscriptionSpec) SetDefaults(ctx context.Context) {
	if ss.AckDeadline == nil {
		ss.AckDeadline = &defaultAckDeadline
	}

	if ss.RetentionDuration == nil {
		ss.RetentionDuration = &defaultRetentionDuration
	}
}
