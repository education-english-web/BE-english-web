package pusherclient

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/education-english-web/BE-english-web/app/domain/entity"
)

func TestNewMember(t *testing.T) {
	memberID := uint64(1)
	userID := uint32(2)
	name := "name"
	email := "email"
	role := "pic"
	colorCode := "color_code"
	want := Member{
		ID:        memberID,
		UserID:    userID,
		Name:      name,
		Email:     email,
		Role:      role,
		ColorCode: colorCode,
	}
	got := NewMember(memberID, userID, name, email, role, colorCode)

	assert.Equal(t, want, got)
}

func TestNewUpdateMemberEventMessage(t *testing.T) {
	oldPIC := Member{
		Name:  "old_pic",
		Email: "old_pic_email",
	}
	newPIC := Member{
		Name:  "new_pic",
		Email: "new_pic_email",
	}
	member := Member{
		Name:  "",
		Email: "",
	}
	want := UpdateMemberEventMessage{
		OldPIC:    &oldPIC,
		NewPIC:    &newPIC,
		Member:    &member,
		Action:    entity.ProposalMemberActionRemoveMember.String(),
		IsDeleted: true,
	}
	got := NewUpdateMemberEventMessage(&newPIC, &oldPIC, &member, entity.ProposalMemberActionRemoveMember, true)

	assert.Equal(t, want, got)
}

func TestNewUpdateMemberEventMessage_ToMap(t *testing.T) {
	oldPIC := Member{
		ID:        uint64(1),
		UserID:    uint32(2),
		Name:      "old_pic",
		Email:     "old_pic_email",
		ColorCode: "red",
	}
	newPIC := Member{
		ID:        uint64(3),
		UserID:    uint32(4),
		Name:      "new_pic",
		Email:     "new_pic_email",
		ColorCode: "green",
	}
	member := Member{
		ID:        uint64(5),
		UserID:    uint32(6),
		Name:      "member",
		Email:     "member_email",
		ColorCode: "blue",
	}
	want := map[string]interface{}{
		"old_pic": map[string]interface{}{
			"id":         float64(1),
			"user_id":    float64(2),
			"name":       "old_pic",
			"email":      "old_pic_email",
			"color_code": "red",
		},
		"new_pic": map[string]interface{}{
			"id":         float64(3),
			"user_id":    float64(4),
			"name":       "new_pic",
			"email":      "new_pic_email",
			"color_code": "green",
		},
		"member": map[string]interface{}{
			"id":         float64(5),
			"user_id":    float64(6),
			"name":       "member",
			"email":      "member_email",
			"color_code": "blue",
		},
		"action":     entity.ProposalMemberActionUpdatePIC.String(),
		"is_deleted": true,
	}
	got := NewUpdateMemberEventMessage(&newPIC, &oldPIC, &member, entity.ProposalMemberActionUpdatePIC, true).ToMap()

	assert.Equal(t, want, got)
}

func TestNewUpdateMembersEventMessage(t *testing.T) {
	members := []Member{
		{
			Name:  "",
			Email: "",
		},
	}
	want := UpdateMembersEventMessage{
		Members: members,
		Action:  entity.ProposalMemberActionRemoveMember.String(),
	}
	got := NewUpdateMembersEventMessage(members, entity.ProposalMemberActionRemoveMember)

	assert.Equal(t, want, got)
}

func TestNewUpdateMembersEventMessage_ToMap(t *testing.T) {
	members := []Member{
		{
			ID:        uint64(5),
			UserID:    uint32(6),
			Name:      "member",
			Email:     "member_email",
			ColorCode: "blue",
		},
	}
	want := map[string]interface{}{
		"members": []interface{}{
			map[string]interface{}{
				"id":         float64(5),
				"user_id":    float64(6),
				"name":       "member",
				"email":      "member_email",
				"color_code": "blue",
			},
		},
		"action": entity.ProposalMemberActionUpdatePIC.String(),
	}
	got := NewUpdateMembersEventMessage(members, entity.ProposalMemberActionUpdatePIC).ToMap()

	assert.Equal(t, want, got)
}

func TestNewUpdateDocumentEventMessage(t *testing.T) {
	document := entity.ProposalMessageDocument{
		ID:              "",
		Name:            "",
		ExternalVersion: nil,
		InternalVersion: nil,
		FileType:        "",
	}
	want := UpdateDocumentEventMessage{
		Document: document,
		Action:   entity.ProposalMessageDocumentActionRemove.String(),
	}
	got := NewUpdateDocumentEventMessage(document, entity.ProposalMessageDocumentActionRemove)

	assert.Equal(t, want, got)
}

func TestNewUpdateDocumentEventMessage_ToMap(t *testing.T) {
	externalVersion := uint32(0)
	internalVersion := uint32(1)
	document := entity.ProposalMessageDocument{
		ID:              "document_id",
		Name:            "document name",
		ExternalVersion: &externalVersion,
		InternalVersion: &internalVersion,
		FileType:        entity.FileTypeDOCX.String(),
	}
	want := map[string]interface{}{
		"document": map[string]interface{}{
			"id":               "document_id",
			"name":             "document name",
			"external_version": float64(0),
			"internal_version": float64(1),
			"file_type":        entity.FileTypeDOCX.String(),
		},
		"action": entity.ProposalMessageDocumentActionAdd.String(),
	}
	got := NewUpdateDocumentEventMessage(document, entity.ProposalMessageDocumentActionAdd).ToMap()

	assert.Equal(t, want, got)
}

func TestNewDeleteMessageEventMessage(t *testing.T) {
	want := DeleteMessageEventMessage{
		messageID:         "message_id",
		isMessageDeleted:  true,
		isDocumentDeleted: true,
	}
	got := NewDeleteMessageEventMessage("message_id", true, true, nil)

	assert.Equal(t, want, got)
}

func TestDeleteMessageEventMessage_ToMap(t *testing.T) {
	attachmentID := "deleted-attachment-id"
	want := map[string]interface{}{
		"id":                    "message_id",
		"is_message_deleted":    true,
		"is_document_deleted":   true,
		"deleted_attachment_id": &attachmentID,
	}
	got := NewDeleteMessageEventMessage("message_id", true, true, &attachmentID).ToMap()

	assert.Equal(t, want, got)
}
