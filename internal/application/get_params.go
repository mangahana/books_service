package application

import "context"

func (u *useCase) GetParams(c context.Context) (map[string]any, error) {
	output := make(map[string]any)

	{
		genres, err := u.repo.GetGenres(c)
		if err != nil {
			return output, err
		}
		output["genres"] = genres
	}

	{
		statuses, err := u.repo.GetStatuses(c)
		if err != nil {
			return output, err
		}
		output["statuses"] = statuses
	}

	{
		types, err := u.repo.GetTypes(c)
		if err != nil {
			return output, err
		}
		output["types"] = types
	}

	{
		formats, err := u.repo.GetFormats(c)
		if err != nil {
			return output, err
		}
		output["formats"] = formats
	}

	return output, nil
}
