package squery

import (
	"encoding/json"
)

type Document struct {
	Name   string                 // name of the entity
	Fields map[string]interface{} // json document itself
}

func (d Document) ToSet() Set {
	return Set([]Document{d})
}

type Set []Document

// Extract ids out of a set. Dedupes.
func (s Set) Ids() []string {
	return []string{}
}

// "Type alias" (not really...)
type M map[string]interface{}

type ApiAccessor interface {
	GetSchemas(entity string) map[string]Schema
}

func NewApi(ai ApiAccessor) Api {
	return Api{}
}

type filterFunc struct {
	endpoint string
	args     []string
}

type Schema struct {
}

type Api struct {
	aa ApiAccessor
}

// Call any endpoint. Use sparingly, should be used for actions which are
// modifying data, not reads. For reads see the Query Api.
//
// Returns Set nevertheless, you know... just in case you are such a rebel.
func (a Api) Call(endpoint string, args ...interface{}) Set {
	return Set{}
}

// Resolves foreign keys
func (a Api) Resolve(entity, fields string, set Set) Set {
	return Set{}
}

type between struct {
	from interface{}
	to   interface{}
}

type Query struct {
	api        Api
	entityName string
	ids        []string
	between    map[string]between       // fieldnames to betweens
	equals     map[string][]interface{} // fieldnames to values
	refersTo   []Query                  // entitynames to queries
}

// Returns a Query. An empty Query represents the whole set.
func (a Api) Query(entityName string) Query {
	return Query{
		api:        a,
		entityName: entityName,
		ids:        []string{},
		between:    map[string]between{},
		equals:     map[string][]interface{}{},
		refersTo:   []Query{},
	}
}

// Special case of Equals
func (q Query) Id(ids ...string) Query {
	q.ids = ids
	return q
}

func (q Query) MarshalJSON() ([]byte, error) {
	v := M{
		"entityName": q.entityName,
	}
	if len(q.ids) > 0 {
		v["ids"] = q.ids
	}
	if len(q.between) > 0 {
		v["between"] = q.between
	}
	if len(q.equals) > 0 {
		v["equals"] = q.equals
	}
	if len(q.refersTo) > 0 {
		v["refersTo"] = q.refersTo
	}
	return json.Marshal(v)
}

func (q Query) Between(fieldName string, from, to interface{}) Query {
	q.between[fieldName] = between{
		from: from,
		to:   to,
	}
	return q
}

func (q Query) Equals(fieldName string, value ...interface{}) Query {
	q.equals[fieldName] = value
	return q
}

// A "foreign key" in the entity
func (q Query) RefersTo(query Query) Query {
	q.refersTo = append(q.refersTo, query)
	return q
}

// Return all hits
func (q Query) All() (Set, error) {
	return Set{}, nil
}

// Classic paging
func (q Query) Page(skip, limit int) (Set, error) {
	return Set{}, nil
}

// Cassandra-like paging
func (q Query) Stream(start, end string, limit int) (Set, error) {
	return Set{}, nil
}
