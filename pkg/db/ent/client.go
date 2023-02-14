// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/third-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/third-manager/pkg/db/ent/contact"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/emailtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/frontendtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/smstemplate"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Contact is the client for interacting with the Contact builders.
	Contact *ContactClient
	// EmailTemplate is the client for interacting with the EmailTemplate builders.
	EmailTemplate *EmailTemplateClient
	// FrontendTemplate is the client for interacting with the FrontendTemplate builders.
	FrontendTemplate *FrontendTemplateClient
	// SMSTemplate is the client for interacting with the SMSTemplate builders.
	SMSTemplate *SMSTemplateClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Contact = NewContactClient(c.config)
	c.EmailTemplate = NewEmailTemplateClient(c.config)
	c.FrontendTemplate = NewFrontendTemplateClient(c.config)
	c.SMSTemplate = NewSMSTemplateClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Contact:          NewContactClient(cfg),
		EmailTemplate:    NewEmailTemplateClient(cfg),
		FrontendTemplate: NewFrontendTemplateClient(cfg),
		SMSTemplate:      NewSMSTemplateClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Contact:          NewContactClient(cfg),
		EmailTemplate:    NewEmailTemplateClient(cfg),
		FrontendTemplate: NewFrontendTemplateClient(cfg),
		SMSTemplate:      NewSMSTemplateClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Contact.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Contact.Use(hooks...)
	c.EmailTemplate.Use(hooks...)
	c.FrontendTemplate.Use(hooks...)
	c.SMSTemplate.Use(hooks...)
}

// ContactClient is a client for the Contact schema.
type ContactClient struct {
	config
}

// NewContactClient returns a client for the Contact from the given config.
func NewContactClient(c config) *ContactClient {
	return &ContactClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contact.Hooks(f(g(h())))`.
func (c *ContactClient) Use(hooks ...Hook) {
	c.hooks.Contact = append(c.hooks.Contact, hooks...)
}

// Create returns a builder for creating a Contact entity.
func (c *ContactClient) Create() *ContactCreate {
	mutation := newContactMutation(c.config, OpCreate)
	return &ContactCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contact entities.
func (c *ContactClient) CreateBulk(builders ...*ContactCreate) *ContactCreateBulk {
	return &ContactCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contact.
func (c *ContactClient) Update() *ContactUpdate {
	mutation := newContactMutation(c.config, OpUpdate)
	return &ContactUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContactClient) UpdateOne(co *Contact) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContact(co))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContactClient) UpdateOneID(id uuid.UUID) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContactID(id))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contact.
func (c *ContactClient) Delete() *ContactDelete {
	mutation := newContactMutation(c.config, OpDelete)
	return &ContactDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContactClient) DeleteOne(co *Contact) *ContactDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ContactClient) DeleteOneID(id uuid.UUID) *ContactDeleteOne {
	builder := c.Delete().Where(contact.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContactDeleteOne{builder}
}

// Query returns a query builder for Contact.
func (c *ContactClient) Query() *ContactQuery {
	return &ContactQuery{
		config: c.config,
	}
}

// Get returns a Contact entity by its id.
func (c *ContactClient) Get(ctx context.Context, id uuid.UUID) (*Contact, error) {
	return c.Query().Where(contact.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContactClient) GetX(ctx context.Context, id uuid.UUID) *Contact {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ContactClient) Hooks() []Hook {
	hooks := c.hooks.Contact
	return append(hooks[:len(hooks):len(hooks)], contact.Hooks[:]...)
}

// EmailTemplateClient is a client for the EmailTemplate schema.
type EmailTemplateClient struct {
	config
}

// NewEmailTemplateClient returns a client for the EmailTemplate from the given config.
func NewEmailTemplateClient(c config) *EmailTemplateClient {
	return &EmailTemplateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `emailtemplate.Hooks(f(g(h())))`.
func (c *EmailTemplateClient) Use(hooks ...Hook) {
	c.hooks.EmailTemplate = append(c.hooks.EmailTemplate, hooks...)
}

// Create returns a builder for creating a EmailTemplate entity.
func (c *EmailTemplateClient) Create() *EmailTemplateCreate {
	mutation := newEmailTemplateMutation(c.config, OpCreate)
	return &EmailTemplateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EmailTemplate entities.
func (c *EmailTemplateClient) CreateBulk(builders ...*EmailTemplateCreate) *EmailTemplateCreateBulk {
	return &EmailTemplateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EmailTemplate.
func (c *EmailTemplateClient) Update() *EmailTemplateUpdate {
	mutation := newEmailTemplateMutation(c.config, OpUpdate)
	return &EmailTemplateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmailTemplateClient) UpdateOne(et *EmailTemplate) *EmailTemplateUpdateOne {
	mutation := newEmailTemplateMutation(c.config, OpUpdateOne, withEmailTemplate(et))
	return &EmailTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmailTemplateClient) UpdateOneID(id uuid.UUID) *EmailTemplateUpdateOne {
	mutation := newEmailTemplateMutation(c.config, OpUpdateOne, withEmailTemplateID(id))
	return &EmailTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EmailTemplate.
func (c *EmailTemplateClient) Delete() *EmailTemplateDelete {
	mutation := newEmailTemplateMutation(c.config, OpDelete)
	return &EmailTemplateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EmailTemplateClient) DeleteOne(et *EmailTemplate) *EmailTemplateDeleteOne {
	return c.DeleteOneID(et.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EmailTemplateClient) DeleteOneID(id uuid.UUID) *EmailTemplateDeleteOne {
	builder := c.Delete().Where(emailtemplate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmailTemplateDeleteOne{builder}
}

// Query returns a query builder for EmailTemplate.
func (c *EmailTemplateClient) Query() *EmailTemplateQuery {
	return &EmailTemplateQuery{
		config: c.config,
	}
}

// Get returns a EmailTemplate entity by its id.
func (c *EmailTemplateClient) Get(ctx context.Context, id uuid.UUID) (*EmailTemplate, error) {
	return c.Query().Where(emailtemplate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmailTemplateClient) GetX(ctx context.Context, id uuid.UUID) *EmailTemplate {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EmailTemplateClient) Hooks() []Hook {
	hooks := c.hooks.EmailTemplate
	return append(hooks[:len(hooks):len(hooks)], emailtemplate.Hooks[:]...)
}

// FrontendTemplateClient is a client for the FrontendTemplate schema.
type FrontendTemplateClient struct {
	config
}

// NewFrontendTemplateClient returns a client for the FrontendTemplate from the given config.
func NewFrontendTemplateClient(c config) *FrontendTemplateClient {
	return &FrontendTemplateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `frontendtemplate.Hooks(f(g(h())))`.
func (c *FrontendTemplateClient) Use(hooks ...Hook) {
	c.hooks.FrontendTemplate = append(c.hooks.FrontendTemplate, hooks...)
}

// Create returns a builder for creating a FrontendTemplate entity.
func (c *FrontendTemplateClient) Create() *FrontendTemplateCreate {
	mutation := newFrontendTemplateMutation(c.config, OpCreate)
	return &FrontendTemplateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FrontendTemplate entities.
func (c *FrontendTemplateClient) CreateBulk(builders ...*FrontendTemplateCreate) *FrontendTemplateCreateBulk {
	return &FrontendTemplateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FrontendTemplate.
func (c *FrontendTemplateClient) Update() *FrontendTemplateUpdate {
	mutation := newFrontendTemplateMutation(c.config, OpUpdate)
	return &FrontendTemplateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FrontendTemplateClient) UpdateOne(ft *FrontendTemplate) *FrontendTemplateUpdateOne {
	mutation := newFrontendTemplateMutation(c.config, OpUpdateOne, withFrontendTemplate(ft))
	return &FrontendTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FrontendTemplateClient) UpdateOneID(id uuid.UUID) *FrontendTemplateUpdateOne {
	mutation := newFrontendTemplateMutation(c.config, OpUpdateOne, withFrontendTemplateID(id))
	return &FrontendTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FrontendTemplate.
func (c *FrontendTemplateClient) Delete() *FrontendTemplateDelete {
	mutation := newFrontendTemplateMutation(c.config, OpDelete)
	return &FrontendTemplateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FrontendTemplateClient) DeleteOne(ft *FrontendTemplate) *FrontendTemplateDeleteOne {
	return c.DeleteOneID(ft.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *FrontendTemplateClient) DeleteOneID(id uuid.UUID) *FrontendTemplateDeleteOne {
	builder := c.Delete().Where(frontendtemplate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FrontendTemplateDeleteOne{builder}
}

// Query returns a query builder for FrontendTemplate.
func (c *FrontendTemplateClient) Query() *FrontendTemplateQuery {
	return &FrontendTemplateQuery{
		config: c.config,
	}
}

// Get returns a FrontendTemplate entity by its id.
func (c *FrontendTemplateClient) Get(ctx context.Context, id uuid.UUID) (*FrontendTemplate, error) {
	return c.Query().Where(frontendtemplate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FrontendTemplateClient) GetX(ctx context.Context, id uuid.UUID) *FrontendTemplate {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FrontendTemplateClient) Hooks() []Hook {
	hooks := c.hooks.FrontendTemplate
	return append(hooks[:len(hooks):len(hooks)], frontendtemplate.Hooks[:]...)
}

// SMSTemplateClient is a client for the SMSTemplate schema.
type SMSTemplateClient struct {
	config
}

// NewSMSTemplateClient returns a client for the SMSTemplate from the given config.
func NewSMSTemplateClient(c config) *SMSTemplateClient {
	return &SMSTemplateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `smstemplate.Hooks(f(g(h())))`.
func (c *SMSTemplateClient) Use(hooks ...Hook) {
	c.hooks.SMSTemplate = append(c.hooks.SMSTemplate, hooks...)
}

// Create returns a builder for creating a SMSTemplate entity.
func (c *SMSTemplateClient) Create() *SMSTemplateCreate {
	mutation := newSMSTemplateMutation(c.config, OpCreate)
	return &SMSTemplateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SMSTemplate entities.
func (c *SMSTemplateClient) CreateBulk(builders ...*SMSTemplateCreate) *SMSTemplateCreateBulk {
	return &SMSTemplateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SMSTemplate.
func (c *SMSTemplateClient) Update() *SMSTemplateUpdate {
	mutation := newSMSTemplateMutation(c.config, OpUpdate)
	return &SMSTemplateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SMSTemplateClient) UpdateOne(st *SMSTemplate) *SMSTemplateUpdateOne {
	mutation := newSMSTemplateMutation(c.config, OpUpdateOne, withSMSTemplate(st))
	return &SMSTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SMSTemplateClient) UpdateOneID(id uuid.UUID) *SMSTemplateUpdateOne {
	mutation := newSMSTemplateMutation(c.config, OpUpdateOne, withSMSTemplateID(id))
	return &SMSTemplateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SMSTemplate.
func (c *SMSTemplateClient) Delete() *SMSTemplateDelete {
	mutation := newSMSTemplateMutation(c.config, OpDelete)
	return &SMSTemplateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SMSTemplateClient) DeleteOne(st *SMSTemplate) *SMSTemplateDeleteOne {
	return c.DeleteOneID(st.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SMSTemplateClient) DeleteOneID(id uuid.UUID) *SMSTemplateDeleteOne {
	builder := c.Delete().Where(smstemplate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SMSTemplateDeleteOne{builder}
}

// Query returns a query builder for SMSTemplate.
func (c *SMSTemplateClient) Query() *SMSTemplateQuery {
	return &SMSTemplateQuery{
		config: c.config,
	}
}

// Get returns a SMSTemplate entity by its id.
func (c *SMSTemplateClient) Get(ctx context.Context, id uuid.UUID) (*SMSTemplate, error) {
	return c.Query().Where(smstemplate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SMSTemplateClient) GetX(ctx context.Context, id uuid.UUID) *SMSTemplate {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SMSTemplateClient) Hooks() []Hook {
	hooks := c.hooks.SMSTemplate
	return append(hooks[:len(hooks):len(hooks)], smstemplate.Hooks[:]...)
}
