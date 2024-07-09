// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
package user

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Create creates a new user.
func Create(ctx context.Context, params *CreateParams) (*clerk.User, error) {
	return getClient().Create(ctx, params)
}

// Get retrieves details about the user.
func Get(ctx context.Context, id string) (*clerk.User, error) {
	return getClient().Get(ctx, id)
}

// Update updates a user.
func Update(ctx context.Context, id string, params *UpdateParams) (*clerk.User, error) {
	return getClient().Update(ctx, id, params)
}

// UpdateProfileImage sets or replaces the users's profile image.
func UpdateProfileImage(ctx context.Context, id string, params *UpdateProfileImageParams) (*clerk.User, error) {
	return getClient().UpdateProfileImage(ctx, id, params)
}

// UpdateMetadata updates the user's metadata by merging the
// provided values with the existing ones.
func UpdateMetadata(ctx context.Context, id string, params *UpdateMetadataParams) (*clerk.User, error) {
	return getClient().UpdateMetadata(ctx, id, params)
}

// Delete deletes a user.
func Delete(ctx context.Context, id string) (*clerk.DeletedResource, error) {
	return getClient().Delete(ctx, id)
}

// List returns a list of users.
func List(ctx context.Context, params *ListParams) (*clerk.UserList, error) {
	return getClient().List(ctx, params)
}

// Count returns the total count of users satisfying the parameters.
func Count(ctx context.Context, params *ListParams) (*TotalCount, error) {
	return getClient().Count(ctx, params)
}

// ListOAuthAccessTokens retrieves a list of the user's access
// tokens for a specific OAuth provider.
func ListOAuthAccessTokens(ctx context.Context, params *ListOAuthAccessTokensParams) (*clerk.OAuthAccessTokenList, error) {
	return getClient().ListOAuthAccessTokens(ctx, params)
}

// DeleteMFA disables a user's multi-factor authentication methods.
func DeleteMFA(ctx context.Context, params *DeleteMFAParams) (*MultifactorAuthentication, error) {
	return getClient().DeleteMFA(ctx, params)
}

// Ban marks the user as banned.
func Ban(ctx context.Context, id string) (*clerk.User, error) {
	return getClient().Ban(ctx, id)
}

// Unban removes the ban for a user.
func Unban(ctx context.Context, id string) (*clerk.User, error) {
	return getClient().Unban(ctx, id)
}

// Lock marks the user as locked.
func Lock(ctx context.Context, id string) (*clerk.User, error) {
	return getClient().Lock(ctx, id)
}

// Unlock removes the lock for a user.
func Unlock(ctx context.Context, id string) (*clerk.User, error) {
	return getClient().Unlock(ctx, id)
}

// ListOrganizationMemberships lists all the user's organization memberships.
func ListOrganizationMemberships(ctx context.Context, id string, params *ListOrganizationMembershipsParams) (*clerk.OrganizationMembershipList, error) {
	return getClient().ListOrganizationMemberships(ctx, id, params)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
