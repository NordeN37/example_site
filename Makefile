all:
	docker buildx build --platform=linux/amd64 -t norden37/example_site .
	docker push norden37/example_site