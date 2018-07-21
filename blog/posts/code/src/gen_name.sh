#!/bin/bash

mkdir -p person
cd person

cat <<EOF > person.go
package person

type Person struct {
    Name string
    Age int
}
EOF
