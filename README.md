# Simple Hello

This project automates the deployment process for a Go application using GitHub Actions and Kubernetes. When you push changes to the Go code, GitHub Actions is triggered, which in turn performs the following steps:

1. **Go Test**: It starts by running a simple Go test to ensure that the code changes meet the required quality standards.

2. **Docker Image Build**: If the Go tests pass successfully, GitHub Actions proceeds to build a Docker image of the code changes and uploads it to DockerHub.

3. **Deployment**: You have two deployment options:
   - **Argo CD**: You can use Argo CD to deploy the updated Docker image to your Kubernetes cluster automatically. Argo CD is a GitOps tool that continuously deploys changes based on your Git repository.
   - **Manual Deployment**: Alternatively, you can manually redeploy your application by rerunning the `deployment.yaml` file in your Kubernetes cluster. This gives you flexibility and control over the deployment process.

## Requirements

Before you can run this project, you'll need to have the following tools installed on your system:

- [Go](https://golang.org/doc/install) - The Go programming language.
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) - A tool to run Kubernetes locally.
- [VirtualBox](https://www.virtualbox.org/wiki/Downloads) - A virtualization tool used by Minikube.
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) - The Kubernetes command-line tool.
- [ArgoCd](https://argo-cd.readthedocs.io/en/stable/#quick-start) - ArgoCd will be used to sync updates we push to DockerHub to deployment.

