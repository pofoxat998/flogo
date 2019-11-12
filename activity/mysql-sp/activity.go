package mysqlcallsp

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	_ = activity.Register(&Activity{})
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
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	ctx.Logger().Debugf("Input: %s", input.sqlStatement)
	
    // db 
	db, err := sql.Open("mysql", input.connectionString)
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare(input.sqlStatement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
    
    _, err = stmtIns.Exec("flogo", 3)
    // db
    
	return true, nil
}
