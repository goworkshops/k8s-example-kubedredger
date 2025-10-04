# Kubebuilder Workshop Exercise

Create and run integration and e2e tests.
This practice session can be quite long.
If short on time, aim for the first 2 items;
If you have time, try to complete all 4.

**IMPORTANT NOTE**: please ignore `test/e2e`.
Please implement the e2e test code in `test/e2eworkshop`.

## 1. Familiarize with ginkgo and gomega

Try to run the existing tests. Be aware that the first run will download
any missing bits and pieces like apiserver and etcd instances. It may
take a while on slow networks.
```
make test
```
**note** the tests are meant to fail at this stage

## 2. Write new integration tests

Take a look at `internal/controller/configuration_controller_test.go`.
Try to implement a controller integration test. The test code has some
minimal scaffolding and suggestion, feel free to ignore as you feel.

Use `make test` to run the tests.
To run just the controller tests, you can do
```
go test ./internal/controller/...
```
as usual.

When writing the test, take into account that
A. unless you delete data from the apiserver (e.g. Configuration objects),
   these are going to persist among testcases.
B. envtest, the environment on which controller integration test run,
   **does not support deletion of namespaces**. Just create and use a
   new namespace for each testcase 

## 3. (stretch) Familiarize with e2e tests using kind

Try to run the existing e2e tests. These have a specific makefile target
because they require a very complex setup. First, you need to setup a real
albeit local cluster against why the test run.
```
make deploy-on-kind
```

this does everything but run the tests: create the container images, create
the cluster, upload the images, deploy the controller in the cluster.

Go look around to familiarize with the test environments.
Which pods do run as baseline? how many nodes?

To cleanup your cluster and run from scratch, just do
```
make cleanup-test-e2e
```

## 4. (stretch) Write an e2e test

Run the e2e existing e2e tests.
If your cluster is up, or if you want to keep the cluster among runs
```
make run-test-e2e-golab
```
**Note** that any polluted state will persist across runs.

To run everything from scratch, including cluster setup and teardown:
```
make test-e2e-golab
```

The test is meant to fail. Try to complete it, or to write a new one

**Hint**: you have a full cluster *and* a docker environment at your disposal.
You can leverage both (e.g. docker exec...).
