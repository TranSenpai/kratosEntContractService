package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"kratosEntContractService/internal/biz"
	entity "kratosEntContractService/internal/ent"
	whereContract "kratosEntContractService/internal/ent/contract"
	models "kratosEntContractService/internal/models"
	"time"
)

// Type conversion the value nil to type concrete (*contractRepo)
// After that we assign some variable to the pointer type (var x Interface = (*ConcreteInterface)(nil))
// Compiler will check if ConcreteInterface satisfy of not (implement all method of interface)
// Use to check if a concrete type fully implement it's Interface
// In this example b/c we dont use the value so let it be a blank variable:
var _ biz.IContractRepo = (*contractRepo)(nil)

func NewContractRepo(data *Data) biz.IContractRepo {
	return &contractRepo{
		data: data,
	}
}

type contractRepo struct {
	data *Data
}

func withTx(ctx context.Context, client *entity.Client, fn func(tx *entity.Tx) error) error {
	if ctx.Err() != nil {
		return GetError(fmt.Errorf("context error before transaction: %w", ctx.Err()))
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return GetError(err)
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return GetError(errors.New(fmt.Sprintf("%w: rolling back transaction: %v", err, rerr)))
		}
		return GetError(err)
	}

	if err := tx.Commit(); err != nil {
		return GetError(err)
	}

	return nil
}

func (cr *contractRepo) CreateContract(ctx context.Context, createContract *entity.Contract) error {
	err := withTx(ctx, cr.data.DB(), func(tx *entity.Tx) error {
		return tx.Contract.Create().
			SetStudentCode(createContract.StudentCode).
			SetFirstName(createContract.FirstName).
			SetLastName(createContract.LastName).
			SetMiddleName(createContract.MiddleName).
			SetEmail(createContract.Email).
			SetSign(createContract.Sign).
			SetPhone(createContract.Phone).
			SetGender(createContract.Gender).
			SetDob(createContract.Dob).
			SetAddress(createContract.Address).
			SetAvatar(createContract.Avatar).
			SetIsActive(createContract.IsActive).
			SetRegistryAt(createContract.RegistryAt).
			SetRoomID(createContract.RoomID).
			SetNotificationChannels(createContract.NotificationChannels).Exec(ctx)
	})
	if err != nil {
		return GetError(err)
	}

	return nil
}

func (cr *contractRepo) UpdateContract(ctx context.Context, updateData map[string]any, filter *models.ContractFilter) error {
	return withTx(ctx, cr.data.DB(), func(tx *entity.Tx) error {
		update := tx.Contract.UpdateOneID(filter.Id.Includes[0])
		if updateData["StudentCode"] != nil {
			if studentCode, ok := updateData["StudentCode"].(string); ok {
				update = update.SetStudentCode(studentCode)
			}
		}
		if updateData["FirstName"] != nil {
			if firstName, ok := updateData["FirstName"].(string); ok {
				update = update.SetFirstName(firstName)
			}
		}
		if updateData["LastName"] != nil {
			if lastName, ok := updateData["LastName"].(string); ok {
				update = update.SetLastName(lastName)
			}
		}
		if updateData["MiddleName"] != nil {
			if middleName, ok := updateData["MiddleName"].(string); ok {
				update = update.SetMiddleName(middleName)
			}
		}
		if updateData["Email"] != nil {
			if email, ok := updateData["Email"].(string); ok {
				update = update.SetEmail(email)
			}
		}
		if updateData["Sign"] != nil {
			if sign, ok := updateData["Sign"].(string); ok {
				update = update.SetSign(sign)
			}
		}
		if updateData["Phone"] != nil {
			if phone, ok := updateData["Phone"].(string); ok {
				update = update.SetPhone(phone)
			}
		}
		if updateData["Gender"] != nil {
			if gender, ok := updateData["Gender"].(uint8); ok {
				update = update.SetGender(gender)
			}
		}
		if updateData["Dob"] != nil {
			if dob, ok := updateData["Dob"].(time.Time); ok {
				update = update.SetDob(dob)
			}
		}
		if updateData["Address"] != nil {
			if address, ok := updateData["Address"].(string); ok {
				update = update.SetAddress(address)
			}
		}
		if updateData["Avatar"] != nil {
			if avatar, ok := updateData["Avatar"].([]byte); ok {
				update = update.SetAvatar(avatar)
			}
		}
		if updateData["IsActive"] != nil {
			if isActive, ok := updateData["IsActive"].(bool); ok {
				update = update.SetIsActive(isActive)
			}
		}
		if updateData["RegistryAt"] != nil {
			if registryAt, ok := updateData["RegistryAt"].(time.Time); ok {
				update = update.SetRegistryAt(registryAt)
			}
		}
		if updateData["RoomID"] != nil {
			if roomID, ok := updateData["RoomID"].(string); ok {
				update = update.SetRoomID(roomID)
			}
		}
		if updateData["NotificationChannels"] != nil {
			if notificationChannels, ok := updateData["NotificationChannels"].(uint8); ok {
				update = update.SetNotificationChannels(notificationChannels)
			}
		}

		return update.Exec(ctx)
	})
}

func (cr *contractRepo) DeleteContract(ctx context.Context, filter *models.ContractFilter) error {
	return withTx(ctx, cr.data.DB(), func(tx *entity.Tx) error {
		return tx.Contract.DeleteOneID(filter.Id.Includes[0]).Exec(ctx)
	})
}

func (cr contractRepo) GetContract(ctx context.Context, filter *models.ContractFilter) (*entity.Contract, error) {
	contract, err := cr.data.DB().Contract.Query().Where(whereContract.IDIn(filter.Id.Includes[0])).First(ctx)
	return contract, err
}

func (cr contractRepo) ListContract(ctx context.Context, filter *models.ContractFilter) ([]entity.Contract, error) {
	query := cr.getQuery(filter, cr.data.DB().Contract.Query())
	if query == nil {
		return nil, GetError(errors.New("nil query"))
	}
	lst, err := query.All(ctx)
	if err != nil {
		return nil, GetError(err)
	}
	var result []entity.Contract
	for _, c := range lst {
		if c != nil {
			result = append(result, *c)
		}
	}

	return result, nil
}

func (cr contractRepo) GetTotalContractRoom(ctx context.Context, roomID string) (models.TotalContractsEachRoom, error) {
	var result models.TotalContractsEachRoom

	count, err := cr.data.DB().Contract.Query().
		Where(whereContract.RoomIDIn(roomID)).
		Count(ctx)

	if err != nil {
		return result, GetError(err)
	}

	result.Total = uint8(count)
	result.RoomID = roomID
	return result, nil
}

func (cr contractRepo) ListTotalContractEachRoom(ctx context.Context, number uint8) ([]models.TotalContractsEachRoom, error) {
	var checks, results []models.TotalContractsEachRoom
	err := cr.data.DB().Contract.Query().
		GroupBy(whereContract.FieldRoomID).
		Aggregate(func(selector *sql.Selector) string {
			return sql.As(sql.Count("*"), "total")
		}).
		Scan(ctx, &checks)
	for _, v := range checks {
		if v.Total >= number {
			results = append(results, v)
		}
	}
	if err != nil {
		return nil, GetError(err)
	}

	return results, err
}
