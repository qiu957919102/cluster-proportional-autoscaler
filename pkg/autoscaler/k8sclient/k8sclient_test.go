/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package k8sclient

import (
	"testing"
)

func TestGetScaleTarget(t *testing.T) {
	testCases := []struct {
		target   string
		expKind  string
		expName  string
		expError bool
	}{
		{
			"deployment/anything",
			"deployment",
			"anything",
			false,
		},
		{
			"replicationcontroller/anotherthing",
			"replicationcontroller",
			"anotherthing",
			false,
		},
		{
			"replicationcontroller",
			"",
			"",
			true,
		},
		{
			"replicaset/anything/what",
			"",
			"",
			true,
		},
	}

	for _, tc := range testCases {
		res, err := getScaleTarget(tc.target, "default")
		if err != nil && !tc.expError {
			t.Errorf("Expect no error, got error for target: %v", tc.target)
			continue
		} else if err == nil && tc.expError {
			t.Errorf("Expect error, got no error for target: %v", tc.target)
			continue
		}
		if res.kind != tc.expKind || res.name != tc.expName {
			t.Errorf("Expect kind: %v, name: %v\ngot kind: %v, name: %v", tc.expKind, tc.expName, res.kind, res.name)
		}
	}
}
