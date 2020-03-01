IMAGE := udamliyanage/actions-api:v1.0.0

test:
	true

image:
	docker build -t $(IMAGE) .

push-image:
	docker push $(IMAGE)


.PHONY: image push-image test