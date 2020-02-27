IMAGE := udamliyanage/device-configuration-service:v2.1.0

test:
	true

image:
	docker build -t $(IMAGE) .

push-image:
	docker push $(IMAGE)


.PHONY: image push-image test