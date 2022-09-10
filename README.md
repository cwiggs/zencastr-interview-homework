#  zencastr-interview-homework
This is the homework assignment for zencastr interview.

There is a Makefile to help automate local development.

There is a github actions yaml file that automatically builds
and pushes the container image to the github container registry when on the `main` branch.

This repo also had a `deployment.yaml` that can be used with podman to deploy a pod with 3 replicas, or deployed using kubernetes.

This repo uses [pre-commit](https://pre-commit.com/) to manage pre-commit hooks.

## TODO
- [x] Build standalone REST app, 1 resource, max 3 routes.
- [x] Create docker container for service
- [x] Create deployment and autoscaling script.
- [x] Create an integration test as a pre-commit hook
- [x] Upload to github and send to zencastr
