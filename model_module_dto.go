/*
 * oKaPIBackend API
 *
 * oKaPIBackend API documentation
 *
 * API version: 6.3.1
 * Contact: software@sma.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ModuleDto struct {
	Id string `json:"id,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	CreateUser string `json:"createUser,omitempty"`
	LastModifiedTime time.Time `json:"lastModifiedTime,omitempty"`
	LastModifiedUser string `json:"lastModifiedUser,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
	ParentModule *ModuleDto `json:"parentModule,omitempty"`
}
