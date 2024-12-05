package pusherclient

import (
	"encoding/json"

	appLog "github.com/education-english-web/BE-english-web/pkg/log"
)

type Member struct {
	ID                    uint64   `json:"id,omitempty"`
	UserID                uint32   `json:"user_id,omitempty"`
	UserIDs               []uint32 `json:"user_ids,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Email                 string   `json:"email,omitempty"`
	Role                  string   `json:"role,omitempty"`
	ColorCode             string   `json:"color_code,omitempty"`
	UserGroupID           string   `json:"user_group_id,omitempty"`
	UserGroupName         string   `json:"user_group_name,omitempty"`
	UserGroupNumberMember uint32   `json:"user_group_number_member,omitempty"`
}

func NewMember(
	id uint64,
	userID uint32,
	name,
	email string,
	role string,
	colorCode string,
) Member {
	return Member{
		ID:        id,
		UserID:    userID,
		Name:      name,
		Email:     email,
		Role:      role,
		ColorCode: colorCode,
	}
}

type UpdateMemberEventMessage struct {
	NewPIC    *Member `json:"new_pic,omitempty"`
	OldPIC    *Member `json:"old_pic,omitempty"`
	Member    *Member `json:"member,omitempty"`
	Action    string  `json:"action"`
	IsDeleted bool    `json:"is_deleted,omitempty"`
}

//func NewUpdateMemberEventMessage(
//	newPIC,
//	oldPIC,
//	member *Member,
//	action entity.ProposalMemberAction,
//	isDeleted bool,
//) UpdateMemberEventMessage {
//	return UpdateMemberEventMessage{
//		NewPIC:    newPIC,
//		OldPIC:    oldPIC,
//		Member:    member,
//		Action:    action.String(),
//		IsDeleted: isDeleted,
//	}
//}

func (e UpdateMemberEventMessage) ToMap() map[string]interface{} {
	var result map[string]interface{}

	bs, err := json.Marshal(e)
	if err != nil {
		appLog.WithError(err).Errorln("UpdateMemberEventMessage marshal error")

		return nil
	}

	err = json.Unmarshal(bs, &result)
	if err != nil {
		appLog.WithError(err).Errorln("UpdateMemberEventMessage unmarshal error")

		return nil
	}

	return result
}

type UpdateMembersEventMessage struct {
	Members []Member `json:"members,omitempty"`
	Action  string   `json:"action"`
}

//func NewUpdateMembersEventMessage(members []Member, action entity.ProposalMemberAction) UpdateMembersEventMessage {
//	return UpdateMembersEventMessage{
//		Members: members,
//		Action:  action.String(),
//	}
//}

func (e UpdateMembersEventMessage) ToMap() map[string]interface{} {
	var result map[string]interface{}

	bs, err := json.Marshal(e)
	if err != nil {
		appLog.WithError(err).Errorln("UpdateMembersEventMessage marshal error")

		return nil
	}

	err = json.Unmarshal(bs, &result)
	if err != nil {
		appLog.WithError(err).Errorln("UpdateMembersEventMessage unmarshal error")

		return nil
	}

	return result
}

//type UpdateDocumentEventMessage struct {
//	Document entity.ProposalMessageDocument `json:"document"`
//	Action   string                         `json:"action"`
//}

//func NewUpdateDocumentEventMessage(
//	document entity.ProposalMessageDocument,
//	action entity.ProposalMessageDocumentAction,
//) UpdateDocumentEventMessage {
//	return UpdateDocumentEventMessage{
//		Document: document,
//		Action:   action.String(),
//	}
//}

//func (e UpdateDocumentEventMessage) ToMap() map[string]interface{} {
//	var result map[string]interface{}
//
//	bs, err := json.Marshal(e)
//	if err != nil {
//		appLog.WithError(err).Errorln("UpdateDocumentEventMessage marshal error")
//
//		return nil
//	}
//
//	err = json.Unmarshal(bs, &result)
//	if err != nil {
//		appLog.WithError(err).Errorln("UpdateDocumentEventMessage unmarshal error")
//
//		return nil
//	}
//
//	return result
//}

type DeleteMessageEventMessage struct {
	messageID           string
	isMessageDeleted    bool
	isDocumentDeleted   bool
	deletedAttachmentID *string
}

func NewDeleteMessageEventMessage(
	messageID string,
	isMessageDeleted,
	isDocumentDeleted bool,
	deletedAttachmentID *string,
) DeleteMessageEventMessage {
	return DeleteMessageEventMessage{
		messageID:           messageID,
		isMessageDeleted:    isMessageDeleted,
		isDocumentDeleted:   isDocumentDeleted,
		deletedAttachmentID: deletedAttachmentID,
	}
}

func (e DeleteMessageEventMessage) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                    e.messageID,
		"is_message_deleted":    e.isMessageDeleted,
		"is_document_deleted":   e.isDocumentDeleted,
		"deleted_attachment_id": e.deletedAttachmentID,
	}
}
