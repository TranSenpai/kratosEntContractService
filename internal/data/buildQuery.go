package data

import (
	entity "kratosEntContractService/internal/ent"
	whereContract "kratosEntContractService/internal/ent/contract"
	models "kratosEntContractService/internal/models"
)

func buildID[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeIDs, excludeIDs []int

	for _, v := range data.Includes {
		if includeID, ok := any(v).(int); ok {
			includeIDs = append(includeIDs, includeID)
		}
	}
	if len(includeIDs) > 0 {
		tx = tx.Where(whereContract.IDIn(includeIDs...))
	}

	for _, v := range data.Excludes {
		if excludeID, ok := any(v).(int); ok {
			excludeIDs = append(excludeIDs, excludeID)
		}
	}
	if len(excludeIDs) > 0 {
		tx = tx.Where(whereContract.IDNotIn(excludeIDs...))
	}

	return tx
}

func buildStudentCode[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeStudentCodes, excludeStudentCodes []string
	for _, v := range data.Includes {
		if includeStudentCode, ok := any(v).(string); ok {
			includeStudentCodes = append(includeStudentCodes, includeStudentCode)
		}
	}
	if len(includeStudentCodes) > 0 {
		tx = tx.Where(whereContract.StudentCodeIn(includeStudentCodes...))
	}

	for _, v := range data.Excludes {
		if excludeStudentCode, ok := any(v).(string); ok {
			excludeStudentCodes = append(excludeStudentCodes, excludeStudentCode)
		}
	}
	if len(excludeStudentCodes) > 0 {
		tx = tx.Where(whereContract.StudentCodeNotIn(excludeStudentCodes...))
	}

	return tx
}

func buildEmail[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeEmails, excludeEmails []string
	for _, v := range data.Includes {
		if includeEmail, ok := any(v).(string); ok {
			includeEmails = append(includeEmails, includeEmail)
		}
	}
	if len(includeEmails) > 0 {
		tx = tx.Where(whereContract.EmailIn(includeEmails...))
	}

	for _, v := range data.Excludes {
		if excludeEmail, ok := any(v).(string); ok {
			excludeEmails = append(excludeEmails, excludeEmail)
		}
	}

	if len(excludeEmails) > 0 {
		tx = tx.Where(whereContract.EmailNotIn(excludeEmails...))
	}

	return tx
}

func buildFirstName[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeFirstNames, excludeFirstNames []string
	for _, v := range data.Includes {
		if includeFirstName, ok := any(v).(string); ok {
			includeFirstNames = append(includeFirstNames, includeFirstName)
		}
	}
	if len(includeFirstNames) > 0 {
		tx = tx.Where(whereContract.FirstNameIn(includeFirstNames...))
	}

	for _, v := range data.Excludes {
		if excludeFirstName, ok := any(v).(string); ok {
			excludeFirstNames = append(excludeFirstNames, excludeFirstName)
		}
	}
	if len(excludeFirstNames) > 0 {
		tx = tx.Where(whereContract.FirstNameNotIn(excludeFirstNames...))
	}

	return tx
}

func buildLastName[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeLastNames, excludeLastNames []string
	for _, v := range data.Includes {
		if includeLastName, ok := any(v).(string); ok {
			includeLastNames = append(includeLastNames, includeLastName)
		}
	}
	if len(includeLastNames) > 0 {
		tx = tx.Where(whereContract.LastNameIn(includeLastNames...))
	}

	for _, v := range data.Excludes {
		if excludeLastName, ok := any(v).(string); ok {
			excludeLastNames = append(excludeLastNames, excludeLastName)
		}
	}
	if len(excludeLastNames) > 0 {
		tx = tx.Where(whereContract.LastNameNotIn(excludeLastNames...))
	}

	return tx
}

func buildMiddleName[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeMiddleNames, excludeMiddleNames []string
	for _, v := range data.Includes {
		if includeMiddleName, ok := any(v).(string); ok {
			includeMiddleNames = append(includeMiddleNames, includeMiddleName)
		}
	}
	if len(includeMiddleNames) > 0 {
		tx = tx.Where(whereContract.MiddleNameIn(includeMiddleNames...))
	}

	for _, v := range data.Excludes {
		if excludeMiddleName, ok := any(v).(string); ok {
			excludeMiddleNames = append(excludeMiddleNames, excludeMiddleName)
		}
	}
	if len(excludeMiddleNames) > 0 {
		tx = tx.Where(whereContract.MiddleNameNotIn(excludeMiddleNames...))
	}

	return tx
}

func buildPhone[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includePhones, excludePhones []string
	for _, v := range data.Includes {
		if includePhone, ok := any(v).(string); ok {
			includePhones = append(includePhones, includePhone)
		}
	}

	if len(includePhones) > 0 {
		tx = tx.Where(whereContract.PhoneIn(includePhones...))
	}

	for _, v := range data.Excludes {
		if excludePhone, ok := any(v).(string); ok {
			excludePhones = append(excludePhones, excludePhone)
		}
	}

	if len(excludePhones) > 0 {
		tx = tx.Where(whereContract.PhoneNotIn(excludePhones...))
	}

	return tx
}

func buildSign[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeSigns, excludeSigns []string
	for _, v := range data.Includes {
		if includeSign, ok := any(v).(string); ok {
			includeSigns = append(includeSigns, includeSign)
		}
	}

	if len(includeSigns) > 0 {
		tx = tx.Where(whereContract.SignIn(includeSigns...))
	}

	for _, v := range data.Excludes {
		if excludeSign, ok := any(v).(string); ok {
			excludeSigns = append(excludeSigns, excludeSign)
		}
	}
	if len(excludeSigns) > 0 {
		tx = tx.Where(whereContract.SignNotIn(excludeSigns...))
	}

	return tx
}

func buildRoomId[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeRoomIds, excludeRoomIds []string
	for _, v := range data.Includes {
		if includeRoomId, ok := any(v).(string); ok {
			includeRoomIds = append(includeRoomIds, includeRoomId)
		}
	}

	if len(includeRoomIds) > 0 {
		tx = tx.Where(whereContract.RoomIDIn(includeRoomIds...))
	}

	for _, v := range data.Excludes {
		if excludeRoomId, ok := any(v).(string); ok {
			excludeRoomIds = append(excludeRoomIds, excludeRoomId)
		}
	}
	if len(excludeRoomIds) > 0 {
		tx = tx.Where(whereContract.RoomIDNotIn(excludeRoomIds...))
	}

	return tx
}

func buildGender[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeGenders, excludeGenders []uint8
	for _, v := range data.Includes {
		if includeGender, ok := any(v).(uint8); ok {
			includeGenders = append(includeGenders, includeGender)
		}
	}
	if len(includeGenders) > 0 {
		tx = tx.Where(whereContract.GenderIn(includeGenders...))
	}
	for _, v := range data.Excludes {
		if excludeGender, ok := any(v).(uint8); ok {
			excludeGenders = append(excludeGenders, excludeGender)
		}
	}
	if len(excludeGenders) > 0 {
		tx = tx.Where(whereContract.GenderNotIn(excludeGenders...))
	}

	return tx
}

func buildAddress[T any](data models.SearchFilter[T], tx *entity.ContractQuery) *entity.ContractQuery {
	var includeAddresses, excludeAddresses []string
	for _, v := range data.Includes {
		if includeAddress, ok := any(v).(string); ok {
			includeAddresses = append(includeAddresses, includeAddress)
		}
	}
	if len(includeAddresses) > 0 {
		tx = tx.Where(whereContract.AddressIn(includeAddresses...))
	}

	for _, v := range data.Excludes {
		if excludeAddress, ok := any(v).(string); ok {
			excludeAddresses = append(excludeAddresses, excludeAddress)
		}
	}
	if len(excludeAddresses) > 0 {
		tx = tx.Where(whereContract.AddressNotIn(excludeAddresses...))
	}

	return tx
}

func buildRegistryAt(filter *models.ContractFilter, tx *entity.ContractQuery) *entity.ContractQuery {
	if filter.RegistryAt.FromTime != nil && !filter.RegistryAt.FromTime.IsZero() {
		tx = tx.Where(whereContract.RegistryAtGTE(*filter.RegistryAt.FromTime))
	}
	if filter.RegistryAt.FromTime != nil && !filter.RegistryAt.FromTime.IsZero() {
		tx = tx.Where(whereContract.RegistryAtLTE(*filter.RegistryAt.ToTime))
	}

	return tx
}

func buildIsActive(filter *models.ContractFilter, tx *entity.ContractQuery) *entity.ContractQuery {
	if filter.IsActive != nil {
		tx = tx.Where(whereContract.IsActive(*filter.IsActive))
	}

	return tx
}

func buildQuery(tx *entity.ContractQuery, field string, filter *models.ContractFilter) *entity.ContractQuery {
	switch field {
	case "id":
		tx = buildID(filter.Id, tx)
	case "student_code":
		tx = buildStudentCode(filter.StudentCode, tx)
	case "email":
		tx = buildEmail(filter.Email, tx)
	case "first_name":
		tx = buildFirstName(filter.FirstName, tx)
	case "last_name":
		tx = buildLastName(filter.LastName, tx)
	case "middle_name":
		tx = buildMiddleName(filter.MiddleName, tx)
	case "phone":
		tx = buildPhone(filter.Phone, tx)
	case "sign":
		tx = buildSign(filter.Sign, tx)
	case "room_id":
		tx = buildRoomId(filter.RoomId, tx)
	case "gender":
		tx = buildGender(filter.Gender, tx)
	case "address":
		tx = buildAddress(filter.Address, tx)
	case "registryAt":
		tx = buildRegistryAt(filter, tx)
	case "buildIsActive":
		tx = buildIsActive(filter, tx)
	}
	return tx
}

func (cr contractRepo) getQuery(filter *models.ContractFilter, query *entity.ContractQuery) *entity.ContractQuery {
	if filter == nil {
		return query
	}
	query = buildQuery(query, "id", filter)
	query = buildQuery(query, "student_code", filter)
	query = buildQuery(query, "email", filter)
	query = buildQuery(query, "first_name", filter)
	query = buildQuery(query, "last_name", filter)
	query = buildQuery(query, "middle_name", filter)
	query = buildQuery(query, "phone", filter)
	query = buildQuery(query, "sign", filter)
	query = buildQuery(query, "room_id", filter)
	query = buildQuery(query, "gender", filter)
	query = buildQuery(query, "address", filter)
	query = buildQuery(query, "registryAt", filter)
	query = buildQuery(query, "buildIsActive", filter)

	return query
}
