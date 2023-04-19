// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package graphql

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/kevinmichaelchen/genqlient-custom-type-omitempty-bug/pkg/postgres"
)

// Boolean expression to filter rows from the table "address". All fields are combined with a logical 'AND'.
type AddressBoolExp struct {
	And                []*AddressBoolExp    `json:"_and,omitempty"`
	Not                *AddressBoolExp      `json:"_not,omitempty"`
	Or                 []*AddressBoolExp    `json:"_or,omitempty"`
	AddressLines       *JsonbComparisonExp  `json:"addressLines,omitempty"`
	AdministrativeArea *StringComparisonExp `json:"administrativeArea,omitempty"`
	Id                 *UuidComparisonExp   `json:"id,omitempty"`
	Locality           *StringComparisonExp `json:"locality,omitempty"`
	Person             *PersonBoolExp       `json:"person,omitempty"`
	PersonId           *UuidComparisonExp   `json:"personId,omitempty"`
	PostalCode         *StringComparisonExp `json:"postalCode,omitempty"`
	RegionCode         *StringComparisonExp `json:"regionCode,omitempty"`
	Sublocality        *StringComparisonExp `json:"sublocality,omitempty"`
}

// GetAnd returns AddressBoolExp.And, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetAnd() []*AddressBoolExp { return v.And }

// GetNot returns AddressBoolExp.Not, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetNot() *AddressBoolExp { return v.Not }

// GetOr returns AddressBoolExp.Or, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetOr() []*AddressBoolExp { return v.Or }

// GetAddressLines returns AddressBoolExp.AddressLines, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetAddressLines() *JsonbComparisonExp { return v.AddressLines }

// GetAdministrativeArea returns AddressBoolExp.AdministrativeArea, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetAdministrativeArea() *StringComparisonExp { return v.AdministrativeArea }

// GetId returns AddressBoolExp.Id, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetId() *UuidComparisonExp { return v.Id }

// GetLocality returns AddressBoolExp.Locality, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetLocality() *StringComparisonExp { return v.Locality }

// GetPerson returns AddressBoolExp.Person, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetPerson() *PersonBoolExp { return v.Person }

// GetPersonId returns AddressBoolExp.PersonId, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetPersonId() *UuidComparisonExp { return v.PersonId }

// GetPostalCode returns AddressBoolExp.PostalCode, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetPostalCode() *StringComparisonExp { return v.PostalCode }

// GetRegionCode returns AddressBoolExp.RegionCode, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetRegionCode() *StringComparisonExp { return v.RegionCode }

// GetSublocality returns AddressBoolExp.Sublocality, and is useful for accessing the field via an interface.
func (v *AddressBoolExp) GetSublocality() *StringComparisonExp { return v.Sublocality }

// unique or primary key constraints on table "address"
type AddressConstraint string

const (
	// unique or primary key constraint on columns "id"
	AddressConstraintAddressPkey AddressConstraint = "address_pkey"
)

// input type for inserting data into table "address"
type AddressInsertInput struct {
	AddressLines       *any                     `json:"addressLines"`
	AdministrativeArea *string                  `json:"administrativeArea"`
	Id                 *string                  `json:"id"`
	Locality           *string                  `json:"locality"`
	Person             *PersonObjRelInsertInput `json:"person,omitempty"`
	PersonId           *string                  `json:"personId"`
	PostalCode         *string                  `json:"postalCode"`
	RegionCode         *string                  `json:"regionCode"`
	Sublocality        *string                  `json:"sublocality"`
}

// GetAddressLines returns AddressInsertInput.AddressLines, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetAddressLines() *any { return v.AddressLines }

// GetAdministrativeArea returns AddressInsertInput.AdministrativeArea, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetAdministrativeArea() *string { return v.AdministrativeArea }

// GetId returns AddressInsertInput.Id, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetId() *string { return v.Id }

// GetLocality returns AddressInsertInput.Locality, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetLocality() *string { return v.Locality }

// GetPerson returns AddressInsertInput.Person, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetPerson() *PersonObjRelInsertInput { return v.Person }

// GetPersonId returns AddressInsertInput.PersonId, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetPersonId() *string { return v.PersonId }

// GetPostalCode returns AddressInsertInput.PostalCode, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetPostalCode() *string { return v.PostalCode }

// GetRegionCode returns AddressInsertInput.RegionCode, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetRegionCode() *string { return v.RegionCode }

// GetSublocality returns AddressInsertInput.Sublocality, and is useful for accessing the field via an interface.
func (v *AddressInsertInput) GetSublocality() *string { return v.Sublocality }

// input type for inserting object relation for remote table "address"
type AddressObjRelInsertInput struct {
	Data *AddressInsertInput `json:"data,omitempty"`
	// upsert condition
	OnConflict *AddressOnConflict `json:"onConflict,omitempty"`
}

// GetData returns AddressObjRelInsertInput.Data, and is useful for accessing the field via an interface.
func (v *AddressObjRelInsertInput) GetData() *AddressInsertInput { return v.Data }

// GetOnConflict returns AddressObjRelInsertInput.OnConflict, and is useful for accessing the field via an interface.
func (v *AddressObjRelInsertInput) GetOnConflict() *AddressOnConflict { return v.OnConflict }

// on_conflict condition type for table "address"
type AddressOnConflict struct {
	Constraint    AddressConstraint     `json:"constraint"`
	UpdateColumns []AddressUpdateColumn `json:"updateColumns"`
	Where         *AddressBoolExp       `json:"where,omitempty"`
}

// GetConstraint returns AddressOnConflict.Constraint, and is useful for accessing the field via an interface.
func (v *AddressOnConflict) GetConstraint() AddressConstraint { return v.Constraint }

// GetUpdateColumns returns AddressOnConflict.UpdateColumns, and is useful for accessing the field via an interface.
func (v *AddressOnConflict) GetUpdateColumns() []AddressUpdateColumn { return v.UpdateColumns }

// GetWhere returns AddressOnConflict.Where, and is useful for accessing the field via an interface.
func (v *AddressOnConflict) GetWhere() *AddressBoolExp { return v.Where }

// update columns of table "address"
type AddressUpdateColumn string

const (
	// column name
	AddressUpdateColumnAddresslines AddressUpdateColumn = "addressLines"
	// column name
	AddressUpdateColumnAdministrativearea AddressUpdateColumn = "administrativeArea"
	// column name
	AddressUpdateColumnId AddressUpdateColumn = "id"
	// column name
	AddressUpdateColumnLocality AddressUpdateColumn = "locality"
	// column name
	AddressUpdateColumnPersonid AddressUpdateColumn = "personId"
	// column name
	AddressUpdateColumnPostalcode AddressUpdateColumn = "postalCode"
	// column name
	AddressUpdateColumnRegioncode AddressUpdateColumn = "regionCode"
	// column name
	AddressUpdateColumnSublocality AddressUpdateColumn = "sublocality"
)

// CreatePersonPerson includes the requested fields of the GraphQL type Person.
// The GraphQL type's documentation follows.
//
// columns and relationships of "person"
type CreatePersonPerson struct {
	Id             string             `json:"id"`
	ValidTimeRange postgres.TimeRange `json:"-"`
}

// GetId returns CreatePersonPerson.Id, and is useful for accessing the field via an interface.
func (v *CreatePersonPerson) GetId() string { return v.Id }

// GetValidTimeRange returns CreatePersonPerson.ValidTimeRange, and is useful for accessing the field via an interface.
func (v *CreatePersonPerson) GetValidTimeRange() postgres.TimeRange { return v.ValidTimeRange }

func (v *CreatePersonPerson) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CreatePersonPerson
		ValidTimeRange json.RawMessage `json:"validTimeRange"`
		graphql.NoUnmarshalJSON
	}
	firstPass.CreatePersonPerson = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.ValidTimeRange
		src := firstPass.ValidTimeRange
		if len(src) != 0 && string(src) != "null" {
			err = postgres.UnmarshalTimeRange(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal CreatePersonPerson.ValidTimeRange: %w", err)
			}
		}
	}
	return nil
}

type __premarshalCreatePersonPerson struct {
	Id string `json:"id"`

	ValidTimeRange json.RawMessage `json:"validTimeRange"`
}

func (v *CreatePersonPerson) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CreatePersonPerson) __premarshalJSON() (*__premarshalCreatePersonPerson, error) {
	var retval __premarshalCreatePersonPerson

	retval.Id = v.Id
	{

		dst := &retval.ValidTimeRange
		src := v.ValidTimeRange
		var err error
		*dst, err = postgres.MarshalTimeRange(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal CreatePersonPerson.ValidTimeRange: %w", err)
		}
	}
	return &retval, nil
}

// CreatePersonResponse is returned by CreatePerson on success.
type CreatePersonResponse struct {
	// insert a single row into the table: "person"
	Person *CreatePersonPerson `json:"person"`
}

// GetPerson returns CreatePersonResponse.Person, and is useful for accessing the field via an interface.
func (v *CreatePersonResponse) GetPerson() *CreatePersonPerson { return v.Person }

type JsonbCastExp struct {
	String *StringComparisonExp `json:"String,omitempty"`
}

// GetString returns JsonbCastExp.String, and is useful for accessing the field via an interface.
func (v *JsonbCastExp) GetString() *StringComparisonExp { return v.String }

// Boolean expression to compare columns of type "jsonb". All fields are combined with logical 'AND'.
type JsonbComparisonExp struct {
	Cast *JsonbCastExp `json:"_cast,omitempty"`
	// is the column contained in the given json value
	ContainedIn *any `json:"_containedIn"`
	// does the column contain the given json value at the top level
	Contains *any `json:"_contains"`
	Eq       *any `json:"_eq"`
	Gt       *any `json:"_gt"`
	Gte      *any `json:"_gte"`
	// does the string exist as a top-level key in the column
	HasKey *string `json:"_hasKey"`
	// do all of these strings exist as top-level keys in the column
	HasKeysAll []string `json:"_hasKeysAll"`
	// do any of these strings exist as top-level keys in the column
	HasKeysAny []string `json:"_hasKeysAny"`
	In         []any    `json:"_in"`
	IsNull     *bool    `json:"_isNull"`
	Lt         *any     `json:"_lt"`
	Lte        *any     `json:"_lte"`
	Neq        *any     `json:"_neq"`
	Nin        []any    `json:"_nin"`
}

// GetCast returns JsonbComparisonExp.Cast, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetCast() *JsonbCastExp { return v.Cast }

// GetContainedIn returns JsonbComparisonExp.ContainedIn, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetContainedIn() *any { return v.ContainedIn }

// GetContains returns JsonbComparisonExp.Contains, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetContains() *any { return v.Contains }

// GetEq returns JsonbComparisonExp.Eq, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetEq() *any { return v.Eq }

// GetGt returns JsonbComparisonExp.Gt, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetGt() *any { return v.Gt }

// GetGte returns JsonbComparisonExp.Gte, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetGte() *any { return v.Gte }

// GetHasKey returns JsonbComparisonExp.HasKey, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetHasKey() *string { return v.HasKey }

// GetHasKeysAll returns JsonbComparisonExp.HasKeysAll, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetHasKeysAll() []string { return v.HasKeysAll }

// GetHasKeysAny returns JsonbComparisonExp.HasKeysAny, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetHasKeysAny() []string { return v.HasKeysAny }

// GetIn returns JsonbComparisonExp.In, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetIn() []any { return v.In }

// GetIsNull returns JsonbComparisonExp.IsNull, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetIsNull() *bool { return v.IsNull }

// GetLt returns JsonbComparisonExp.Lt, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetLt() *any { return v.Lt }

// GetLte returns JsonbComparisonExp.Lte, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetLte() *any { return v.Lte }

// GetNeq returns JsonbComparisonExp.Neq, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetNeq() *any { return v.Neq }

// GetNin returns JsonbComparisonExp.Nin, and is useful for accessing the field via an interface.
func (v *JsonbComparisonExp) GetNin() []any { return v.Nin }

// Boolean expression to filter rows from the table "person". All fields are combined with a logical 'AND'.
type PersonBoolExp struct {
	And            []*PersonBoolExp        `json:"_and,omitempty"`
	Not            *PersonBoolExp          `json:"_not,omitempty"`
	Or             []*PersonBoolExp        `json:"_or,omitempty"`
	Address        *AddressBoolExp         `json:"address,omitempty"`
	Id             *UuidComparisonExp      `json:"id,omitempty"`
	Name           *StringComparisonExp    `json:"name,omitempty"`
	ValidTimeRange *TstzrangeComparisonExp `json:"validTimeRange,omitempty"`
}

// GetAnd returns PersonBoolExp.And, and is useful for accessing the field via an interface.
func (v *PersonBoolExp) GetAnd() []*PersonBoolExp { return v.And }

// GetNot returns PersonBoolExp.Not, and is useful for accessing the field via an interface.
func (v *PersonBoolExp) GetNot() *PersonBoolExp { return v.Not }

// GetOr returns PersonBoolExp.Or, and is useful for accessing the field via an interface.
func (v *PersonBoolExp) GetOr() []*PersonBoolExp { return v.Or }

// GetAddress returns PersonBoolExp.Address, and is useful for accessing the field via an interface.
func (v *PersonBoolExp) GetAddress() *AddressBoolExp { return v.Address }

// GetId returns PersonBoolExp.Id, and is useful for accessing the field via an interface.
func (v *PersonBoolExp) GetId() *UuidComparisonExp { return v.Id }

// GetName returns PersonBoolExp.Name, and is useful for accessing the field via an interface.
func (v *PersonBoolExp) GetName() *StringComparisonExp { return v.Name }

// GetValidTimeRange returns PersonBoolExp.ValidTimeRange, and is useful for accessing the field via an interface.
func (v *PersonBoolExp) GetValidTimeRange() *TstzrangeComparisonExp { return v.ValidTimeRange }

// unique or primary key constraints on table "person"
type PersonConstraint string

const (
	// unique or primary key constraint on columns "id"
	PersonConstraintPersonPkey PersonConstraint = "person_pkey"
)

// input type for inserting data into table "person"
type PersonInsertInput struct {
	Address        *AddressObjRelInsertInput `json:"address,omitempty"`
	Id             *string                   `json:"id"`
	Name           *string                   `json:"name"`
	ValidTimeRange *postgres.TimeRange       `json:"-"`
}

// GetAddress returns PersonInsertInput.Address, and is useful for accessing the field via an interface.
func (v *PersonInsertInput) GetAddress() *AddressObjRelInsertInput { return v.Address }

// GetId returns PersonInsertInput.Id, and is useful for accessing the field via an interface.
func (v *PersonInsertInput) GetId() *string { return v.Id }

// GetName returns PersonInsertInput.Name, and is useful for accessing the field via an interface.
func (v *PersonInsertInput) GetName() *string { return v.Name }

// GetValidTimeRange returns PersonInsertInput.ValidTimeRange, and is useful for accessing the field via an interface.
func (v *PersonInsertInput) GetValidTimeRange() *postgres.TimeRange { return v.ValidTimeRange }

func (v *PersonInsertInput) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*PersonInsertInput
		ValidTimeRange json.RawMessage `json:"validTimeRange"`
		graphql.NoUnmarshalJSON
	}
	firstPass.PersonInsertInput = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.ValidTimeRange
		src := firstPass.ValidTimeRange
		if len(src) != 0 && string(src) != "null" {
			*dst = new(postgres.TimeRange)
			err = postgres.UnmarshalTimeRange(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal PersonInsertInput.ValidTimeRange: %w", err)
			}
		}
	}
	return nil
}

type __premarshalPersonInsertInput struct {
	Address *AddressObjRelInsertInput `json:"address"`

	Id *string `json:"id"`

	Name *string `json:"name"`

	ValidTimeRange json.RawMessage `json:"validTimeRange"`
}

func (v *PersonInsertInput) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *PersonInsertInput) __premarshalJSON() (*__premarshalPersonInsertInput, error) {
	var retval __premarshalPersonInsertInput

	retval.Address = v.Address
	retval.Id = v.Id
	retval.Name = v.Name
	{

		dst := &retval.ValidTimeRange
		src := v.ValidTimeRange
		if src != nil {
			var err error
			*dst, err = postgres.MarshalTimeRange(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal PersonInsertInput.ValidTimeRange: %w", err)
			}
		}
	}
	return &retval, nil
}

// input type for inserting object relation for remote table "person"
type PersonObjRelInsertInput struct {
	Data *PersonInsertInput `json:"data,omitempty"`
	// upsert condition
	OnConflict *PersonOnConflict `json:"onConflict,omitempty"`
}

// GetData returns PersonObjRelInsertInput.Data, and is useful for accessing the field via an interface.
func (v *PersonObjRelInsertInput) GetData() *PersonInsertInput { return v.Data }

// GetOnConflict returns PersonObjRelInsertInput.OnConflict, and is useful for accessing the field via an interface.
func (v *PersonObjRelInsertInput) GetOnConflict() *PersonOnConflict { return v.OnConflict }

// on_conflict condition type for table "person"
type PersonOnConflict struct {
	Constraint    PersonConstraint     `json:"constraint"`
	UpdateColumns []PersonUpdateColumn `json:"updateColumns"`
	Where         *PersonBoolExp       `json:"where,omitempty"`
}

// GetConstraint returns PersonOnConflict.Constraint, and is useful for accessing the field via an interface.
func (v *PersonOnConflict) GetConstraint() PersonConstraint { return v.Constraint }

// GetUpdateColumns returns PersonOnConflict.UpdateColumns, and is useful for accessing the field via an interface.
func (v *PersonOnConflict) GetUpdateColumns() []PersonUpdateColumn { return v.UpdateColumns }

// GetWhere returns PersonOnConflict.Where, and is useful for accessing the field via an interface.
func (v *PersonOnConflict) GetWhere() *PersonBoolExp { return v.Where }

// update columns of table "person"
type PersonUpdateColumn string

const (
	// column name
	PersonUpdateColumnId PersonUpdateColumn = "id"
	// column name
	PersonUpdateColumnName PersonUpdateColumn = "name"
	// column name
	PersonUpdateColumnValidtimerange PersonUpdateColumn = "validTimeRange"
)

// Boolean expression to compare columns of type "String". All fields are combined with logical 'AND'.
type StringComparisonExp struct {
	Eq  *string `json:"_eq"`
	Gt  *string `json:"_gt"`
	Gte *string `json:"_gte"`
	// does the column match the given case-insensitive pattern
	Ilike *string  `json:"_ilike"`
	In    []string `json:"_in"`
	// does the column match the given POSIX regular expression, case insensitive
	Iregex *string `json:"_iregex"`
	IsNull *bool   `json:"_isNull"`
	// does the column match the given pattern
	Like *string `json:"_like"`
	Lt   *string `json:"_lt"`
	Lte  *string `json:"_lte"`
	Neq  *string `json:"_neq"`
	// does the column NOT match the given case-insensitive pattern
	Nilike *string  `json:"_nilike"`
	Nin    []string `json:"_nin"`
	// does the column NOT match the given POSIX regular expression, case insensitive
	Niregex *string `json:"_niregex"`
	// does the column NOT match the given pattern
	Nlike *string `json:"_nlike"`
	// does the column NOT match the given POSIX regular expression, case sensitive
	Nregex *string `json:"_nregex"`
	// does the column NOT match the given SQL regular expression
	Nsimilar *string `json:"_nsimilar"`
	// does the column match the given POSIX regular expression, case sensitive
	Regex *string `json:"_regex"`
	// does the column match the given SQL regular expression
	Similar *string `json:"_similar"`
}

// GetEq returns StringComparisonExp.Eq, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetEq() *string { return v.Eq }

// GetGt returns StringComparisonExp.Gt, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetGt() *string { return v.Gt }

// GetGte returns StringComparisonExp.Gte, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetGte() *string { return v.Gte }

// GetIlike returns StringComparisonExp.Ilike, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetIlike() *string { return v.Ilike }

// GetIn returns StringComparisonExp.In, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetIn() []string { return v.In }

// GetIregex returns StringComparisonExp.Iregex, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetIregex() *string { return v.Iregex }

// GetIsNull returns StringComparisonExp.IsNull, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetIsNull() *bool { return v.IsNull }

// GetLike returns StringComparisonExp.Like, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetLike() *string { return v.Like }

// GetLt returns StringComparisonExp.Lt, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetLt() *string { return v.Lt }

// GetLte returns StringComparisonExp.Lte, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetLte() *string { return v.Lte }

// GetNeq returns StringComparisonExp.Neq, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetNeq() *string { return v.Neq }

// GetNilike returns StringComparisonExp.Nilike, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetNilike() *string { return v.Nilike }

// GetNin returns StringComparisonExp.Nin, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetNin() []string { return v.Nin }

// GetNiregex returns StringComparisonExp.Niregex, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetNiregex() *string { return v.Niregex }

// GetNlike returns StringComparisonExp.Nlike, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetNlike() *string { return v.Nlike }

// GetNregex returns StringComparisonExp.Nregex, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetNregex() *string { return v.Nregex }

// GetNsimilar returns StringComparisonExp.Nsimilar, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetNsimilar() *string { return v.Nsimilar }

// GetRegex returns StringComparisonExp.Regex, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetRegex() *string { return v.Regex }

// GetSimilar returns StringComparisonExp.Similar, and is useful for accessing the field via an interface.
func (v *StringComparisonExp) GetSimilar() *string { return v.Similar }

// Boolean expression to compare columns of type "tstzrange". All fields are combined with logical 'AND'.
type TstzrangeComparisonExp struct {
	Eq     *postgres.TimeRange  `json:"-"`
	Gt     *postgres.TimeRange  `json:"-"`
	Gte    *postgres.TimeRange  `json:"-"`
	In     []postgres.TimeRange `json:"-"`
	IsNull *bool                `json:"_isNull"`
	Lt     *postgres.TimeRange  `json:"-"`
	Lte    *postgres.TimeRange  `json:"-"`
	Neq    *postgres.TimeRange  `json:"-"`
	Nin    []postgres.TimeRange `json:"-"`
}

// GetEq returns TstzrangeComparisonExp.Eq, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetEq() *postgres.TimeRange { return v.Eq }

// GetGt returns TstzrangeComparisonExp.Gt, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetGt() *postgres.TimeRange { return v.Gt }

// GetGte returns TstzrangeComparisonExp.Gte, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetGte() *postgres.TimeRange { return v.Gte }

// GetIn returns TstzrangeComparisonExp.In, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetIn() []postgres.TimeRange { return v.In }

// GetIsNull returns TstzrangeComparisonExp.IsNull, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetIsNull() *bool { return v.IsNull }

// GetLt returns TstzrangeComparisonExp.Lt, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetLt() *postgres.TimeRange { return v.Lt }

// GetLte returns TstzrangeComparisonExp.Lte, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetLte() *postgres.TimeRange { return v.Lte }

// GetNeq returns TstzrangeComparisonExp.Neq, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetNeq() *postgres.TimeRange { return v.Neq }

// GetNin returns TstzrangeComparisonExp.Nin, and is useful for accessing the field via an interface.
func (v *TstzrangeComparisonExp) GetNin() []postgres.TimeRange { return v.Nin }

func (v *TstzrangeComparisonExp) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*TstzrangeComparisonExp
		Eq  json.RawMessage   `json:"_eq"`
		Gt  json.RawMessage   `json:"_gt"`
		Gte json.RawMessage   `json:"_gte"`
		In  []json.RawMessage `json:"_in"`
		Lt  json.RawMessage   `json:"_lt"`
		Lte json.RawMessage   `json:"_lte"`
		Neq json.RawMessage   `json:"_neq"`
		Nin []json.RawMessage `json:"_nin"`
		graphql.NoUnmarshalJSON
	}
	firstPass.TstzrangeComparisonExp = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Eq
		src := firstPass.Eq
		if len(src) != 0 && string(src) != "null" {
			*dst = new(postgres.TimeRange)
			err = postgres.UnmarshalTimeRange(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal TstzrangeComparisonExp.Eq: %w", err)
			}
		}
	}

	{
		dst := &v.Gt
		src := firstPass.Gt
		if len(src) != 0 && string(src) != "null" {
			*dst = new(postgres.TimeRange)
			err = postgres.UnmarshalTimeRange(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal TstzrangeComparisonExp.Gt: %w", err)
			}
		}
	}

	{
		dst := &v.Gte
		src := firstPass.Gte
		if len(src) != 0 && string(src) != "null" {
			*dst = new(postgres.TimeRange)
			err = postgres.UnmarshalTimeRange(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal TstzrangeComparisonExp.Gte: %w", err)
			}
		}
	}

	{
		dst := &v.In
		src := firstPass.In
		*dst = make(
			[]postgres.TimeRange,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = postgres.UnmarshalTimeRange(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"Unable to unmarshal TstzrangeComparisonExp.In: %w", err)
				}
			}
		}
	}

	{
		dst := &v.Lt
		src := firstPass.Lt
		if len(src) != 0 && string(src) != "null" {
			*dst = new(postgres.TimeRange)
			err = postgres.UnmarshalTimeRange(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal TstzrangeComparisonExp.Lt: %w", err)
			}
		}
	}

	{
		dst := &v.Lte
		src := firstPass.Lte
		if len(src) != 0 && string(src) != "null" {
			*dst = new(postgres.TimeRange)
			err = postgres.UnmarshalTimeRange(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal TstzrangeComparisonExp.Lte: %w", err)
			}
		}
	}

	{
		dst := &v.Neq
		src := firstPass.Neq
		if len(src) != 0 && string(src) != "null" {
			*dst = new(postgres.TimeRange)
			err = postgres.UnmarshalTimeRange(
				src, *dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal TstzrangeComparisonExp.Neq: %w", err)
			}
		}
	}

	{
		dst := &v.Nin
		src := firstPass.Nin
		*dst = make(
			[]postgres.TimeRange,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			if len(src) != 0 && string(src) != "null" {
				err = postgres.UnmarshalTimeRange(
					src, dst)
				if err != nil {
					return fmt.Errorf(
						"Unable to unmarshal TstzrangeComparisonExp.Nin: %w", err)
				}
			}
		}
	}
	return nil
}

type __premarshalTstzrangeComparisonExp struct {
	Eq json.RawMessage `json:"_eq"`

	Gt json.RawMessage `json:"_gt"`

	Gte json.RawMessage `json:"_gte"`

	In []json.RawMessage `json:"_in"`

	IsNull *bool `json:"_isNull"`

	Lt json.RawMessage `json:"_lt"`

	Lte json.RawMessage `json:"_lte"`

	Neq json.RawMessage `json:"_neq"`

	Nin []json.RawMessage `json:"_nin"`
}

func (v *TstzrangeComparisonExp) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *TstzrangeComparisonExp) __premarshalJSON() (*__premarshalTstzrangeComparisonExp, error) {
	var retval __premarshalTstzrangeComparisonExp

	{

		dst := &retval.Eq
		src := v.Eq
		if src != nil {
			var err error
			*dst, err = postgres.MarshalTimeRange(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.Eq: %w", err)
			}
		}
	}
	{

		dst := &retval.Gt
		src := v.Gt
		if src != nil {
			var err error
			*dst, err = postgres.MarshalTimeRange(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.Gt: %w", err)
			}
		}
	}
	{

		dst := &retval.Gte
		src := v.Gte
		if src != nil {
			var err error
			*dst, err = postgres.MarshalTimeRange(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.Gte: %w", err)
			}
		}
	}
	{

		dst := &retval.In
		src := v.In
		*dst = make(
			[]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			var err error
			*dst, err = postgres.MarshalTimeRange(
				&src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.In: %w", err)
			}
		}
	}
	retval.IsNull = v.IsNull
	{

		dst := &retval.Lt
		src := v.Lt
		if src != nil {
			var err error
			*dst, err = postgres.MarshalTimeRange(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.Lt: %w", err)
			}
		}
	}
	{

		dst := &retval.Lte
		src := v.Lte
		if src != nil {
			var err error
			*dst, err = postgres.MarshalTimeRange(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.Lte: %w", err)
			}
		}
	}
	{

		dst := &retval.Neq
		src := v.Neq
		if src != nil {
			var err error
			*dst, err = postgres.MarshalTimeRange(
				src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.Neq: %w", err)
			}
		}
	}
	{

		dst := &retval.Nin
		src := v.Nin
		*dst = make(
			[]json.RawMessage,
			len(src))
		for i, src := range src {
			dst := &(*dst)[i]
			var err error
			*dst, err = postgres.MarshalTimeRange(
				&src)
			if err != nil {
				return nil, fmt.Errorf(
					"Unable to marshal TstzrangeComparisonExp.Nin: %w", err)
			}
		}
	}
	return &retval, nil
}

// Boolean expression to compare columns of type "uuid". All fields are combined with logical 'AND'.
type UuidComparisonExp struct {
	Eq     *string  `json:"_eq"`
	Gt     *string  `json:"_gt"`
	Gte    *string  `json:"_gte"`
	In     []string `json:"_in"`
	IsNull *bool    `json:"_isNull"`
	Lt     *string  `json:"_lt"`
	Lte    *string  `json:"_lte"`
	Neq    *string  `json:"_neq"`
	Nin    []string `json:"_nin"`
}

// GetEq returns UuidComparisonExp.Eq, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetEq() *string { return v.Eq }

// GetGt returns UuidComparisonExp.Gt, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetGt() *string { return v.Gt }

// GetGte returns UuidComparisonExp.Gte, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetGte() *string { return v.Gte }

// GetIn returns UuidComparisonExp.In, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetIn() []string { return v.In }

// GetIsNull returns UuidComparisonExp.IsNull, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetIsNull() *bool { return v.IsNull }

// GetLt returns UuidComparisonExp.Lt, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetLt() *string { return v.Lt }

// GetLte returns UuidComparisonExp.Lte, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetLte() *string { return v.Lte }

// GetNeq returns UuidComparisonExp.Neq, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetNeq() *string { return v.Neq }

// GetNin returns UuidComparisonExp.Nin, and is useful for accessing the field via an interface.
func (v *UuidComparisonExp) GetNin() []string { return v.Nin }

// __CreatePersonInput is used internally by genqlient
type __CreatePersonInput struct {
	Person *PersonInsertInput `json:"person,omitempty"`
}

// GetPerson returns __CreatePersonInput.Person, and is useful for accessing the field via an interface.
func (v *__CreatePersonInput) GetPerson() *PersonInsertInput { return v.Person }

func CreatePerson(
	ctx context.Context,
	client graphql.Client,
	person *PersonInsertInput,
) (*CreatePersonResponse, error) {
	req := &graphql.Request{
		OpName: "CreatePerson",
		Query: `
mutation CreatePerson ($person: PersonInsertInput!) {
	person: insertPersonOne(object: $person, onConflict: {constraint:person_pkey,updateColumns:[id]}) {
		id
		validTimeRange
	}
}
`,
		Variables: &__CreatePersonInput{
			Person: person,
		},
	}
	var err error

	var data CreatePersonResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
