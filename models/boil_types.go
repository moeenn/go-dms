// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

type DeliveryStatus string

// Enum values for DeliveryStatus
const (
	DeliveryStatusPendingApproval DeliveryStatus = "pending-approval"
	DeliveryStatusApproved        DeliveryStatus = "approved"
	DeliveryStatusArrived         DeliveryStatus = "arrived"
	DeliveryStatusCompleted       DeliveryStatus = "completed"
	DeliveryStatusCancelled       DeliveryStatus = "cancelled"
)

func AllDeliveryStatus() []DeliveryStatus {
	return []DeliveryStatus{
		DeliveryStatusPendingApproval,
		DeliveryStatusApproved,
		DeliveryStatusArrived,
		DeliveryStatusCompleted,
		DeliveryStatusCancelled,
	}
}

func (e DeliveryStatus) IsValid() error {
	switch e {
	case DeliveryStatusPendingApproval, DeliveryStatusApproved, DeliveryStatusArrived, DeliveryStatusCompleted, DeliveryStatusCancelled:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e DeliveryStatus) String() string {
	return string(e)
}

func (e DeliveryStatus) Ordinal() int {
	switch e {
	case DeliveryStatusPendingApproval:
		return 0
	case DeliveryStatusApproved:
		return 1
	case DeliveryStatusArrived:
		return 2
	case DeliveryStatusCompleted:
		return 3
	case DeliveryStatusCancelled:
		return 4

	default:
		panic(errors.New("enum is not valid"))
	}
}

type Weekday string

// Enum values for Weekday
const (
	WeekdayMonday    Weekday = "monday"
	WeekdayTuesday   Weekday = "tuesday"
	WeekdayWednesday Weekday = "wednesday"
	WeekdayThursday  Weekday = "thursday"
	WeekdayFriday    Weekday = "friday"
	WeekdaySaturday  Weekday = "saturday"
	WeekdaySunday    Weekday = "sunday"
)

func AllWeekday() []Weekday {
	return []Weekday{
		WeekdayMonday,
		WeekdayTuesday,
		WeekdayWednesday,
		WeekdayThursday,
		WeekdayFriday,
		WeekdaySaturday,
		WeekdaySunday,
	}
}

func (e Weekday) IsValid() error {
	switch e {
	case WeekdayMonday, WeekdayTuesday, WeekdayWednesday, WeekdayThursday, WeekdayFriday, WeekdaySaturday, WeekdaySunday:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e Weekday) String() string {
	return string(e)
}

func (e Weekday) Ordinal() int {
	switch e {
	case WeekdayMonday:
		return 0
	case WeekdayTuesday:
		return 1
	case WeekdayWednesday:
		return 2
	case WeekdayThursday:
		return 3
	case WeekdayFriday:
		return 4
	case WeekdaySaturday:
		return 5
	case WeekdaySunday:
		return 6

	default:
		panic(errors.New("enum is not valid"))
	}
}

type SiteResourceType string

// Enum values for SiteResourceType
const (
	SiteResourceTypeForklift      SiteResourceType = "forklift"
	SiteResourceTypeCrane         SiteResourceType = "crane"
	SiteResourceTypeOffloadingBay SiteResourceType = "offloading-bay"
	SiteResourceTypeTrolley       SiteResourceType = "trolley"
)

func AllSiteResourceType() []SiteResourceType {
	return []SiteResourceType{
		SiteResourceTypeForklift,
		SiteResourceTypeCrane,
		SiteResourceTypeOffloadingBay,
		SiteResourceTypeTrolley,
	}
}

func (e SiteResourceType) IsValid() error {
	switch e {
	case SiteResourceTypeForklift, SiteResourceTypeCrane, SiteResourceTypeOffloadingBay, SiteResourceTypeTrolley:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e SiteResourceType) String() string {
	return string(e)
}

func (e SiteResourceType) Ordinal() int {
	switch e {
	case SiteResourceTypeForklift:
		return 0
	case SiteResourceTypeCrane:
		return 1
	case SiteResourceTypeOffloadingBay:
		return 2
	case SiteResourceTypeTrolley:
		return 3

	default:
		panic(errors.New("enum is not valid"))
	}
}

type UserRole string

// Enum values for UserRole
const (
	UserRoleAdmin         UserRole = "admin"
	UserRoleSubcontractor UserRole = "subcontractor"
	UserRoleSiteManager   UserRole = "site-manager"
	UserRoleGatekeeper    UserRole = "gatekeeper"
)

func AllUserRole() []UserRole {
	return []UserRole{
		UserRoleAdmin,
		UserRoleSubcontractor,
		UserRoleSiteManager,
		UserRoleGatekeeper,
	}
}

func (e UserRole) IsValid() error {
	switch e {
	case UserRoleAdmin, UserRoleSubcontractor, UserRoleSiteManager, UserRoleGatekeeper:
		return nil
	default:
		return errors.New("enum is not valid")
	}
}

func (e UserRole) String() string {
	return string(e)
}

func (e UserRole) Ordinal() int {
	switch e {
	case UserRoleAdmin:
		return 0
	case UserRoleSubcontractor:
		return 1
	case UserRoleSiteManager:
		return 2
	case UserRoleGatekeeper:
		return 3

	default:
		panic(errors.New("enum is not valid"))
	}
}
