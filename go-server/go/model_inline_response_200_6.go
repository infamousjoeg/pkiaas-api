/*
 * CyberArk PKIaaS API
 *
 * API for CyberArk's Public-Key Infrastructure (PKI)-as-a-Service
 *
 * API version: 1.0.2
 * Contact: joe.garcia@cyberark.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse2006 struct {

	Count string `json:"count,omitempty"`

	Records []InlineResponse2006Records `json:"records,omitempty"`
}
