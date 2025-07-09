package entity

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

type Contract struct {
	ent.Schema
}

func (Contract) Fields() []ent.Field {
	return []ent.Field{
		//The id field is builtin in the schema and does not need declaration. In SQL-based databases,
		//its type defaults to int (but can be changed with a codegen option)
		//and auto-incremented in the database.
		field.String("student_code").NotEmpty().MaxLen(10),
		field.String("first_name").MaxLen(128).NotEmpty(),
		field.String("last_name").MaxLen(128).NotEmpty(),
		field.String("middle_name").MaxLen(128),
		field.String("email").MaxLen(128).NotEmpty(),
		field.String("sign").MaxLen(128),
		field.String("phone").MaxLen(10).NotEmpty(),
		field.Uint8("gender"),
		field.Time("dob").Optional(),
		field.String("address"),
		field.Bytes("avatar"),
		field.Bool("is_active").Default(false),
		field.Time("registry_at").Default(time.Now()),
		field.String("room_id").MaxLen(5),
		field.Uint8("notification_channels"),
	}
}

func (Contract) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("student_code"),
		index.Fields("first_name", "last_name").Annotations(
			entsql.PrefixColumn("first_name", 16),
			entsql.PrefixColumn("last_name", 16)),
		index.Fields("email"),
		index.Fields("sign"),
		index.Fields("phone"),
		index.Fields("is_active"),
		index.Fields("room_id"),
	}
}
