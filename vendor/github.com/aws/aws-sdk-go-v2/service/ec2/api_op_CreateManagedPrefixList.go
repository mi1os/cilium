// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateManagedPrefixListInput struct {
	_ struct{} `type:"structure"`

	// The IP address type.
	//
	// Valid Values: IPv4 | IPv6
	//
	// AddressFamily is a required field
	AddressFamily *string `type:"string" required:"true"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	//
	// Constraints: Up to 255 UTF-8 characters in length.
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more entries for the prefix list.
	Entries []AddPrefixListEntry `locationName:"Entry" type:"list"`

	// The maximum number of entries for the prefix list.
	//
	// MaxEntries is a required field
	MaxEntries *int64 `type:"integer" required:"true"`

	// A name for the prefix list.
	//
	// Constraints: Up to 255 characters in length. The name cannot start with com.amazonaws.
	//
	// PrefixListName is a required field
	PrefixListName *string `type:"string" required:"true"`

	// The tags to apply to the prefix list during creation.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateManagedPrefixListInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateManagedPrefixListInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateManagedPrefixListInput"}

	if s.AddressFamily == nil {
		invalidParams.Add(aws.NewErrParamRequired("AddressFamily"))
	}

	if s.MaxEntries == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaxEntries"))
	}

	if s.PrefixListName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrefixListName"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateManagedPrefixListOutput struct {
	_ struct{} `type:"structure"`

	// Information about the prefix list.
	PrefixList *ManagedPrefixList `locationName:"prefixList" type:"structure"`
}

// String returns the string representation
func (s CreateManagedPrefixListOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateManagedPrefixList = "CreateManagedPrefixList"

// CreateManagedPrefixListRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a managed prefix list. You can specify one or more entries for the
// prefix list. Each entry consists of a CIDR block and an optional description.
//
// You must specify the maximum number of entries for the prefix list. The maximum
// number of entries cannot be changed later.
//
//    // Example sending a request using CreateManagedPrefixListRequest.
//    req := client.CreateManagedPrefixListRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateManagedPrefixList
func (c *Client) CreateManagedPrefixListRequest(input *CreateManagedPrefixListInput) CreateManagedPrefixListRequest {
	op := &aws.Operation{
		Name:       opCreateManagedPrefixList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateManagedPrefixListInput{}
	}

	req := c.newRequest(op, input, &CreateManagedPrefixListOutput{})

	return CreateManagedPrefixListRequest{Request: req, Input: input, Copy: c.CreateManagedPrefixListRequest}
}

// CreateManagedPrefixListRequest is the request type for the
// CreateManagedPrefixList API operation.
type CreateManagedPrefixListRequest struct {
	*aws.Request
	Input *CreateManagedPrefixListInput
	Copy  func(*CreateManagedPrefixListInput) CreateManagedPrefixListRequest
}

// Send marshals and sends the CreateManagedPrefixList API request.
func (r CreateManagedPrefixListRequest) Send(ctx context.Context) (*CreateManagedPrefixListResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateManagedPrefixListResponse{
		CreateManagedPrefixListOutput: r.Request.Data.(*CreateManagedPrefixListOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateManagedPrefixListResponse is the response type for the
// CreateManagedPrefixList API operation.
type CreateManagedPrefixListResponse struct {
	*CreateManagedPrefixListOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateManagedPrefixList request.
func (r *CreateManagedPrefixListResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
