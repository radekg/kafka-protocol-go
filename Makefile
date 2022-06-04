CURRENT_DIR=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: generate-request-types
generate-request-types:
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ProduceRequest > $(CURRENT_DIR)messages/00_produce_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=FetchRequest > $(CURRENT_DIR)messages/01_fetch_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ListOffsetsRequest > $(CURRENT_DIR)messages/02_listoffsets_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=MetadataRequest > $(CURRENT_DIR)messages/03_metadata_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=LeaderAndIsrRequest > $(CURRENT_DIR)messages/04_leaderandisr_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=StopReplicaRequest > $(CURRENT_DIR)messages/05_stopreplica_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=UpdateMetadataRequest > $(CURRENT_DIR)messages/06_updatemetadata_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ControlledShutdownRequest > $(CURRENT_DIR)messages/07_controlledshutdown_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetCommitRequest > $(CURRENT_DIR)messages/08_offsetcommit_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetFetchRequest > $(CURRENT_DIR)messages/09_offsetfetch_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=FindCoordinatorRequest > $(CURRENT_DIR)messages/10_findcoordinator_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=JoinGroupRequest > $(CURRENT_DIR)messages/11_joingroup_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=HeartbeatRequest > $(CURRENT_DIR)messages/12_heartbeat_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=LeaveGroupRequest > $(CURRENT_DIR)messages/13_leavegroup_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=SyncGroupRequest > $(CURRENT_DIR)messages/14_syncgroup_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeGroupsRequest > $(CURRENT_DIR)messages/15_describegroups_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ListGroupsRequest > $(CURRENT_DIR)messages/16_listgroups_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=SaslHandshakeRequest > $(CURRENT_DIR)messages/17_saslhandshake_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ApiVersionsRequest > $(CURRENT_DIR)messages/18_apiversions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=CreateTopicsRequest > $(CURRENT_DIR)messages/19_createtopics_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DeleteTopicsRequest > $(CURRENT_DIR)messages/20_deletetopics_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DeleteRecordsRequest > $(CURRENT_DIR)messages/21_deleterecords_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=InitProducerIdRequest > $(CURRENT_DIR)messages/22_initproducerid_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetForLeaderEpochRequest > $(CURRENT_DIR)messages/23_offsetforleaderepoch_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AddPartitionsToTxnRequest > $(CURRENT_DIR)messages/24_addpartitionstotxn_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AddOffsetsToTxnRequest > $(CURRENT_DIR)messages/25_addoffsetstotxn_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=EndTxnRequest > $(CURRENT_DIR)messages/26_endtxn_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=WriteTxnMarkersRequest > $(CURRENT_DIR)messages/27_writetxnmarkers_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=TxnOffsetCommitRequest > $(CURRENT_DIR)messages/28_txnoffsetcommit_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeAclsRequest > $(CURRENT_DIR)messages/29_describeacls_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=CreateAclsRequest > $(CURRENT_DIR)messages/30_createacls_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DeleteAclsRequest > $(CURRENT_DIR)messages/31_deleteacls_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeConfigsRequest > $(CURRENT_DIR)messages/32_describeconfigs_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterConfigsRequest > $(CURRENT_DIR)messages/33_alterconfigs_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterReplicaLogDirsRequest > $(CURRENT_DIR)messages/34_alterreplicalogdirs_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeLogDirsRequest > $(CURRENT_DIR)messages/35_describelogdirs_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=SaslAuthenticateRequest > $(CURRENT_DIR)messages/36_saslauthenticate_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=CreatePartitionsRequest > $(CURRENT_DIR)messages/37_createpartitions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=CreateDelegationTokenRequest > $(CURRENT_DIR)messages/38_createdelegationtoken_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=RenewDelegationTokenRequest > $(CURRENT_DIR)messages/39_renewdelegationtoken_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ExpireDelegationTokenRequest > $(CURRENT_DIR)messages/40_expiredelegationtoken_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeDelegationTokenRequest > $(CURRENT_DIR)messages/41_describedelegationtoken_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DeleteGroupsRequest > $(CURRENT_DIR)messages/42_deletegroups_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ElectLeadersRequest > $(CURRENT_DIR)messages/43_electleaders_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=IncrementalAlterConfigsRequest > $(CURRENT_DIR)messages/44_incrementalalterconfigs_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterPartitionReassignmentsRequest > $(CURRENT_DIR)messages/45_alterpartitionreassignments_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ListPartitionReassignmentsRequest > $(CURRENT_DIR)messages/46_listpartitionreassignments_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetDeleteRequest > $(CURRENT_DIR)messages/47_offsetdelete_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeClientQuotasRequest > $(CURRENT_DIR)messages/48_describeclientquotas_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterClientQuotasRequest > $(CURRENT_DIR)messages/49_alterclientquotas_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeUserScramCredentialsRequest > $(CURRENT_DIR)messages/50_describeuserscramcredentials_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterUserScramCredentialsRequest > $(CURRENT_DIR)messages/51_alteruserscramcredentials_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AlterPartitionRequest > $(CURRENT_DIR)messages/56_alterpartitions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=UpdateFeaturesRequest > $(CURRENT_DIR)messages/57_updatefeatures_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeClusterRequest > $(CURRENT_DIR)messages/60_describecluster_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeProducersRequest > $(CURRENT_DIR)messages/61_describeproducers_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=DescribeTransactionsRequest > $(CURRENT_DIR)messages/65_describetransactions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ListTransactionsRequest > $(CURRENT_DIR)messages/66_listtransactions_requests.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=AllocateProducerIdsRequest > $(CURRENT_DIR)messages/67_allocateproducerids_requests.go
	go fmt $(CURRENT_DIR)/messages/...

.PHONY: generate-response-types
generate-response-types:
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ProduceResponse > $(CURRENT_DIR)messages/00_produce_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=FetchResponse > $(CURRENT_DIR)messages/01_fetch_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ListOffsetsResponse > $(CURRENT_DIR)messages/02_listoffsets_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=MetadataResponse > $(CURRENT_DIR)messages/03_metadata_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=LeaderAndIsrResponse > $(CURRENT_DIR)messages/04_leaderandisr_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=StopReplicaResponse > $(CURRENT_DIR)messages/05_stopreplica_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=UpdateMetadataResponse > $(CURRENT_DIR)messages/06_updatemetadata_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=ControlledShutdownResponse > $(CURRENT_DIR)messages/07_controlledshutdown_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetCommitResponse > $(CURRENT_DIR)messages/08_offsetcommit_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=OffsetFetchResponse > $(CURRENT_DIR)messages/09_offsetfetch_responses.go
	go run ./codegen/... --kafka-source-root="${KAFKA_SOURCE_ROOT}" --gen-type=FindCoordinatorResponse > $(CURRENT_DIR)messages/10_findcoordinator_responses.go
	go fmt $(CURRENT_DIR)/messages/...