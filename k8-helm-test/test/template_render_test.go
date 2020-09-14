package test

import (
	"path/filepath"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkv1 "k8s.io/api/networking/v1beta1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelmRenderedDeploymentTemplate(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../k8-test")
	releaseName := "helm-test"
	require.NoError(t, err)

	options := &helm.Options{
		ValuesFiles: []string{"../k8-test/values.yaml"},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/deployment.yaml"})

	var renderedDeployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &renderedDeployment)

	assert.Equal(t, renderedDeployment.GetObjectMeta().GetName(), "test-app-deployment")
	assert.Equal(t, renderedDeployment.Spec.Template.Spec.Containers[0].Image, "nginx")
	assert.Equal(t, renderedDeployment.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort, int32(3000))
}

func TestHelmRenderedServiceTemplate(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../k8-test")
	releaseName := "helm-test"
	require.NoError(t, err)

	options := &helm.Options{
		ValuesFiles: []string{"../k8-test/values.yaml"},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/service.yaml"})

	var renderedService corev1.Service
	helm.UnmarshalK8SYaml(t, output, &renderedService)

	assert.Equal(t, renderedService.GetObjectMeta().GetName(), "test-app-cluster-ip")
	assert.Equal(t, renderedService.Spec.Ports[0].Port, int32(80))
	assert.Equal(t, renderedService.Spec.Ports[0].TargetPort.IntVal, int32(3000))
	assert.Equal(t, renderedService.Spec.Selector["app"], "test-app-deployment")
}

func TestHelmRenderedIngressTemplate(t *testing.T) {
	t.Parallel()

	helmChartPath, err := filepath.Abs("../k8-test")
	releaseName := "helm-test"
	require.NoError(t, err)

	options := &helm.Options{
		ValuesFiles: []string{"../k8-test/values.yaml"},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/ingress.yaml"})

	var renderedIngress networkv1.Ingress
	helm.UnmarshalK8SYaml(t, output, &renderedIngress)

	assert.Equal(t, renderedIngress.GetObjectMeta().GetName(), "test-app-ingress")
	assert.Equal(t, renderedIngress.Spec.Rules[0].HTTP.Paths[0].Backend.ServiceName, "test-app-cluster-ip")
	assert.Equal(t, renderedIngress.Spec.Rules[0].HTTP.Paths[0].Backend.ServicePort.IntVal, int32(80))

}
