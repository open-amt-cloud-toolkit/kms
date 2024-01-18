/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// NewMediaAccessDevice returns a new instance of the MediaAccessDevice struct.
func NewMediaAccessDeviceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Device {
	return Device{
		base:   message.NewBaseWithClient(wsmanMessageCreator, CIM_MediaAccessDevice, client),
		client: client,
	}
}

// TODO: Figure out how to call GET requiring resourceURIs and Selectors
// Get retrieves the representation of the instance

// Enumerates the instances of this class
func (device Device) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: device.base.Enumerate(),
		},
	}

	err = device.base.Execute(response.Message)
	if err != nil {
		return
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return

}

// Pulls instances of this class, following an Enumerate operation
func (device Device) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: device.base.Pull(enumerationContext),
		},
	}
	err = device.base.Execute(response.Message)
	if err != nil {
		return
	}
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}
