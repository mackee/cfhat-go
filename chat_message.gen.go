// Code generated by github.com/mackee/go-sqlla/v2/cmd/sqlla - DO NOT EDIT.
package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"database/sql"

	"github.com/mackee/go-sqlla/v2"
)

type chatMessageSQL struct {
	where sqlla.Where
}

func NewChatMessageSQL() chatMessageSQL {
	q := chatMessageSQL{}
	return q
}

var chatMessageAllColumns = []string{
	"`id`", "`message`", "`created_at`",
}

type chatMessageSelectSQL struct {
	chatMessageSQL
	Columns     []string
	order       string
	limit       *uint64
	offset      *uint64
	tableAlias  string
	joinClauses []string

	additionalWhereClause     string
	additionalWhereClauseArgs []interface{}

	groupByColumns []string

	isForUpdate bool
}

func (q chatMessageSQL) Select() chatMessageSelectSQL {
	return chatMessageSelectSQL{
		q,
		chatMessageAllColumns,
		"",
		nil,
		nil,
		"",
		nil,
		"",
		nil,
		nil,
		false,
	}
}

func (q chatMessageSelectSQL) Or(qs ...chatMessageSelectSQL) chatMessageSelectSQL {
	ws := make([]sqlla.Where, 0, len(qs))
	for _, q := range qs {
		ws = append(ws, q.where)
	}
	q.where = append(q.where, sqlla.ExprOr(ws))
	return q
}

func (q chatMessageSelectSQL) Limit(l uint64) chatMessageSelectSQL {
	q.limit = &l
	return q
}

func (q chatMessageSelectSQL) Offset(o uint64) chatMessageSelectSQL {
	q.offset = &o
	return q
}

func (q chatMessageSelectSQL) ForUpdate() chatMessageSelectSQL {
	q.isForUpdate = true
	return q
}

func (q chatMessageSelectSQL) TableAlias(alias string) chatMessageSelectSQL {
	q.tableAlias = "`" + alias + "`"
	return q
}

func (q chatMessageSelectSQL) SetColumns(columns ...string) chatMessageSelectSQL {
	q.Columns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(.`") {
			q.Columns = append(q.Columns, column)
		} else {
			q.Columns = append(q.Columns, "`"+column+"`")
		}
	}
	return q
}

func (q chatMessageSelectSQL) JoinClause(clause string) chatMessageSelectSQL {
	q.joinClauses = append(q.joinClauses, clause)
	return q
}

func (q chatMessageSelectSQL) AdditionalWhereClause(clause string, args ...interface{}) chatMessageSelectSQL {
	q.additionalWhereClause = clause
	q.additionalWhereClauseArgs = args
	return q
}

func (q chatMessageSelectSQL) appendColumnPrefix(column string) string {
	if q.tableAlias == "" || strings.ContainsAny(column, "(.") {
		return column
	}
	return q.tableAlias + "." + column
}

func (q chatMessageSelectSQL) GroupBy(columns ...string) chatMessageSelectSQL {
	q.groupByColumns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(.`") {
			q.groupByColumns = append(q.groupByColumns, column)
		} else {
			q.groupByColumns = append(q.groupByColumns, "`"+column+"`")
		}
	}
	return q
}

func (q chatMessageSelectSQL) ID(v ChatMessageID, exprs ...sqlla.Operator) chatMessageSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprInt64{Value: int64(v), Op: op, Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageSelectSQL) IDIn(vs ...ChatMessageID) chatMessageSelectSQL {
	_vs := make([]int64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, int64(v))
	}
	where := sqlla.ExprMultiInt64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`id`")}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) chatMessageSelectSQL {
	v := ChatMessageID(pk)
	return q.ID(v, exprs...)
}

func (q chatMessageSelectSQL) OrderByID(order sqlla.Order) chatMessageSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`id`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q chatMessageSelectSQL) Message(v string, exprs ...sqlla.Operator) chatMessageSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprString{Value: v, Op: op, Column: q.appendColumnPrefix("`message`")}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageSelectSQL) MessageIn(vs ...string) chatMessageSelectSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`message`")}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageSelectSQL) OrderByMessage(order sqlla.Order) chatMessageSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`message`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q chatMessageSelectSQL) CreatedAt(v int64, exprs ...sqlla.Operator) chatMessageSelectSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprInt64{Value: v, Op: op, Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageSelectSQL) CreatedAtIn(vs ...int64) chatMessageSelectSQL {
	where := sqlla.ExprMultiInt64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("`created_at`")}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageSelectSQL) OrderByCreatedAt(order sqlla.Order) chatMessageSelectSQL {
	q.order = " ORDER BY " + q.appendColumnPrefix("`created_at`")
	if order == sqlla.Asc {
		q.order += " ASC"
	} else {
		q.order += " DESC"
	}

	return q
}

func (q chatMessageSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	tableName := "chat_message"
	if q.tableAlias != "" {
		tableName = tableName + " AS " + q.tableAlias
		pcs := make([]string, 0, len(q.Columns))
		for _, column := range q.Columns {
			pcs = append(pcs, q.appendColumnPrefix(column))
		}
		columns = strings.Join(pcs, ", ")
	}
	query := "SELECT " + columns + " FROM " + tableName
	if len(q.joinClauses) > 0 {
		jc := strings.Join(q.joinClauses, " ")
		query += " " + jc
	}
	if wheres != "" {
		query += " WHERE" + wheres
	}
	if q.additionalWhereClause != "" {
		query += " " + q.additionalWhereClause
		if len(q.additionalWhereClauseArgs) > 0 {
			vs = append(vs, q.additionalWhereClauseArgs...)
		}
	}
	if len(q.groupByColumns) > 0 {
		query += " GROUP BY "
		gbcs := make([]string, 0, len(q.groupByColumns))
		for _, column := range q.groupByColumns {
			gbcs = append(gbcs, q.appendColumnPrefix(column))
		}
		query += strings.Join(gbcs, ", ")
	}
	query += q.order
	if q.limit != nil {
		query += " LIMIT " + strconv.FormatUint(*q.limit, 10)
	}
	if q.offset != nil {
		query += " OFFSET " + strconv.FormatUint(*q.offset, 10)
	}

	if q.isForUpdate {
		query += " FOR UPDATE"
	}

	return query + ";", vs, nil
}

func (s ChatMessage) Select() chatMessageSelectSQL {
	return NewChatMessageSQL().Select().ID(s.ID)
}
func (q chatMessageSelectSQL) Single(db sqlla.DB) (ChatMessage, error) {
	q.Columns = chatMessageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return ChatMessage{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q chatMessageSelectSQL) SingleContext(ctx context.Context, db sqlla.DB) (ChatMessage, error) {
	q.Columns = chatMessageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return ChatMessage{}, err
	}

	row := db.QueryRowContext(ctx, query, args...)
	return q.Scan(row)
}

func (q chatMessageSelectSQL) All(db sqlla.DB) ([]ChatMessage, error) {
	rs := make([]ChatMessage, 0, 10)
	q.Columns = chatMessageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q chatMessageSelectSQL) AllContext(ctx context.Context, db sqlla.DB) ([]ChatMessage, error) {
	rs := make([]ChatMessage, 0, 10)
	q.Columns = chatMessageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q chatMessageSelectSQL) Scan(s sqlla.Scanner) (ChatMessage, error) {
	var row ChatMessage
	err := s.Scan(
		&row.ID,
		&row.Message,
		&row.CreatedAt,
	)
	return row, err
}

type chatMessageUpdateSQL struct {
	chatMessageSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q chatMessageSQL) Update() chatMessageUpdateSQL {
	return chatMessageUpdateSQL{
		chatMessageSQL: q,
		setMap:         sqlla.SetMap{},
	}
}

func (q chatMessageUpdateSQL) SetID(v ChatMessageID) chatMessageUpdateSQL {
	q.setMap["`id`"] = v
	return q
}

func (q chatMessageUpdateSQL) WhereID(v ChatMessageID, exprs ...sqlla.Operator) chatMessageUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprInt64{Value: int64(v), Op: op, Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageUpdateSQL) WhereIDIn(vs ...ChatMessageID) chatMessageUpdateSQL {
	_vs := make([]int64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, int64(v))
	}
	where := sqlla.ExprMultiInt64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageUpdateSQL) SetMessage(v string) chatMessageUpdateSQL {
	q.setMap["`message`"] = v
	return q
}

func (q chatMessageUpdateSQL) WhereMessage(v string, exprs ...sqlla.Operator) chatMessageUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprString{Value: v, Op: op, Column: "`message`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageUpdateSQL) WhereMessageIn(vs ...string) chatMessageUpdateSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`message`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageUpdateSQL) SetCreatedAt(v int64) chatMessageUpdateSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q chatMessageUpdateSQL) WhereCreatedAt(v int64, exprs ...sqlla.Operator) chatMessageUpdateSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprInt64{Value: v, Op: op, Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageUpdateSQL) WhereCreatedAtIn(vs ...int64) chatMessageUpdateSQL {
	where := sqlla.ExprMultiInt64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = ChatMessage{}
	if t, ok := s.(chatMessageDefaultUpdateHooker); ok {
		q, err = t.DefaultUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	setColumns, svs, err := q.setMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	wheres, wvs, err := q.where.ToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "UPDATE chat_message SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (s ChatMessage) Update() chatMessageUpdateSQL {
	return NewChatMessageSQL().Update().WhereID(s.ID)
}

func (q chatMessageUpdateSQL) Exec(db sqlla.DB) ([]ChatMessage, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.chatMessageSQL

	return qq.Select().All(db)
}

func (q chatMessageUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]ChatMessage, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.chatMessageSQL

	return qq.Select().AllContext(ctx, db)
}

type chatMessageDefaultUpdateHooker interface {
	DefaultUpdateHook(chatMessageUpdateSQL) (chatMessageUpdateSQL, error)
}

type chatMessageInsertSQL struct {
	chatMessageSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q chatMessageSQL) Insert() chatMessageInsertSQL {
	return chatMessageInsertSQL{
		chatMessageSQL: q,
		setMap:         sqlla.SetMap{},
	}
}

func (q chatMessageInsertSQL) ValueID(v ChatMessageID) chatMessageInsertSQL {
	q.setMap["`id`"] = v
	return q
}

func (q chatMessageInsertSQL) ValueMessage(v string) chatMessageInsertSQL {
	q.setMap["`message`"] = v
	return q
}

func (q chatMessageInsertSQL) ValueCreatedAt(v int64) chatMessageInsertSQL {
	q.setMap["`created_at`"] = v
	return q
}

func (q chatMessageInsertSQL) ToSql() (string, []interface{}, error) {
	query, vs, err := q.chatMessageInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	return query + ";", vs, nil
}

func (q chatMessageInsertSQL) chatMessageInsertSQLToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = ChatMessage{}
	if t, ok := s.(chatMessageDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	qs, vs, err := q.setMap.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "INSERT INTO chat_message " + qs

	return query, vs, nil
}

func (q chatMessageInsertSQL) OnDuplicateKeyUpdate() chatMessageInsertOnDuplicateKeyUpdateSQL {
	return chatMessageInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

func (q chatMessageInsertSQL) Exec(db sqlla.DB) (ChatMessage, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return ChatMessage{}, err
	}
	result, err := db.Exec(query, args...)
	if err != nil {
		return ChatMessage{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return ChatMessage{}, err
	}
	return NewChatMessageSQL().Select().PkColumn(id).Single(db)
}

func (q chatMessageInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (ChatMessage, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return ChatMessage{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return ChatMessage{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return ChatMessage{}, err
	}
	return NewChatMessageSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q chatMessageInsertSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type chatMessageDefaultInsertHooker interface {
	DefaultInsertHook(chatMessageInsertSQL) (chatMessageInsertSQL, error)
}

type chatMessageInsertSQLToSqler interface {
	chatMessageInsertSQLToSql() (string, []interface{}, error)
}

type chatMessageInsertOnDuplicateKeyUpdateSQL struct {
	insertSQL               chatMessageInsertSQLToSqler
	onDuplicateKeyUpdateMap sqlla.SetMap
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateID(v ChatMessageID) chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = v
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateID(v sqlla.SetMapRawValue) chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = v
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) SameOnUpdateID() chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`id`"] = sqlla.SetMapRawValue("VALUES(`id`)")
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateMessage(v string) chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`message`"] = v
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateMessage(v sqlla.SetMapRawValue) chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`message`"] = v
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) SameOnUpdateMessage() chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`message`"] = sqlla.SetMapRawValue("VALUES(`message`)")
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) ValueOnUpdateCreatedAt(v int64) chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) RawValueOnUpdateCreatedAt(v sqlla.SetMapRawValue) chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = v
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) SameOnUpdateCreatedAt() chatMessageInsertOnDuplicateKeyUpdateSQL {
	q.onDuplicateKeyUpdateMap["`created_at`"] = sqlla.SetMapRawValue("VALUES(`created_at`)")
	return q
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = ChatMessage{}
	if t, ok := s.(chatMessageDefaultInsertOnDuplicateKeyUpdateHooker); ok {
		q, err = t.DefaultInsertOnDuplicateKeyUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}

	query, vs, err := q.insertSQL.chatMessageInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	os, ovs, err := q.onDuplicateKeyUpdateMap.ToUpdateSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	query += " ON DUPLICATE KEY UPDATE" + os
	vs = append(vs, ovs...)

	return query + ";", vs, nil
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) (ChatMessage, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return ChatMessage{}, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return ChatMessage{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return ChatMessage{}, err
	}
	return NewChatMessageSQL().Select().PkColumn(id).SingleContext(ctx, db)
}

func (q chatMessageInsertOnDuplicateKeyUpdateSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type chatMessageDefaultInsertOnDuplicateKeyUpdateHooker interface {
	DefaultInsertOnDuplicateKeyUpdateHook(chatMessageInsertOnDuplicateKeyUpdateSQL) (chatMessageInsertOnDuplicateKeyUpdateSQL, error)
}

type chatMessageBulkInsertSQL struct {
	insertSQLs []chatMessageInsertSQL
}

func (q chatMessageSQL) BulkInsert() *chatMessageBulkInsertSQL {
	return &chatMessageBulkInsertSQL{
		insertSQLs: []chatMessageInsertSQL{},
	}
}

func (q *chatMessageBulkInsertSQL) Append(iqs ...chatMessageInsertSQL) {
	q.insertSQLs = append(q.insertSQLs, iqs...)
}

func (q *chatMessageBulkInsertSQL) chatMessageInsertSQLToSql() (string, []interface{}, error) {
	if len(q.insertSQLs) == 0 {
		return "", []interface{}{}, fmt.Errorf("sqlla: This chatMessageBulkInsertSQL's InsertSQL was empty")
	}
	iqs := make([]chatMessageInsertSQL, len(q.insertSQLs))
	copy(iqs, q.insertSQLs)

	var s interface{} = ChatMessage{}
	if t, ok := s.(chatMessageDefaultInsertHooker); ok {
		for i, iq := range iqs {
			var err error
			iq, err = t.DefaultInsertHook(iq)
			if err != nil {
				return "", []interface{}{}, err
			}
			iqs[i] = iq
		}
	}

	sms := make(sqlla.SetMaps, 0, len(q.insertSQLs))
	for _, iq := range q.insertSQLs {
		sms = append(sms, iq.setMap)
	}

	query, vs, err := sms.ToInsertSql()
	if err != nil {
		return "", []interface{}{}, err
	}

	return "INSERT INTO `chat_message` " + query, vs, nil
}

func (q *chatMessageBulkInsertSQL) ToSql() (string, []interface{}, error) {
	query, vs, err := q.chatMessageInsertSQLToSql()
	if err != nil {
		return "", []interface{}{}, err
	}
	return query + ";", vs, nil
}

func (q *chatMessageBulkInsertSQL) OnDuplicateKeyUpdate() chatMessageInsertOnDuplicateKeyUpdateSQL {
	return chatMessageInsertOnDuplicateKeyUpdateSQL{
		insertSQL:               q,
		onDuplicateKeyUpdateMap: sqlla.SetMap{},
	}
}

func (q *chatMessageBulkInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type chatMessageDeleteSQL struct {
	chatMessageSQL
}

func (q chatMessageSQL) Delete() chatMessageDeleteSQL {
	return chatMessageDeleteSQL{
		q,
	}
}

func (q chatMessageDeleteSQL) ID(v ChatMessageID, exprs ...sqlla.Operator) chatMessageDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprInt64{Value: int64(v), Op: op, Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageDeleteSQL) IDIn(vs ...ChatMessageID) chatMessageDeleteSQL {
	_vs := make([]int64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, int64(v))
	}
	where := sqlla.ExprMultiInt64{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`id`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageDeleteSQL) Message(v string, exprs ...sqlla.Operator) chatMessageDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprString{Value: v, Op: op, Column: "`message`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageDeleteSQL) MessageIn(vs ...string) chatMessageDeleteSQL {
	where := sqlla.ExprMultiString{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`message`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageDeleteSQL) CreatedAt(v int64, exprs ...sqlla.Operator) chatMessageDeleteSQL {
	var op sqlla.Operator
	if len(exprs) == 0 {
		op = sqlla.OpEqual
	} else {
		op = exprs[0]
	}
	where := sqlla.ExprInt64{Value: v, Op: op, Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageDeleteSQL) CreatedAtIn(vs ...int64) chatMessageDeleteSQL {
	where := sqlla.ExprMultiInt64{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "`created_at`"}
	q.where = append(q.where, where)
	return q
}

func (q chatMessageDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, vs, err := q.where.ToSql()
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM chat_message"
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func (q chatMessageDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (q chatMessageDeleteSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
func (s ChatMessage) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewChatMessageSQL().Delete().ID(s.ID).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (s ChatMessage) DeleteContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := NewChatMessageSQL().Delete().ID(s.ID).ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
