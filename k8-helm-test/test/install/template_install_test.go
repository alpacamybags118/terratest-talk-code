package test

import (
	"crypto/tls"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gruntwork-io/terratest/modules/helm"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestHelmBasicExampleDeployment(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../../k8-test")
	require.NoError(t, err)

	namespaceName := fmt.Sprintf("k8-test-%s", strings.ToLower(random.UniqueId()))

	kubectlOptions := k8s.NewKubectlOptions("", "", namespaceName)

	k8s.CreateNamespace(t, kubectlOptions, namespaceName)
	defer k8s.DeleteNamespace(t, kubectlOptions, namespaceName)

	options := &helm.Options{
		KubectlOptions: kubectlOptions,
		ValuesFiles:    []string{"../../k8-test/values.yaml"},
	}

	releaseName := fmt.Sprintf(
		"k8-test-app-%s",
		strings.ToLower(random.UniqueId()),
	)
	defer helm.Delete(t, options, releaseName, true)

	helm.Install(t, options, helmChartPath, releaseName)

	serviceName := "test-app-cluster-ip"

	k8s.WaitUntilServiceAvailable(t, kubectlOptions, serviceName, 10, 1*time.Second)

	tlsConfig := tls.Config{}

	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		"http://localhost",
		&tlsConfig,
		30,
		10*time.Second,
		func(statusCode int, body string) bool {
			return statusCode == 200
		},
	)
}
