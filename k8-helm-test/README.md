# k8-helm-test

A example of testing using terratest to test a helm chart, both simple template validation and template deployment. It creates the following:

1. Deployment of a single container with single port
2. A service that exposes the container on a specified service port, mapped to the container port
3. An ingress configuration that routes traffic to the service.

# Requirements

You will need the following to run these examples

1. A Kubernetes cluster. These examples were developed using the Kubernetes cluster included with [Docker Desktop](https://docs.docker.com/get-docker/)
2. [Helm](https://helm.sh/docs/intro/install/)
3. Your Kubernetes cluster will need some kind of ingress controller installed. These examples were developed using the [nginx ingress controller](https://kubernetes.github.io/ingress-nginx/deploy/)

## Running Helm

You can run the helm chart by running `helm install {your-release-name} ./k8-test`. This will use the default values in the `values.yaml` file, which runs an httpd container.

## Running Tests

You will need to have `go` installed in order to run the test.
To run tests, follow these steps:

1. Run `cd test`
2. Run `go mod init k8-helm-test`
3. Run `go test -v -timeout 30m`