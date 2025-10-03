# Kubebuilder Workshop Exercise

Follow these steps to create a Kubernetes controller using Kubebuilder:

## 1. Download and Install Kubebuilder

```bash
# download kubebuilder and install locally.
curl -L -o kubebuilder "https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)"
chmod +x kubebuilder && sudo mv kubebuilder /usr/local/bin/
```

## 2. Initialize the Project

Generate the project with scaffolding:

```bash
kubebuilder init --domain golab.io --repo golab.io/kubedredger
```

## 3. Create an API

Generate an API:

```bash
kubebuilder create api --group workshop --version v1alpha1 \
--kind Configuration
```

When asked, say yes to create resource and create controller:

```bash
INFO Create Resource [y/n]
y
INFO Create Controller [y/n]
y
```

## 4. Explore the Project Structure

Look around and check the structure of the project to understand the generated files and directories.

## 5. Generate Manifests

Generate the manifests:

```bash
make manifests
```

## 6. Fill the API Fields

Add fields to your API. For example, add a `Filename` field:

```go
Filename string `json:"filename"`
```

## 7. Regenerate Manifests and Compare

Regenerate the manifests and see the differences:

```bash
make manifests
```

## 8. Add Default Values

Add a default value to the filename field using kubebuilder annotations:

```go
// +kubebuilder:default:=32
```

## 9. Final Manifest Generation

Regenerate the manifests one more time and observe the differences:

```bash
make manifests
```

## 10. Explore RBAC Annotations

Check the RBAC annotations in `internal/controller/configuration_controller.go:36-38` and find the corresponding generated manifests that were created from these annotations.
