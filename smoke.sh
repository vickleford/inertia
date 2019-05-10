#!/usr/bin/env bash

actual=$(go run main.go \
	-set image=jabba:thehut \
	-set noun=sith \
	-set adjective=fat \
	-b64set "something=secrets of the force" \
	rendering/testdata/simple.tpl)

expected="jabba:thehut has a big ole fat sith inside and knows c2VjcmV0cyBvZiB0aGUgZm9yY2U="

if [[ "$actual" != "$expected" ]]; then
	echo "FAILURE!"
	echo "Wanted: $expected"
	echo "Got: $actual"
	exit 1
fi

echo "PASS"
