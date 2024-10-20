package application

import "context"

func (u *useCase) GetPages(c context.Context, chapterID string) ([]string, error) {
	return u.repo.GetPages(c, chapterID)
}
