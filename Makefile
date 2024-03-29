
# Image URL to use all building/pushing image targets
IMG ?= controller:latest
IMGVERSION=0.1.19

all: test manager

# Run tests
test: generate fmt vet manifests
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager sigs.k8s.io/cluster-api-provider-helloworld/cmd/manager

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install: manifests
	kubectl apply -f config/crds

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	cat provider-components.yaml | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
	kustomize build config/default/ > provider-components.yaml
	echo "---" >> provider-components.yaml
	kustomize build vendor/sigs.k8s.io/cluster-api/config/default/ >> provider-components.yaml
	sed -i 's/cluster-api-provider-helloworld-controller-manager-metrics-service/cluster-api-provider-hw-controller-manager-metrics-service/g' provider-components.yaml
	sed -i 's/himani93:cluster-api-provider-hw/himani93:cluster-api-provider-hw:'"$IMGVERSION"'/g' provider-components.yaml

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	go generate ./pkg/... ./cmd/...

# Build the docker image
docker-build: test
	docker build . -t himani93/cluster-api-provider-hw:${IMGVERSION}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"himani93/cluster-api-provider-hw"'@' ./config/default/manager_image_patch.yaml

# Push the docker image
docker-push:
	docker push himani93/cluster-api-provider-hw:${IMGVERSION}

# Build and push docker image
magic:
	bash ./magic.sh
