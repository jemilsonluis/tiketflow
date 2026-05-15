package event

import (
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/jemilsonluis/internals/_shared/types"
	"github.com/jemilsonluis/internals/modules/events/domain/dto"
	"github.com/jemilsonluis/internals/modules/events/domain/entity"
	"github.com/jemilsonluis/internals/modules/events/domain/enum"
)

func TestNewEntity(t *testing.T) {
	faker := faker.New()

	//arrange
	params := dto.CreateEventDTO{
		Title:       faker.Lorem().Sentence(3),
		Description: faker.Lorem().Paragraph(2),
		Capacity:    faker.IntBetween(1, 1024),
		Status:      enum.EventStatusCanceled,
		OrganizerId: faker.UUID().V4(),
		StartsAt:    faker.Time().Past(),
		EndsAt:      faker.Time().Future().UTC(),
		Banner:      faker.Internet().URL(),
		EventTypeId: faker.UUID().V4(),
		Location: types.LocationType{
			Country:     faker.Address().Country(),
			Province:    faker.Address().City(),
			Municipal:   faker.Address().StreetName(),
			Adress:      faker.Address().Address(),
			MoreDetails: faker.Lorem().Sentence(2),
		},
	}

	//act
	_, err := entity.NewEventEntity(params)
	if err != nil {
		t.Errorf("Test failed: %v", err)
	}

	//assert
}

func TestNewEntityInvalidCapacity(t *testing.T) {
	faker := faker.New()

	//arrange
	params := dto.CreateEventDTO{
		Title:       faker.Lorem().Sentence(3),
		Description: faker.Lorem().Paragraph(2),
		Capacity:    -1,
		Status:      enum.EventStatusPublished,
		OrganizerId: faker.UUID().V4(),
		StartsAt:    faker.Time().Past().UTC(),
		EndsAt:      faker.Time().Future().UTC(),
		Banner:      faker.Internet().URL(),
		EventTypeId: faker.UUID().V4(),
		Location: types.LocationType{
			Country:     faker.Address().Country(),
			Province:    faker.Address().City(),
			Municipal:   faker.Address().StreetName(),
			Adress:      faker.Address().Address(),
			MoreDetails: faker.Lorem().Sentence(2),
		},
	}

	//act
	_, err := entity.NewEventEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected capacity error but got nil")
	}

	//assert
}

func TestNewEntityWithDiferentDates(t *testing.T) {
	faker := faker.New()

	//arrange
	params := dto.CreateEventDTO{
		Title:       faker.Lorem().Sentence(3),
		Description: faker.Lorem().Paragraph(2),
		Capacity:    faker.IntBetween(1, 1024),
		Status:      enum.EventStatusPublished,
		OrganizerId: faker.UUID().V4(),
		StartsAt:    faker.Time().Future().UTC(),
		EndsAt:      faker.Time().Past().UTC(),
		Banner:      faker.Internet().URL(),
		EventTypeId: faker.UUID().V4(),
		Location: types.LocationType{
			Country:     faker.Address().Country(),
			Province:    faker.Address().City(),
			Municipal:   faker.Address().StreetName(),
			Adress:      faker.Address().Address(),
			MoreDetails: faker.Lorem().Sentence(2),
		},
	}

	//act
	_, err := entity.NewEventEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected startsAt error but got nil")
	}

	//assert
}

func TestNewEntityWithoutLocation(t *testing.T) {
	faker := faker.New()

	//arrange
	params := dto.CreateEventDTO{
		Title:       faker.Lorem().Sentence(3),
		Description: faker.Lorem().Paragraph(2),
		Capacity:    faker.IntBetween(1, 1024),
		Status:      enum.EventStatusPublished,
		OrganizerId: faker.UUID().V4(),
		StartsAt:    faker.Time().Past().UTC(),
		EndsAt:      faker.Time().Future().UTC(),
		EventTypeId: faker.UUID().V4(),
		Banner:      faker.Internet().URL(),
	}

	//act
	_, err := entity.NewEventEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected Location error but got nil")
	}

	//assert
}

func TestNewEntityWithInvalidStatus(t *testing.T) {
	faker := faker.New()

	//arrange
	params := dto.CreateEventDTO{
		Title:       faker.Lorem().Sentence(3),
		Description: faker.Lorem().Paragraph(2),
		Capacity:    faker.IntBetween(1, 1024),
		Status:      "INVALID_STATUS",
		OrganizerId: faker.UUID().V4(),
		StartsAt:    faker.Time().Past().UTC(),
		EndsAt:      faker.Time().Future().UTC(),
		EventTypeId: faker.UUID().V4(),
		Banner:      faker.Internet().URL(),
		Location: types.LocationType{
			Country:     faker.Address().Country(),
			Province:    faker.Address().City(),
			Municipal:   faker.Address().StreetName(),
			Adress:      faker.Address().Address(),
			MoreDetails: faker.Lorem().Sentence(2),
		},
	}

	//act
	_, err := entity.NewEventEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected status error but got nil")
	}

	//assert
}
