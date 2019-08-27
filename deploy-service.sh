#! /bin/bash
set -eu

REPO=$(echo $1 | cut -d "#" -f 1)
BRANCH=$(echo $1 | cut -d "#" -f 2)
echo "Deploying $REPO from branch $BRANCH to $(kubectl config current-context)"

SHA=$(git ls-remote https://github.com/giantswarm/$REPO | grep $BRANCH | awk '{ print $1 }')
if [ -z "$SHA" ]
then
  echo "Error: branch '$BRANCH' doesn't exist on remote repository"
  exit 1
fi

BASE_URL=https://quay.io/cnr/api/v1/packages/giantswarm
REQUEST=$BASE_URL/$REPO-chart/channels/wip-$SHA

STATUS=$(curl -I -XGET $REQUEST  2> /dev/null | head -n 1| cut -d$' ' -f2)
if [ "$STATUS" -eq 404 ]
then
  echo -e "Error: Channel wip-$SHA hasn't been pushed to quay.io.\nCheck https://circleci.com/gh/giantswarm/$REPO if your build was successful"
  exit 1
fi

PATCH='{"spec":{"chart":{"channel":"wip-'$SHA'"}},"metadata":{"labels":{"giantswarm.io/managed-by":"'$USER'"}}}'
if kubectl cluster-info 1> /dev/null ; then
  kubectl patch chartconfigs.core.giantswarm.io $REPO-chart --type merge -n giantswarm --patch $PATCH
else
  echo "Error: Can't connect to cluster via kubectl, make sure you created the kubeconfig of the tenant cluster"
  exit 1
fi

