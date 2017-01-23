// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package specs

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestSwagger_Parse(t *testing.T) {
	filePath, err := filepath.Abs("fixtures/qingstor_sample/2016-01-06/swagger/api_v2.0.json")
	assert.Nil(t, err)
	swagger := &Swagger{
		FilePath: filePath,
	}

	err = swagger.Parse("v2.0")
	assert.Nil(t, err)

	assert.Equal(t, "QingStor", swagger.Data.Service.Name)
	assert.Equal(t, "Bucket", swagger.Data.SubServices["Bucket"].Name)
	assert.Equal(t, 5, len(swagger.Data.SubServices["Bucket"].Operations))

	owner := swagger.Data.SubServices["Bucket"].Operations["ListObjects"].Response.Elements.Properties["owner"]
	assert.Equal(t, "object", owner.Type)
	//assert.Equal(t, "owner", owner.ExtraType)

	listBuckets := swagger.Data.Service.Operations["ListBuckets"]
	location := listBuckets.Request.Headers.Properties["Location"]
	assert.Equal(t, "Location", location.Name)
	assert.Equal(t, "string", location.Type)

	bucket := swagger.Data.CustomizedTypes["bucket"]
	assert.Equal(t, "bucket", bucket.Name)
	assert.Equal(t, "Bucket", bucket.Description)
	assert.Equal(t, "object", bucket.Type)
}

func TestSwaggerParse(t *testing.T) {
	swagger := &Swagger{
		FilePath: "./swagger.json",
	}

	err := swagger.Parse("v2.0")
	assert.Nil(t, err)
	service := swagger.Data.Service
	subServices := swagger.Data.SubServices
	customizedTypes := swagger.Data.CustomizedTypes
	//service
	fmt.Println("---------------service start--------------")
	fmt.Println(service.APIVersion,service.Name,service.Description)
	fmt.Println(service.Properties.ID,service.Properties.Name,service.Properties.Description,service.Properties.Type,
	service.Properties.ExtraType,service.Properties.Format,service.Properties.Default,service.Properties.IsRequired)

	for _,v := range service.Properties.Enum {
		fmt.Println(v)
	}
	for k,v := range service.Properties.Properties {
		fmt.Println(k,v)
	}

	fmt.Println("---------------service end--------------")

	fmt.Println("---------------subservices start--------------")
	//subservices
	for k,v := range subServices {
		fmt.Println(k,v.ID,v.Name,v.Properties)
		for kk,vv := range v.Operations {
			fmt.Println(kk,vv.ID,vv.Name,vv.Description,vv.DocumentationURL)
			fmt.Println(vv.Request.Method,vv.Request.URI,vv.Request.Params,vv.Request.Headers,vv.Request.Elements,vv.Request.Body)
			for kkk,vvv := range vv.Response.StatusCodes {
				fmt.Println(kkk,vvv.Description)
			}
			fmt.Println(vv.Response.Headers,vv.Response.Elements,vv.Response.Body)
		}
	}
	fmt.Println("---------------subservices end--------------")

	fmt.Println("---------------customized start--------------")
	//customizedTypes
	for k,v := range customizedTypes {
		fmt.Println(k,v.ID,v.Name,v.Description,v.Type,v.ExtraType,v.Format,v.Default,v.IsRequired)
		for _,vv := range v.Enum {
			fmt.Println(vv)
		}
		for kk,vv := range v.Properties {
			fmt.Println(kk,vv)
		}
	}

	fmt.Println("---------------customized end--------------")
}