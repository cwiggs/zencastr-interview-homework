.PHONY: docker/build
docker/build:
	docker build --tag zencastr-interview-homework .

.PHONY: docker/run
docker/run:
	docker run -it --rm -p 1234:1234 zencastr-interview-homework
