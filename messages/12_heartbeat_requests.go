package messages

import "github.com/radekg/kafka-protocol-go/schema"

const (
	FieldNameHeartbeatGenerationId    = "generation_id"
	FieldNameHeartbeatGroupId         = "group_id"
	FieldNameHeartbeatMemberId        = "member_id"
	FieldNameHeartbeatGroupInstanceId = "group_instance_id"
)

func initHeartbeatRequests() []schema.Schema {

	schemas := []schema.Schema{}

	heartbeatV0 := schema.NewSchema("heartbeat_v0",
		&schema.Mfield{Name: FieldNameHeartbeatGroupId, Ty: schema.TypeStr},
		&schema.Mfield{Name: FieldNameHeartbeatGenerationId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeartbeatMemberId, Ty: schema.TypeStr})

	schemas = append(schemas, heartbeatV0)
	schemas = append(schemas, heartbeatV0)
	schemas = append(schemas, heartbeatV0)

	heartbeatV3 := schema.NewSchema("heartbeat_v3",
		&schema.Mfield{Name: FieldNameHeartbeatGroupId, Ty: schema.TypeStr},
		&schema.Mfield{Name: FieldNameHeartbeatGenerationId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeartbeatMemberId, Ty: schema.TypeStr},
		&schema.Mfield{Name: FieldNameHeartbeatGroupInstanceId, Ty: schema.TypeNullableStr})

	schemas = append(schemas, heartbeatV3)

	heartbeatV4 := schema.NewSchema("heartbeat_v4",
		&schema.Mfield{Name: FieldNameHeartbeatGroupId, Ty: schema.TypeCompactStr},
		&schema.Mfield{Name: FieldNameHeartbeatGenerationId, Ty: schema.TypeInt32},
		&schema.Mfield{Name: FieldNameHeartbeatMemberId, Ty: schema.TypeCompactStr},
		&schema.Mfield{Name: FieldNameHeartbeatGroupInstanceId, Ty: schema.TypeCompactNullableStr},
		&schema.SchemaTaggedFields{Name: FieldNameTagBuffer})

	schemas = append(schemas, heartbeatV4)

	return schemas

}
