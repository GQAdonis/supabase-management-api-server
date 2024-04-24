// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"tribemedia.io/m/ent/authconfig"
	"tribemedia.io/m/ent/branch"
	"tribemedia.io/m/ent/customhostname"
	"tribemedia.io/m/ent/function"
	"tribemedia.io/m/ent/networkban"
	"tribemedia.io/m/ent/organization"
	"tribemedia.io/m/ent/pgsodiumconfig"
	"tribemedia.io/m/ent/project"
	"tribemedia.io/m/ent/schema"
	"tribemedia.io/m/ent/secret"
	"tribemedia.io/m/ent/typescripttype"
	"tribemedia.io/m/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authconfigFields := schema.AuthConfig{}.Fields()
	_ = authconfigFields
	// authconfigDescID is the schema descriptor for id field.
	authconfigDescID := authconfigFields[0].Descriptor()
	// authconfig.DefaultID holds the default value on creation for the id field.
	authconfig.DefaultID = authconfigDescID.Default.(func() uuid.UUID)
	branchFields := schema.Branch{}.Fields()
	_ = branchFields
	// branchDescID is the schema descriptor for id field.
	branchDescID := branchFields[0].Descriptor()
	// branch.DefaultID holds the default value on creation for the id field.
	branch.DefaultID = branchDescID.Default.(func() uuid.UUID)
	customhostnameFields := schema.CustomHostname{}.Fields()
	_ = customhostnameFields
	// customhostnameDescID is the schema descriptor for id field.
	customhostnameDescID := customhostnameFields[0].Descriptor()
	// customhostname.DefaultID holds the default value on creation for the id field.
	customhostname.DefaultID = customhostnameDescID.Default.(func() uuid.UUID)
	functionFields := schema.Function{}.Fields()
	_ = functionFields
	// functionDescID is the schema descriptor for id field.
	functionDescID := functionFields[0].Descriptor()
	// function.DefaultID holds the default value on creation for the id field.
	function.DefaultID = functionDescID.Default.(func() uuid.UUID)
	networkbanFields := schema.NetworkBan{}.Fields()
	_ = networkbanFields
	// networkbanDescCreatedAt is the schema descriptor for created_at field.
	networkbanDescCreatedAt := networkbanFields[4].Descriptor()
	// networkban.DefaultCreatedAt holds the default value on creation for the created_at field.
	networkban.DefaultCreatedAt = networkbanDescCreatedAt.Default.(func() time.Time)
	// networkbanDescID is the schema descriptor for id field.
	networkbanDescID := networkbanFields[0].Descriptor()
	// networkban.DefaultID holds the default value on creation for the id field.
	networkban.DefaultID = networkbanDescID.Default.(func() uuid.UUID)
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[0].Descriptor()
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = organizationDescName.Validators[0].(func(string) error)
	pgsodiumconfigFields := schema.PgsodiumConfig{}.Fields()
	_ = pgsodiumconfigFields
	// pgsodiumconfigDescEnabled is the schema descriptor for enabled field.
	pgsodiumconfigDescEnabled := pgsodiumconfigFields[2].Descriptor()
	// pgsodiumconfig.DefaultEnabled holds the default value on creation for the enabled field.
	pgsodiumconfig.DefaultEnabled = pgsodiumconfigDescEnabled.Default.(bool)
	// pgsodiumconfigDescID is the schema descriptor for id field.
	pgsodiumconfigDescID := pgsodiumconfigFields[0].Descriptor()
	// pgsodiumconfig.DefaultID holds the default value on creation for the id field.
	pgsodiumconfig.DefaultID = pgsodiumconfigDescID.Default.(func() uuid.UUID)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescID is the schema descriptor for id field.
	projectDescID := projectFields[0].Descriptor()
	// project.DefaultID holds the default value on creation for the id field.
	project.DefaultID = projectDescID.Default.(func() uuid.UUID)
	secretFields := schema.Secret{}.Fields()
	_ = secretFields
	// secretDescID is the schema descriptor for id field.
	secretDescID := secretFields[0].Descriptor()
	// secret.DefaultID holds the default value on creation for the id field.
	secret.DefaultID = secretDescID.Default.(func() uuid.UUID)
	typescripttypeFields := schema.TypeScriptType{}.Fields()
	_ = typescripttypeFields
	// typescripttypeDescID is the schema descriptor for id field.
	typescripttypeDescID := typescripttypeFields[0].Descriptor()
	// typescripttype.DefaultID holds the default value on creation for the id field.
	typescripttype.DefaultID = typescripttypeDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}