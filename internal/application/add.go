package application

import (
	"books_service/internal/core/cerror"
	"books_service/internal/core/dto"
	"books_service/internal/core/models"
	"context"
	"slices"
)

const CREATE_BOOK_PERMISSION = "CREATE_BOOK_PERMISSION"

func (u *useCase) Add(c context.Context, user models.User, dto dto.AddBook) error {
	team, err := u.teamsService.GetTeam(c, dto.OwnerTeamID)
	if err != nil {
		return err
	}

	if team.OwnerId != user.ID {
		member, err := u.teamsService.GetMember(c, team.ID, user.ID)
		if err != nil {
			return cerror.New(cerror.FORBIDDEN, "you're not a team member'")
		}

		if !slices.Contains(member.Permissions, CREATE_BOOK_PERMISSION) {
			return cerror.New(cerror.FORBIDDEN, "you don't have permission")
		}
	}

	_, err = u.repo.GetOneByLink(c, dto.Link)
	if err == nil {
		return cerror.New(cerror.LINK_ALREADY_USE, "link is already in use")
	}

	// check params
	if err := u.repo.IsGenresExists(c, dto.Genres); err != nil {
		return err
	}

	if err := u.repo.IsPersonsExists(c, dto.Authors); err != nil {
		return err
	}

	if err := u.repo.IsPersonsExists(c, dto.Artists); err != nil {
		return err
	}

	if err := u.repo.IsStatusExists(c, dto.StatusId); err != nil {
		return err
	}

	if err := u.repo.IsTypeExists(c, dto.TypeId); err != nil {
		return err
	}

	return u.repo.CreateBook(c, &dto)
}
