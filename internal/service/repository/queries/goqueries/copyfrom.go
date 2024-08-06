// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: copyfrom.go

package goqueries

import (
	"context"
)

// iteratorForAddTagsToSmth implements pgx.CopyFromSource.
type iteratorForAddTagsToSmth struct {
	rows                 []AddTagsToSmthParams
	skippedFirstNextCall bool
}

func (r *iteratorForAddTagsToSmth) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForAddTagsToSmth) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].SmthID,
		r.rows[0].TagID,
		r.rows[0].UserID,
	}, nil
}

func (r iteratorForAddTagsToSmth) Err() error {
	return nil
}

func (q *Queries) AddTagsToSmth(ctx context.Context, db DBTX, arg []AddTagsToSmthParams) (int64, error) {
	return db.CopyFrom(ctx, []string{"smth_to_tags"}, []string{"smth_id", "tag_id", "user_id"}, &iteratorForAddTagsToSmth{rows: arg})
}
