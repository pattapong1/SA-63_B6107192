// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/pattapong1/app/ent/club"
	"github.com/pattapong1/app/ent/location"
	"github.com/pattapong1/app/ent/predicate"
)

// LocationUpdate is the builder for updating Location entities.
type LocationUpdate struct {
	config
	hooks      []Hook
	mutation   *LocationMutation
	predicates []predicate.Location
}

// Where adds a new predicate for the builder.
func (lu *LocationUpdate) Where(ps ...predicate.Location) *LocationUpdate {
	lu.predicates = append(lu.predicates, ps...)
	return lu
}

// SetCLUBELOCATIONNAME sets the CLUBE_LOCATION_NAME field.
func (lu *LocationUpdate) SetCLUBELOCATIONNAME(s string) *LocationUpdate {
	lu.mutation.SetCLUBELOCATIONNAME(s)
	return lu
}

// SetCLUBELOCATIONADDRESS sets the CLUBE_LOCATION_ADDRESS field.
func (lu *LocationUpdate) SetCLUBELOCATIONADDRESS(s string) *LocationUpdate {
	lu.mutation.SetCLUBELOCATIONADDRESS(s)
	return lu
}

// AddClubIDs adds the club edge to Club by ids.
func (lu *LocationUpdate) AddClubIDs(ids ...int) *LocationUpdate {
	lu.mutation.AddClubIDs(ids...)
	return lu
}

// AddClub adds the club edges to Club.
func (lu *LocationUpdate) AddClub(c ...*Club) *LocationUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.AddClubIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lu *LocationUpdate) Mutation() *LocationMutation {
	return lu.mutation
}

// RemoveClubIDs removes the club edge to Club by ids.
func (lu *LocationUpdate) RemoveClubIDs(ids ...int) *LocationUpdate {
	lu.mutation.RemoveClubIDs(ids...)
	return lu
}

// RemoveClub removes club edges to Club.
func (lu *LocationUpdate) RemoveClub(c ...*Club) *LocationUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return lu.RemoveClubIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (lu *LocationUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := lu.mutation.CLUBELOCATIONNAME(); ok {
		if err := location.CLUBELOCATIONNAMEValidator(v); err != nil {
			return 0, &ValidationError{Name: "CLUBE_LOCATION_NAME", err: fmt.Errorf("ent: validator failed for field \"CLUBE_LOCATION_NAME\": %w", err)}
		}
	}
	if v, ok := lu.mutation.CLUBELOCATIONADDRESS(); ok {
		if err := location.CLUBELOCATIONADDRESSValidator(v); err != nil {
			return 0, &ValidationError{Name: "CLUBE_LOCATION_ADDRESS", err: fmt.Errorf("ent: validator failed for field \"CLUBE_LOCATION_ADDRESS\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LocationUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LocationUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LocationUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LocationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   location.Table,
			Columns: location.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: location.FieldID,
			},
		},
	}
	if ps := lu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CLUBELOCATIONNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: location.FieldCLUBELOCATIONNAME,
		})
	}
	if value, ok := lu.mutation.CLUBELOCATIONADDRESS(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: location.FieldCLUBELOCATIONADDRESS,
		})
	}
	if nodes := lu.mutation.RemovedClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ClubTable,
			Columns: []string{location.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ClubTable,
			Columns: []string{location.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LocationUpdateOne is the builder for updating a single Location entity.
type LocationUpdateOne struct {
	config
	hooks    []Hook
	mutation *LocationMutation
}

// SetCLUBELOCATIONNAME sets the CLUBE_LOCATION_NAME field.
func (luo *LocationUpdateOne) SetCLUBELOCATIONNAME(s string) *LocationUpdateOne {
	luo.mutation.SetCLUBELOCATIONNAME(s)
	return luo
}

// SetCLUBELOCATIONADDRESS sets the CLUBE_LOCATION_ADDRESS field.
func (luo *LocationUpdateOne) SetCLUBELOCATIONADDRESS(s string) *LocationUpdateOne {
	luo.mutation.SetCLUBELOCATIONADDRESS(s)
	return luo
}

// AddClubIDs adds the club edge to Club by ids.
func (luo *LocationUpdateOne) AddClubIDs(ids ...int) *LocationUpdateOne {
	luo.mutation.AddClubIDs(ids...)
	return luo
}

// AddClub adds the club edges to Club.
func (luo *LocationUpdateOne) AddClub(c ...*Club) *LocationUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.AddClubIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (luo *LocationUpdateOne) Mutation() *LocationMutation {
	return luo.mutation
}

// RemoveClubIDs removes the club edge to Club by ids.
func (luo *LocationUpdateOne) RemoveClubIDs(ids ...int) *LocationUpdateOne {
	luo.mutation.RemoveClubIDs(ids...)
	return luo
}

// RemoveClub removes club edges to Club.
func (luo *LocationUpdateOne) RemoveClub(c ...*Club) *LocationUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luo.RemoveClubIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (luo *LocationUpdateOne) Save(ctx context.Context) (*Location, error) {
	if v, ok := luo.mutation.CLUBELOCATIONNAME(); ok {
		if err := location.CLUBELOCATIONNAMEValidator(v); err != nil {
			return nil, &ValidationError{Name: "CLUBE_LOCATION_NAME", err: fmt.Errorf("ent: validator failed for field \"CLUBE_LOCATION_NAME\": %w", err)}
		}
	}
	if v, ok := luo.mutation.CLUBELOCATIONADDRESS(); ok {
		if err := location.CLUBELOCATIONADDRESSValidator(v); err != nil {
			return nil, &ValidationError{Name: "CLUBE_LOCATION_ADDRESS", err: fmt.Errorf("ent: validator failed for field \"CLUBE_LOCATION_ADDRESS\": %w", err)}
		}
	}

	var (
		err  error
		node *Location
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LocationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LocationUpdateOne) SaveX(ctx context.Context) *Location {
	l, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return l
}

// Exec executes the query on the entity.
func (luo *LocationUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LocationUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LocationUpdateOne) sqlSave(ctx context.Context) (l *Location, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   location.Table,
			Columns: location.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: location.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Location.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := luo.mutation.CLUBELOCATIONNAME(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: location.FieldCLUBELOCATIONNAME,
		})
	}
	if value, ok := luo.mutation.CLUBELOCATIONADDRESS(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: location.FieldCLUBELOCATIONADDRESS,
		})
	}
	if nodes := luo.mutation.RemovedClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ClubTable,
			Columns: []string{location.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.ClubIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ClubTable,
			Columns: []string{location.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	l = &Location{config: luo.config}
	_spec.Assign = l.assignValues
	_spec.ScanValues = l.scanValues()
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{location.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return l, nil
}
