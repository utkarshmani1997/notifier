
# Output registry and image names for operator image
# Set env to override this value
ifeq (${IMAGE_ORG}, )
  IMAGE_ORG:=utkarshmani1997
endif

export IMAGE_ORG

# Output app name and its image name and tag
APP_NAME=notifier
APP_TAG=ci

PACKAGES = $(shell go list ./... | grep -v 'vendor')

.PHONY: all
all:
	@echo "Available commands:"
	@echo "  build                           - build app source code"
	@echo "  image                           - build app container image"
	@echo "  push                            - push app to dockerhub registry (${IMAGE_ORG})"
	@echo ""
	@make print-variables --no-print-directory

.PHONY: print-variables
print-variables:
	@echo "Testing variables:"
	@echo " Produced Image: ${APP_NAME}:${APP_TAG}"
	@echo " IMAGE_ORG: ${IMAGE_ORG}"

.get:
	rm -rf ./build/bin/
	go mod download

deps: .get
	go mod vendor

test:
	@echo "--> Running go test" ;
	go test -v --cover $(PACKAGES)

build: deps test
	@echo "--> Build binary $(APP_NAME) ..."
	GOOS=linux go build -o ./build/bin/$(APP_NAME)-apiserver ./cmd/apiserver/main.go

jfrog: test
	@echo "--> Build artifacts $(APP_NAME) ..."
	jfrog rt go build -o ./build/bin/$(APP_NAME)-apiserver ./cmd/apiserver/main.go --build-name=$(BUILD_NAME) --build-number=$(BUILD_NUMBER)

jfrog-publish:
	@echo "--> Publish build to artifactory ..."
	jfrog rt bp $(BUILD_NAME) $(BUILD_NUMBER)

image: build
	@echo "--> Build image $(IMAGE_ORG)/$(APP_NAME):$(APP_TAG) ..."
	docker build -f ./build/Dockerfile -t $(IMAGE_ORG)/$(APP_NAME):$(APP_TAG) .

push: image
	@echo "--> Push image $(IMAGE_ORG)/$(APP_NAME):$(APP_TAG) ..."
	docker push $(IMAGE_ORG)/$(APP_NAME):$(APP_TAG)
