# logging-library-go

Logging package for a standard logging interface across microservices.
Includes context information such as request ids, span ids, stack trace and other unstructured context data useful for tracking and debugging purposes.

## Installation
To install this package, use `go get`:

    go get github.com/rokwire/logging-library-go

This will then make the following packages available to you:

    github.com/rokwire/logging-library-go/logs
    github.com/rokwire/logging-library-go/errors
    github.com/rokwire/logging-library-go/logutils

Import the `logging-library-go/logs` package into your code using this template:

```go
package yours

import (
  ...

  "github.com/rokwire/logging-library-go/logs"
)

func main() {
	var logger = logs.NewLogger("example", nil)
	logger.SetLevel(logs.Debug)

    ...
}
```

### Staying up to date
To update logging-library-go to the latest version, use `go get -u github.com/rokwire/logging-library-go`.

## Packages
There are three packages provided by this libarary: `logs`, `errors`, and `logutils`.

### `logs` 
The `logs` package provides the `Logger` and `Log` types. The `Logger` object provides the base configurations for the entire application, while the `Log` object carries state related to a specific request. 

### `errors`
The `errors` package provides the `Error` type which expands upon the functionality provided by the typical `error` primitive provided by Golang. For example, additional context such as a trace of wrapped errors is automatically maintained when using the `Wrap` functions. Various components of this chain can then be accessed through the convenience functions provided by this package.

### `logutils`
The `logutils` package contains constants and standard utilities shared by the `logs` and `errors` packages. 

## Error Wrappers
There are several convenience functions to help standardize the error generation process.

```go
//NewError returns an error containing the provided message
func NewError(message string) error

//NewErrorf returns an error containing the formatted message
func NewErrorf(message string, args ...interface{}) error 

//WrapErrorf returns an error containing the provided message and error
func WrapError(message string, err error) error 

//WrapErrorf returns an error containing the formatted message and provided error
func WrapErrorf(format string, err error, args ...interface{}) error
```

These functions should be used in place of `fmt.Errorf` and `errors.New`. They provide several key benefits.
1. **Consistent formatting:** When using these functions, the provided messages will be formatted in one standard format. This will make it easier to read and follow logs throughout and across services. 
2. **Context:** These functions will all automatically include information about the function that is generating the error to help trace the path of the call when the errors are logged at a higher level.
3. **Convenience:** This provides on central package that can be imported to create errors. It also provides convenience functions to wrap existing errors with more context which should be a common practice with our logging approach.

## Logging Helpers
There are several convenience functions that will help perform logging in common situations.

The `LogError` function can be used to log a message along with an existing `error` object
```go
//LogError prints the log at error level with given message and error
//	Returns combined error message as string
func (l *Log) LogError(message string, err error) string
```

The following functions manage logging and writing HTTP responses conveniently 

```go
//RequestSuccess sets the provided success message as the HTTP response, sets standard headers, and stores the message
// 	to the log context
func (l *Log) RequestSuccessMessage(w http.ResponseWriter, message string)

//RequestSuccessJSON sets the provided JSON as the HTTP response body, sets standard headers, and stores the request status
// 	to the log context
func (l *Log) RequestSuccessJSON(w http.ResponseWriter, responseJSON []byte)

//RequestError logs the provided message and error and sets it as the HTTP response
//	Params:
//		w: The http response writer for the active request
//		msg: The error message
//		err: The error received from the application
//		code: The HTTP response code to be set
//		showDetails: Only provide 'msg' not 'err' in HTTP response when false
func (l *Log) RequestError(w http.ResponseWriter, msg string, err error, code int, showDetails bool)
```

## Message Templates
This library includes two standardized templates/grammars for messages, as well as a dictionary of commonly used terms. The intention of providing this is to help keep the logs very consistent and easy to interpret when adding new functionality with new logs. 

### Data Template
The "data" message template can be used to describe common statuses of a specified data element. 

**Pattern:** `{data status} {type}: {args}`
**Example:** `Invalid query param: id=test_id`

### Action Template
The "action" message template can be used to describe common actions performed on a specified data type. 

**Pattern:** `{action status} {action} {type} for {args}`
**Example:** `Error marshalling organization for id=test_id`

### Message Template Parameters
Below are definitions and examples for the template parameters references above.

#### Data Status:
Data statuses describe the data element and are represented by the `logDataStatus` type.
- `StatusValid` ("Valid"), `StatusFound` ("Found"), `StatusInvalid` ("Invalid"), `MissingStatus` ("Missing")

#### Action Status:
Action statuses describe the the action and are represented by the `logActionStatus` type.
- `StatusSuccess` ("Success"), `StatusError` ("Error")

#### Action:
Actions represent the action taken on the data element and are represented by the `LogAction` type.
- Eg. `ActionFind` ("finding"), `ActionMarshal` ("marshalling"), `ActionInitialize` ("initializing"), `ActionSend` ("sending")... etc.

Many common actions are defined in the logging library and these should be used when possible to maintain standardization. If you cannot construct an accurate message with the provided defined actions, you may provide your own action verb (ending in -ing) to describe the situation. 

#### Type:
Types are representations of the data type that the status applies to represented by the `LogData` type. 
- Eg. `TypeQueryParam` ("query param"), `TypeRequest` ("request"), "organization", "user"... etc

There are several common types that will be reused across applications defined in the logging library, however each application should define types to represent various models specific to its context.

#### Args:
Args are arbitrary parameters which can be included to provide additional information about the data or action represented by the `logArgs` interface. There are three types of `logArgs`: `FieldArgs` (`map[string]string`), `ListArgs` (`[]string`), and `StringArgs` (`string`). Most commonly, these will be variable name and value pairs (`FieldArgs`).
- Eg. `FieldArgs{"id": "test_id", "name": "test_name"}`, `ListArgs{"id", "name"}`, `StringArgs("id")`... etc

### Message Template Helper Functions
There are several convenience functions to help log or create an error from these templates.

Note: `nil` "args" params are ok 

Errors:
```go
//DataMessage generates a message string for a data element
func DataMessage(status logDataStatus, dataType LogData, args logArgs) string

//DataError generates an error for a data element
func DataError(status logDataStatus, dataType LogData, args logArgs) error

//WrapDataError wraps an error for a data element
func WrapDataError(status logDataStatus, dataType LogData, args logArgs, err error) error

//ActionMessage generates a message string for an action
func ActionMessage(status logActionStatus, action LogAction, dataType LogData, args logArgs) string

//ActionError generates an error for an action
func ActionError(action LogAction, dataType LogData, args logArgs) error

//WrapActionError wraps an error for an action
func WrapActionError(action LogAction, dataType LogData, args logArgs, err error) error
```

Logs:
```go
//LogData logs and returns a data message at the designated level
func (l *Log) LogData(level logLevel, status logDataStatus, dataType LogData, args logArgs) string

//WarnData logs and returns a data message for the given error at the warn level
func (l *Log) WarnData(status logDataStatus, dataType LogData, err error) string

//ErrorData logs and returns a data message for the given error at the error level
func (l *Log) ErrorData(status logDataStatus, dataType LogData, err error) string

//RequestErrorData logs a data message and error and sets it as the HTTP response
func (l *Log) RequestErrorData(w http.ResponseWriter, status logDataStatus, dataType LogData, args logArgs, err error, code int, showDetails bool)

//LogAction logs and returns an action message at the designated level
func (l *Log) LogAction(level logLevel, status logActionStatus, action LogAction, dataType LogData, args logArgs) string

//WarnAction logs and returns an action message for the given error at the warn level
func (l *Log) WarnAction(action LogAction, dataType LogData, err error) string

//ErrorAction logs and returns an action message for the given error at the error level
func (l *Log) ErrorAction(action LogAction, dataType LogData, err error) string

//RequestSuccessAction sets the provided success action message as the HTTP response, sets standard headers, and stores the message
func (l *Log) RequestSuccessAction(w http.ResponseWriter, action LogAction, dataType LogData, args logArgs)

//RequestErrorAction logs an action message and error and sets it as the HTTP response
func (l *Log) RequestErrorAction(w http.ResponseWriter, action LogAction, dataType LogData, args logArgs, err error, code int, showDetails bool)
```

## Other Conventions
There are several recommended conventions for the use of this library:

### Internal functions do not write logs unless necessary
Internal functions (core, storage, auth...etc) should not log to the console in general. They should instead return an error to be logged at the API handler level. Using the error wrapping functions in this libarary will make sure that the relevant context is not lost along the way. 

Exceptions to this rule include warnings, where it is important that a log is generated indicating that an error occurred, but it is not a critical error which prevented successful execution. When this happens the error should not be returned, so a `Warn` function should be called on the `Log` object with the relevant information. Debug logging statements are also an exception here, for example printing the contents of an object at a specific point to keep a record in the dev environment. Finally, on some occasions, `Info` logs can be printed in core functions to indicate that a specific action occurred... etc.

### Use the Log object whenever possible
When it is necessary to write to the logs, the `Log` object should be used over the `Logger` object (or any other logging library/package) whenever possible. `Log` objects contain additional information and ensure that any printed logs are properly associated with the request being handled. They also allow you to store context to be logged upon the success or failure of the request. 

For example, if a non-critical issue occurs in a storage function and we want to log a warning without returning an error, the storage function should include a `*log.Log` in the function params. No `Logger` object should be stored and made available to internal functions outside of the initialization context.

Exceptions to this rule include initialization when there is no request being processed and therefore no `Log` object. Timer based functions are also in this category. In these cases the `Logger` should be used instead.

### Errors should almost always be wrapped at every level
When an error is received from a function call and returned, one of the error wrapping helpers should be used to provide additional context in a message. This also will ensure that the chain of function calls is preserved within the `Error` object. 

To get started, take a look at `example/app.go`
