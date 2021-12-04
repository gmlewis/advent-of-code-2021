#!/bin/bash -ex
for i in day*; do
  pushd "${i}" && ./run-all.sh && popd
done
