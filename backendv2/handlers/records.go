package handlers

import (
	"context"
	"net/http"

	"github.com/bedminer1/echoserver/dbiface"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

// HDBRecord contains data of a HDB resale unit
type HDBRecord struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Month string `json:"month" bson:"month"`
	Town string `json:"town" bson:"town"`
	FlatType string `json:"flatType" bson:"flatType"`
	FloorArea int `json:"floorArea" bson:"floorArea"`
	Model string `json:"model" bson:"model"`
	LeaseStart int `json:"leaseStart" bson:"leaseStart"`
	RemainingLease string `json:"remainingLease" bson:"remainingLease"`
	Price int `json:"price" bson:"price"`
}

// RecordHandler pass in col(reference to mongodb collection) as attribute
type RecordHandler struct {
	Col dbiface.CollectionAPI
}

// RecordValidator class with validate method
type RecordValidator struct {
	validator *validator.Validate
}

// Validate method that validates a record
func (record *RecordValidator) Validate(i interface{}) error {
	return record.validator.Struct(i)
}

func (h *RecordHandler) GetRecords(c echo.Context) error {
	return c.JSON(http.StatusOK, "get")
}

// insertProducts generates IDs and inserts products into mongo col
func insertRecords(ctx context.Context, records []HDBRecord, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, record := range records {
		record.ID = primitive.NewObjectID()
		insertID, err := collection.InsertOne(ctx, record)
		if err != nil {
			log.Errorf("Unable to insert %v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

// CreateRecords create records on mongodb and responds with IDs of records
func (h *RecordHandler) CreateRecords(c echo.Context) error {
	var records []HDBRecord
	c.Echo().Validator = &RecordValidator{validator: v}

	// bind echoContext to records
	if err := c.Bind(&records); err != nil {
		log.Errorf("Unable to bind: %v", err)
		return err
	}

	// validate records
	for _, record := range records {
		if err := c.Validate(record); err != nil {
			log.Errorf("Unable to validate record %+v, %v", record, err)
			return err
		}
	}

	// jsonFileRead insert here
	// readJsonFileWriteRecords(&records)

	IDs, err := insertRecords(context.Background(), records, h.Col)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, IDs)
}

// Read from jsonFile and write to records 
// func readJsonFileWriteRecords(records *[]HDBRecord) {
// 	selectedYear := "2015"
// 	jsonFile, err := os.Open(fmt.Sprintf("data/%s.json", selectedYear))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer jsonFile.Close()

// 	byteValue, _ := io.ReadAll(jsonFile)
// 	json.Unmarshal(byteValue, &records)
// }	