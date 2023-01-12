// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/breed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dog"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofilebreed"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/dogprofileowner"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/image"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/profile"
	"github.com/OSBC-LLC/togo-subgraph-main/ent/user"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     uuid.UUID `json:"id,omitempty"`     // node id.
	Type   string    `json:"type,omitempty"`   // node type.
	Fields []*Field  `json:"fields,omitempty"` // node fields.
	Edges  []*Edge   `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string      `json:"type,omitempty"` // edge type.
	Name string      `json:"name,omitempty"` // edge name.
	IDs  []uuid.UUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (b *Breed) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     b.ID,
		Type:   "Breed",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(b.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(b.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(b.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	return node, nil
}

func (d *Dog) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     d.ID,
		Type:   "Dog",
		Fields: make([]*Field, 9),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(d.FullName); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "full_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.Age); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "int",
		Name:  "age",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.WeightLbs); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "float64",
		Name:  "weight_lbs",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.WeightKgs); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "float64",
		Name:  "weight_kgs",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.Size); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "size",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.Birthday); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "birthday",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.DogImgID); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "uuid.UUID",
		Name:  "dog_img_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(d.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "DogProfileOwner",
		Name: "ownerProfiles",
	}
	err = d.QueryOwnerProfiles().
		Select(dogprofileowner.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (dpb *DogProfileBreed) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     dpb.ID,
		Type:   "DogProfileBreed",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(dpb.BreedID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "uuid.UUID",
		Name:  "breed_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dpb.DogID); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "uuid.UUID",
		Name:  "dog_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dpb.Percentage); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "float64",
		Name:  "percentage",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dpb.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dpb.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	return node, nil
}

func (dpo *DogProfileOwner) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     dpo.ID,
		Type:   "DogProfileOwner",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(dpo.OwnerID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "uuid.UUID",
		Name:  "owner_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dpo.DogID); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "uuid.UUID",
		Name:  "dog_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dpo.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(dpo.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "owner",
	}
	err = dpo.QueryOwner().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Dog",
		Name: "dog",
	}
	err = dpo.QueryDog().
		Select(dog.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (i *Image) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     i.ID,
		Type:   "Image",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(i.URL); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "url",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Width); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "int",
		Name:  "width",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Height); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "height",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.Type); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(i.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	return node, nil
}

func (pr *Profile) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pr.ID,
		Type:   "Profile",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(pr.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.Description); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pr.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "users",
	}
	err = pr.QueryUsers().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(u.FirstName); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "first_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.LastName); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "last_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UserImageID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "uuid.UUID",
		Name:  "user_image_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.ProfileID); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "uuid.UUID",
		Name:  "profile_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Profile",
		Name: "profile",
	}
	err = u.QueryProfile().
		Select(profile.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "DogProfileOwner",
		Name: "dogProfiles",
	}
	err = u.QueryDogProfiles().
		Select(dogprofileowner.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id uuid.UUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case breed.Table:
		query := c.Breed.Query().
			Where(breed.ID(id))
		query, err := query.CollectFields(ctx, "Breed")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case dog.Table:
		query := c.Dog.Query().
			Where(dog.ID(id))
		query, err := query.CollectFields(ctx, "Dog")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case dogprofilebreed.Table:
		query := c.DogProfileBreed.Query().
			Where(dogprofilebreed.ID(id))
		query, err := query.CollectFields(ctx, "DogProfileBreed")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case dogprofileowner.Table:
		query := c.DogProfileOwner.Query().
			Where(dogprofileowner.ID(id))
		query, err := query.CollectFields(ctx, "DogProfileOwner")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case image.Table:
		query := c.Image.Query().
			Where(image.ID(id))
		query, err := query.CollectFields(ctx, "Image")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case profile.Table:
		query := c.Profile.Query().
			Where(profile.ID(id))
		query, err := query.CollectFields(ctx, "Profile")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case breed.Table:
		query := c.Breed.Query().
			Where(breed.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Breed")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case dog.Table:
		query := c.Dog.Query().
			Where(dog.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Dog")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case dogprofilebreed.Table:
		query := c.DogProfileBreed.Query().
			Where(dogprofilebreed.IDIn(ids...))
		query, err := query.CollectFields(ctx, "DogProfileBreed")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case dogprofileowner.Table:
		query := c.DogProfileOwner.Query().
			Where(dogprofileowner.IDIn(ids...))
		query, err := query.CollectFields(ctx, "DogProfileOwner")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case image.Table:
		query := c.Image.Query().
			Where(image.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Image")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case profile.Table:
		query := c.Profile.Query().
			Where(profile.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Profile")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
