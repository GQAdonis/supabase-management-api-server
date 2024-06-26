// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"tribemedia.io/m/ent/organization"
	"tribemedia.io/m/ent/predicate"
	"tribemedia.io/m/ent/provider"
)

// ProviderUpdate is the builder for updating Provider entities.
type ProviderUpdate struct {
	config
	hooks    []Hook
	mutation *ProviderMutation
}

// Where appends a list predicates to the ProviderUpdate builder.
func (pu *ProviderUpdate) Where(ps ...predicate.Provider) *ProviderUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetType sets the "type" field.
func (pu *ProviderUpdate) SetType(s string) *ProviderUpdate {
	pu.mutation.SetType(s)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableType(s *string) *ProviderUpdate {
	if s != nil {
		pu.SetType(*s)
	}
	return pu
}

// SetMetadataXML sets the "metadata_xml" field.
func (pu *ProviderUpdate) SetMetadataXML(s string) *ProviderUpdate {
	pu.mutation.SetMetadataXML(s)
	return pu
}

// SetNillableMetadataXML sets the "metadata_xml" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableMetadataXML(s *string) *ProviderUpdate {
	if s != nil {
		pu.SetMetadataXML(*s)
	}
	return pu
}

// ClearMetadataXML clears the value of the "metadata_xml" field.
func (pu *ProviderUpdate) ClearMetadataXML() *ProviderUpdate {
	pu.mutation.ClearMetadataXML()
	return pu
}

// SetMetadataURL sets the "metadata_url" field.
func (pu *ProviderUpdate) SetMetadataURL(s string) *ProviderUpdate {
	pu.mutation.SetMetadataURL(s)
	return pu
}

// SetNillableMetadataURL sets the "metadata_url" field if the given value is not nil.
func (pu *ProviderUpdate) SetNillableMetadataURL(s *string) *ProviderUpdate {
	if s != nil {
		pu.SetMetadataURL(*s)
	}
	return pu
}

// ClearMetadataURL clears the value of the "metadata_url" field.
func (pu *ProviderUpdate) ClearMetadataURL() *ProviderUpdate {
	pu.mutation.ClearMetadataURL()
	return pu
}

// SetDomains sets the "domains" field.
func (pu *ProviderUpdate) SetDomains(s []string) *ProviderUpdate {
	pu.mutation.SetDomains(s)
	return pu
}

// AppendDomains appends s to the "domains" field.
func (pu *ProviderUpdate) AppendDomains(s []string) *ProviderUpdate {
	pu.mutation.AppendDomains(s)
	return pu
}

// SetAttributeMapping sets the "attribute_mapping" field.
func (pu *ProviderUpdate) SetAttributeMapping(m map[string]string) *ProviderUpdate {
	pu.mutation.SetAttributeMapping(m)
	return pu
}

// ClearAttributeMapping clears the value of the "attribute_mapping" field.
func (pu *ProviderUpdate) ClearAttributeMapping() *ProviderUpdate {
	pu.mutation.ClearAttributeMapping()
	return pu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (pu *ProviderUpdate) SetOrganizationID(id int) *ProviderUpdate {
	pu.mutation.SetOrganizationID(id)
	return pu
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (pu *ProviderUpdate) SetNillableOrganizationID(id *int) *ProviderUpdate {
	if id != nil {
		pu = pu.SetOrganizationID(*id)
	}
	return pu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (pu *ProviderUpdate) SetOrganization(o *Organization) *ProviderUpdate {
	return pu.SetOrganizationID(o.ID)
}

// Mutation returns the ProviderMutation object of the builder.
func (pu *ProviderUpdate) Mutation() *ProviderMutation {
	return pu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (pu *ProviderUpdate) ClearOrganization() *ProviderUpdate {
	pu.mutation.ClearOrganization()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProviderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProviderUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProviderUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(provider.FieldType, field.TypeString, value)
	}
	if value, ok := pu.mutation.MetadataXML(); ok {
		_spec.SetField(provider.FieldMetadataXML, field.TypeString, value)
	}
	if pu.mutation.MetadataXMLCleared() {
		_spec.ClearField(provider.FieldMetadataXML, field.TypeString)
	}
	if value, ok := pu.mutation.MetadataURL(); ok {
		_spec.SetField(provider.FieldMetadataURL, field.TypeString, value)
	}
	if pu.mutation.MetadataURLCleared() {
		_spec.ClearField(provider.FieldMetadataURL, field.TypeString)
	}
	if value, ok := pu.mutation.Domains(); ok {
		_spec.SetField(provider.FieldDomains, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedDomains(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, provider.FieldDomains, value)
		})
	}
	if value, ok := pu.mutation.AttributeMapping(); ok {
		_spec.SetField(provider.FieldAttributeMapping, field.TypeJSON, value)
	}
	if pu.mutation.AttributeMappingCleared() {
		_spec.ClearField(provider.FieldAttributeMapping, field.TypeJSON)
	}
	if pu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provider.OrganizationTable,
			Columns: []string{provider.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provider.OrganizationTable,
			Columns: []string{provider.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProviderUpdateOne is the builder for updating a single Provider entity.
type ProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProviderMutation
}

// SetType sets the "type" field.
func (puo *ProviderUpdateOne) SetType(s string) *ProviderUpdateOne {
	puo.mutation.SetType(s)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableType(s *string) *ProviderUpdateOne {
	if s != nil {
		puo.SetType(*s)
	}
	return puo
}

// SetMetadataXML sets the "metadata_xml" field.
func (puo *ProviderUpdateOne) SetMetadataXML(s string) *ProviderUpdateOne {
	puo.mutation.SetMetadataXML(s)
	return puo
}

// SetNillableMetadataXML sets the "metadata_xml" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableMetadataXML(s *string) *ProviderUpdateOne {
	if s != nil {
		puo.SetMetadataXML(*s)
	}
	return puo
}

// ClearMetadataXML clears the value of the "metadata_xml" field.
func (puo *ProviderUpdateOne) ClearMetadataXML() *ProviderUpdateOne {
	puo.mutation.ClearMetadataXML()
	return puo
}

// SetMetadataURL sets the "metadata_url" field.
func (puo *ProviderUpdateOne) SetMetadataURL(s string) *ProviderUpdateOne {
	puo.mutation.SetMetadataURL(s)
	return puo
}

// SetNillableMetadataURL sets the "metadata_url" field if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableMetadataURL(s *string) *ProviderUpdateOne {
	if s != nil {
		puo.SetMetadataURL(*s)
	}
	return puo
}

// ClearMetadataURL clears the value of the "metadata_url" field.
func (puo *ProviderUpdateOne) ClearMetadataURL() *ProviderUpdateOne {
	puo.mutation.ClearMetadataURL()
	return puo
}

// SetDomains sets the "domains" field.
func (puo *ProviderUpdateOne) SetDomains(s []string) *ProviderUpdateOne {
	puo.mutation.SetDomains(s)
	return puo
}

// AppendDomains appends s to the "domains" field.
func (puo *ProviderUpdateOne) AppendDomains(s []string) *ProviderUpdateOne {
	puo.mutation.AppendDomains(s)
	return puo
}

// SetAttributeMapping sets the "attribute_mapping" field.
func (puo *ProviderUpdateOne) SetAttributeMapping(m map[string]string) *ProviderUpdateOne {
	puo.mutation.SetAttributeMapping(m)
	return puo
}

// ClearAttributeMapping clears the value of the "attribute_mapping" field.
func (puo *ProviderUpdateOne) ClearAttributeMapping() *ProviderUpdateOne {
	puo.mutation.ClearAttributeMapping()
	return puo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (puo *ProviderUpdateOne) SetOrganizationID(id int) *ProviderUpdateOne {
	puo.mutation.SetOrganizationID(id)
	return puo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (puo *ProviderUpdateOne) SetNillableOrganizationID(id *int) *ProviderUpdateOne {
	if id != nil {
		puo = puo.SetOrganizationID(*id)
	}
	return puo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (puo *ProviderUpdateOne) SetOrganization(o *Organization) *ProviderUpdateOne {
	return puo.SetOrganizationID(o.ID)
}

// Mutation returns the ProviderMutation object of the builder.
func (puo *ProviderUpdateOne) Mutation() *ProviderMutation {
	return puo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (puo *ProviderUpdateOne) ClearOrganization() *ProviderUpdateOne {
	puo.mutation.ClearOrganization()
	return puo
}

// Where appends a list predicates to the ProviderUpdate builder.
func (puo *ProviderUpdateOne) Where(ps ...predicate.Provider) *ProviderUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProviderUpdateOne) Select(field string, fields ...string) *ProviderUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Provider entity.
func (puo *ProviderUpdateOne) Save(ctx context.Context) (*Provider, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProviderUpdateOne) SaveX(ctx context.Context) *Provider {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProviderUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProviderUpdateOne) sqlSave(ctx context.Context) (_node *Provider, err error) {
	_spec := sqlgraph.NewUpdateSpec(provider.Table, provider.Columns, sqlgraph.NewFieldSpec(provider.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Provider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, provider.FieldID)
		for _, f := range fields {
			if !provider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != provider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(provider.FieldType, field.TypeString, value)
	}
	if value, ok := puo.mutation.MetadataXML(); ok {
		_spec.SetField(provider.FieldMetadataXML, field.TypeString, value)
	}
	if puo.mutation.MetadataXMLCleared() {
		_spec.ClearField(provider.FieldMetadataXML, field.TypeString)
	}
	if value, ok := puo.mutation.MetadataURL(); ok {
		_spec.SetField(provider.FieldMetadataURL, field.TypeString, value)
	}
	if puo.mutation.MetadataURLCleared() {
		_spec.ClearField(provider.FieldMetadataURL, field.TypeString)
	}
	if value, ok := puo.mutation.Domains(); ok {
		_spec.SetField(provider.FieldDomains, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedDomains(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, provider.FieldDomains, value)
		})
	}
	if value, ok := puo.mutation.AttributeMapping(); ok {
		_spec.SetField(provider.FieldAttributeMapping, field.TypeJSON, value)
	}
	if puo.mutation.AttributeMappingCleared() {
		_spec.ClearField(provider.FieldAttributeMapping, field.TypeJSON)
	}
	if puo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provider.OrganizationTable,
			Columns: []string{provider.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provider.OrganizationTable,
			Columns: []string{provider.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Provider{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
