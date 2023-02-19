package graph

import (
	"bipper-backend/graph/model"
	"bipper-backend/models"
)

func ToUserResponse(entity *models.User) *models.User {
	return &models.User{
		ID:           entity.ID,
		Email:        entity.Email,
		Name:         entity.Name,
		Username:     entity.Username,
		CreatedAt:    entity.CreatedAt,
		Description:  entity.Description,
		ProfileImg:   entity.ProfileImg,
		BannerImg:    entity.BannerImg,
		Dob:          entity.Dob,
		Location:     entity.Location,
		Url:          entity.Url,
		PinnedBeepID: entity.PinnedBeepID,
		Protected:    entity.Protected,
		Config:       entity.Config,
	}
}

func ToUserEdgeResponse(entity *models.User) *model.UserEdge {
	node := ToUserResponse(entity)
	return &model.UserEdge{
		Cursor: string(entity.ID),
		Node:   node,
	}
}

func ToUserConnectionResponse(edges []*model.UserEdge, pageInfo *model.PageInfo) *model.UserConnection {

	return &model.UserConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}
}

func ToBeepResponse(entity *models.Beep) *models.Beep {
	return &models.Beep{
		ID:                entity.ID,
		UserID:            entity.UserID,
		Text:              entity.Text,
		ConversationID:    entity.ConversationID,
		ReferenceID:       entity.ReferenceID,
		CreatedAt:         entity.CreatedAt,
		Sensitive:         entity.Sensitive,
		Config:            entity.Config,
		Location:          entity.Location,
		ImpressionCount:   entity.ImpressionCount,
		RebeepCount:       entity.RebeepCount,
		QuoteCount:        entity.QuoteCount,
		LikeCount:         entity.LikeCount,
		ReplyCount:        entity.ReplyCount,
		UrlClickCount:     entity.UrlClickCount,
		ProfileClickCount: entity.ProfileClickCount,
		DetailsClickCount: entity.DetailsClickCount,
		VideoViewCount:    entity.VideoViewCount,
	}
}

func ToBeepEdgeResponse(entity *models.Beep) *model.BeepEdge {
	node := ToBeepResponse(entity)
	return &model.BeepEdge{
		Cursor: string(entity.ID),
		Node:   node,
	}
}

func ToBeepConnectionResponse(edges []*model.BeepEdge, pageInfo *model.PageInfo) *model.BeepConnection {

	return &model.BeepConnection{
		Edges:    edges,
		PageInfo: pageInfo,
	}
}
