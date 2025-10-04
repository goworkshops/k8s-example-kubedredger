/*
Copyright 2025.

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

package controller

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	//	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes/scheme"

	"golab.io/kubedredger/internal/configfile"
)

const (
	confSnippet = "answer=42\n"
)

func NewFakeConfigurationReconciler() (*ConfigurationReconciler, string, func() error, error) {
	dir, err := os.MkdirTemp("", "kubedredger-ctrl-test")
	if err != nil {
		return nil, "", func() error { return nil }, err
	}
	GinkgoLogr.Info("created temporary directory", "path", dir)
	cleanup := func() error {
		return os.RemoveAll(dir)
	}
	rec := ConfigurationReconciler{
		Client:  k8sClient,
		Scheme:  scheme.Scheme,
		ConfMgr: configfile.NewManager(dir),
	}
	return &rec, dir, cleanup, nil
}

var _ = Describe("Configuration Controller", func() {
	//	var testNamespace *v1.Namespace

	Context("When reconciling a resource", func() {
		/*
			var cleanup func() error
			var reconciler *ConfigurationReconciler
			var configRoot string

			BeforeEach(func() {
				var err error
				reconciler, configRoot, cleanup, err = NewFakeConfigurationReconciler()
				Expect(err).ToNot(HaveOccurred())

					// see: https://book.kubebuilder.io/reference/envtest.html?highlight=envtest#namespace-usage-limitation
					ns := &v1.Namespace{
						ObjectMeta: metav1.ObjectMeta{
							GenerateName: "workshop-",
						},
					}
					Expect(reconciler.Client.Create(ctx, ns)).To(Succeed())
					testNamespace = ns
			})

			AfterEach(func() {
				// intentionally not try to delete namespaces.
				// see: https://book.kubebuilder.io/reference/envtest.html?highlight=envtest#namespace-usage-limitation
				Expect(cleanup()).To(Succeed())
			})
		*/

		When("handling the configuration", func() {
			It("creates the configuration from scratch", func(ctx context.Context) {
				/* example:
				conf := &workshopv1alpha1.Configuration{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: testNamespace.Name,
						Name:      "test-create",
					},
					Spec: workshopv1alpha1.ConfigurationSpec{
						Filename: "foo.conf",
						Content:  "foo=bar\nbaz=42\n",
						Create:   true,
					},
				}
				*/
				Fail("TODO: implement the test")
			})
		})
	})
})
