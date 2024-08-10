package source

import (
	PhoeniciaDigitalDatabase "Phoenicia-Digital-Base-API/base/database"
	PhoeniciaDigitalUtils "Phoenicia-Digital-Base-API/base/utils"
	"encoding/json"
	"net/http"
)

type ContactInfo struct {
	PhoneNumber *string `json:"phoneNumber"`
	Email       *string `json:"email"`
}

type CustomerMessage struct {
	Name    *string `json:"name"`
	Email   *string `json:"email"`
	Message *string `json:"message"`
}

func PostContactInfoToDatabase(w http.ResponseWriter, r *http.Request) PhoeniciaDigitalUtils.PhoeniciaDigitalResponse {
	var contactInfo ContactInfo

	if err := json.NewDecoder(r.Body).Decode(&contactInfo); err != nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
	}

	// Ensure that either phoneNumber or email is provided
	if contactInfo.PhoneNumber == nil && contactInfo.Email == nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusFailedDependency, Quote: "At least one of phone number or email must be provided."}
	}

	if _, err := PhoeniciaDigitalDatabase.Postgres.SecureExecSQL("PostContactInfoToDatabase", contactInfo.PhoneNumber, contactInfo.Email); err != nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
	}

	return PhoeniciaDigitalUtils.ApiSuccess{Code: http.StatusAccepted, Quote: "We will contact you as soon as possible!"}
}

func PostNewMessageToDatabase(w http.ResponseWriter, r *http.Request) PhoeniciaDigitalUtils.PhoeniciaDigitalResponse {
	var messageData CustomerMessage

	if err := json.NewDecoder(r.Body).Decode(&messageData); err != nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
	}

	if messageData.Name == nil || messageData.Email == nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusFailedDependency, Quote: "Name & Email fields are required"}
	}

	if _, err := PhoeniciaDigitalDatabase.Postgres.SecureExecSQL("PostNewMessageToDatabase", messageData.Name, messageData.Email, messageData.Message); err != nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
	}

	return PhoeniciaDigitalUtils.ApiSuccess{Code: http.StatusAccepted, Quote: "We will contact you as soon as possible!"}
}

func GetContactInfoFromDatabase(w http.ResponseWriter, r *http.Request) PhoeniciaDigitalUtils.PhoeniciaDigitalResponse {
	var contactInfos []ContactInfo
	var contactInfo ContactInfo

	if stmt, err := PhoeniciaDigitalDatabase.Postgres.PrepareSQL("GetContactInfoFromDatabase"); err != nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
	} else {
		if rows, err := stmt.Query(); err != nil {
			return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
		} else {
			defer rows.Close()

			for rows.Next() {
				if err := rows.Scan(&contactInfo.PhoneNumber, &contactInfo.Email); err != nil {
					return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
				}
				contactInfos = append(contactInfos, contactInfo)
			}

			if err := rows.Err(); err != nil {
				return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
			}
		}
	}

	return PhoeniciaDigitalUtils.ApiSuccess{Code: http.StatusOK, Quote: contactInfos}
}

func GetCustomerMessagesFromDatabase(w http.ResponseWriter, r *http.Request) PhoeniciaDigitalUtils.PhoeniciaDigitalResponse {
	var messages []CustomerMessage
	var messageData CustomerMessage

	if stmt, err := PhoeniciaDigitalDatabase.Postgres.PrepareSQL("GetCustomerMessagesFromDatabase"); err != nil {
		return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
	} else {
		if rows, err := stmt.Query(); err != nil {
			return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
		} else {
			defer rows.Close()

			for rows.Next() {
				if err := rows.Scan(&messageData.Name, &messageData.Email, &messageData.Message); err != nil {
					return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
				}
				messages = append(messages, messageData)
			}

			if err := rows.Err(); err != nil {
				return PhoeniciaDigitalUtils.ApiError{Code: http.StatusInternalServerError, Quote: err.Error()}
			}
		}
	}

	return PhoeniciaDigitalUtils.ApiSuccess{Code: http.StatusOK, Quote: messages}
}
