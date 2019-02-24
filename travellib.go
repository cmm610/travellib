package travellib

import (
	"log"
	"cloud.google.com/go/firestore"
	"context"
	"net/http"
)

type TripInfo struct {
	Users []User `json:"users"`
	Name string `json:"name"`
	Activities []Activity `json:"activities"`
	Lodging []Lodging `json:"lodging"`
	Uid string
	CreatedAt int64
	UpdatedAt int64
}
type User struct {
	Uid string `json:"uid"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}
type Trip struct {
	Uid string `json:"uid"`
	Name string `json:"name"`
}
type FullUserInfo struct {
	CreationTime string `json:"creationTime"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Profile string `json:"profile"`
	Email string `json:"email"`
	Uid string `json:"uid"`
	RefreshToken string `json:"refreshToken"`
	EmailVerified bool`json:"emailVerified"`
	IsAnonymous bool`json:"isAnonymous"`
	Trips []Trip `json:"trips"`
}
type Geometry struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type Location struct {
	Name string `json:"name"`
	Geometry Geometry `json:"geometry"`
	Rank string `json:"rank"`
	MapsId string `json:"mapsId"`
	MapsUrl string `json:"mapsUrl"`
	Photos []string `json:"photos"`
	Category string `json:"category"`
}
type Activity struct {
	Name string `json:"name"`
	Geometry Geometry `json:"geometry"`
	Rank string `json:"rank"`
	MapsId string `json:"mapsId"`
	MapsUrl string `json:"mapsUrl"`
	Photos []string `json:"photos"`
	Category string `json:"category"`
}
type Lodging struct {
	Name string `json:"name"`
	Geometry Geometry `json:"geometry"`
	MapsId string `json:"mapsId"`
	MapsUrl string `json:"mapsUrl"`
	Photos []string `json:"photos"`
	Category string `json:"category"`
}

var ctx context.Context
var client *firestore.Client

func InitFirestore() (firestore.Client, context.Context) {
	projectID := "traveller-3"
	if ctx == nil {
		ctx = context.Background()
	}
	if client == nil {
		// Pre-declare an err variable to avoid shadowing client.
		var err error
		client, err = firestore.NewClient(ctx, projectID)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
	}

	return *client, ctx
}


func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	(*w).Header().Set("Content-Type", "application/json")
}
