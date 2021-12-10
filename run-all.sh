#!/bin/bash -ex
go test ./...

for i in day*; do
  pushd "${i}" && ./run-all.sh && popd || exit 1
done

echo "Success!"
