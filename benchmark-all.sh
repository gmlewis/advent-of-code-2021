#!/bin/bash -ex
go test ./...

for i in day*/part*; do
  pushd "${i}" && go test ./... -bench=. && popd
done

echo "Success!"
