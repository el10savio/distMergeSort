# distMergeSort

An implementation of mergesort distributed across nodes used to sort large sets.

## Introduction

Merge sort partitions sets so that they can be recursively sorted and then merged back to form a single sorted set. They can be either split across several threads in a node or across multiple nodes across a network.

## But why not just use QuickSort?

Quicksort is indeed much faster than merge sort, but is a hassle when the sets to be sorted gets too big (>100M elements) that it does not fit in a single node's memory. The following implementation aims to tackle this by splitting these large sets and distributing them across nodes in a network. A trick here is we dont recursively split them until it reaches two elements, but split them until they reach a certain count so that they can then be sorted using quicksort and then be sent back to the sender node to be merged.

## Example

```
$ curl -X POST localhost:8001/sort -d {"values": [5,4,3,2,1]} => [1,2,3,4,5]
```

## Steps

To provision the cluster:

```
$ git clone https://github.com/el10savio/distMergeSort
$ cd distMergeSort
$ make provision
```

This creates a 3 node sort cluster established in their own docker network.

To view the status of the cluster

```
$ make info
```

This provides information on the cluster and its associated ports to access each node. An example of the output seen in `make info` would be like below:

```
d3fd26dd4df3  sort  "/go/bin/sort"  2 hours ago  Up 2 hours  0.0.0.0:8004->8080/tcp  peer-1
8830feb6cd68  sort  "/go/bin/sort"  2 hours ago  Up 2 hours  0.0.0.0:8003->8080/tcp  peer-0
```


Now we can also send requests to sort values from any peer node using its port allocated.

```
$ curl -i -X POST localhost:<peer-port>/sort -d {"values": <values>}
```

In the logs for each peer docker container, we can see the logs of the peer nodes partitioning and sorting the sent list.

To tear down the cluster and remove the built docker images:

```
$ make clean
```

This is not certain to clean up all the locally created docker images at times. You can do a `docker rmi` to manually delete them.

## Testing

To provision the cluster and run automated end to end tests you can use `make e2e`. This uses BATS bash testing to run curl requests to each node and asserts the output received.

```
$ make e2e
Running E2E Testing On Sort Cluster
bash scripts/tests.sh
Provisioning Cluster
Cluster Sanity Tests
1..2
ok 1 Check Replicas Count
ok 2 Check Replicas Are Available
Sort Tests
1..3
ok 1 Sort Empty List
ok 2 Sort Basic List
ok 3 Sort Basic List II
Large Sort Tests
1..1
ok 1 Sort Large List
Tearing Down Cluster
```
