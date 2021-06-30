# integration-tests
This repository contains files related to integration tests for devfile.  

**NOTE:**  
All related issues are being tracked under the main devfile API repo https://github.com/devfile/api with the label `area/integration-tests`

# CI integration tests
Periodic integration tests are configured with [OpenShift CI system](https://docs.ci.openshift.org/docs/how-tos/onboarding-a-new-component/) and runs ODO and ODC tests with devfile.

## ODO CI tests
[ODO integration test cases](./scripts/odo/features/odo-devfile.feature)

| OCP version   |      Prow Test Status    |
|----------|:-------------:|
| 4.8 | [devfile-integration-tests-main-v4.8.odo](https://prow.ci.openshift.org/?job=periodic-ci-devfile-integration-tests-main-v4.8.odo-integration-devfile-odo-periodic) |
| 4.7 | [devfile-integration-tests-main-v4.7.odo](https://prow.ci.openshift.org/?job=periodic-ci-devfile-integration-tests-main-v4.7.odo-integration-devfile-odo-periodic) |
| 4.6 | [devfile-integration-tests-main-v4.6.odo](https://prow.ci.openshift.org/?job=periodic-ci-devfile-integration-tests-main-v4.6.odo-integration-devfile-odo-periodic) |

## ODC CI tests
[ODC integration test cases](./scripts/console/frontend/packages/dev-console/integration-tests/features/addFlow/create-from-devfile.feature)

| OCP version   |      Prow Test Status    |
|----------|:-------------:|
| 4.8 | [devfile-integration-tests-main-v4.8.console](https://prow.ci.openshift.org/?job=periodic-ci-devfile-integration-tests-main-v4.8.console-e2e-gcp-console-periodic) |
| 4.7 | [devfile-integration-tests-main-v4.7.console](https://prow.ci.openshift.org/?job=periodic-ci-devfile-integration-tests-main-v4.7.console-e2e-gcp-console-periodic) |

# Local integration tests
**NOTE**: This document covers setup for macOS specifically however the similar steps can be used for other OSes.

## ODO Tests
Tests in this repository are based on [ODO integration tests](https://github.com/openshift/odo/blob/main/docs/dev/test-architecture.adoc#integration-and-e2e-tests) and it mainly focuses on custom test cases with devfile service.

### Prerequisites
- Go 1.13 and Ginkgo latest version
- git
- [OpenShift Cluster](https://github.com/openshift/odo/blob/main/docs/dev/test-architecture.adoc#integration-and-e2e-tests).  e.g. crc environment for 4.* local cluster 
- [Optional] [xunit-viewer](https://www.npmjs.com/package/xunit-viewer)
  and [jrm](https://www.npmjs.com/package/junit-report-merger?activeTab=readme) : required to get performance test results in a merged html format in addition to `junit*.xml`.

### Run tests
1. cd `local/odo/tests`
1. run `./odo-integration-tests.sh`  : it runs `odo catlog command` test by default. In order to run other test cases, modify `./odo-integration-tests.sh` by enabling other test options. e.g. `make test-cmd-devfile-create`

### Performance tests for ODO
1. Open `Makefile` and remove `--skipMeasurements` option from `GINKGO_FLAGS_ALL` flag.
1. Add [Ginkgo's Measure block](https://onsi.github.io/ginkgo/#benchmark-tests) for the spec that you want to measure the performance. A sample `Measure` block can be found from `https://github.com/devfile/integration-tests/tree/main/local/odo/tests/integration/devfile/cmd_devfile_catalog_test.go`
1. [Optional] In order to get test results in html format, uncomment `jrm...` and `xunit-viewer...` command calls from `odo-integration-tests.sh`. 
1. run `./odo-integration-tests.sh`
   
    A sample performance test output in console.
   
    ![alt text](./docs/images/perf_measure_sample.png "Performance test result")

    A sample test output in html format

   ![alt text](./docs/images/perf_html_sample.png "Performance test result")

## ODC Tests
Tests in this repository are based on [ODC integration tests](https://github.com/openshift/console#integration-tests). 

### Prerequisites
- Cypress integration tests are implemented in [Cypress.io](https://www.cypress.io/).

See https://github.com/openshift/console#integration-tests for more detail on how to run devfile test with ODC. 
