package kendaql

import (
	pbkql "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/kendaql"
)

type isCol interface {
	isCol() (*pbkql.Col, error)
}

// LogicGP is kandaql logical group.
type LogicGP struct {
	lg *pbkql.LogicalGroup
	e  error
}

func (lg *LogicGP) logicalGroupNode() *pbkql.LogicalGroupNode {
	return &pbkql.LogicalGroupNode{
		Field: &pbkql.LogicalGroupNode_Group{Group: lg.lg},
	}
}

func (lg *LogicGP) logicalGroup() *pbkql.LogicalGroup {
	return lg.lg
}

func (lg *LogicGP) err() error {
	return lg.e
}

// And is a grammatical conjunction used to connect one or more nodes.
func And(fields ...isLogicalGroup) *LogicGP {
	lg := setGroupNode(fields)
	lg.lg.Logical = pbkql.Logical_AND
	return lg
}

// Or is a grammatical conjunction used to connect one or more nodes.
func Or(fields ...isLogicalGroup) *LogicGP {
	lg := setGroupNode(fields)
	lg.lg.Logical = pbkql.Logical_OR
	return lg
}

func setGroupNode(fields []isLogicalGroup) *LogicGP {
	lg := &LogicGP{
		lg: &pbkql.LogicalGroup{
			Fields: make([]*pbkql.LogicalGroupNode, len(fields)),
		},
	}

	for i, field := range fields {
		if err := field.err(); err != nil {
			lg.e = err
			return lg
		}
		lg.lg.Fields[i] = field.logicalGroupNode()
	}
	return lg
}

// Query is a request message of Query endpoint for kendaql service.
type Query struct {
	q   *pbkql.Query
	err error
}

// Select creates a Query message.
func Select(columns ...isCol) *Query {
	rp := &Query{
		q: &pbkql.Query{
			Selects: make([]*pbkql.Col, len(columns)),
		},
	}
	for i, column := range columns {
		col, err := column.isCol()
		if err != nil {
			rp.err = err
			return rp
		}
		rp.q.Selects[i] = col
	}

	return rp
}

// From specifies what table we care.
func (q *Query) From(table *Tb) *Query {
	if q.err != nil {
		return q
	}

	q.q.From = table.tb

	return q
}

// Join specifies relational tables.
func (q *Query) Join(relationalTables ...*RelationalTable) *Query {
	if q.err != nil {
		return q
	}

	q.q.Joins = make([]*pbkql.Join, len(relationalTables))

	for i, relationalTable := range relationalTables {
		if relationalTable.err != nil {
			q.err = relationalTable.err
			return q
		}
		q.q.Joins[i] = relationalTable.join
	}

	return q
}

// Where describes a condition clause for Query message.
func (q *Query) Where(clause isLogicalGroup) *Query {
	if q.err != nil {
		return q
	}

	if err := clause.err(); err != nil {
		q.err = err
		return q
	}
	q.q.Where = clause.logicalGroup()

	return q
}

// GroupBy groups rows that have the same values into summary rows.
func (q *Query) GroupBy(columns ...*Col) *Query {
	if q.err != nil {
		return q
	}

	q.q.GroupBys = make([]*pbkql.Column, len(columns))

	for i, column := range columns {
		q.q.GroupBys[i] = column.col
	}

	return q
}

// Having describes a condition clause with aggregate function for Query message.
func (q *Query) Having(clause isLogicalGroup) *Query {
	if q.err != nil {
		return q
	}

	if err := clause.err(); err != nil {
		q.err = err
		return q
	}
	q.q.Having = clause.logicalGroup()

	return q
}

// OrderBy is used to sort the result-set in ascending or descending order.
func (q *Query) OrderBy(clauses ...*pbkql.OrderBy) *Query {
	if q.err != nil {
		return q
	}

	q.q.OrderBys = make([]*pbkql.OrderBy, len(clauses))

	for i, clause := range clauses {
		q.q.OrderBys[i] = clause
	}

	return q
}

// ToProtoMessage generates a Query message for Query request.
func (q *Query) ToProtoMessage() (*pbkql.Query, error) {
	if q.err != nil {
		return nil, q.err
	}
	return q.q, nil
}

// Mutation is a request message of Mutation endpoint for kendaql servce.
type Mutation struct {
	m   *pbkql.Mutation
	err error
}

func newMutation(typ pbkql.MutationType, table string) *Mutation {
	return &Mutation{
		m: &pbkql.Mutation{
			Type: typ,
			Table: &pbkql.Table{
				Name: table,
			},
		},
	}
}

// Insert creates a insert message for Mutation endpoint, it is used to insert new records in a table.
func Insert(table string) *Mutation {
	return newMutation(pbkql.MutationType_INSERT, table)
}

// Update creates a update message for Mutation endpoint, it is used to modify the existing records in a table.
func Update(table string) *Mutation {
	return newMutation(pbkql.MutationType_UPDATE, table)
}

// Delete creates a delete message for Mutation endpoint, it is used to delete existing records in a table.
func Delete(table string) *Mutation {
	return newMutation(pbkql.MutationType_DELETE, table)
}

// Values is used to assign the values for Insert or Update functions.
func (m *Mutation) Values(columns ...*ColumnWithValue) *Mutation {
	if m.err != nil {
		return m
	}

	m.m.Fields = make([]*pbkql.Field, len(columns))

	for i, column := range columns {
		if column.err != nil {
			m.err = column.err
			return m
		}
		m.m.Fields[i] = column.mf
	}
	return m
}

// Where describes a condition clause for Mutation message.
func (m *Mutation) Where(clause isLogicalGroup) *Mutation {
	if m.err != nil {
		return m
	}

	if err := clause.err(); err != nil {
		m.err = err
		return m
	}
	m.m.Filter = clause.logicalGroup()

	return m
}

// ToProtoMessage generates a Mutation message for Mutation request.
func (m *Mutation) ToProtoMessage() (*pbkql.Mutation, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.m, nil
}
