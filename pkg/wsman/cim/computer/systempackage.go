/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package computer

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewComputerSystemPackage returns a new instance of the ComputerSystemPackage struct.
func NewComputerSystemPackageWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SystemPackage {
	return SystemPackage{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_ComputerSystemPackage, client),
		client: client,
	}
}

// Get retrieves the representation of the instance
func (systemPackage SystemPackage) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: systemPackage.base.Get(nil),
		},
	}

	err = systemPackage.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Enumerate returns an enumeration context which is used in a subsequent Pull call
func (systemPackage SystemPackage) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: systemPackage.base.Enumerate(),
		},
	}

	err = systemPackage.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (systemPackage SystemPackage) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: systemPackage.base.Pull(enumerationContext),
		},
	}
	err = systemPackage.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
