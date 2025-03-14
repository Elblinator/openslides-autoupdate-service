package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestGroupModeA(t *testing.T) {
	var g collection.Group

	testCase(
		"no perms",
		t,
		g.Modes("A"),
		false,
		`---
		group/1/meeting_id: 30
		meeting/30/id: 30
		meeting/30/committee_id: 404
		`,
	)

	testCase(
		"public access enabled",
		t,
		g.Modes("A"),
		true,
		`---
		organization/1/enable_anonymous: true
		group/1/meeting_id: 30
		meeting/30/enable_anonymous: true
		`,
	)

	testCase(
		"In meeting",
		t,
		g.Modes("A"),
		true,
		`---
		group/1:
			meeting_id: 30
			meeting_user_ids: [50]
		meeting_user/50:
			user_id: 5
			meeting_id: 30
		meeting/30:
			group_ids: [1]
			committee_id: 2
		user/5/meeting_user_ids: [50]
		`,
		withRequestUser(5),
	)

	testCase(
		"Organization Manager in locked meeting",
		t,
		g.Modes("A"),
		true,
		`---
		group/1:
			meeting_id: 30
		meeting/30:
			group_ids: [1]
			committee_id: 2
			locked_from_inside: true
		user/5/organization_management_level: can_manage_organization
		`,
		withRequestUser(5),
	)
}
