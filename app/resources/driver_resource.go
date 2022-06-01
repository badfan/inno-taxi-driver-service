package resources

import (
	"context"

	"github.com/badfan/inno-taxi-driver-service/app/models"
	"github.com/badfan/inno-taxi-driver-service/app/models/sqlc"
)

func (r *Resource) GetDriverIDByPhone(ctx context.Context, phone string) (int, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetDriverIDByPhone(ctx, phone)
	if err != nil {
		r.logger.Errorf("error occured while getting driver's id by phone: %s", err.Error())
		return 0, err
	}

	return int(res), nil
}

func (r *Resource) CreateDriver(ctx context.Context, driver *models.Driver) (int, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.CreateDriver(ctx, sqlc.CreateDriverParams{
		Name:        driver.Name,
		PhoneNumber: driver.PhoneNumber,
		Email:       driver.Email,
		Password:    driver.Password,
		TaxiType:    sqlc.TaxiTypes(driver.TaxiType),
	})
	if err != nil {
		r.logger.Errorf("error occured while creating driver: %s", err.Error())
		return 0, err
	}

	return int(res.ID), nil
}

func (r *Resource) GetDriverByPhoneAndPassword(ctx context.Context, phone, password string) (*models.Driver, error) {
	queries := sqlc.New(r.Db)

	driver, err := queries.GetDriverByPhoneAndPassword(ctx, sqlc.GetDriverByPhoneAndPasswordParams{
		PhoneNumber: phone,
		Password:    password,
	})
	if err != nil {
		r.logger.Errorf("error occured while getting driver by phone and password: %s", err.Error())
		return nil, err
	}

	res := sqlcDriverConvert(&driver)

	return res, nil
}

func (r *Resource) GetDriverRatingByID(ctx context.Context, id int) (float32, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetDriverRatingByID(ctx, int32(id))
	if err != nil {
		r.logger.Errorf("error occured while getting driver's rating by id: %s", err.Error())
		return 0, err
	}

	return res, nil
}

func (r *Resource) GetDriverStatusByID(ctx context.Context, id int) (bool, error) {
	queries := sqlc.New(r.Db)

	res, err := queries.GetDriverStatusByID(ctx, int32(id))
	if err != nil {
		r.logger.Errorf("error occured while getting driver's status by id: %s", err.Error())
		return false, err
	}

	return res, nil
}

func sqlcDriverConvert(source *sqlc.Driver) *models.Driver {
	res := &models.Driver{
		ID:           source.ID,
		DriverUuid:   source.DriverUuid,
		Name:         source.Name,
		PhoneNumber:  source.PhoneNumber,
		Email:        source.Email,
		Password:     source.Password,
		TaxiType:     models.TaxiTypes(source.TaxiType),
		DriverRating: source.DriverRating,
		CreatedAt:    source.CreatedAt,
		UpdatedAt:    source.UpdatedAt,
	}

	return res
}
