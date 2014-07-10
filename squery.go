package squery

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
	GetSchema(entity string) Schema
	Call(endpoint string, args ...interface{}) Set
	// CallWith(endpoint string, doc interface{}) Set
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

type Query struct {
	api        Api
	entityName string
}

// Returns a Query. An empty Query represents the whole set.
func (a Api) Query(entityName string) Query {
	return Query{
		api:        a,
		entityName: entityName,
	}
}

// Id in
func (q Query) Id(ids ...string) Query {
	return Query{}
}

// Field between
func (q Query) Between(fieldName string, from, to interface{}) Query {
	return Query{}
}

// Field equals
func (q Query) Equals(fieldName string, value ...interface{}) Query {
	return Query{}
}

// Looks like Equals but in fact this can used relations stored
// separately of any of the two involved entities.
func (e Query) Fk(entity string, entityIds ...string) Query {
	return Query{}
}

// Produces a result set
func (q Query) All() Set {
	return Set{}
}
