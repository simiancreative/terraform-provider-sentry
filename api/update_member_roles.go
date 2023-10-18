package api

import (
	"context"
	"fmt"

	"github.com/jianyuan/go-sentry/v2/sentry"
)

type TeamMemberRole struct {
	Team string `json:"teamSlug"`
	Role string `json:"role"`
}

type OrganizationMemberRole struct {
	Role  string           `json:"orgRole"`
	Teams []TeamMemberRole `json:"teamRoles,omitempty"`
}

func UpdateMemberRoles(c *sentry.Client, ctx context.Context, organizationSlug string, memberID string, params *OrganizationMemberRole) (*sentry.OrganizationMember, *sentry.Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/%v/", organizationSlug, memberID)
	req, err := c.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	member := new(sentry.OrganizationMember)
	resp, err := c.Do(ctx, req, member)
	if err != nil {
		return nil, resp, err
	}
	return member, resp, nil
}
