package kendaql

import (
	pbkql "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/kendaql"
)

// Tb is the table.
type Tb struct {
	tb *pbkql.Table
}

// Table create a table.
func Table(name string) *Tb {
	return &Tb{
		tb: &pbkql.Table{
			Name: name,
		},
	}
}

// WithAlias defines the alias of the table.
func (tb *Tb) WithAlias(alias string) *Tb {
	tb.tb.Alias = alias
	return tb
}

// JoinThisOn make the table be a relational table with specific table.
func (tb *Tb) JoinThisOn(pairs ...*Col) *RelationalTable {
	rTable := &RelationalTable{
		join: &pbkql.Join{
			Type: pbkql.JoinType_JOIN,
		},
	}
	return rTable.on(tb, pairs...)
}

// LeftJoinThisOn make the table be a relational table with specific table.
func (tb *Tb) LeftJoinThisOn(pairs ...*Col) *RelationalTable {
	rTable := &RelationalTable{
		join: &pbkql.Join{
			Type: pbkql.JoinType_LEFT,
		},
	}
	return rTable.on(tb, pairs...)
}

// RightJoinThisOn make the table be a relational table with specific table.
func (tb *Tb) RightJoinThisOn(pairs ...*Col) *RelationalTable {
	rTable := &RelationalTable{
		join: &pbkql.Join{
			Type: pbkql.JoinType_RIGHT,
		},
	}
	return rTable.on(tb, pairs...)
}

// InnerJoinThisOn make the table be a relational table with specific table.
func (tb *Tb) InnerJoinThisOn(pairs ...*Col) *RelationalTable {
	rTable := &RelationalTable{
		join: &pbkql.Join{
			Type: pbkql.JoinType_INNER,
		},
	}
	return rTable.on(tb, pairs...)
}

// LeftOuterJoinThisOn make the table be a relational table with specific table.
func (tb *Tb) LeftOuterJoinThisOn(pairs ...*Col) *RelationalTable {
	rTable := &RelationalTable{
		join: &pbkql.Join{
			Type: pbkql.JoinType_LEFT_OUTER,
		},
	}
	return rTable.on(tb, pairs...)
}

// RightOuterJoinThisOn make the table be a relational table with specific table.
func (tb *Tb) RightOuterJoinThisOn(pairs ...*Col) *RelationalTable {
	rTable := &RelationalTable{
		join: &pbkql.Join{
			Type: pbkql.JoinType_RIGHT_OUTER,
		},
	}
	return rTable.on(tb, pairs...)
}

// FullOuterJoinThisOn make the table be a relational table with specific table.
func (tb *Tb) FullOuterJoinThisOn(pairs ...*Col) *RelationalTable {
	rTable := &RelationalTable{
		join: &pbkql.Join{
			Type: pbkql.JoinType_FULL_OUTER,
		},
	}
	return rTable.on(tb, pairs...)
}

// CrossJoinThis make the table be a relational table with specific table.
func (tb *Tb) CrossJoinThis() *RelationalTable {
	return &RelationalTable{
		join: &pbkql.Join{
			Type:  pbkql.JoinType_CROSS,
			Table: tb.tb,
		},
	}
}

// RelationalTable is used to be represented what table it is related.
type RelationalTable struct {
	join *pbkql.Join
	err  error
}

// on describes the relationship between two tables.
func (tb *RelationalTable) on(table *Tb, pairs ...*Col) *RelationalTable {
	if tb.err != nil {
		return tb
	}

	tb.join.Table = table.tb

	pairSize := len(pairs)
	if pairSize == 0 {
		tb.err = &callerError{caller: "on", err: ErrNoParameter}
		return tb
	} else if pairSize%2 == 1 {
		tb.err = &callerError{caller: "on", err: ErrOddParameter}
		return tb
	}

	size := pairSize / 2
	tb.join.Ons = make([]*pbkql.JoinOn, size)

	for i := 0; i < size; i++ {
		ii := 2 * i
		tb.join.Ons[i] = &pbkql.JoinOn{
			From: pairs[ii].col,
			To:   pairs[ii+1].col,
		}
	}

	return tb
}
