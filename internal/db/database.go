package database

import (
	"time"
)

type EventType string

const (
	EventWedding EventType = "Wedding"
	EventCorporate EventType = "Corporate"
	EventPrivate    EventType = "Private Party"
	EventSchool   EventType = "School Event"
	EventOther    EventType = "Other"
)

type PackageType string

const (
	PackageBasic  PackageType = "Bronze"
	PackageStandard PackageType = "Silver"
	PackagePremium  PackageType = "Gold"
	PackageLuxury   PackageType = "Platinum"
	PackageCustom   PackageType = "Custom"
)

type AddOnType string

const (
	RehearsalDinner AddOnType = "Rehearsal Dinner"
	BridalShower    AddOnType = "Bridal Shower"
	EngagementParty AddOnType = "Engagement Party"
	AfterParty      AddOnType = "After Party"
	Ceremony     AddOnType = "Ceremony"
	CocktailHour AddOnType = "Cocktail Hour"
	TentLighting AddOnType = "Tent Lighting"
	PhotoBooth   AddOnType = "Photo Booth"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "Pending"
	PaymentCompleted PaymentStatus = "Completed"
	PaymentFailed    PaymentStatus = "Failed"
	PaymentRefunded  PaymentStatus = "Refunded"
)

type Client struct {
	ID           int
	FirstName   string
	LastName    string
	CompanyName string
	Email       string
	Phone       string
	Address     string `db:"st_address"`
	City        string
	State       string
	Zip         string
	DateAdded   time.Time
}

type Event struct {
	ID		  int
	ClientID     int
	EventDate   time.Time
	EventName   string
	EventType   EventType
	StartTime    time.Time
	EndTime      time.Time
	Location	 string
	PackageType  PackageType
	GuestCount   int
	AddOns       []EventAddOn
	DepositAmount float64
	DepositReceived bool
	TotalPrice   float64
	PaymentStatus PaymentStatus
	PaymentDate   time.Time
	DepositDate   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type EventAddOn struct {
	EventID	 int
	AddOnID int
	Name   AddOnType
	Price float64
}
