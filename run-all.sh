#!/bin/bash -ex
go test ./...

for i in day*; do
  pushd "${i}" && ./run-all.sh && popd
done

echo "Success!"
