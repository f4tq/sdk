/***************************************************************************
 *
 * AVI CONFIDENTIAL
 * __________________
 *
 * [2013] - [2017] Avi Networks Incorporated
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is, and remains the property
 * of Avi Networks Incorporated and its suppliers, if any. The intellectual
 * and technical concepts contained herein are proprietary to Avi Networks
 * Incorporated, and its suppliers and are covered by U.S. and Foreign
 * Patents, patents in process, and are protected by trade secret or
 * copyright law, and other laws. Dissemination of this information or
 * reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Avi Networks Incorporated.
 */

package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// WafProfileClient is a client for avi WafProfile resource
type WafProfileClient struct {
	aviSession *session.AviSession
}

// NewWafProfileClient creates a new client for WafProfile resource
func NewWafProfileClient(aviSession *session.AviSession) *WafProfileClient {
	return &WafProfileClient{aviSession: aviSession}
}

func (client *WafProfileClient) getAPIPath(uuid string) string {
	path := "api/wafprofile"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of WafProfile objects
func (client *WafProfileClient) GetAll() ([]*models.WafProfile, error) {
	var plist []*models.WafProfile
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist)
	return plist, err
}

// Get an existing WafProfile by uuid
func (client *WafProfileClient) Get(uuid string) (*models.WafProfile, error) {
	var obj *models.WafProfile
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj)
	return obj, err
}

// GetByName - Get an existing WafProfile by name
func (client *WafProfileClient) GetByName(name string) (*models.WafProfile, error) {
	var obj *models.WafProfile
	err := client.aviSession.GetObjectByName("wafprofile", name, &obj)
	return obj, err
}

// Create a new WafProfile object
func (client *WafProfileClient) Create(obj *models.WafProfile) (*models.WafProfile, error) {
	var robj *models.WafProfile
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj)
	return robj, err
}

// Update an existing WafProfile object
func (client *WafProfileClient) Update(obj *models.WafProfile) (*models.WafProfile, error) {
	var robj *models.WafProfile
	path := client.getAPIPath(obj.UUID)
	err := client.aviSession.Put(path, obj, &robj)
	return robj, err
}

// Delete an existing WafProfile object with a given UUID
func (client *WafProfileClient) Delete(uuid string) error {
	return client.aviSession.Delete(client.getAPIPath(uuid))
}

// DeleteByName - Delete an existing WafProfile object with a given name
func (client *WafProfileClient) DeleteByName(name string) error {
	res, err := client.GetByName(name)
	if err != nil {
		return err
	}
	return client.Delete(res.UUID)
}
