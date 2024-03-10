/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package service

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"AssociatedPowerManagementService\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"AssociatedPowerManagementService\":{\"AvailableRequestedPowerStates\":0,\"PowerState\":0,\"ServiceProvided\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelelctorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selector\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Name\":\"\",\"Value\":\"\"}}}},\"UserOfService\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Address\":\"\",\"ReferenceParameters\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ResourceURI\":\"\",\"SelelctorSet\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Selector\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"Name\":\"\",\"Value\":\"\"}}}}}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			PullResponse: PullResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    associatedpowermanagementservice: []\nenumerateresponse:\n    enumerationcontext: \"\"\nassociatedpowermanagementservice:\n    availablerequestedpowerstates: 0\n    powerstate: 0\n    serviceprovided:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selelctorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selector:\n                    xmlname:\n                        space: \"\"\n                        local: \"\"\n                    name: \"\"\n                    value: \"\"\n    userofservice:\n        xmlname:\n            space: \"\"\n            local: \"\"\n        address: \"\"\n        referenceparameters:\n            xmlname:\n                space: \"\"\n                local: \"\"\n            resourceuri: \"\"\n            selelctorset:\n                xmlname:\n                    space: \"\"\n                    local: \"\"\n                selector:\n                    xmlname:\n                        space: \"\"\n                        local: \"\"\n                    name: \"\"\n                    value: \"\"\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestNegativeAvailableToElement(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/service/availabletoelement",
	}
	elementUnderTest := NewServiceAvailableToElementWithClient(wsmanMessageCreator, &client)

	t.Run("cim_* Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//ENUMERATES
			{
				"should handle error making cim_ServiceAvailableToElement Enumerate call",
				CIM_ServiceAvailableToElement,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "DD020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should handle error making cim_ServiceAvailableToElement Pull call",
				CIM_ServiceAvailableToElement,
				wsmantesting.PULL,
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						AssociatedPowerManagementService: []CIM_AssociatedPowerManagementService{
							{
								AvailableRequestedPowerStates: 10,
								PowerState:                    2,
								ServiceProvided: ServiceProvided{
									XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/08/addressing", Local: "ServiceProvided"},
									Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
									ReferenceParameters: ReferenceParameters{
										XMLName:     xml.Name{Space: "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd", Local: "ReferenceParameters"},
										ResourceURI: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService",
										SelelctorSet: message.SelectorSet{
											Selector: message.Selector{
												Name:  "CreationClassName",
												Value: "CIM_PowerManagementService",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
