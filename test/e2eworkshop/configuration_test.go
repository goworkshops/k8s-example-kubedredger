package e2e

import (
	"context"
	"time"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"golab.io/kubedredger/api/v1alpha1"
)

var _ = ginkgo.Describe("Configuration E2E", func() {
	var (
		configuration *v1alpha1.Configuration
		testNamespace string
		confName      string
		//		confRoot      string
	)

	ginkgo.BeforeEach(func() {
		testNamespace = "golab-kubedredger"
		//		confRoot = "/tmp/config.d"
		//		confName = "kubedredger.conf"

		configuration = &v1alpha1.Configuration{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-config",
				Namespace: testNamespace,
			},
			Spec: v1alpha1.ConfigurationSpec{
				Filename: confName,
				Content:  "test content for e2e",
				Create:   true,
			},
		}
	})

	ginkgo.AfterEach(func() {
		ctx := context.Background()

		if configuration == nil {
			return // nothing to do
		}
		_ = cl.Delete(ctx, configuration)

		ginkgo.By("ensuring configuration is removed")
		Eventually(func() bool {
			err := cl.Get(ctx, client.ObjectKeyFromObject(configuration), configuration)
			return apierrors.IsNotFound(err)
		}, time.Minute, time.Second).Should(BeTrue())
	})

	ginkgo.It("should create configuration and verify status and file creation", func() {
		ginkgo.Fail("TODO: implement the test")
	})
})
