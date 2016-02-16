package mocks

import "github.com/stretchr/testify/mock"

import "github.com/aws/aws-sdk-go/aws/request"
import "github.com/aws/aws-sdk-go/service/sqs"

type SQSAPI struct {
	mock.Mock
}

// AddPermissionRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) AddPermissionRequest(_a0 *sqs.AddPermissionInput) (*request.Request, *sqs.AddPermissionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.AddPermissionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.AddPermissionOutput
	if rf, ok := ret.Get(1).(func(*sqs.AddPermissionInput) *sqs.AddPermissionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.AddPermissionOutput)
		}
	}

	return r0, r1
}

// AddPermission provides a mock function with given fields: _a0
func (_m *SQSAPI) AddPermission(_a0 *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.AddPermissionOutput
	if rf, ok := ret.Get(0).(func(*sqs.AddPermissionInput) *sqs.AddPermissionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.AddPermissionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.AddPermissionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangeMessageVisibilityRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) ChangeMessageVisibilityRequest(_a0 *sqs.ChangeMessageVisibilityInput) (*request.Request, *sqs.ChangeMessageVisibilityOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.ChangeMessageVisibilityInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.ChangeMessageVisibilityOutput
	if rf, ok := ret.Get(1).(func(*sqs.ChangeMessageVisibilityInput) *sqs.ChangeMessageVisibilityOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.ChangeMessageVisibilityOutput)
		}
	}

	return r0, r1
}

// ChangeMessageVisibility provides a mock function with given fields: _a0
func (_m *SQSAPI) ChangeMessageVisibility(_a0 *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.ChangeMessageVisibilityOutput
	if rf, ok := ret.Get(0).(func(*sqs.ChangeMessageVisibilityInput) *sqs.ChangeMessageVisibilityOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.ChangeMessageVisibilityOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.ChangeMessageVisibilityInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChangeMessageVisibilityBatchRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) ChangeMessageVisibilityBatchRequest(_a0 *sqs.ChangeMessageVisibilityBatchInput) (*request.Request, *sqs.ChangeMessageVisibilityBatchOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.ChangeMessageVisibilityBatchInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.ChangeMessageVisibilityBatchOutput
	if rf, ok := ret.Get(1).(func(*sqs.ChangeMessageVisibilityBatchInput) *sqs.ChangeMessageVisibilityBatchOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.ChangeMessageVisibilityBatchOutput)
		}
	}

	return r0, r1
}

// ChangeMessageVisibilityBatch provides a mock function with given fields: _a0
func (_m *SQSAPI) ChangeMessageVisibilityBatch(_a0 *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.ChangeMessageVisibilityBatchOutput
	if rf, ok := ret.Get(0).(func(*sqs.ChangeMessageVisibilityBatchInput) *sqs.ChangeMessageVisibilityBatchOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.ChangeMessageVisibilityBatchOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.ChangeMessageVisibilityBatchInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateQueueRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) CreateQueueRequest(_a0 *sqs.CreateQueueInput) (*request.Request, *sqs.CreateQueueOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.CreateQueueInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.CreateQueueOutput
	if rf, ok := ret.Get(1).(func(*sqs.CreateQueueInput) *sqs.CreateQueueOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.CreateQueueOutput)
		}
	}

	return r0, r1
}

// CreateQueue provides a mock function with given fields: _a0
func (_m *SQSAPI) CreateQueue(_a0 *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.CreateQueueOutput
	if rf, ok := ret.Get(0).(func(*sqs.CreateQueueInput) *sqs.CreateQueueOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.CreateQueueOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.CreateQueueInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessageRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) DeleteMessageRequest(_a0 *sqs.DeleteMessageInput) (*request.Request, *sqs.DeleteMessageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.DeleteMessageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.DeleteMessageOutput
	if rf, ok := ret.Get(1).(func(*sqs.DeleteMessageInput) *sqs.DeleteMessageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.DeleteMessageOutput)
		}
	}

	return r0, r1
}

// DeleteMessage provides a mock function with given fields: _a0
func (_m *SQSAPI) DeleteMessage(_a0 *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.DeleteMessageOutput
	if rf, ok := ret.Get(0).(func(*sqs.DeleteMessageInput) *sqs.DeleteMessageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.DeleteMessageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.DeleteMessageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessageBatchRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) DeleteMessageBatchRequest(_a0 *sqs.DeleteMessageBatchInput) (*request.Request, *sqs.DeleteMessageBatchOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.DeleteMessageBatchInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.DeleteMessageBatchOutput
	if rf, ok := ret.Get(1).(func(*sqs.DeleteMessageBatchInput) *sqs.DeleteMessageBatchOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.DeleteMessageBatchOutput)
		}
	}

	return r0, r1
}

// DeleteMessageBatch provides a mock function with given fields: _a0
func (_m *SQSAPI) DeleteMessageBatch(_a0 *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.DeleteMessageBatchOutput
	if rf, ok := ret.Get(0).(func(*sqs.DeleteMessageBatchInput) *sqs.DeleteMessageBatchOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.DeleteMessageBatchOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.DeleteMessageBatchInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteQueueRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) DeleteQueueRequest(_a0 *sqs.DeleteQueueInput) (*request.Request, *sqs.DeleteQueueOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.DeleteQueueInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.DeleteQueueOutput
	if rf, ok := ret.Get(1).(func(*sqs.DeleteQueueInput) *sqs.DeleteQueueOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.DeleteQueueOutput)
		}
	}

	return r0, r1
}

// DeleteQueue provides a mock function with given fields: _a0
func (_m *SQSAPI) DeleteQueue(_a0 *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.DeleteQueueOutput
	if rf, ok := ret.Get(0).(func(*sqs.DeleteQueueInput) *sqs.DeleteQueueOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.DeleteQueueOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.DeleteQueueInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQueueAttributesRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) GetQueueAttributesRequest(_a0 *sqs.GetQueueAttributesInput) (*request.Request, *sqs.GetQueueAttributesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.GetQueueAttributesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.GetQueueAttributesOutput
	if rf, ok := ret.Get(1).(func(*sqs.GetQueueAttributesInput) *sqs.GetQueueAttributesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.GetQueueAttributesOutput)
		}
	}

	return r0, r1
}

// GetQueueAttributes provides a mock function with given fields: _a0
func (_m *SQSAPI) GetQueueAttributes(_a0 *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.GetQueueAttributesOutput
	if rf, ok := ret.Get(0).(func(*sqs.GetQueueAttributesInput) *sqs.GetQueueAttributesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.GetQueueAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.GetQueueAttributesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQueueUrlRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) GetQueueUrlRequest(_a0 *sqs.GetQueueUrlInput) (*request.Request, *sqs.GetQueueUrlOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.GetQueueUrlInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.GetQueueUrlOutput
	if rf, ok := ret.Get(1).(func(*sqs.GetQueueUrlInput) *sqs.GetQueueUrlOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.GetQueueUrlOutput)
		}
	}

	return r0, r1
}

// GetQueueUrl provides a mock function with given fields: _a0
func (_m *SQSAPI) GetQueueUrl(_a0 *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.GetQueueUrlOutput
	if rf, ok := ret.Get(0).(func(*sqs.GetQueueUrlInput) *sqs.GetQueueUrlOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.GetQueueUrlOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.GetQueueUrlInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDeadLetterSourceQueuesRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) ListDeadLetterSourceQueuesRequest(_a0 *sqs.ListDeadLetterSourceQueuesInput) (*request.Request, *sqs.ListDeadLetterSourceQueuesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.ListDeadLetterSourceQueuesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.ListDeadLetterSourceQueuesOutput
	if rf, ok := ret.Get(1).(func(*sqs.ListDeadLetterSourceQueuesInput) *sqs.ListDeadLetterSourceQueuesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.ListDeadLetterSourceQueuesOutput)
		}
	}

	return r0, r1
}

// ListDeadLetterSourceQueues provides a mock function with given fields: _a0
func (_m *SQSAPI) ListDeadLetterSourceQueues(_a0 *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.ListDeadLetterSourceQueuesOutput
	if rf, ok := ret.Get(0).(func(*sqs.ListDeadLetterSourceQueuesInput) *sqs.ListDeadLetterSourceQueuesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.ListDeadLetterSourceQueuesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.ListDeadLetterSourceQueuesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListQueuesRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) ListQueuesRequest(_a0 *sqs.ListQueuesInput) (*request.Request, *sqs.ListQueuesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.ListQueuesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.ListQueuesOutput
	if rf, ok := ret.Get(1).(func(*sqs.ListQueuesInput) *sqs.ListQueuesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.ListQueuesOutput)
		}
	}

	return r0, r1
}

// ListQueues provides a mock function with given fields: _a0
func (_m *SQSAPI) ListQueues(_a0 *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.ListQueuesOutput
	if rf, ok := ret.Get(0).(func(*sqs.ListQueuesInput) *sqs.ListQueuesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.ListQueuesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.ListQueuesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PurgeQueueRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) PurgeQueueRequest(_a0 *sqs.PurgeQueueInput) (*request.Request, *sqs.PurgeQueueOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.PurgeQueueInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.PurgeQueueOutput
	if rf, ok := ret.Get(1).(func(*sqs.PurgeQueueInput) *sqs.PurgeQueueOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.PurgeQueueOutput)
		}
	}

	return r0, r1
}

// PurgeQueue provides a mock function with given fields: _a0
func (_m *SQSAPI) PurgeQueue(_a0 *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.PurgeQueueOutput
	if rf, ok := ret.Get(0).(func(*sqs.PurgeQueueInput) *sqs.PurgeQueueOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.PurgeQueueOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.PurgeQueueInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReceiveMessageRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) ReceiveMessageRequest(_a0 *sqs.ReceiveMessageInput) (*request.Request, *sqs.ReceiveMessageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.ReceiveMessageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.ReceiveMessageOutput
	if rf, ok := ret.Get(1).(func(*sqs.ReceiveMessageInput) *sqs.ReceiveMessageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.ReceiveMessageOutput)
		}
	}

	return r0, r1
}

// ReceiveMessage provides a mock function with given fields: _a0
func (_m *SQSAPI) ReceiveMessage(_a0 *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.ReceiveMessageOutput
	if rf, ok := ret.Get(0).(func(*sqs.ReceiveMessageInput) *sqs.ReceiveMessageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.ReceiveMessageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.ReceiveMessageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovePermissionRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) RemovePermissionRequest(_a0 *sqs.RemovePermissionInput) (*request.Request, *sqs.RemovePermissionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.RemovePermissionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.RemovePermissionOutput
	if rf, ok := ret.Get(1).(func(*sqs.RemovePermissionInput) *sqs.RemovePermissionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.RemovePermissionOutput)
		}
	}

	return r0, r1
}

// RemovePermission provides a mock function with given fields: _a0
func (_m *SQSAPI) RemovePermission(_a0 *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.RemovePermissionOutput
	if rf, ok := ret.Get(0).(func(*sqs.RemovePermissionInput) *sqs.RemovePermissionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.RemovePermissionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.RemovePermissionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessageRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) SendMessageRequest(_a0 *sqs.SendMessageInput) (*request.Request, *sqs.SendMessageOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.SendMessageInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.SendMessageOutput
	if rf, ok := ret.Get(1).(func(*sqs.SendMessageInput) *sqs.SendMessageOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.SendMessageOutput)
		}
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: _a0
func (_m *SQSAPI) SendMessage(_a0 *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.SendMessageOutput
	if rf, ok := ret.Get(0).(func(*sqs.SendMessageInput) *sqs.SendMessageOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.SendMessageOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.SendMessageInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessageBatchRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) SendMessageBatchRequest(_a0 *sqs.SendMessageBatchInput) (*request.Request, *sqs.SendMessageBatchOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.SendMessageBatchInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.SendMessageBatchOutput
	if rf, ok := ret.Get(1).(func(*sqs.SendMessageBatchInput) *sqs.SendMessageBatchOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.SendMessageBatchOutput)
		}
	}

	return r0, r1
}

// SendMessageBatch provides a mock function with given fields: _a0
func (_m *SQSAPI) SendMessageBatch(_a0 *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.SendMessageBatchOutput
	if rf, ok := ret.Get(0).(func(*sqs.SendMessageBatchInput) *sqs.SendMessageBatchOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.SendMessageBatchOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.SendMessageBatchInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetQueueAttributesRequest provides a mock function with given fields: _a0
func (_m *SQSAPI) SetQueueAttributesRequest(_a0 *sqs.SetQueueAttributesInput) (*request.Request, *sqs.SetQueueAttributesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*sqs.SetQueueAttributesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *sqs.SetQueueAttributesOutput
	if rf, ok := ret.Get(1).(func(*sqs.SetQueueAttributesInput) *sqs.SetQueueAttributesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*sqs.SetQueueAttributesOutput)
		}
	}

	return r0, r1
}

// SetQueueAttributes provides a mock function with given fields: _a0
func (_m *SQSAPI) SetQueueAttributes(_a0 *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *sqs.SetQueueAttributesOutput
	if rf, ok := ret.Get(0).(func(*sqs.SetQueueAttributesInput) *sqs.SetQueueAttributesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqs.SetQueueAttributesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sqs.SetQueueAttributesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
