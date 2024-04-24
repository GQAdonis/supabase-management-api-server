// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// Activate implements activate operation.
	//
	// Activates a custom hostname for a project.
	//
	// POST /v1/projects/{ref}/custom-hostname/activate
	Activate(ctx context.Context, params ActivateParams) (ActivateRes, error)
	// ActivateVanitySubdomainPlease implements activateVanitySubdomainPlease operation.
	//
	// Activates a vanity subdomain for a project.
	//
	// POST /v1/projects/{ref}/vanity-subdomain/activate
	ActivateVanitySubdomainPlease(ctx context.Context, req *VanitySubdomainBody, params ActivateVanitySubdomainPleaseParams) (ActivateVanitySubdomainPleaseRes, error)
	// ApplyNetworkRestrictions implements applyNetworkRestrictions operation.
	//
	// Updates project's network restrictions.
	//
	// POST /v1/projects/{ref}/network-restrictions/apply
	ApplyNetworkRestrictions(ctx context.Context, req *NetworkRestrictionsRequest, params ApplyNetworkRestrictionsParams) (ApplyNetworkRestrictionsRes, error)
	// Authorize implements authorize operation.
	//
	// Authorize user through oauth.
	//
	// GET /v1/oauth/authorize
	Authorize(ctx context.Context, params AuthorizeParams) error
	// CheckServiceHealth implements checkServiceHealth operation.
	//
	// Gets project's service health status.
	//
	// GET /v1/projects/{ref}/health
	CheckServiceHealth(ctx context.Context, params CheckServiceHealthParams) (CheckServiceHealthRes, error)
	// CheckVanitySubdomainAvailability implements checkVanitySubdomainAvailability operation.
	//
	// Checks vanity subdomain availability.
	//
	// POST /v1/projects/{ref}/vanity-subdomain/check-availability
	CheckVanitySubdomainAvailability(ctx context.Context, req *VanitySubdomainBody, params CheckVanitySubdomainAvailabilityParams) (CheckVanitySubdomainAvailabilityRes, error)
	// CreateBranch implements createBranch operation.
	//
	// Creates a database branch from the specified project.
	//
	// POST /v1/projects/{ref}/branches
	CreateBranch(ctx context.Context, req *CreateBranchBody, params CreateBranchParams) (CreateBranchRes, error)
	// CreateCustomHostnameConfig implements createCustomHostnameConfig operation.
	//
	// Updates project's custom hostname configuration.
	//
	// POST /v1/projects/{ref}/custom-hostname/initialize
	CreateCustomHostnameConfig(ctx context.Context, req *UpdateCustomHostnameBody, params CreateCustomHostnameConfigParams) (CreateCustomHostnameConfigRes, error)
	// CreateFunction implements createFunction operation.
	//
	// Creates a function and adds it to the specified project.
	//
	// POST /v1/projects/{ref}/functions
	CreateFunction(ctx context.Context, req *V1CreateFunctionBody, params CreateFunctionParams) (CreateFunctionRes, error)
	// CreateOrganization implements createOrganization operation.
	//
	// Create an organization.
	//
	// POST /v1/organizations
	CreateOrganization(ctx context.Context, req *CreateOrganizationBodyV1) (CreateOrganizationRes, error)
	// CreateProject implements createProject operation.
	//
	// Create a project.
	//
	// POST /v1/projects
	CreateProject(ctx context.Context, req *V1CreateProjectBody) (*V1ProjectResponse, error)
	// CreateProviderForProject implements createProviderForProject operation.
	//
	// Creates a new SSO provider.
	//
	// POST /v1/projects/{ref}/config/auth/sso/providers
	CreateProviderForProject(ctx context.Context, req *CreateProviderBody, params CreateProviderForProjectParams) (CreateProviderForProjectRes, error)
	// CreateSecrets implements createSecrets operation.
	//
	// Creates multiple secrets and adds them to the specified project.
	//
	// POST /v1/projects/{ref}/secrets
	CreateSecrets(ctx context.Context, req []CreateSecretBody, params CreateSecretsParams) (CreateSecretsRes, error)
	// CreateTPAForProject implements createTPAForProject operation.
	//
	// Creates a new third-party auth integration.
	//
	// POST /v1/projects/{ref}/config/auth/third-party-auth
	CreateTPAForProject(ctx context.Context, req *CreateThirdPartyAuthBody, params CreateTPAForProjectParams) (CreateTPAForProjectRes, error)
	// DeleteBranch implements deleteBranch operation.
	//
	// Deletes the specified database branch.
	//
	// DELETE /v1/branches/{branch_id}
	DeleteBranch(ctx context.Context, params DeleteBranchParams) (DeleteBranchRes, error)
	// DeleteFunction implements deleteFunction operation.
	//
	// Deletes a function with the specified slug from the specified project.
	//
	// DELETE /v1/projects/{ref}/functions/{function_slug}
	DeleteFunction(ctx context.Context, params DeleteFunctionParams) (DeleteFunctionRes, error)
	// DeleteProject implements deleteProject operation.
	//
	// Deletes the given project.
	//
	// DELETE /v1/projects/{ref}
	DeleteProject(ctx context.Context, params DeleteProjectParams) (DeleteProjectRes, error)
	// DeleteSecrets implements deleteSecrets operation.
	//
	// Deletes all secrets with the given names from the specified project.
	//
	// DELETE /v1/projects/{ref}/secrets
	DeleteSecrets(ctx context.Context, req []string, params DeleteSecretsParams) (DeleteSecretsRes, error)
	// DeleteTPAForProject implements deleteTPAForProject operation.
	//
	// Removes a third-party auth integration.
	//
	// DELETE /v1/projects/{ref}/config/auth/third-party-auth/{tpa_id}
	DeleteTPAForProject(ctx context.Context, params DeleteTPAForProjectParams) (DeleteTPAForProjectRes, error)
	// DisableBranch implements disableBranch operation.
	//
	// Disables preview branching for the specified project.
	//
	// DELETE /v1/projects/{ref}/branches
	DisableBranch(ctx context.Context, params DisableBranchParams) (DisableBranchRes, error)
	// GetBackups implements getBackups operation.
	//
	// Lists all backups.
	//
	// GET /v1/projects/{ref}/database/backups
	GetBackups(ctx context.Context, params GetBackupsParams) (GetBackupsRes, error)
	// GetBranchDetails implements getBranchDetails operation.
	//
	// Fetches configurations of the specified database branch.
	//
	// GET /v1/branches/{branch_id}
	GetBranchDetails(ctx context.Context, params GetBranchDetailsParams) (GetBranchDetailsRes, error)
	// GetBranches implements getBranches operation.
	//
	// Returns all database branches of the specified project.
	//
	// GET /v1/projects/{ref}/branches
	GetBranches(ctx context.Context, params GetBranchesParams) (GetBranchesRes, error)
	// GetBuckets implements getBuckets operation.
	//
	// Lists all buckets.
	//
	// GET /v1/projects/{ref}/storage/buckets
	GetBuckets(ctx context.Context, params GetBucketsParams) (GetBucketsRes, error)
	// GetConfig implements getConfig operation.
	//
	// Gets project's Postgres config.
	//
	// GET /v1/projects/{ref}/config/database/postgres
	GetConfig(ctx context.Context, params GetConfigParams) (GetConfigRes, error)
	// GetCustomHostnameConfig implements getCustomHostnameConfig operation.
	//
	// Gets project's custom hostname config.
	//
	// GET /v1/projects/{ref}/custom-hostname
	GetCustomHostnameConfig(ctx context.Context, params GetCustomHostnameConfigParams) (GetCustomHostnameConfigRes, error)
	// GetFunction implements getFunction operation.
	//
	// Retrieves a function with the specified slug and project.
	//
	// GET /v1/projects/{ref}/functions/{function_slug}
	GetFunction(ctx context.Context, params GetFunctionParams) (GetFunctionRes, error)
	// GetFunctionBody implements getFunctionBody operation.
	//
	// Retrieves a function body for the specified slug and project.
	//
	// GET /v1/projects/{ref}/functions/{function_slug}/body
	GetFunctionBody(ctx context.Context, params GetFunctionBodyParams) (GetFunctionBodyRes, error)
	// GetFunctions implements getFunctions operation.
	//
	// Returns all functions you've previously added to the specified project.
	//
	// GET /v1/projects/{ref}/functions
	GetFunctions(ctx context.Context, params GetFunctionsParams) (GetFunctionsRes, error)
	// GetNetworkBans implements getNetworkBans operation.
	//
	// Gets project's network bans.
	//
	// POST /v1/projects/{ref}/network-bans/retrieve
	GetNetworkBans(ctx context.Context, params GetNetworkBansParams) (GetNetworkBansRes, error)
	// GetNetworkRestrictions implements getNetworkRestrictions operation.
	//
	// Gets project's network restrictions.
	//
	// GET /v1/projects/{ref}/network-restrictions
	GetNetworkRestrictions(ctx context.Context, params GetNetworkRestrictionsParams) (GetNetworkRestrictionsRes, error)
	// GetOrganization implements getOrganization operation.
	//
	// Gets information about the organization.
	//
	// GET /v1/organizations/{slug}
	GetOrganization(ctx context.Context, params GetOrganizationParams) (*V1OrganizationSlugResponse, error)
	// GetOrganizations implements getOrganizations operation.
	//
	// Returns a list of organizations that you currently belong to.
	//
	// GET /v1/organizations
	GetOrganizations(ctx context.Context) (GetOrganizationsRes, error)
	// GetPgsodiumConfig implements getPgsodiumConfig operation.
	//
	// Gets project's pgsodium config.
	//
	// GET /v1/projects/{ref}/pgsodium
	GetPgsodiumConfig(ctx context.Context, params GetPgsodiumConfigParams) (GetPgsodiumConfigRes, error)
	// GetPostgRESTConfig implements getPostgRESTConfig operation.
	//
	// Gets project's postgrest config.
	//
	// GET /v1/projects/{ref}/postgrest
	GetPostgRESTConfig(ctx context.Context, params GetPostgRESTConfigParams) (GetPostgRESTConfigRes, error)
	// GetProjectApiKeys implements getProjectApiKeys operation.
	//
	// Get project api keys.
	//
	// GET /v1/projects/{ref}/api-keys
	GetProjectApiKeys(ctx context.Context, params GetProjectApiKeysParams) (GetProjectApiKeysRes, error)
	// GetProjects implements getProjects operation.
	//
	// Returns a list of all projects you've previously created.
	//
	// GET /v1/projects
	GetProjects(ctx context.Context) ([]V1ProjectResponse, error)
	// GetProviderById implements getProviderById operation.
	//
	// Gets a SSO provider by its UUID.
	//
	// GET /v1/projects/{ref}/config/auth/sso/providers/{provider_id}
	GetProviderById(ctx context.Context, params GetProviderByIdParams) (GetProviderByIdRes, error)
	// GetReadOnlyModeStatus implements getReadOnlyModeStatus operation.
	//
	// Returns project's readonly mode status.
	//
	// GET /v1/projects/{ref}/readonly
	GetReadOnlyModeStatus(ctx context.Context, params GetReadOnlyModeStatusParams) (GetReadOnlyModeStatusRes, error)
	// GetSecrets implements getSecrets operation.
	//
	// Returns all secrets you've previously added to the specified project.
	//
	// GET /v1/projects/{ref}/secrets
	GetSecrets(ctx context.Context, params GetSecretsParams) (GetSecretsRes, error)
	// GetSnippet implements getSnippet operation.
	//
	// Gets a specific SQL snippet.
	//
	// GET /v1/snippets/{id}
	GetSnippet(ctx context.Context, params GetSnippetParams) (GetSnippetRes, error)
	// GetSslEnforcementConfig implements getSslEnforcementConfig operation.
	//
	// Get project's SSL enforcement configuration.
	//
	// GET /v1/projects/{ref}/ssl-enforcement
	GetSslEnforcementConfig(ctx context.Context, params GetSslEnforcementConfigParams) (GetSslEnforcementConfigRes, error)
	// GetTPAForProject implements getTPAForProject operation.
	//
	// Get a third-party integration.
	//
	// GET /v1/projects/{ref}/config/auth/third-party-auth/{tpa_id}
	GetTPAForProject(ctx context.Context, params GetTPAForProjectParams) (GetTPAForProjectRes, error)
	// GetTypescriptTypes implements getTypescriptTypes operation.
	//
	// Returns the TypeScript types of your schema for use with supabase-js.
	//
	// GET /v1/projects/{ref}/types/typescript
	GetTypescriptTypes(ctx context.Context, params GetTypescriptTypesParams) (GetTypescriptTypesRes, error)
	// GetUpgradeStatus implements getUpgradeStatus operation.
	//
	// Gets the latest status of the project's upgrade.
	//
	// GET /v1/projects/{ref}/upgrade/status
	GetUpgradeStatus(ctx context.Context, params GetUpgradeStatusParams) (GetUpgradeStatusRes, error)
	// GetV1AuthConfig implements getV1AuthConfig operation.
	//
	// Gets project's auth config.
	//
	// GET /v1/projects/{ref}/config/auth
	GetV1AuthConfig(ctx context.Context, params GetV1AuthConfigParams) (GetV1AuthConfigRes, error)
	// GetVanitySubdomainConfig implements getVanitySubdomainConfig operation.
	//
	// Gets current vanity subdomain config.
	//
	// GET /v1/projects/{ref}/vanity-subdomain
	GetVanitySubdomainConfig(ctx context.Context, params GetVanitySubdomainConfigParams) (GetVanitySubdomainConfigRes, error)
	// ListAllProviders implements listAllProviders operation.
	//
	// Lists all SSO providers.
	//
	// GET /v1/projects/{ref}/config/auth/sso/providers
	ListAllProviders(ctx context.Context, params ListAllProvidersParams) (ListAllProvidersRes, error)
	// ListSnippets implements listSnippets operation.
	//
	// Lists SQL snippets for the logged in user.
	//
	// GET /v1/snippets
	ListSnippets(ctx context.Context, params ListSnippetsParams) (ListSnippetsRes, error)
	// ListTPAForProject implements listTPAForProject operation.
	//
	// Lists all third-party auth integrations.
	//
	// GET /v1/projects/{ref}/config/auth/third-party-auth
	ListTPAForProject(ctx context.Context, params ListTPAForProjectParams) (ListTPAForProjectRes, error)
	// RemoveCustomHostnameConfig implements removeCustomHostnameConfig operation.
	//
	// Deletes a project's custom hostname configuration.
	//
	// DELETE /v1/projects/{ref}/custom-hostname
	RemoveCustomHostnameConfig(ctx context.Context, params RemoveCustomHostnameConfigParams) (RemoveCustomHostnameConfigRes, error)
	// RemoveNetworkBan implements removeNetworkBan operation.
	//
	// Remove network bans.
	//
	// DELETE /v1/projects/{ref}/network-bans
	RemoveNetworkBan(ctx context.Context, req *RemoveNetworkBanRequest, params RemoveNetworkBanParams) (RemoveNetworkBanRes, error)
	// RemoveProviderById implements removeProviderById operation.
	//
	// Removes a SSO provider by its UUID.
	//
	// DELETE /v1/projects/{ref}/config/auth/sso/providers/{provider_id}
	RemoveProviderById(ctx context.Context, params RemoveProviderByIdParams) (RemoveProviderByIdRes, error)
	// RemoveReadReplica implements removeReadReplica operation.
	//
	// Remove a read replica.
	//
	// POST /v1/projects/{ref}/read-replicas/remove
	RemoveReadReplica(ctx context.Context, req *RemoveReadReplicaBody, params RemoveReadReplicaParams) (RemoveReadReplicaRes, error)
	// RemoveVanitySubdomainConfig implements removeVanitySubdomainConfig operation.
	//
	// Deletes a project's vanity subdomain configuration.
	//
	// DELETE /v1/projects/{ref}/vanity-subdomain
	RemoveVanitySubdomainConfig(ctx context.Context, params RemoveVanitySubdomainConfigParams) (RemoveVanitySubdomainConfigRes, error)
	// ResetBranch implements resetBranch operation.
	//
	// Resets the specified database branch.
	//
	// POST /v1/branches/{branch_id}/reset
	ResetBranch(ctx context.Context, params ResetBranchParams) (ResetBranchRes, error)
	// Reverify implements reverify operation.
	//
	// Attempts to verify the DNS configuration for project's custom hostname configuration.
	//
	// POST /v1/projects/{ref}/custom-hostname/reverify
	Reverify(ctx context.Context, params ReverifyParams) (ReverifyRes, error)
	// SetUpReadReplica implements setUpReadReplica operation.
	//
	// Set up a read replica.
	//
	// POST /v1/projects/{ref}/read-replicas/setup
	SetUpReadReplica(ctx context.Context, req *SetUpReadReplicaBody, params SetUpReadReplicaParams) (SetUpReadReplicaRes, error)
	// TemporarilyDisableReadonlyMode implements temporarilyDisableReadonlyMode operation.
	//
	// Disables project's readonly mode for the next 15 minutes.
	//
	// POST /v1/projects/{ref}/readonly/temporary-disable
	TemporarilyDisableReadonlyMode(ctx context.Context, params TemporarilyDisableReadonlyModeParams) (TemporarilyDisableReadonlyModeRes, error)
	// Token implements token operation.
	//
	// Exchange auth code for user's access and refresh token.
	//
	// POST /v1/oauth/token
	Token(ctx context.Context, req *OAuthTokenBody) (*OAuthTokenResponse, error)
	// UpdateBranch implements updateBranch operation.
	//
	// Updates the configuration of the specified database branch.
	//
	// PATCH /v1/branches/{branch_id}
	UpdateBranch(ctx context.Context, req *UpdateBranchBody, params UpdateBranchParams) (UpdateBranchRes, error)
	// UpdateConfig implements updateConfig operation.
	//
	// Updates project's Postgres config.
	//
	// PUT /v1/projects/{ref}/config/database/postgres
	UpdateConfig(ctx context.Context, req *UpdatePostgresConfigBody, params UpdateConfigParams) (UpdateConfigRes, error)
	// UpdateFunction implements updateFunction operation.
	//
	// Updates a function with the specified slug and project.
	//
	// PATCH /v1/projects/{ref}/functions/{function_slug}
	UpdateFunction(ctx context.Context, req *V1UpdateFunctionBody, params UpdateFunctionParams) (UpdateFunctionRes, error)
	// UpdatePgsodiumConfig implements updatePgsodiumConfig operation.
	//
	// Updates project's pgsodium config. Updating the root_key can cause all data encrypted with the
	// older key to become inaccessible.
	//
	// PUT /v1/projects/{ref}/pgsodium
	UpdatePgsodiumConfig(ctx context.Context, req *UpdatePgsodiumConfigBody, params UpdatePgsodiumConfigParams) (UpdatePgsodiumConfigRes, error)
	// UpdatePostgRESTConfig implements updatePostgRESTConfig operation.
	//
	// Updates project's postgrest config.
	//
	// PATCH /v1/projects/{ref}/postgrest
	UpdatePostgRESTConfig(ctx context.Context, req *UpdatePostgrestConfigBody, params UpdatePostgRESTConfigParams) (UpdatePostgRESTConfigRes, error)
	// UpdateProviderById implements updateProviderById operation.
	//
	// Updates a SSO provider by its UUID.
	//
	// PUT /v1/projects/{ref}/config/auth/sso/providers/{provider_id}
	UpdateProviderById(ctx context.Context, req *UpdateProviderBody, params UpdateProviderByIdParams) (UpdateProviderByIdRes, error)
	// UpdateSslEnforcementConfig implements updateSslEnforcementConfig operation.
	//
	// Update project's SSL enforcement configuration.
	//
	// PUT /v1/projects/{ref}/ssl-enforcement
	UpdateSslEnforcementConfig(ctx context.Context, req *SslEnforcementRequest, params UpdateSslEnforcementConfigParams) (UpdateSslEnforcementConfigRes, error)
	// UpdateV1AuthConfig implements updateV1AuthConfig operation.
	//
	// Updates a project's auth config.
	//
	// PATCH /v1/projects/{ref}/config/auth
	UpdateV1AuthConfig(ctx context.Context, req *UpdateAuthConfigBody, params UpdateV1AuthConfigParams) (UpdateV1AuthConfigRes, error)
	// UpgradeEligibilityInformation implements upgradeEligibilityInformation operation.
	//
	// Returns the project's eligibility for upgrades.
	//
	// GET /v1/projects/{ref}/upgrade/eligibility
	UpgradeEligibilityInformation(ctx context.Context, params UpgradeEligibilityInformationParams) (UpgradeEligibilityInformationRes, error)
	// UpgradeProject implements upgradeProject operation.
	//
	// Upgrades the project's Postgres version.
	//
	// POST /v1/projects/{ref}/upgrade
	UpgradeProject(ctx context.Context, req *UpgradeDatabaseBody, params UpgradeProjectParams) (UpgradeProjectRes, error)
	// V1EnableDatabaseWebhooks implements v1EnableDatabaseWebhooks operation.
	//
	// Enables Database Webhooks on the project.
	//
	// POST /v1/projects/{ref}/database/webhooks/enable
	V1EnableDatabaseWebhooks(ctx context.Context, params V1EnableDatabaseWebhooksParams) (V1EnableDatabaseWebhooksRes, error)
	// V1GetPgbouncerConfig implements v1GetPgbouncerConfig operation.
	//
	// Get project's pgbouncer config.
	//
	// GET /v1/projects/{ref}/config/database/pgbouncer
	V1GetPgbouncerConfig(ctx context.Context, params V1GetPgbouncerConfigParams) (V1GetPgbouncerConfigRes, error)
	// V1ListOrganizationMembers implements v1ListOrganizationMembers operation.
	//
	// List members of an organization.
	//
	// GET /v1/organizations/{slug}/members
	V1ListOrganizationMembers(ctx context.Context, params V1ListOrganizationMembersParams) ([]V1OrganizationMemberResponse, error)
	// V1RestorePitr implements v1RestorePitr operation.
	//
	// Restores a PITR backup for a database.
	//
	// POST /v1/projects/{ref}/database/backups/restore-pitr
	V1RestorePitr(ctx context.Context, req *V1RestorePitrBody, params V1RestorePitrParams) error
	// V1RunQuery implements v1RunQuery operation.
	//
	// Run sql query.
	//
	// POST /v1/projects/{ref}/database/query
	V1RunQuery(ctx context.Context, req *V1RunQueryBody, params V1RunQueryParams) (V1RunQueryRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}