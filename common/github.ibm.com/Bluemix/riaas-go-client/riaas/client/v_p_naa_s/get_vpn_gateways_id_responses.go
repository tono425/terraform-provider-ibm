// Code generated by go-swagger; DO NOT EDIT.

package v_p_naa_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetVpnGatewaysIDReader is a Reader for the GetVpnGatewaysID structure.
type GetVpnGatewaysIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVpnGatewaysIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVpnGatewaysIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetVpnGatewaysIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVpnGatewaysIDOK creates a GetVpnGatewaysIDOK with default headers values
func NewGetVpnGatewaysIDOK() *GetVpnGatewaysIDOK {
	return &GetVpnGatewaysIDOK{}
}

/*GetVpnGatewaysIDOK handles this case with default header values.

The VPN gateway was retrieved successfully.
*/
type GetVpnGatewaysIDOK struct {
	Payload *models.VPNGateway
}

func (o *GetVpnGatewaysIDOK) Error() string {
	return fmt.Sprintf("[GET /vpn_gateways/{id}][%d] getVpnGatewaysIdOK  %+v", 200, o.Payload)
}

func (o *GetVpnGatewaysIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VPNGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVpnGatewaysIDNotFound creates a GetVpnGatewaysIDNotFound with default headers values
func NewGetVpnGatewaysIDNotFound() *GetVpnGatewaysIDNotFound {
	return &GetVpnGatewaysIDNotFound{}
}

/*GetVpnGatewaysIDNotFound handles this case with default header values.

A VPN gateway with the specified identifier could not be found.
*/
type GetVpnGatewaysIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetVpnGatewaysIDNotFound) Error() string {
	return fmt.Sprintf("[GET /vpn_gateways/{id}][%d] getVpnGatewaysIdNotFound  %+v", 404, o.Payload)
}

func (o *GetVpnGatewaysIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}