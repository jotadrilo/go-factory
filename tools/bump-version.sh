#!/bin/bash

root=$(git rev-parse --show-toplevel)
version="${root}/VERSION"
readme="${root}/README.md"

old_version_string="$(cat "$version")"
major=$(cut -d. -f1 "$version")
minor=$(cut -d. -f2 "$version")
patch=$(cut -d. -f3 "$version")
new_version_string="$major.$minor.$((patch+1))"

# Update VERSION
echo "$new_version_string" >"$version"

# Update README.md
gsed -i "s;download/${old_version_string}/go-factory_${old_version_string}_;download/${new_version_string}/go-factory_${new_version_string}_;g" "$readme"

git commit -am "Bump VERSION to $new_version_string"
