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

func NewApi() Api {
	return Api{}
}

type filterFunc struct {
	endpoint String
	args     []string
}

type Schema struct {
}

type Api struct {
	sl schemaLoader
}

type SchemaLoader interface {
	Load(entity string) Schema
}

// Call any endpoint
func (a Api) Call(endpoint string, args ...interface{}) Set {
	q.fcs = append(filterFunc{
		endpoint: endpoint,
		args:     args,
	})
	return q
}

type Entity struct {
	api        Api
	entityName entityName
}

// Read by Id
func (a Api) Entity(entityName string) Entity {
	return Entity{
		api:        a,
		entityName: entityName,
	}
}

func (e Entity) Read(ids ...string) Set {

}

func (e Entity) ListAll() Set {

}

// Resolves foreign keys
func (e Entity) Resolve(fields string, set Set) Set {

}
