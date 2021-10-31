#!/usr/bin/env bash

echo "Provisioning Cluster"
scripts_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
bash "${scripts_dir}"/provision.sh 3 >/dev/null

echo "Cluster Sanity Tests"
bats --tap "${scripts_dir}"/bats/cluster-sanity.bats

echo "Sort Tests"
bats --tap "${scripts_dir}"/bats/sort.bats

echo "Large Sort Tests"
bats --tap "${scripts_dir}"/bats/sort-large.bats

echo "Tearing Down Cluster"
bash "${scripts_dir}"/teardown.sh >/dev/null
