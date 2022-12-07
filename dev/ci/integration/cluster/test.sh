#!/usr/bin/env bash
set -euo pipefail

# setup DIR for easier pathing test dir
test_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)""

# cd to repo root
root_dir="$(dirname "${BASH_SOURCE[0]}")/../../../.."
cd "$root_dir"
root_dir=$(pwd)

export NAMESPACE="cluster-ci-$BUILDKITE_BUILD_NUMBER-$BUILDKITE_JOB_ID"

# Capture information about the state of the test cluster
function cluster_capture_state() {
  # Get some more verbobe information about what is running.
  set -x

  echo "--- dump diagnostics"
  # The reason have the grep here and filter out the otel-agents in Pending state is due to how otel-agents
  # are scheduled. The otel agent is deployed using a DaemonSet, which means on every node, k8s will schedule a
  # otel-agent, even if the node isn't running anything else. For this QA scenario we don't want to run anything
  # more than what we want, so if you look in deploy-sourcegraph/overlays/otel-agent-patch.yaml you'll see
  # the otel-agent DaemonSet is patched with a podAffinity. PodAffinity ensures that a Pod will only be scheduled
  # that matches a certain condition, if the Pod doesn't match it's status will be PENDING - hence we filter them
  # out
  # Get overview of all pods
  kubectl get pods | grep -v -e "otel-agent-.*Pending"

  # Get specifics of pods
  kubectl describe pods >"$root_dir/describe_pods.log" 2>&1

  # Get logs for some deployments
  IFS=' ' read -ra deployments <<< "$(kubectl get deployments -o=jsonpath='{.items[*].metadata.name}')"
  for dep in "${deployments[@]}"; do
    kubectl logs "deployment/$dep" --all-containers --previous >"$root_dir/$dep.log" 2>&1
  done
  set +x
}

# Cleanup the cluster
function cluster_cleanup() {
  cluster_capture_state || true
  kubectl delete namespace "$NAMESPACE"
}

function cluster_setup() {
  gcloud container clusters get-credentials default-buildkite --zone=us-central1-c --project=sourcegraph-ci

  echo "--- create namespace"
  kubectl create ns "$NAMESPACE" -oyaml --dry-run=client | kubectl apply -f -
  trap cluster_cleanup exit

  echo "--- create storageclass"
  kubectl apply -f "$test_dir/storageClass.yaml"
  kubectl config set-context --current --namespace="$NAMESPACE"
  kubectl config current-context
  echo "--- wait for namespace to come up and check pods"
  sleep 15 # wait for namespace to come up
  kubectl get -n "$NAMESPACE" pods

  echo "--- rewrite manifests"
  pushd "$test_dir/deploy-sourcegraph"
  set +e
  set +o pipefail
  # See $DOCKER_CLUSTER_IMAGES_TXT in pipeline-steps.go for env var
  # replace all docker image tags with previously built candidate images
  while IFS= read -r line; do
    echo "$line"
    grep -lr './base/' -e "index.docker.io/sourcegraph/$line" --include \*.yaml | xargs sed -i -E "s#index.docker.io/sourcegraph/$line:.*#us.gcr.io/sourcegraph-dev/$line:$CANDIDATE_VERSION#g"
  done < <(printf '%s\n' "$DOCKER_CLUSTER_IMAGES_TXT")

  echo "--- create cluster"
  ./overlay-generate-cluster.sh low-resource generated-cluster
  kubectl apply -n "$NAMESPACE" --recursive --validate -f generated-cluster
  popd
  echo "--- wait for ready"
  sleep 15 #add in a small wait for all pods to be rolled out by the replication controller
  kubectl get pods -n "$NAMESPACE"
  time kubectl wait --for=condition=Ready -l app=sourcegraph-frontend pod --timeout=5m -n "$NAMESPACE"
  set -e
  set -o pipefail
}

function test_setup() {

  set +x +u
  # shellcheck disable=SC1091
  source /root/.profile

  dev/ci/integration/setup-deps.sh

  sleep 15
  export SOURCEGRAPH_BASE_URL="http://sourcegraph-frontend.$NAMESPACE.svc.cluster.local:30080"
  curl "$SOURCEGRAPH_BASE_URL"

  # setup admin users, etc
  pushd internal/cmd/init-sg
  go build
  ./init-sg initSG -baseurl="$SOURCEGRAPH_BASE_URL"
  popd

  # Load variables set up by init-server, disabling `-x` to avoid printing variables, setting +u to avoid blowing up on ubound ones
  set +x +u
  # shellcheck disable=SC1091
  source /root/.sg_envrc
  set -u

  echo "--- TEST: Checking Sourcegraph instance is accessible"

  curl --fail "$SOURCEGRAPH_BASE_URL"
  curl --fail "$SOURCEGRAPH_BASE_URL/healthz"
}

function e2e() {
  pushd client/web
  echo "$SOURCEGRAPH_BASE_URL"
  echo "--- TEST: Running tests"
  yarn run test:regression:core
  popd
}

# main
cluster_setup
test_setup
set +o pipefail
# special exit code to capture e2e failures
e2e || exit 123
