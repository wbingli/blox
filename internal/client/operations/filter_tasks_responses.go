package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/aws/amazon-ecs-event-stream-handler/internal/models"
)

// FilterTasksReader is a Reader for the FilterTasks structure.
type FilterTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFilterTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewFilterTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFilterTasksOK creates a FilterTasksOK with default headers values
func NewFilterTasksOK() *FilterTasksOK {
	return &FilterTasksOK{}
}

/*FilterTasksOK handles this case with default header values.

Filter tasks - success
*/
type FilterTasksOK struct {
	Payload []*models.TaskModel
}

func (o *FilterTasksOK) Error() string {
	return fmt.Sprintf("[GET /tasks/filter][%d] filterTasksOK  %+v", 200, o.Payload)
}

func (o *FilterTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterTasksInternalServerError creates a FilterTasksInternalServerError with default headers values
func NewFilterTasksInternalServerError() *FilterTasksInternalServerError {
	return &FilterTasksInternalServerError{}
}

/*FilterTasksInternalServerError handles this case with default header values.

Filter tasks - unexpected error
*/
type FilterTasksInternalServerError struct {
	Payload string
}

func (o *FilterTasksInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tasks/filter][%d] filterTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *FilterTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
