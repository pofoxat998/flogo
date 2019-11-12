package sample

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	connectionString string `md:"connectionString,required"`
	sqlStatement    string `md:"sqlStatement"`     // sql Statement
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["connectionString"])
	r.connectionString = strVal
	strVal, _ := coerce.ToString(values["sqlStatement"])
	r.sqlStatement = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"connectionString": r.connectionString,
		"sqlStatement": r.sqlStatement,
	}
}
