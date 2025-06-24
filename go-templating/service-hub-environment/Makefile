REGISTRY_NAME = servicehubregistry
IMG_NAME = service_hub_environment
DATE = $(shell date +%Y%m%d)

build-multiarch-image:
	if ! docker buildx ls | grep -q svchubbuilder; then \
		docker buildx create --name svchubbuilder --driver docker-container --bootstrap --use; \
	else \
		docker buildx use svchubbuilder; \
	fi
	docker buildx build --platform linux/amd64,linux/arm64 -t $(IMG_NAME) -f Dockerfile .. --load
	docker buildx use --builder default
	docker buildx rm --builder svchubbuilder
	docker save -o $(IMG_NAME)-image.tar $(IMG_NAME)

single-arch-image:
	docker build -t $(IMG_NAME) -f Dockerfile ..
	docker save -o $(IMG_NAME)-image.tar $(IMG_NAME)

push-image:
	az acr login --name $(REGISTRY_NAME)
	docker load -i $(IMG_NAME)-image.tar
	docker tag $(IMG_NAME) $(REGISTRY_NAME).azurecr.io/$(IMG_NAME):$(DATE)
	docker push $(REGISTRY_NAME).azurecr.io/$(IMG_NAME):$(DATE)
	rm $(IMG_NAME)-image.tar;
