#!/bin/bash

version=$(cat version.txt)
git tag "v${version}"
git push && git push --tags