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

type Body4 struct {

	CommonName string `json:"commonName,omitempty"`

	KeyAlgo string `json:"keyAlgo,omitempty"`

	KeyBits string `json:"keyBits,omitempty"`

	TimeToLive string `json:"timeToLive,omitempty"`

	Subject string `json:"subject,omitempty"`

	AltNames []string `json:"altNames,omitempty"`

	SelfSigned bool `json:"selfSigned,omitempty"`
}
