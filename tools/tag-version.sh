#!/bin/bash

root=$(git rev-parse --show-toplevel)
version="$(cat "${root}/VERSION")"
git tag "${version}" -m "${version}"
