all:
	docker build -t norden37/example_site .
	docker push norden37/example_site