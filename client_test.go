package accountapi

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var client = NewAccountApi("http://accountapi:8080")

var country string = "GB"
var classification = "Personal"
var version = int64(0)
var account = AccountData{
	ID:             "0568dec3-e8af-433b-906e-d9e616b30843",
	OrganisationID: "c50b6049-ef55-4c5b-8923-403758fc7170",
	Type:           "accounts",
	Attributes: &AccountAttributes{
		Country:               &country,
		BaseCurrency:          "GBP",
		BankID:                "400302",
		BankIDCode:            "GBDSC",
		AccountNumber:         "10000004",
		Iban:                  "GB28NWBK40030212764204",
		Bic:                   "NWBKGB42",
		Name:                  []string{"Samantha Holder"},
		AccountClassification: &classification},
	Version: &version}

func TestAddGetRemove(t *testing.T) {
	res, err := client.Add(&account)

	assert.Nil(t, err)
	assert.Equal(t, account, *res)

	res, err = client.Get(account.ID)

	assert.Nil(t, err)
	assert.Equal(t, account, *res)

	err = client.Remove(&account)
	assert.Nil(t, err)
}

func TestAddWrongRequest(t *testing.T) {
	_, err := client.Add(&AccountData{})

	if assert.NotNil(t, err) {
		assert.Contains(t, err.Error(), "validation failure list")
	}
}

func TestGetInvalidId(t *testing.T) {
	_, err := client.Get("invalid")

	if assert.NotNil(t, err) {
		assert.Equal(t, "id is not a valid uuid", err.Error())
	}
}

func TestGetNonExistentAccount(t *testing.T) {
	uuid := uuid.NewString()

	_, err := client.Get(uuid)

	if assert.NotNil(t, err) {
		assert.Equal(t, "record "+uuid+" does not exist", err.Error())
	}
}

func TestRemoveInvalidId(t *testing.T) {
	invalidAccount := account
	invalidAccount.ID = "invalid"

	err := client.Remove(&invalidAccount)

	if assert.NotNil(t, err) {
		assert.Equal(t, "id is not a valid uuid", err.Error())
	}
}

func TestRemoveNonExistentAccount(t *testing.T) {
	err := client.Remove(&account)

	if assert.NotNil(t, err) {
		assert.Equal(t, "record "+account.ID+" does not exist", err.Error())
	}
}
