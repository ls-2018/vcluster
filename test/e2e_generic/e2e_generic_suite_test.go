package e2egeneric

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/loft-sh/vcluster/cmd/vclusterctl/log"
	"github.com/loft-sh/vcluster/test/framework"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	// Enable cloud provider auth
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	// Register tests
	_ "github.com/loft-sh/vcluster/test/e2e_generic/clusterscope"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
}

func TestRunE2ETests(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	gomega.RegisterFailHandler(ginkgo.Fail)
	err := framework.CreateFramework(context.Background(), scheme)
	if err != nil {
		log.GetInstance().Fatalf("Error setting up framework: %v", err)
	}

	var _ = ginkgo.AfterSuite(func() {
		err = framework.DefaultFramework.Cleanup()
		if err != nil {
			log.GetInstance().Warnf("Error executing testsuite cleanup: %v", err)
		}
	})

	ginkgo.RunSpecs(t, "Vcluster e2e suite")
}
