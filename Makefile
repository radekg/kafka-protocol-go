.PHONY: generate-request-types
generate-request-types:
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ProduceRequest > messages/00_produce_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=FetchRequest > messages/01_fetch_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ListOffsetsRequest > messages/02_listoffsets_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=MetadataRequest > messages/03_metadata_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetCommitRequest > messages/08_offsetcommit_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetFetchRequest > messages/09_offsetfetch_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=FindCoordinatorRequest > messages/10_findcoordinator_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=HeartbeatRequest > messages/12_heartbeat_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ApiVersionsRequest > messages/18_apiversions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeUserScramCredentialsRequest > messages/50_describeuserscramcredentials_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterUserScramCredentialsRequest > messages/51_alteruserscramcredentials_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterPartitionRequest > messages/56_alterpartitions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=UpdateFeaturesRequest > messages/57_updatefeatures_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeClusterRequest > messages/60_describecluster_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeProducersRequest > messages/61_describeproducers_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeTransactionsRequest > messages/65_describetransactions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ListTransactionsRequest > messages/66_listtransactions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AllocateProducerIdsRequest > messages/67_allocateproducerids_requests.go