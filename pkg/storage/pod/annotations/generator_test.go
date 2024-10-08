/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2024 The KubeVirt Authors.
 *
 */

package annotations_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"kubevirt.io/kubevirt/pkg/libvmi"
	"kubevirt.io/kubevirt/pkg/storage/pod/annotations"
)

var _ = Describe("Annotations Generator", func() {
	It("Should generate storage annotations", func() {
		const (
			testNamespace = "testns"
			vmiName       = "testvmi"
		)

		generator := annotations.Generator{}
		annotations, err := generator.Generate(libvmi.New(libvmi.WithNamespace(testNamespace), libvmi.WithName(vmiName)))
		Expect(err).NotTo(HaveOccurred())

		expectedAnnotations := map[string]string{
			"pre.hook.backup.velero.io/container":  "compute",
			"pre.hook.backup.velero.io/command":    "[\"/usr/bin/virt-freezer\", \"--freeze\", \"--name\", \"testvmi\", \"--namespace\", \"testns\"]",
			"post.hook.backup.velero.io/container": "compute",
			"post.hook.backup.velero.io/command":   "[\"/usr/bin/virt-freezer\", \"--unfreeze\", \"--name\", \"testvmi\", \"--namespace\", \"testns\"]",
		}

		Expect(annotations).To(Equal(expectedAnnotations))
	})
})
