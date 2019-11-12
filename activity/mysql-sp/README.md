# mysql-sp
This activity allows you to call mysql store procedure

## Installation

### Flogo CLI
```bash
flogo install github.com/pofoxat998/flogo/tree/master/activity/mysql-sp
```

## Configuration

### Input:
| Name     			  | Type   | Description
|:---      			  | :---   | :---    
| connectionString    | string | connection string to mysql server
| sqlStatement 		  | string | sql statement with CALL

## Examples
The below example logs a message 'test message':

```json
{
  "id": "log_message",
  "name": "Log Message",
  "activity": {
    "ref": "github.com/project-flogo/contrib/activity/log",
    "input": {
      "message": "test message",
      "addDetails": "false"
    }
  }
}
```