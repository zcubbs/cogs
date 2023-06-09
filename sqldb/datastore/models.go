// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package datastore

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// app stores data about applications that interact with the system
type App struct {
	// The Unique ID for the table.
	AppID uuid.UUID
	// The unique application External ID to be given to outside callers.
	AppExtlID string
	// The Foreign key for the organization that the app belongs to.
	OrgID uuid.UUID
	// The application name is a short name for the application.
	AppName string
	// The application description is several sentences to describe the application.
	AppDescription string
	// unique identifier representing authorization provider (e.g. Google, Github, etc.)
	AuthProviderID sql.NullInt32
	// Unique identifer of client ID given by an authentication provider. For example, GCP supports cross-client identity - see https://developers.google.com/identity/protocols/oauth2/cross-client-identity for a great explanation.
	AuthProviderClientID sql.NullString
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

type AppApiKey struct {
	// app_key is a hash of a key given to an person for an app
	ApiKey string
	// foreign key to app table
	AppID           uuid.UUID
	DeactvDate      time.Time
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

// The auth table stores which user has authenticated through an Oauth2 provider.
type Auth struct {
	// The unique id given to the authorization.
	AuthID uuid.UUID
	UserID uuid.UUID
	// Unique ID given to an authorization provider.
	AuthProviderID int64
	// Unique code given to an authorization provider (e.g. google).
	AuthProviderCd string
	// External ID (given by authorization provider) which represents the Oauth2 client which authenticated the user
	AuthProviderClientID sql.NullString
	// Unique ID given by the authorization provider which represents the person.
	AuthProviderPersonID string
	// Oauth2 access token given by the authorization provider.
	AuthProviderAccessToken string
	// OAuth2 refresh token given by the authorization provider.
	AuthProviderRefreshToken sql.NullString
	// Expiration of access token given by the authorization provider. Is not a perfect precision instrument as some providers do not give an exact time, but rather seconds until expiration, which means the value is calculated relative to the server time.
	AuthProviderAccessTokenExpiry time.Time
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// Authentication Provider (e.g. Google, Github, Apple, Facebook, etc.)
type AuthProvider struct {
	// Unique ID representing the authentication provider.
	AuthProviderID int64
	// Short code representing the authentication provider (e.g., google, github, apple, etc.)
	AuthProviderCd string
	// Longer description of authentication provider
	AuthProviderDesc string
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// The movie table stores details about a movie.
type Movie struct {
	// The unique ID given to the movie.
	MovieID uuid.UUID
	// A unique ID given to the movie which can be used externally.
	ExtlID string
	// The title of the movie.
	Title string
	// The movie rating (PG, PG-13, R, etc.)
	Rated sql.NullString
	// The date the movie was released.
	Released sql.NullTime
	// The movie run time in minutes.
	RunTime sql.NullInt32
	// The movie director.
	Director sql.NullString
	// The movie writer.
	Writer sql.NullString
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

type Org struct {
	// Organization ID - Unique ID for table
	OrgID uuid.UUID
	// Organization Unique External ID to be given to outside callers.
	OrgExtlID string
	// Organization Name - a short name for the organization
	OrgName string
	// Organization Description - several sentences to describe the organization
	OrgDescription string
	// Foreign Key to org_kind table.
	OrgKindID uuid.UUID
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// Organization Kind is a reference table denoting an organization's (org) classification. Examples are Genesis, Test, Standard
type OrgKind struct {
	// Organization Kind ID - pk for table
	OrgKindID uuid.UUID
	// A short code denoting the organization kind
	OrgKindExtlID string
	// A longer descriptor of the organization kind
	OrgKindDesc string
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// The permission table stores an approval of a mode of access to a resource.
type Permission struct {
	// The unique ID for the table.
	PermissionID uuid.UUID
	// Unique External ID to be given to outside callers.
	PermissionExtlID string
	// A human-readable string which represents a resource (e.g. an HTTP route or document, etc.).
	Resource string
	// A string representing the action taken on the resource (e.g. POST, GET, edit, etc.)
	Operation string
	// A description of what the permission is granting, e.g. "grants ability to edit a billing document".
	PermissionDescription string
	// A boolean denoting whether the permission is active (true) or not (false).
	Active bool
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

type Person struct {
	// The user ID is the unique ID for user (pk for table)
	PersonID uuid.UUID
	// The unique user external ID to be given to outside callers.
	PersonExtlID string
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// The role table stores a job function or title which defines an authority level.
type Role struct {
	// The unique ID for the table.
	RoleID uuid.UUID
	// Unique External ID to be given to outside callers.
	RoleExtlID string
	// A human-readable code which represents the role.
	RoleCd string
	// A longer description of the role.
	RoleDescription string
	// A boolean denoting whether the role is active (true) or not (false).
	Active bool
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// The role_permission table stores which roles have which permissions.
type RolePermission struct {
	// The unique role which can have 1 to many permissions set in this table.
	RoleID uuid.UUID
	// The unique permission that is being given to the role.
	PermissionID uuid.UUID
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// users stores data about users that interact with the system. A user is a person who utilizes a computer or network service." In the context of this project, given that we allow Persons to authenticate with multiple providers, a User is akin to a user (Wikipedia - "The word persona derives from Latin, where it originally referred to a theatrical mask. On the social web, users develop virtual personas as online identities.") and as such, a Person can have one to many Users (for instance, I can have a GitHub user and a Google user, but I am just one Person). As a general, practical matter, most operations are considered at the User level. For instance, roles are assigned at the user level instead of the Person level, which allows for more fine-grained access control. Architecture note: All tables are to be singular, however, because user is a reserved word, the rules are broken here. It is unfortunate, but the alternatives are no better.
type User struct {
	UserID     uuid.UUID
	UserExtlID string
	PersonID   uuid.UUID
	NamePrefix sql.NullString
	FirstName  string
	MiddleName sql.NullString
	LastName   string
	NameSuffix sql.NullString
	Nickname   sql.NullString
	// Primary email for the user
	Email           sql.NullString
	CompanyName     sql.NullString
	CompanyDept     sql.NullString
	JobTitle        sql.NullString
	BirthDate       sql.NullTime
	BirthYear       sql.NullInt64
	BirthMonth      sql.NullInt64
	BirthDay        sql.NullInt64
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

// The users_lang_prefs table stores the list of language tag preferences for the user.
type UsersLangPref struct {
	// The user having a language tag preference.
	UserID uuid.UUID
	// The BCP 47 Language Tag which identifies a language both spoken and written.
	LanguageTag string
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

type UsersOrg struct {
	// Unique identifier for a user's association with an organization
	UsersOrgID uuid.UUID
	// Organization ID foreign key to org table
	OrgID uuid.UUID
	// User ID foreign key to users table
	UserID uuid.UUID
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}

// The users_role table stores which users have which role(s) in which organization(s).
type UsersRole struct {
	// The user which is being given a role (within an organization).
	UserID uuid.UUID
	// The role which can have one to many users set in this table.
	RoleID uuid.UUID
	// The organization to which the role and user are associated.
	OrgID uuid.UUID
	// The application which created this record.
	CreateAppID uuid.UUID
	// The user which created this record.
	CreateUserID uuid.NullUUID
	// The timestamp when this record was created.
	CreateTimestamp time.Time
	// The application which performed the most recent update to this record.
	UpdateAppID uuid.UUID
	// The user which performed the most recent update to this record.
	UpdateUserID uuid.NullUUID
	// The timestamp when the record was updated most recently.
	UpdateTimestamp time.Time
}
