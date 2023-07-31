// Code generated by queryx, DO NOT EDIT.

package db

import (
	"database/sql"
	"errors"
	"fmt"
	"snap/db/queryx"
)

type CardQuery struct {
	adapter         *queryx.Adapter
	schema          *queryx.Schema
	queries         Queries
	selectStatement *queryx.SelectStatement
	preload         map[string]bool
	err             error
}

func NewCardQuery(adapter *queryx.Adapter, schema *queryx.Schema, queries Queries) *CardQuery {
	s := queryx.NewSelect().Select("cards.*").From("cards")
	return &CardQuery{
		adapter:         adapter,
		schema:          schema,
		queries:         queries,
		selectStatement: s,
		preload:         make(map[string]bool),
	}
}

func (q *CardQuery) Create(change *queryx.CardChange) (*Card, error) {
	if q.err != nil {
		return nil, q.err
	}

	record := &Card{
		schema:  q.schema,
		queries: q.queries,
	}
	now := queryx.Now("2006-01-02 15:04:05.000")
	if !change.CreatedAt.Set {
		change.SetCreatedAt(now)
	}
	if !change.UpdatedAt.Set {
		change.SetUpdatedAt(now)
	}
	columns, values := change.Changes()
	query, args := queryx.NewInsert().
		Into("cards").
		Columns(columns...).
		Values(values...).
		Returning("id,name,cost,power,ability,def_id,created_at,updated_at").ToSQL()
	err := q.adapter.QueryOne(query, args...).Scan(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (q *CardQuery) InsertAll(changes []*queryx.CardChange) (int64, error) {
	if q.err != nil {
		return 0, q.err
	}

	if len(changes) == 0 {
		return 0, ErrInsertAllEmptyChanges
	}
	now := queryx.Now("2006-01-02 15:04:05.000")
	for _, change := range changes {
		if !change.CreatedAt.Set {
			change.SetCreatedAt(now)
		}
		if !change.UpdatedAt.Set {
			change.SetUpdatedAt(now)
		}
	}

	s := queryx.NewInsert().Into("cards")
	for i, change := range changes {
		columns, values := change.Changes()
		if i == 0 {
			s.Columns(columns...)
		}
		s.Values(values...)
	}

	query, args := s.ToSQL()
	return q.adapter.Exec(query, args...)
}

func (q *CardQuery) Delete(id int64) (int64, error) {
	query, args := queryx.NewDelete().From("cards").Where(q.schema.CardID.EQ(id)).ToSQL()
	result, err := q.adapter.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result, err
}

func (q *CardQuery) Find(id int64) (*Card, error) {
	res, err := q.Where(q.schema.CardID.EQ(id)).First()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, sql.ErrNoRows
	}
	res.schema = q.schema
	res.queries = q.queries
	return res, err
}

func (q *CardQuery) FindBy(where *queryx.Clause) (*Card, error) {
	return q.Where(where).First()
}

func (q *CardQuery) FindBySQL(query string, args ...interface{}) ([]*Card, error) {
	var cardList []Card
	cards := make([]*Card, 0)
	err := q.adapter.Query(query, args...).Scan(&cardList)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(cardList); i++ {
		cards = append(cards, &cardList[i])
	}
	return cards, nil
}

func (q *CardQuery) Where(clause *queryx.Clause) *CardQuery {
	q.selectStatement.Where(clause)
	q.err = clause.Err()
	return q
}

func (q *CardQuery) Select(selection ...string) *CardQuery {
	q.selectStatement.Select(selection...)
	return q
}

func (q *CardQuery) Limit(limit int) *CardQuery {
	q.selectStatement.Limit(limit)
	return q
}

func (q *CardQuery) Offset(offset int) *CardQuery {
	q.selectStatement.Offset(offset)
	return q
}

func (q *CardQuery) Group(group string) *CardQuery {
	q.selectStatement.GroupBy(group)
	return q
}

func (q *CardQuery) Having(having string) *CardQuery {
	q.selectStatement.Having(having)
	return q
}

func (q *CardQuery) Joins(joins string) *CardQuery {
	q.selectStatement.Join(joins)
	return q
}

func (q *CardQuery) Order(order ...string) *CardQuery {
	q.selectStatement.Order(order...)
	return q
}

func (q *CardQuery) PreloadDeckCards() *CardQuery {
	q.preload["deck_cards"] = true
	return q
}

func (q *CardQuery) preloadDeckCards(rows []*Card) error {
	ids := []int64{}
	for _, r := range rows {
		ids = append(ids, r.ID)
	}
	rows1, err := q.queries.QueryDeckCard().Where(q.schema.DeckCardCardID.In(ids)).All()
	if err != nil {
		return err
	}

	m := make(map[int64][]*DeckCard)
	for _, r := range rows1 {
		m[r.CardID.Val] = append(m[r.CardID.Val], r)
	}
	for _, r := range rows {
		if m[r.ID] != nil {
			r.DeckCards = m[r.ID]
		} else {
			r.DeckCards = make([]*DeckCard, 0)
		}
	}

	return nil
}

func (q *CardQuery) PreloadDeck() *CardQuery {
	q.preload["deck"] = true
	return q
}

func (q *CardQuery) preloadDeck(rows []*Card) error {
	ids := []int64{}
	for _, r := range rows {
		ids = append(ids, r.ID)
	}
	rows1, err := q.queries.QueryDeckCard().Where(q.schema.DeckCardCardID.In(ids)).All()
	if err != nil {
		return err
	}
	m1 := make(map[int64][]*DeckCard)
	for _, r := range rows1 {
		m1[r.CardID.Val] = append(m1[r.CardID.Val], r)
	}
	for _, r := range rows {
		if m1[r.ID] != nil {
			r.DeckCards = m1[r.ID]
		} else {
			r.DeckCards = make([]*DeckCard, 0)
		}
	}

	ids1 := []int64{}
	for _, r := range rows1 {
		ids1 = append(ids1, r.DeckID.Val)
	}
	rows2, err := q.queries.QueryDeck().Where(q.schema.DeckID.In(ids1)).All()
	if err != nil {
		return err
	}
	m2 := make(map[int64]*Deck)
	for _, r := range rows2 {
		m2[r.ID] = r
	}
	for _, r := range rows1 {
		r.Deck = m2[r.DeckID.Val]
	}

	m3 := make(map[int64][]*Deck)
	for _, r := range rows1 {
		m3[r.CardID.Val] = append(m3[r.CardID.Val], r.Deck)
	}
	for _, r := range rows {
		if m3[r.ID] != nil {
			r.Deck = m3[r.ID]
		} else {
			r.Deck = make([]*Deck, 0)
		}
	}

	return nil
}

func (q *CardQuery) All() ([]*Card, error) {
	if q.err != nil {
		return nil, q.err
	}
	var rows []Card
	cards := make([]*Card, 0)
	query, args := q.selectStatement.ToSQL()
	err := q.adapter.Query(query, args...).Scan(&rows)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return cards, nil
	}

	for i := range rows {
		rows[i].schema = q.schema
		rows[i].queries = q.queries
		cards = append(cards, &rows[i])
	}
	if q.preload["deck_cards"] {
		if err := q.preloadDeckCards(cards); err != nil {
			return nil, err
		}
	}
	if q.preload["deck"] {
		if err := q.preloadDeck(cards); err != nil {
			return nil, err
		}
	}

	return cards, nil
}

func (q *CardQuery) First() (*Card, error) {
	q.Limit(1)
	results, err := q.All()
	if err != nil {
		return nil, err
	}
	if len(results) > 0 {
		return results[0], nil
	}

	return nil, nil
}

func (q *CardQuery) Count() (int64, error) {
	var res struct {
		Count int64 `db:"count"`
	}
	q.selectStatement.Select("count(*)")
	query, args := q.selectStatement.ToSQL()
	if err := q.adapter.QueryOne(query, args...).Scan(&res); err != nil {
		return 0, err
	}

	return res.Count, nil
}

func (q *CardQuery) Avg(v string) (float64, error) {
	var res struct {
		Avg sql.NullFloat64 `db:"avg"`
	}
	q.selectStatement.Select(fmt.Sprintf("avg(%+v)", v))
	query, args := q.selectStatement.ToSQL()
	if err := q.adapter.QueryOne(query, args...).Scan(&res); err != nil {
		return 0, err
	}

	return res.Avg.Float64, nil
}

func (q *CardQuery) Sum(v string) (float64, error) {
	var res struct {
		Sum sql.NullFloat64 `db:"sum"`
	}
	q.selectStatement.Select(fmt.Sprintf("sum(%+v)", v))
	query, args := q.selectStatement.ToSQL()
	if err := q.adapter.QueryOne(query, args...).Scan(&res); err != nil {
		return 0, err
	}

	return res.Sum.Float64, nil
}

func (q *CardQuery) Max(v string) (float64, error) {
	var res struct {
		Max sql.NullFloat64 `db:"max"`
	}
	q.selectStatement.Select(fmt.Sprintf("max(%+v)", v))
	query, args := q.selectStatement.ToSQL()
	if err := q.adapter.QueryOne(query, args...).Scan(&res); err != nil {
		return 0, err
	}

	return res.Max.Float64, nil
}

func (q *CardQuery) Min(v string) (float64, error) {
	var res struct {
		Min sql.NullFloat64 `db:"min"`
	}
	q.selectStatement.Select(fmt.Sprintf("min(%+v)", v))
	query, args := q.selectStatement.ToSQL()
	if err := q.adapter.QueryOne(query, args...).Scan(&res); err != nil {
		return 0, err
	}

	return res.Min.Float64, nil
}

func (q *CardQuery) Exists() (bool, error) {
	q.selectStatement.Select("1 AS one").Limit(1)
	// select 1 as one from users limit 1
	query, args := q.selectStatement.ToSQL()
	var res struct {
		One int64 `db:"one"`
	}
	if err := q.adapter.QueryOne(query, args...).Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return res.One == 1, nil
}

func (q *CardQuery) UpdateAll(change *queryx.CardChange) (int64, error) {
	if q.err != nil {
		return 0, q.err
	}
	now := queryx.Now("2006-01-02 15:04:05.000")
	if !change.UpdatedAt.Set {
		change.SetUpdatedAt(now)
	}
	columns, values := change.Changes()
	query, args := q.selectStatement.Update().Columns(columns...).Values(values...).ToSQL()
	result, err := q.adapter.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result, err
}

func (q *CardQuery) DeleteAll() (int64, error) {
	query, args := q.selectStatement.Delete().ToSQL()
	result, err := q.adapter.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result, err
}
