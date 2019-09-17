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

actual=$(go run main.go \
    -set image=praise \
    -set nount=the \
    -set adjective=sun \
    -b64set "something=lookatme" \
    rendering/testdata/simple.tpl)

expected="praise has a big ole sun the inside and knows aHR0cHM6Ly9ncm5oLnNlLzcwOWI5MzBlMQ=="

if [[ "${actual}" == "${expected}" ]]; then
    echo "FAILURE!"
    echo "There was something seriously wrong, we got something completely unexpected"
    echo "Expected: ${expected}"
    echo "Got: ${actual}"
    exit 1
fi

actual=$(go run main.go 2>&1 | grep Usage)

expected="Usage: inertia [ -set templatekey=value ... ] [ -b64set templatekey=value ... ] /path/to/template"

if [[ "$actual" != "$expected" ]]; then
	echo "FAILURE!"
	echo "Wanted: $expected"
	echo "Got: $actual"
	exit 1
fi

echo "PASS"
