package application

import (
	"books_service/internal/core/cerror"
	"books_service/internal/core/dto"
	"books_service/internal/core/models"
	"context"
	"slices"
)

const UPLOAD_CHAPTER_PERMISSION = "UPLOAD_CHAPTER"

func (u *useCase) CreateDraft(c context.Context, user *models.User, dto *dto.CreateDraft) (string, error) {
	team, err := u.teamsService.GetOne(c, dto.TeamId)
	if err != nil {
		return "", err
	}

	if team.OwnerId != user.ID {
		member, err := u.teamsService.GetMember(c, team.ID, user.ID)
		if err != nil {
			return "", cerror.New(cerror.FORBIDDEN, "you're not a team member")
		}

		if !slices.Contains(member.Permissions, UPLOAD_CHAPTER_PERMISSION) {
			return "", cerror.New(cerror.FORBIDDEN, "you don't have permission")
		}
	}

	book, err := u.repo.GetOneByID(c, dto.BookID)
	if err != nil {
		return "", err
	}

	if book.OwnerTeamID != 0 {
		if dto.TeamId != book.OwnerTeamID {
			return "", cerror.New(cerror.TEAM_NOT_OWNER, "this team is not a book's owner")
		}
	}

	id, err := u.repo.GetDraftChapterID(c, user.ID, dto)
	if err == nil {
		return id, nil
	}

	return u.repo.CreateDraft(c, user.ID, team.Name, dto)
}
