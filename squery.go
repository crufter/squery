package squery

type Document struct {
	Name   string                 // name of the entity
	Fields map[string]interface{} // json document itself
}

func (d Document) ToSet() Set {
	return Set([]Document{v})
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
	Call(endpoint, args ...interface{}) Set
	// CallWith(endpoint string, doc interface{}) Set
}

func NewApi(ai ApiAccessor) Api {
	return Api{}
}

type filterFunc struct {
	endpoint String
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
	q.fcs = append(filterFunc{
		endpoint: endpoint,
		args:     args,
	})
	return q
}

// Resolves foreign keys
func (a Api) Resolve(entity, fields string, set Set) Set {

}

type Entity struct {
	api        Api
	entityName entityName
}

// Returns a Query. An empty Query represents the whole set.
func (a Api) Query(entityName string) Query {
	return Entity{
		api:        a,
		entityName: entityName,
	}
}

// Id in
func (q Query) Id(ids ...string) Query {

}

// Field between
func (q Query) Between(fieldName string, from, to interface{}) Query {

}

// Field equals
func (q Query) Equals(fieldName string, value ...interface{}) Query {

}

// Looks like Equals but in fact this can used relations stored
// separately of any of the two involved entities.
func (e Query) Fk(entity string, entityIds ...string) Query {

}

// Produces a result set
func (q Query) All() Set {

}
