/*
 * CyberArk Identity Service API
 *
 * API for CyberArk's Identity Service
 *
 * API version: 1.0.4
 * Contact: joe.garcia@cyberark.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SshTemplateRequest struct {

	TemplateName string `json:"templateName"`

	CertType string `json:"certType"`

	MaxTTL int32 `json:"maxTTL"`

	AllowedHosts []string `json:"allowedHosts,omitempty"`

	AllowedPrincipals []string `json:"allowedPrincipals,omitempty"`

	PermittedCriticalOptions []SshcertificatecreateCriticalOptions `json:"permittedCriticalOptions,omitempty"`

	PermittedExtensions []string `json:"permittedExtensions,omitempty"`
}
