package mock_entities

import (
	"time"

	"github.com/seanyudhistira91/technical-test-YC-W22/entity"
)

var MockUsers = func() entity.Users {
	return entity.Users{
		ID:           1,
		PhoneNumber:  "0822222222",
		Email:        "test@test.com",
		IsPremium:    false,
		HashPassword: "hashMock123",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Profile: &entity.Profiles{
			ID:        1,
			UserId:    1,
			Name:      "my name",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

}

var MockUsersWithHashPassword = func() entity.Users {
	return entity.Users{
		ID:           1,
		PhoneNumber:  "0822222222",
		Email:        "test@test.com",
		IsPremium:    false,
		HashPassword: "$2a$10$jRpjSujPjx435iyNdiL4d.mvGRg9RoV5YVoEpMTMShBa.ZknlPhqm",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Profile: &entity.Profiles{
			ID:        1,
			UserId:    1,
			Name:      "my name",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

}
