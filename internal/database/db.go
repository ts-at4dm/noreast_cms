package database

import (
	"time"
)

type EventType string

const (
	EventWedding    EventType = "Wedding"
	EventCorporate  EventType = "Corporate"
	EventPrivate    EventType = "Private Party"
	EventSchool     EventType = "School Event"
	EventOther      EventType = "Other"
)

type PackageType string

const (
	PackageBasic    PackageType = "Bronze"
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
	Ceremony        AddOnType = "Ceremony"
	CocktailHour    AddOnType = "Cocktail Hour"
	TentLighting    AddOnType = "Tent Lighting"
	PhotoBooth      AddOnType = "Photo Booth"
)

type Client struct {
	ID          int       `db:"id"`
	FirstName   string    `db:"firstname"`
	LastName    string    `db:"lastname"`
	CompanyName string    `db:"company_name"`
	Email       string    `db:"email"`
	Phone       string    `db:"phone"`
	Address     string    `db:"st_address"`
	City        string    `db:"city"`
	State       string    `db:"state"`
	Zip         string    `db:"zip"`
	DateAdded   time.Time `db:"date_added"`
}

type Event struct {
	ID              int           `db:"id"`
	ClientID        int           `db:"client_id"`
	EventDate       time.Time     `db:"event_date"`
	EventName       string        `db:"event_name"`
	EventType       EventType     `db:"event_type"`
	StartTime       time.Time     `db:"start_time"`
	EndTime         time.Time     `db:"end_time"`
	Location        string        `db:"event_location"`
	PackageType     PackageType   `db:"package"`
	GuestCount      int           `db:"guest_count"`
	DepositAmount   float64       `db:"deposit_amount"`
	DepositReceived bool          `db:"deposit_received"`
	TotalPrice      float64       `db:"total_price"`
	PaymentReceived bool          `db:"payment_received"`
	PaymentDate     *time.Time    `db:"payment_date"` // Nullable
	Notes           string        `db:"notes"`
	AddOns          []EventAddOn  `db:"-"` // Not stored directly in events table
}

type EventAddOn struct {
	ID      int       `db:"id"`
	EventID int       `db:"event_id"`
	AddOnID int       `db:"addon_id"`
	Price   float64   `db:"price"`
	Name    AddOnType `db:"-"` // Populated via JOIN
}

type AddOn struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}