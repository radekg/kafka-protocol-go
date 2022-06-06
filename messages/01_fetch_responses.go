package messages

import "github.com/radekg/kafka-protocol-go/schema"

func init1FetchResponse() []schema.Schema {

	return []schema.Schema{
		// Message: FetchResponse, API Key: 1, Version: 0
		schema.NewSchema("FetchResponse:v0",
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v0",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v0",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 1
		schema.NewSchema("FetchResponse:v1",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v1",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v1",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 2
		schema.NewSchema("FetchResponse:v2",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v2",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v2",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 3
		schema.NewSchema("FetchResponse:v3",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v3",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v3",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 4
		schema.NewSchema("FetchResponse:v4",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v4",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v4",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v4",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 5
		schema.NewSchema("FetchResponse:v5",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v5",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v5",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v5",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 6
		schema.NewSchema("FetchResponse:v6",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v6",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v6",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v6",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 7
		schema.NewSchema("FetchResponse:v7",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFetchResponseSessionId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v7",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v7",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v7",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 8
		schema.NewSchema("FetchResponse:v8",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFetchResponseSessionId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v8",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v8",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v8",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 9
		schema.NewSchema("FetchResponse:v9",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFetchResponseSessionId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v9",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v9",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v9",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 10
		schema.NewSchema("FetchResponse:v10",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFetchResponseSessionId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v10",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v10",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v10",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 11
		schema.NewSchema("FetchResponse:v11",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFetchResponseSessionId, Ty: schema.TypeInt32},
			&schema.Array{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v11",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStr},
				&schema.Array{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v11",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.Array{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v11",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPreferredReadReplica, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesNullable},
				)},
			)},
		),

		// Message: FetchResponse, API Key: 1, Version: 12
		schema.NewSchema("FetchResponse:v12",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFetchResponseSessionId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v12",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopic, Ty: schema.TypeStrCompact},
				&schema.ArrayCompact{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v12",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.ArrayCompact{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v12",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
						&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsTags},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPreferredReadReplica, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsTags},
					/** Applicable tags:

						0: DivergingEpoch =
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsDivergingEpoch, Ty: schema.NewSchema("DivergingEpoch:v12",
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsDivergingEpochEpoch, Ty: schema.TypeInt32},
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsDivergingEpochEndOffset, Ty: schema.TypeInt64},
							&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsDivergingEpochTags},

						)},

						1: CurrentLeader =
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsCurrentLeader, Ty: schema.NewSchema("CurrentLeader:v12",
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderId, Ty: schema.TypeInt32},
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderEpoch, Ty: schema.TypeInt32},
							&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsCurrentLeaderTags},

						)},

						2: SnapshotId =
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsSnapshotId, Ty: schema.NewSchema("SnapshotId:v12",
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsSnapshotIdEndOffset, Ty: schema.TypeInt64},
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsSnapshotIdEpoch, Ty: schema.TypeInt32},
							&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsSnapshotIdTags},

						)},

					**/

				)},
				&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldFetchResponseTags},
		),

		// Message: FetchResponse, API Key: 1, Version: 13
		schema.NewSchema("FetchResponse:v13",
			&schema.Mfield{Name: FieldFetchResponseThrottleTimeMs, Ty: schema.TypeInt32},
			&schema.Mfield{Name: FieldFetchResponseErrorCode, Ty: schema.TypeInt16},
			&schema.Mfield{Name: FieldFetchResponseSessionId, Ty: schema.TypeInt32},
			&schema.ArrayCompact{Name: FieldFetchResponseResponses, Ty: schema.NewSchema("Responses:v13",
				&schema.Mfield{Name: FieldFetchResponseResponsesTopicId, Ty: schema.TypeUuid},
				&schema.ArrayCompact{Name: FieldFetchResponseResponsesPartitions, Ty: schema.NewSchema("Partitions:v13",
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPartitionIndex, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsErrorCode, Ty: schema.TypeInt16},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsHighWatermark, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLastStableOffset, Ty: schema.TypeInt64},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsLogStartOffset, Ty: schema.TypeInt64},
					&schema.ArrayCompact{Name: FieldFetchResponseResponsesPartitionsAbortedTransactions, Ty: schema.NewSchema("AbortedTransactions:v13",
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId, Ty: schema.TypeInt64},
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset, Ty: schema.TypeInt64},
						&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsAbortedTransactionsTags},
					)},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsPreferredReadReplica, Ty: schema.TypeInt32},
					&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsRecords, Ty: schema.TypeBytesCompactNullable},
					&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsTags},
					/** Applicable tags:

						0: DivergingEpoch =
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsDivergingEpoch, Ty: schema.NewSchema("DivergingEpoch:v13",
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsDivergingEpochEpoch, Ty: schema.TypeInt32},
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsDivergingEpochEndOffset, Ty: schema.TypeInt64},
							&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsDivergingEpochTags},

						)},

						1: CurrentLeader =
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsCurrentLeader, Ty: schema.NewSchema("CurrentLeader:v13",
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderId, Ty: schema.TypeInt32},
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderEpoch, Ty: schema.TypeInt32},
							&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsCurrentLeaderTags},

						)},

						2: SnapshotId =
						&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsSnapshotId, Ty: schema.NewSchema("SnapshotId:v13",
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsSnapshotIdEndOffset, Ty: schema.TypeInt64},
							&schema.Mfield{Name: FieldFetchResponseResponsesPartitionsSnapshotIdEpoch, Ty: schema.TypeInt32},
							&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesPartitionsSnapshotIdTags},

						)},

					**/

				)},
				&schema.SchemaTaggedFields{Name: FieldFetchResponseResponsesTags},
			)},
			&schema.SchemaTaggedFields{Name: FieldFetchResponseTags},
		),
	}

}

const (

	// FieldFetchResponseErrorCode is: The top level response error code.
	FieldFetchResponseErrorCode = "ErrorCode"

	// FieldFetchResponseResponses is: The response topics.
	FieldFetchResponseResponses = "Responses"

	// FieldFetchResponseResponsesPartitions is: The topic partitions.
	FieldFetchResponseResponsesPartitions = "Partitions"

	// FieldFetchResponseResponsesPartitionsAbortedTransactions is: The aborted transactions.
	FieldFetchResponseResponsesPartitionsAbortedTransactions = "AbortedTransactions"

	// FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset is: The first offset in the aborted transaction.
	FieldFetchResponseResponsesPartitionsAbortedTransactionsFirstOffset = "FirstOffset"

	// FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId is: The producer id associated with the aborted transaction.
	FieldFetchResponseResponsesPartitionsAbortedTransactionsProducerId = "ProducerId"

	// FieldFetchResponseResponsesPartitionsAbortedTransactionsTags is: The tagged fields.
	FieldFetchResponseResponsesPartitionsAbortedTransactionsTags = "Tags"

	// FieldFetchResponseResponsesPartitionsCurrentLeader is:
	FieldFetchResponseResponsesPartitionsCurrentLeader = "CurrentLeader"

	// FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderEpoch is: The latest known leader epoch
	FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderEpoch = "LeaderEpoch"

	// FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderId is: The ID of the current leader or -1 if the leader is unknown.
	FieldFetchResponseResponsesPartitionsCurrentLeaderLeaderId = "LeaderId"

	// FieldFetchResponseResponsesPartitionsCurrentLeaderTags is: The tagged fields.
	FieldFetchResponseResponsesPartitionsCurrentLeaderTags = "Tags"

	// FieldFetchResponseResponsesPartitionsDivergingEpoch is: In case divergence is detected based on the `LastFetchedEpoch` and `FetchOffset` in the request, this field indicates the largest epoch and its end offset such that subsequent records are known to diverge
	FieldFetchResponseResponsesPartitionsDivergingEpoch = "DivergingEpoch"

	// FieldFetchResponseResponsesPartitionsDivergingEpochEndOffset is:
	FieldFetchResponseResponsesPartitionsDivergingEpochEndOffset = "EndOffset"

	// FieldFetchResponseResponsesPartitionsDivergingEpochEpoch is:
	FieldFetchResponseResponsesPartitionsDivergingEpochEpoch = "Epoch"

	// FieldFetchResponseResponsesPartitionsDivergingEpochTags is: The tagged fields.
	FieldFetchResponseResponsesPartitionsDivergingEpochTags = "Tags"

	// FieldFetchResponseResponsesPartitionsErrorCode is: The error code, or 0 if there was no fetch error.
	FieldFetchResponseResponsesPartitionsErrorCode = "ErrorCode"

	// FieldFetchResponseResponsesPartitionsHighWatermark is: The current high water mark.
	FieldFetchResponseResponsesPartitionsHighWatermark = "HighWatermark"

	// FieldFetchResponseResponsesPartitionsLastStableOffset is: The last stable offset (or LSO) of the partition. This is the last offset such that the state of all transactional records prior to this offset have been decided (ABORTED or COMMITTED)
	FieldFetchResponseResponsesPartitionsLastStableOffset = "LastStableOffset"

	// FieldFetchResponseResponsesPartitionsLogStartOffset is: The current log start offset.
	FieldFetchResponseResponsesPartitionsLogStartOffset = "LogStartOffset"

	// FieldFetchResponseResponsesPartitionsPartitionIndex is: The partition index.
	FieldFetchResponseResponsesPartitionsPartitionIndex = "PartitionIndex"

	// FieldFetchResponseResponsesPartitionsPreferredReadReplica is: The preferred read replica for the consumer to use on its next fetch request
	FieldFetchResponseResponsesPartitionsPreferredReadReplica = "PreferredReadReplica"

	// FieldFetchResponseResponsesPartitionsRecords is: The record data.
	FieldFetchResponseResponsesPartitionsRecords = "Records"

	// FieldFetchResponseResponsesPartitionsSnapshotId is: In the case of fetching an offset less than the LogStartOffset, this is the end offset and epoch that should be used in the FetchSnapshot request.
	FieldFetchResponseResponsesPartitionsSnapshotId = "SnapshotId"

	// FieldFetchResponseResponsesPartitionsSnapshotIdEndOffset is:
	FieldFetchResponseResponsesPartitionsSnapshotIdEndOffset = "EndOffset"

	// FieldFetchResponseResponsesPartitionsSnapshotIdEpoch is:
	FieldFetchResponseResponsesPartitionsSnapshotIdEpoch = "Epoch"

	// FieldFetchResponseResponsesPartitionsSnapshotIdTags is: The tagged fields.
	FieldFetchResponseResponsesPartitionsSnapshotIdTags = "Tags"

	// FieldFetchResponseResponsesPartitionsTags is: The tagged fields.
	FieldFetchResponseResponsesPartitionsTags = "Tags"

	// FieldFetchResponseResponsesTags is: The tagged fields.
	FieldFetchResponseResponsesTags = "Tags"

	// FieldFetchResponseResponsesTopic is: The topic name.
	FieldFetchResponseResponsesTopic = "Topic"

	// FieldFetchResponseResponsesTopicId is: The unique topic ID
	FieldFetchResponseResponsesTopicId = "TopicId"

	// FieldFetchResponseSessionId is: The fetch session ID, or 0 if this is not part of a fetch session.
	FieldFetchResponseSessionId = "SessionId"

	// FieldFetchResponseTags is: The tagged fields.
	FieldFetchResponseTags = "Tags"

	// FieldFetchResponseThrottleTimeMs is: The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota.
	FieldFetchResponseThrottleTimeMs = "ThrottleTimeMs"
)

// Generated from Apache Kafka source code file: clients/src/main/resources/common/message/FetchResponse.json
const originalFetchResponseInput = `// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

{
  "apiKey": 1,
  "type": "response",
  "name": "FetchResponse",
  //
  // Version 1 adds throttle time.
  //
  // Version 2 and 3 are the same as version 1. 
  //
  // Version 4 adds features for transactional consumption.
  //
  // Version 5 adds LogStartOffset to indicate the earliest available offset of
  // partition data that can be consumed.
  //
  // Starting in version 6, we may return KAFKA_STORAGE_ERROR as an error code.
  //
  // Version 7 adds incremental fetch request support.
  //
  // Starting in version 8, on quota violation, brokers send out responses before throttling.
  //
  // Version 9 is the same as version 8.
  //
  // Version 10 indicates that the response data can use the ZStd compression
  // algorithm, as described in KIP-110.
  // Version 12 adds support for flexible versions, epoch detection through the 'TruncationOffset' field,
  // and leader discovery through the 'CurrentLeader' field
  //
  // Version 13 replaces the topic name field with topic ID (KIP-516).
  "validVersions": "0-13",
  "flexibleVersions": "12+",
  "fields": [
    { "name": "ThrottleTimeMs", "type": "int32", "versions": "1+", "ignorable": true,
      "about": "The duration in milliseconds for which the request was throttled due to a quota violation, or zero if the request did not violate any quota." },
    { "name": "ErrorCode", "type": "int16", "versions": "7+", "ignorable": true,
      "about": "The top level response error code." },
    { "name": "SessionId", "type": "int32", "versions": "7+", "default": "0", "ignorable": false,
      "about": "The fetch session ID, or 0 if this is not part of a fetch session." },
    { "name": "Responses", "type": "[]FetchableTopicResponse", "versions": "0+",
      "about": "The response topics.", "fields": [
      { "name": "Topic", "type": "string", "versions": "0-12", "ignorable": true, "entityType": "topicName",
        "about": "The topic name." },
      { "name": "TopicId", "type": "uuid", "versions": "13+", "ignorable": true, "about": "The unique topic ID"},
      { "name": "Partitions", "type": "[]PartitionData", "versions": "0+",
        "about": "The topic partitions.", "fields": [
        { "name": "PartitionIndex", "type": "int32", "versions": "0+",
          "about": "The partition index." },
        { "name": "ErrorCode", "type": "int16", "versions": "0+",
          "about": "The error code, or 0 if there was no fetch error." },
        { "name": "HighWatermark", "type": "int64", "versions": "0+",
          "about": "The current high water mark." },
        { "name": "LastStableOffset", "type": "int64", "versions": "4+", "default": "-1", "ignorable": true,
          "about": "The last stable offset (or LSO) of the partition. This is the last offset such that the state of all transactional records prior to this offset have been decided (ABORTED or COMMITTED)" },
        { "name": "LogStartOffset", "type": "int64", "versions": "5+", "default": "-1", "ignorable": true,
          "about": "The current log start offset." },
        { "name": "DivergingEpoch", "type": "EpochEndOffset", "versions": "12+", "taggedVersions": "12+", "tag": 0,
          "about": "In case divergence is detected based on the 'LastFetchedEpoch' and 'FetchOffset' in the request, this field indicates the largest epoch and its end offset such that subsequent records are known to diverge",
          "fields": [
            { "name": "Epoch", "type": "int32", "versions": "12+", "default": "-1" },
            { "name": "EndOffset", "type": "int64", "versions": "12+", "default": "-1" }
        ]},
        { "name": "CurrentLeader", "type": "LeaderIdAndEpoch",
          "versions": "12+", "taggedVersions": "12+", "tag": 1, "fields": [
          { "name": "LeaderId", "type": "int32", "versions": "12+", "default": "-1", "entityType": "brokerId",
            "about": "The ID of the current leader or -1 if the leader is unknown."},
          { "name": "LeaderEpoch", "type": "int32", "versions": "12+", "default": "-1",
            "about": "The latest known leader epoch"}
        ]},
        { "name": "SnapshotId", "type": "SnapshotId",
          "versions": "12+", "taggedVersions": "12+", "tag": 2,
          "about": "In the case of fetching an offset less than the LogStartOffset, this is the end offset and epoch that should be used in the FetchSnapshot request.",
          "fields": [
            { "name": "EndOffset", "type": "int64", "versions": "0+", "default": "-1" },
            { "name": "Epoch", "type": "int32", "versions": "0+", "default": "-1" }
        ]},
        { "name": "AbortedTransactions", "type": "[]AbortedTransaction", "versions": "4+", "nullableVersions": "4+", "ignorable": true,
          "about": "The aborted transactions.",  "fields": [
          { "name": "ProducerId", "type": "int64", "versions": "4+", "entityType": "producerId",
            "about": "The producer id associated with the aborted transaction." },
          { "name": "FirstOffset", "type": "int64", "versions": "4+",
            "about": "The first offset in the aborted transaction." }
        ]},
        { "name": "PreferredReadReplica", "type": "int32", "versions": "11+", "default": "-1", "ignorable": false, "entityType": "brokerId",
          "about": "The preferred read replica for the consumer to use on its next fetch request"},
        { "name": "Records", "type": "records", "versions": "0+", "nullableVersions": "0+", "about": "The record data."}
      ]}
    ]}
  ]
}
`
