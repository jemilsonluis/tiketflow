package tiket

import (
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/jemilsonluis/internals/modules/tikets/domain/dto"
	"github.com/jemilsonluis/internals/modules/tikets/domain/entity"
	"github.com/jemilsonluis/internals/modules/tikets/domain/enum"
)

func TestNewTiketEntity(t *testing.T) {
	faker := faker.New()

	params := dto.CreateTiketDTO{
		Code:      "1234",
		UserId:    faker.UUID().V4(),
		EventId:   faker.UUID().V4(),
		TiketType: enum.TiketTypeStudent,
		ExpiresAt: faker.Time().Future(),
	}

	_, err := entity.NewTiketEntity(params)
	if err != nil {
		t.Errorf("Test failed: expected new tiket entity")
	}
}

func TestNewEntityEmptyCode(t *testing.T) {
	faker := faker.New()

	params := dto.CreateTiketDTO{
		Code:      "",
		UserId:    faker.UUID().V4(),
		EventId:   faker.UUID().V4(),
		TiketType: enum.TiketTypeRegular,
		ExpiresAt: faker.Time().Future(),
	}

	_, err := entity.NewTiketEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected code error but got nil")
	}
}

func TestNewEntityEmptyUserId(t *testing.T) {
	faker := faker.New()

	params := dto.CreateTiketDTO{
		Code:      "qw21",
		UserId:    "",
		EventId:   faker.UUID().V4(),
		TiketType: enum.TiketTypeRegular,
		ExpiresAt: faker.Time().Future(),
	}

	_, err := entity.NewTiketEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected userId error but got nil")
	}
}

func TestNewEntityEmptyEventId(t *testing.T) {
	faker := faker.New()

	params := dto.CreateTiketDTO{
		Code:      "12345",
		UserId:    faker.UUID().V4(),
		EventId:   "",
		TiketType: enum.TiketTypeRegular,
		ExpiresAt: faker.Time().Future(),
	}

	_, err := entity.NewTiketEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected eventId error but got nil")
	}
}

func TestNewEntityEmptyTyketType(t *testing.T) {
	faker := faker.New()

	params := dto.CreateTiketDTO{
		Code:      "123456",
		UserId:    faker.UUID().V4(),
		EventId:   faker.UUID().V4(),
		TiketType: "",
		ExpiresAt: faker.Time().Future(),
	}

	_, err := entity.NewTiketEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected tiketType error but got nil")
	}
}

func TestNewEntityEmptyExpiresAt(t *testing.T) {
	faker := faker.New()

	params := dto.CreateTiketDTO{
		Code:      "1234",
		UserId:    faker.UUID().V4(),
		EventId:   faker.UUID().V4(),
		TiketType: enum.TiketTypeRegular,
		ExpiresAt: faker.Time().Past(),
	}

	_, err := entity.NewTiketEntity(params)
	if err == nil {
		t.Errorf("Test failed: expected expiresAt error but got nil")
	}
}
