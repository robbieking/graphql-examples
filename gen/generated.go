// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package gen

import (
	"context"
	"encoding/json"

	"github.com/Khan/genqlient/graphql"
)
import "github.com/Khan/genqlient/graphql"

// CreateVehicleCreateVehicle includes the requested fields of the GraphQL type Vehicle.
type CreateVehicleCreateVehicle struct {
	Vehicle `json:"-"`
}

// GetReg returns CreateVehicleCreateVehicle.Reg, and is useful for accessing the field via an interface.
func (v *CreateVehicleCreateVehicle) GetReg() string { return v.Vehicle.Reg }

func (v *CreateVehicleCreateVehicle) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*CreateVehicleCreateVehicle
		graphql.NoUnmarshalJSON
	}
	firstPass.CreateVehicleCreateVehicle = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	err = json.Unmarshal(
		b, &v.Vehicle)
	if err != nil {
		return err
	}
	return nil
}

type __premarshalCreateVehicleCreateVehicle struct {
	Reg string `json:"reg"`
}

func (v *CreateVehicleCreateVehicle) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *CreateVehicleCreateVehicle) __premarshalJSON() (*__premarshalCreateVehicleCreateVehicle, error) {
	var retval __premarshalCreateVehicleCreateVehicle

	retval.Reg = v.Vehicle.Reg
	return &retval, nil
}

// CreateVehicleResponse is returned by CreateVehicle on success.
type CreateVehicleResponse struct {
	CreateVehicle CreateVehicleCreateVehicle `json:"createVehicle"`
}

// GetCreateVehicle returns CreateVehicleResponse.CreateVehicle, and is useful for accessing the field via an interface.
func (v *CreateVehicleResponse) GetCreateVehicle() CreateVehicleCreateVehicle { return v.CreateVehicle }

// Vehicle includes the GraphQL fields of Vehicle requested by the fragment Vehicle.
type Vehicle struct {
	Reg string `json:"reg"`
}

// GetReg returns Vehicle.Reg, and is useful for accessing the field via an interface.
func (v *Vehicle) GetReg() string { return v.Reg }

// __CreateVehicleInput is used internally by genqlient
type __CreateVehicleInput struct {
	Reg string `json:"reg"`
}

// GetReg returns __CreateVehicleInput.Reg, and is useful for accessing the field via an interface.
func (v *__CreateVehicleInput) GetReg() string { return v.Reg }

func CreateVehicle(
	ctx context.Context,
	client graphql.Client,
	reg string,
) (*CreateVehicleResponse, error) {
	req := &graphql.Request{
		OpName: "CreateVehicle",
		Query: `
mutation CreateVehicle ($reg: String) {
	createVehicle(reg: $reg) {
		... Vehicle
	}
}
fragment Vehicle on Vehicle {
	reg
}
`,
		Variables: &__CreateVehicleInput{
			Reg: reg,
		},
	}
	var err error

	var data CreateVehicleResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}