
# Setting Up KubeVirt on Ubuntu and Fedora with Docker Hub as Registry

This guide will help you get KubeVirt running on your local Kubernetes cluster using Minikube. We'll walk through the process of building, pushing, and deploying KubeVirt images to Docker Hub.

## 1. Prerequisites

a) Ensure that your processor supports **nested virtualization** for [Ubuntu](https://ubuntu.com/server/docs/how-to-enable-nested-virtualization) or [Fedora](https://docs.fedoraproject.org/en-US/quick-docs/using-nested-virtualization-in-kvm/).

b) [Docker Hub account](https://hub.docker.com/).

c) Install the following tools:
   1) **Minikube**: If you're using Ubuntu and have already installed Minikube, you can skip this step.
   2) **Docker**: Required for building KubeVirt images.
   3) **kubectl**: The Kubernetes command-line tool for interacting with your cluster.

## 2. Start Minikube Cluster

Start Minikube with the Docker driver and allocate necessary resources for KubeVirt.

```bash
minikube start --driver=docker
```

## 3. Configure Docker Registry for Image Storage

To store your KubeVirt images, configure a Docker registry (e.g., Docker Hub). If you have a Docker Hub account, set the following environment variables:

```bash
export DOCKER_PREFIX=your-docker-hub-username
export DOCKER_TAG=mybuild
```

## 4. Build and Push KubeVirt Images

- First, sign in to Docker through the CLI using the following command:`docker login` [Docker documentation](https://docs.docker.com/accounts/create-account/#sign-in)
- Build the KubeVirt manifests, Docker images, and push them to your Docker registry:

```bash
make && make push && make manifests
```

If you encounter timeouts while fetching certain modules, increase the Bazel puller timeout:

```bash
export PULLER_TIMEOUT=10000
```

## 5. Deploy KubeVirt to Minikube

After the images are built and pushed, deploy KubeVirt to your Minikube cluster:

```bash
kubectl create -f _out/manifests/release/kubevirt-operator.yaml
kubectl create -f _out/manifests/release/kubevirt-cr.yaml
```

This will create the KubeVirt operator and its custom resource in your cluster.

## 6. Verify the Deployment

Check that all KubeVirt pods are running correctly:

```bash
kubectl get pods -n kubevirt
```

All pods should be in the `Running` state. If any pod is not running, inspect the logs to troubleshoot:

```bash
kubectl logs -n kubevirt <pod-name>
```

## 7. Create a Test Virtual Machine

With KubeVirt installed, you can create a test virtual machine using the example YAML file:

```bash
kubectl create -f examples/vmi-ephemeral.yaml
```

## 8. Check the Virtual Machine Status

After creating the virtual machine, verify that it is running:

```bash
kubectl get vmis
```

You should see the virtual machine in the `Running` phase.

---

And that's it! You now have KubeVirt up and running on your local cluster with Docker Hub as the image registry.
