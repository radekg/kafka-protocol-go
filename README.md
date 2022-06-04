# Apache Kafka® Protocol for go and tinygo

This is a fork of the Apache Kafka® protocol implementation from [grepplabs/kafka-proxy](https://github.com/grepplabs/kafka-proxy). This fantastic work is a great donor because it:

- Uses minimal reflection.
- Does not use struct tags.
- Does not use goroutines.
- Does not use channels.

Which results in:

- Being able to compile with [tinygo-org/tinygo](https://github.com/tinygo-org/tinygo).
- Because it can be compiled with _tinygo_, it can be compiled with [tetratelabs/proxy-wasm-go-sdk](https://github.com/tetratelabs/proxy-wasm-go-sdk).
  - Which means that it can be used as a foundation for Envoy WebAssembly Layer 7 proxies.

## Why forking?

The upstream Kafka Proxy protocol implements sufficient encoding and decoding to support its own functionality, but does not implement everything else required to parse messages beyond request headers, response headers, `Metadata`, and `FindCoordinator`. The maintainer of Kafka Proxy isn't keen on supporting anything in the protocol encoder/decoder that does not find a direct use in the upstream project.

The goal of this fork is to add support for all Apache Kafka® messages.

## Message support

- [x] Headers

### General notes

- Compact arrays have been added in Apache Kafka® 2.4, hence it will be crucial to test with 2.3 and 2.4 client to get the differences between arrays and compact arrays right.

| Requests                                 | Responses                                |
| ---------------------------------------- | ---------------------------------------- |
| - [x] `0`: Produce                       | - [ ] `0`: Produce                       |
| - [x] `1`: Fetch                         | - [ ] `1`: Fetch                         |
| - [x] `2`: ListOffsets                   | - [ ] `2`: ListOffsets                   |
| - [x] `3`: Metadata                      | - [ ] `3`: Metadata                      |
| - [x] `3`: Metadata                      | - [ ] `3`: Metadata                      |
| - [x] `5`: StopReplica                   | - [ ] `5`: StopReplica                   |
| - [x] `6`: UpdateMetadata                | - [ ] `6`: UpdateMetadata                |
| - [x] `7`: ControlledShutdown            | - [ ] `7`: ControlledShutdown            |
| - [x] `8`: OffsetCommit                  | - [ ] `8`: OffsetCommit                  |
| - [x] `9`: OffsetFetch                   | - [ ] `9`: OffsetFetch                   |
| - [x] `10`: FindCoordinator              | - [ ] `10`: FindCoordinator              |
| - [x] `11`: JoinGroup                    | - [ ] `11`: JoinGroup                    |
| - [x] `12`: Heartbeat                    | - [ ] `12`: Heartbeat                    |
| - [x] `13`: LeaveGroup                   | - [ ] `13`: LeaveGroup                   |
| - [x] `14`: SyncGroup                    | - [ ] `14`: SyncGroup                    |
| - [x] `15`: DescribeGroups               | - [ ] `15`: DescribeGroups               |
| - [x] `16`: ListGroups                   | - [ ] `16`: ListGroups                   |
| - [x] `17`: SaslHandshake                | - [ ] `17`: SaslHandshake                |
| - [x] `18`: ApiVersions                  | - [ ] `18`: ApiVersions                  |
| - [x] `19`: CreateTopics                 | - [ ] `19`: CreateTopics                 |
| - [x] `20`: DeleteTopics                 | - [ ] `20`: DeleteTopics                 |
| - [x] `21`: DeleteRecords                | - [ ] `21`: DeleteRecords                |
| - [x] `22`: InitProducerId               | - [ ] `22`: InitProducerId               |
| - [x] `23`: OffsetForLeaderEpoch         | - [ ] `23`: OffsetForLeaderEpoch         |
| - [x] `24`: AddPartitionsToTxn           | - [ ] `24`: AddPartitionsToTxn           |
| - [x] `25`: AddOffsetsToTxn              | - [ ] `25`: AddOffsetsToTxn              |
| - [x] `26`: EndTxn                       | - [ ] `26`: EndTxn                       |
| - [x] `27`: WriteTxnMarkers              | - [ ] `27`: WriteTxnMarkers              |
| - [x] `28`: TxnOffsetCommit              | - [ ] `28`: TxnOffsetCommit              |
| - [x] `29`: DescribeAcls                 | - [ ] `29`: DescribeAcls                 |
| - [x] `30`: CreateAcls                   | - [ ] `30`: CreateAcls                   |
| - [x] `31`: DeleteAcls                   | - [ ] `31`: DeleteAcls                   |
| - [x] `32`: DescribeConfigs              | - [ ] `32`: DescribeConfigs              |
| - [x] `33`: AlterConfigs                 | - [ ] `33`: AlterConfigs                 |
| - [x] `34`: AlterReplicaLogDirs          | - [ ] `34`: AlterReplicaLogDirs          |
| - [x] `35`: DescribeLogDirs              | - [ ] `35`: DescribeLogDirs              |
| - [x] `36`: SaslAuthenticate             | - [ ] `36`: SaslAuthenticate             |
| - [x] `37`: CreatePartitions             | - [ ] `37`: CreatePartitions             |
| - [x] `38`: CreateDelegationToken        | - [ ] `38`: CreateDelegationToken        |
| - [x] `39`: RenewDelegationToken         | - [ ] `39`: RenewDelegationToken         |
| - [x] `40`: ExpireDelegationToken        | - [ ] `40`: ExpireDelegationToken        |
| - [x] `41`: DescribeDelegationToken      | - [ ] `41`: DescribeDelegationToken      |
| - [x] `42`: DeleteGroups                 | - [ ] `42`: DeleteGroups                 |
| - [x] `43`: ElectLeaders                 | - [ ] `43`: ElectLeaders                 |
| - [x] `44`: IncrementalAlterConfigs      | - [ ] `44`: IncrementalAlterConfigs      |
| - [x] `45`: AlterPartitionReassignments  | - [ ] `45`: AlterPartitionReassignments  |
| - [x] `46`: ListPartitionReassignments   | - [ ] `46`: ListPartitionReassignments   |
| - [x] `47`: OffsetDelete                 | - [ ] `47`: OffsetDelete                 |
| - [x] `48`: DescribeClientQuotas         | - [ ] `48`: DescribeClientQuotas         |
| - [x] `49`: AlterClientQuotas            | - [ ] `49`: AlterClientQuotas            |
| - [x] `50`: DescribeUserScramCredentials | - [ ] `50`: DescribeUserScramCredentials |
| - [x] `51`: AlterUserScramCredentials    | - [ ] `51`: AlterUserScramCredentials    |
| - [x] `56`: AlterPartition               | - [ ] `56`: AlterPartition               |
| - [x] `57`: UpdateFeatures               | - [ ] `57`: UpdateFeatures               |
| - [x] `60`: DescribeCluster              | - [ ] `60`: DescribeCluster              |
| - [x] `61`: DescribeProducers            | - [ ] `61`: DescribeProducers            |
| - [x] `65`: DescribeTransactions         | - [ ] `65`: DescribeTransactions         |
| - [x] `66`: ListTransactions             | - [ ] `66`: ListTransactions             |
| - [x] `67`: AllocateProducerIds          | - [ ] `67`: AllocateProducerIds          |

## License 

Apache 2.0.

## Run tests

```sh
go test -count=1 -v ./...
```

## Generating messages

1. Setup working environment:

```sh
export KAFKA_SOURCE_ROOT="${HOME}/dev/my/kafka"
```

2. Close Kafka sources somewhere on disk and checkout the release you want to use:

```sh
mkdir -p "${KAFKA_SOURCE_ROOT}"
cd "${KAFKA_SOURCE_ROOT}"
git clone https://github.com/apache/kafka.git .
git checkout 3.2.0
cd -
```

3. Run the generator command:

```sh
make generate-request-types
```
