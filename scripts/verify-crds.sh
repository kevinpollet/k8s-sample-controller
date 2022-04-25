#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
DIFF_ROOT=${SCRIPT_ROOT}/deploy/crds
TMP_DIFF_ROOT=${SCRIPT_ROOT}/_tmp/deploy/crds
_tmp=${SCRIPT_ROOT}/_tmp

cleanup() {
  rm -rf "${_tmp}"
}
trap "cleanup" EXIT SIGINT

cleanup

mkdir -p "${TMP_DIFF_ROOT}"
cp -a "${DIFF_ROOT}"/* "${TMP_DIFF_ROOT}"

"${SCRIPT_ROOT}/scripts/update-crds.sh"
echo "Diffing ${DIFF_ROOT} against freshly generated CRDs"
ret=0
diff -Naupr "${DIFF_ROOT}" "${TMP_DIFF_ROOT}" || ret=$?
cp -a "${TMP_DIFF_ROOT}"/* "${DIFF_ROOT}"
if [[ $ret -eq 0 ]]; then
  echo "${DIFF_ROOT} up to date."
else
  echo "${DIFF_ROOT} is out of date. Please run scripts/update-crds.sh"
  exit 1
fi
