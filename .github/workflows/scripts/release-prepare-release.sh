#!/bin/bash -ex

if [ "${CANDIDATE_STABLE}" == "" ] && [ "${CANDIDATE_BETA}" == "" ]; then
    echo "One of CANDIDATE_STABLE or CANDIDATE_BETA must be set"
    exit 1
fi

RELEASE_VERSION=v${CANDIDATE_STABLE}/v${CANDIDATE_BETA}
if [ "${CANDIDATE_STABLE}" == "" ]; then
    RELEASE_VERSION=v${CURRENT_STABLE}/v${CANDIDATE_BETA}
elif [ "${CANDIDATE_BETA}" == "" ]; then
    RELEASE_VERSION=v${CANDIDATE_STABLE}/v${CURRENT_BETA}
fi

make chlog-update VERSION="${RELEASE_VERSION}"
git config user.name opentelemetrybot
git config user.email 107717825+opentelemetrybot@users.noreply.github.com
BRANCH="prepare-release-prs/${CANDIDATE_BETA}"
git checkout -b "${BRANCH}"
git add --all
git commit -m "Changelog update ${RELEASE_VERSION}"

if [ "${CANDIDATE_STABLE}" != "" ]; then
    make prepare-release GH=none PREVIOUS_VERSION="${CURRENT_STABLE}" RELEASE_CANDIDATE="${CANDIDATE_STABLE}" MODSET=stable
fi
if [ "${CANDIDATE_BETA}" != "" ]; then
    make prepare-release GH=none PREVIOUS_VERSION="${CURRENT_BETA}" RELEASE_CANDIDATE="${CANDIDATE_BETA}" MODSET=beta
fi
git push origin "${BRANCH}"

gh pr create --title "[chore] Prepare release ${RELEASE_VERSION}" --body "
The following commands were run to prepare this release:
- make chlog-update VERSION=${RELEASE_VERSION}
- make prepare-release GH=none PREVIOUS_VERSION=${CURRENT_STABLE} RELEASE_CANDIDATE=${CANDIDATE_STABLE} MODSET=stable
- make prepare-release GH=none PREVIOUS_VERSION=${CURRENT_BETA} RELEASE_CANDIDATE=${CANDIDATE_BETA} MODSET=beta
"
