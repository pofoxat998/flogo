package mysqlcallsp

import (
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"

    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	_ = activity.Register(&Activity{})
}

type Input struct {
	connectionString    string `md:"connectionString"`     // connection String
	sqlStatement    string `md:"sqlStatement"`     // sql Statement
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"connectionString":    i.connectionString,
		"sqlStatement": i.sqlStatement,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.connectionString, err = coerce.ToString(values["connectionString"])
	if err != nil {
		return err
	}
	i.sqlStatement, err = coerce.ToString(values["sqlStatement"])
	if err != nil {
		return err
	}

	return nil
}

var activityMd = activity.ToMetadata(&Input{})

// Activity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	ctx.GetInputObject(input)

	if ctx.Logger().DebugEnabled() {
		ctx.Logger().Debug("connection to db")
	}

    // db 
	db, err := sql.Open("mysql", input.connectionString)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare(sqlStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
    
    _, err = stmtIns.Exec("flogo", 3)
    // db
    
	return true, nil
}
