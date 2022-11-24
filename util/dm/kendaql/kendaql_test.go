package kendaql

import (
	"reflect"
	"testing"

	pbkql "gitlab.kenda.com.tw/kenda/commons/v2/proto/golang/dm/kendaql"
)

func TestQuery(t *testing.T) {
	tests := []struct {
		name  string
		field *Query
		want  *pbkql.Query
	}{
		{
			"simple",
			Select(
				Column("id").ForTable("emp"),
				Column("name").ForTable("dep"),
				Column("name").ForTable("emp"),
				Column("description").ForTable("gen"),
				Column("address"),
			).From(
				Table("employee").WithAlias("emp"),
			).Join(
				Table("department").WithAlias("dep").JoinThisOn(
					Column("id").ForTable("dep"), Column("department_id").ForTable("emp")),
				Table("gender").WithAlias("gen").LeftJoinThisOn(
					Column("id").ForTable("gen"), Column("gender").ForTable("emp")),
			).Where(
				Column("department_id").ForTable("emp").Strings("A1234").Equal(),
			).OrderBy(
				Column("id").ForTable("emp").Desc(),
			),
			&pbkql.Query{
				Selects: []*pbkql.Col{
					&pbkql.Col{Col: &pbkql.Col_Column{
						Column: &pbkql.Column{TableAlias: "emp", Name: "id"},
					}},
					&pbkql.Col{Col: &pbkql.Col_Column{
						Column: &pbkql.Column{TableAlias: "dep", Name: "name"},
					}},
					&pbkql.Col{Col: &pbkql.Col_Column{
						Column: &pbkql.Column{TableAlias: "emp", Name: "name"},
					}},
					&pbkql.Col{Col: &pbkql.Col_Column{
						Column: &pbkql.Column{TableAlias: "gen", Name: "description"},
					}},
					&pbkql.Col{Col: &pbkql.Col_Column{
						Column: &pbkql.Column{TableAlias: "", Name: "address"},
					}},
				},
				From: &pbkql.Table{Name: "employee", Alias: "emp"},
				Joins: []*pbkql.Join{
					&pbkql.Join{
						Type:  pbkql.JoinType_JOIN,
						Table: &pbkql.Table{Name: "department", Alias: "dep"},
						Ons: []*pbkql.JoinOn{
							&pbkql.JoinOn{
								From: &pbkql.Column{TableAlias: "dep", Name: "id"},
								To:   &pbkql.Column{TableAlias: "emp", Name: "department_id"},
							},
						},
					},
					&pbkql.Join{
						Type:  pbkql.JoinType_LEFT,
						Table: &pbkql.Table{Name: "gender", Alias: "gen"},
						Ons: []*pbkql.JoinOn{
							&pbkql.JoinOn{
								From: &pbkql.Column{TableAlias: "gen", Name: "id"},
								To:   &pbkql.Column{TableAlias: "emp", Name: "gender"},
							},
						},
					},
				},
				Where: &pbkql.LogicalGroup{
					Logical: pbkql.Logical_AND,
					Fields: []*pbkql.LogicalGroupNode{
						&pbkql.LogicalGroupNode{
							Field: &pbkql.LogicalGroupNode_Condition{
								Condition: &pbkql.ConditionField{
									Condition: pbkql.Condition_EQUAL,
									Field: &pbkql.Field{
										Column: &pbkql.Col{Col: &pbkql.Col_Column{
											Column: &pbkql.Column{TableAlias: "emp", Name: "department_id"},
										}},
										Values: [][]byte{[]byte("A1234")},
									},
								},
							},
						},
					},
				},
				OrderBys: []*pbkql.OrderBy{
					&pbkql.OrderBy{
						Type:   pbkql.Order_DESCENDING,
						Column: &pbkql.Column{TableAlias: "emp", Name: "id"},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := tt.field.ToProtoMessage()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(msg, tt.want) {
				t.Errorf("got = %v, want %v", msg, tt.want)
			}
		})
	}
}

func TestAnd(t *testing.T) {
	type args struct {
		fields []isLogicalGroup
	}
	tests := []struct {
		name string
		args args
		want *LogicGP
	}{
		{
			"result sql : subno = '1' AND depno = 'M2200' ",
			args{
				[]isLogicalGroup{
					Column("subno").Strings("1").Equal(),
					Column("depno").Strings("M2200").Equal(),
				},
			},
			&LogicGP{
				lg: &pbkql.LogicalGroup{
					Logical: pbkql.Logical_AND,
					Fields: []*pbkql.LogicalGroupNode{
						{
							Field: &pbkql.LogicalGroupNode_Condition{
								Condition: &pbkql.ConditionField{
									Condition: pbkql.Condition_EQUAL,
									Field: &pbkql.Field{
										Column: &pbkql.Col{Col: &pbkql.Col_Column{
											Column: &pbkql.Column{Name: "subno"},
										}},
										Values: [][]byte{[]byte("1")},
									},
								},
							},
						},
						{
							Field: &pbkql.LogicalGroupNode_Condition{
								Condition: &pbkql.ConditionField{
									Condition: pbkql.Condition_EQUAL,
									Field: &pbkql.Field{
										Column: &pbkql.Col{Col: &pbkql.Col_Column{
											Column: &pbkql.Column{Name: "depno"},
										}},
										Values: [][]byte{[]byte("M2200")},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			"result sql : date BETWEEN '1911-01-01' AND '2020-01-01' AND depno = 'M2200' ",
			args{
				[]isLogicalGroup{
					Column("date").Strings("1911-01-01", "2020-01-01").Between(),
					Column("depno").Strings("M2200").Equal(),
				},
			},
			&LogicGP{
				lg: &pbkql.LogicalGroup{
					Logical: pbkql.Logical_AND,
					Fields: []*pbkql.LogicalGroupNode{
						{
							Field: &pbkql.LogicalGroupNode_Condition{
								Condition: &pbkql.ConditionField{
									Condition: pbkql.Condition_BETWEEN,
									Field: &pbkql.Field{
										Column: &pbkql.Col{Col: &pbkql.Col_Column{
											Column: &pbkql.Column{Name: "date"},
										}},
										Values: [][]byte{[]byte("1911-01-01"), []byte("2020-01-01")},
									},
								},
							},
						},
						{
							Field: &pbkql.LogicalGroupNode_Condition{
								Condition: &pbkql.ConditionField{
									Condition: pbkql.Condition_EQUAL,
									Field: &pbkql.Field{
										Column: &pbkql.Col{Col: &pbkql.Col_Column{
											Column: &pbkql.Column{Name: "depno"},
										}},
										Values: [][]byte{[]byte("M2200")},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			" result sql : ( a > '1' OR b < '2' ) AND ( c = '3' OR d != '4' ) ",
			args{
				[]isLogicalGroup{
					Or(
						Column("a").Strings("1").Greater(),
						Column("b").Strings("2").Less(),
					),
					Or(
						Column("c").Strings("3").Equal(),
						Column("d").Strings("4").NotEqual(),
					),
				},
			},
			&LogicGP{
				lg: &pbkql.LogicalGroup{
					Logical: pbkql.Logical_AND,
					Fields: []*pbkql.LogicalGroupNode{
						{
							Field: &pbkql.LogicalGroupNode_Group{
								Group: &pbkql.LogicalGroup{
									Logical: pbkql.Logical_OR,
									Fields: []*pbkql.LogicalGroupNode{
										{
											Field: &pbkql.LogicalGroupNode_Condition{
												Condition: &pbkql.ConditionField{
													Condition: pbkql.Condition_GREATER,
													Field: &pbkql.Field{
														Column: &pbkql.Col{Col: &pbkql.Col_Column{
															Column: &pbkql.Column{Name: "a"},
														}},
														Values: [][]byte{[]byte("1")},
													},
												},
											},
										},
										{
											Field: &pbkql.LogicalGroupNode_Condition{
												Condition: &pbkql.ConditionField{
													Condition: pbkql.Condition_LESS,
													Field: &pbkql.Field{
														Column: &pbkql.Col{Col: &pbkql.Col_Column{
															Column: &pbkql.Column{Name: "b"},
														}},
														Values: [][]byte{[]byte("2")},
													},
												},
											},
										},
									},
								},
							},
						},
						{
							Field: &pbkql.LogicalGroupNode_Group{
								Group: &pbkql.LogicalGroup{
									Logical: pbkql.Logical_OR,
									Fields: []*pbkql.LogicalGroupNode{
										{
											Field: &pbkql.LogicalGroupNode_Condition{
												Condition: &pbkql.ConditionField{
													Condition: pbkql.Condition_EQUAL,
													Field: &pbkql.Field{
														Column: &pbkql.Col{Col: &pbkql.Col_Column{
															Column: &pbkql.Column{Name: "c"},
														}},
														Values: [][]byte{[]byte("3")},
													},
												},
											},
										},
										{
											Field: &pbkql.LogicalGroupNode_Condition{
												Condition: &pbkql.ConditionField{
													Condition: pbkql.Condition_NOT_EQUAL,
													Field: &pbkql.Field{
														Column: &pbkql.Col{Col: &pbkql.Col_Column{
															Column: &pbkql.Column{Name: "d"},
														}},
														Values: [][]byte{[]byte("4")},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := And(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_From(t *testing.T) {
	type fields struct {
		q *pbkql.Query
	}
	type args struct {
		tb *Tb
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			" FROM peremp a ",
			fields{&pbkql.Query{}},
			args{
				Table("peremp").WithAlias("a"),
			},
			&Query{
				q: &pbkql.Query{
					From: &pbkql.Table{
						Name:  "peremp",
						Alias: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				q: tt.fields.q,
			}
			if got := q.From(tt.args.tb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.From() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Join(t *testing.T) {
	type fields struct {
		q *pbkql.Query
	}
	type args struct {
		j []*RelationalTable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			" inner join seccod b on a.subno = b.subno and a.tt = b.tt ",
			fields{&pbkql.Query{}},
			args{
				[]*RelationalTable{
					Table("seccod").WithAlias("b").InnerJoinThisOn(
						Column("subno").ForTable("a"), Column("subno").ForTable("b"),
						Column("tt").ForTable("a"), Column("tt").ForTable("b"),
					),
				},
			},
			&Query{
				q: &pbkql.Query{
					Joins: []*pbkql.Join{
						&pbkql.Join{
							Type: pbkql.JoinType_INNER,
							Table: &pbkql.Table{
								Name:  "seccod",
								Alias: "b",
							},
							Ons: []*pbkql.JoinOn{
								&pbkql.JoinOn{
									From: &pbkql.Column{
										TableAlias: "a",
										Name:       "subno",
									},
									To: &pbkql.Column{
										TableAlias: "b",
										Name:       "subno",
									},
								},
								&pbkql.JoinOn{
									From: &pbkql.Column{
										TableAlias: "a",
										Name:       "tt",
									},
									To: &pbkql.Column{
										TableAlias: "b",
										Name:       "tt",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				q: tt.fields.q,
			}
			if got := q.Join(tt.args.j...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery_Where(t *testing.T) {
	type fields struct {
		q *pbkql.Query
	}
	type args struct {
		w isLogicalGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			" where subno = '1' ",
			fields{&pbkql.Query{}},
			args{
				Column("subno").Strings("1").Equal(),
			},
			&Query{
				q: &pbkql.Query{
					Where: &pbkql.LogicalGroup{
						Logical: pbkql.Logical_AND,
						Fields: []*pbkql.LogicalGroupNode{
							{
								Field: &pbkql.LogicalGroupNode_Condition{
									Condition: &pbkql.ConditionField{
										Condition: pbkql.Condition_EQUAL,
										Field: &pbkql.Field{
											Column: &pbkql.Col{Col: &pbkql.Col_Column{
												Column: &pbkql.Column{Name: "subno"},
											}},
											Values: [][]byte{[]byte("1")},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			" where subno = '1' and depno = 'M2200' ",
			fields{&pbkql.Query{}},
			args{
				And(
					Column("subno").Strings("1").Equal(),
					Column("depno").Strings("M2200").Equal(),
				),
			},
			&Query{
				q: &pbkql.Query{
					Where: &pbkql.LogicalGroup{
						Logical: pbkql.Logical_AND,
						Fields: []*pbkql.LogicalGroupNode{
							{
								Field: &pbkql.LogicalGroupNode_Condition{
									Condition: &pbkql.ConditionField{
										Condition: pbkql.Condition_EQUAL,
										Field: &pbkql.Field{
											Column: &pbkql.Col{Col: &pbkql.Col_Column{
												Column: &pbkql.Column{Name: "subno"},
											}},
											Values: [][]byte{[]byte("1")},
										},
									},
								},
							},
							{
								Field: &pbkql.LogicalGroupNode_Condition{
									Condition: &pbkql.ConditionField{
										Condition: pbkql.Condition_EQUAL,
										Field: &pbkql.Field{
											Column: &pbkql.Col{Col: &pbkql.Col_Column{
												Column: &pbkql.Column{Name: "depno"},
											}},
											Values: [][]byte{[]byte("M2200")},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				q: tt.fields.q,
			}
			if got := q.Where(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Where() = %v, want %v", *got.q, *tt.want.q)
			}
		})
	}
}

func TestQuery_Having(t *testing.T) {
	expectedValue, err := ToValue(10)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		q *pbkql.Query
	}
	type args struct {
		h isLogicalGroup
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			" having MAX(price) > 10 ",
			fields{&pbkql.Query{}},
			args{
				Max("price").ForTable("t").WithExpectedType(pbkql.DataType_INT).Ints(10).Greater(),
			},
			&Query{
				q: &pbkql.Query{
					Having: &pbkql.LogicalGroup{
						Logical: pbkql.Logical_AND,
						Fields: []*pbkql.LogicalGroupNode{
							{
								Field: &pbkql.LogicalGroupNode_Condition{
									Condition: &pbkql.ConditionField{
										Condition: pbkql.Condition_GREATER,
										Field: &pbkql.Field{
											Column: &pbkql.Col{Col: &pbkql.Col_Function{
												Function: &pbkql.Function{
													Name: "MAX",
													Parameters: []*pbkql.Function_Param{
														&pbkql.Function_Param{Value: &pbkql.Function_Param_ColumnValue{ColumnValue: &pbkql.Column{
															TableAlias: "t",
															Name:       "price",
														}}},
													},
													ExpectedType: pbkql.DataType_INT},
											}},
											Values: [][]byte{expectedValue},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				q: tt.fields.q,
			}
			if got := q.Having(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.Having() = %v, want %v", *got.q, *tt.want.q)
			}
		})
	}
}

func TestQuery_GroupBy(t *testing.T) {
	type fields struct {
		q *pbkql.Query
	}
	type args struct {
		g []*Col
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			" group by subno,name ",
			fields{&pbkql.Query{}},
			args{
				[]*Col{
					Column("subno"),
					Column("name"),
				},
			},
			&Query{
				q: &pbkql.Query{
					GroupBys: []*pbkql.Column{
						&pbkql.Column{
							Name: "subno",
						},
						&pbkql.Column{
							Name: "name",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				q: tt.fields.q,
			}
			if got := q.GroupBy(tt.args.g...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.GroupBy() = %v, want %v", *got.q, *tt.want.q)
			}
		})
	}
}

func TestQuery_Order(t *testing.T) {
	type fields struct {
		q *pbkql.Query
	}
	type args struct {
		o []*pbkql.OrderBy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Query
	}{
		{
			" order by a.subno desc ",
			fields{&pbkql.Query{}},
			args{
				[]*pbkql.OrderBy{
					Column("subno").ForTable("a").Desc(),
				},
			},
			&Query{
				q: &pbkql.Query{
					OrderBys: []*pbkql.OrderBy{
						&pbkql.OrderBy{
							Type: pbkql.Order_DESCENDING,
							Column: &pbkql.Column{
								TableAlias: "a",
								Name:       "subno",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Query{
				q: tt.fields.q,
			}
			if got := q.OrderBy(tt.args.o...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query.OrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutation_Values(t *testing.T) {
	type fields struct {
		m *pbkql.Mutation
	}
	type args struct {
		fd []*ColumnWithValue
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Mutation
	}{
		{
			" insert into xxx (fa,fb,fc) values ('test1','abc','3.14') ",
			fields{&pbkql.Mutation{}},
			args{
				[]*ColumnWithValue{
					Column("fa").Strings("test1"),
					Column("fb").Strings("abc"),
					Column("fc").Strings("3.14"),
				},
			},
			&Mutation{
				m: &pbkql.Mutation{
					Fields: []*pbkql.Field{
						&pbkql.Field{
							Column: &pbkql.Col{Col: &pbkql.Col_Column{
								Column: &pbkql.Column{Name: "fa"},
							}},
							Values: [][]byte{[]byte("test1")},
						},
						&pbkql.Field{
							Column: &pbkql.Col{Col: &pbkql.Col_Column{
								Column: &pbkql.Column{Name: "fb"},
							}},
							Values: [][]byte{[]byte("abc")},
						},
						&pbkql.Field{
							Column: &pbkql.Col{Col: &pbkql.Col_Column{
								Column: &pbkql.Column{Name: "fc"},
							}},
							Values: [][]byte{[]byte("3.14")},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mutation{
				m: tt.fields.m,
			}
			if got := m.Values(tt.args.fd...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mutation.Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
