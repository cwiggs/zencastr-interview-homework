#  zencastr-interview-homework
This is the homework assignment for zencastr interview.

There is a Makefile to help automate local development.  There is a github actions yaml file that automatically builds
and pushes the container image to the github container registry when on the `main` branch.

This repo uses [pre-commit](https://pre-commit.com/) to add pre-commit hooks.

## TODO
- [x] Build standalone REST app, 1 resource, max 3 routes.
- [x] Create docker container for service
- [] Create deployment and autoscaling script.
- [x] Create an integration test as a pre-commit hook
- [] Upload to github and send to zencastr
