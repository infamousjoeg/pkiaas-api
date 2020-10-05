/*
 * CyberArk Identity as a Service API
 *
 * API for CyberArk's Identity as a Service (IDaaS)
 *
 * API version: 1.0.4
 * Contact: joe.garcia@cyberark.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"GenerateIntermediateCACSR",
		strings.ToUpper("Post"),
		"/ca/generate",
		GenerateIntermediateCACSR,
	},

	Route{
		"GetCA",
		strings.ToUpper("Get"),
		"/ca/certificate",
		GetCA,
	},

	Route{
		"GetCAChain",
		strings.ToUpper("Get"),
		"/ca/chain",
		GetCAChain,
	},

	Route{
		"SetCAChain",
		strings.ToUpper("Post"),
		"/ca/chain/set",
		SetCAChain,
	},

	Route{
		"SetIntermediateCA",
		strings.ToUpper("Post"),
		"/ca/set",
		SetIntermediateCA,
	},

	Route{
		"CreatePKICertificate",
		strings.ToUpper("Post"),
		"/pki/certificate/create",
		CreatePKICertificate,
	},

	Route{
		"CreateSSHCertificate",
		strings.ToUpper("Post"),
		"/ssh/certificate/create",
		CreateSSHCertificate,
	},

	Route{
		"GetCertificate",
		strings.ToUpper("Get"),
		"/pki/certificate/{serialNumber}",
		GetCertificate,
	},

	Route{
		"ListCertificates",
		strings.ToUpper("Get"),
		"/pki/certificates",
		ListCertificates,
	},

	Route{
		"PurgeCertificates",
		strings.ToUpper("Post"),
		"/pki/purge",
		PurgeCertificates,
	},

	Route{
		"RevokeCertificate",
		strings.ToUpper("Post"),
		"/pki/certificate/revoke",
		RevokeCertificate,
	},

	Route{
		"SignCertificate",
		strings.ToUpper("Post"),
		"/pki/certificate/sign",
		SignCertificate,
	},

	Route{
		"GetCRL",
		strings.ToUpper("Get"),
		"/pki/crl",
		GetCRL,
	},

	Route{
		"CreatePKITemplate",
		strings.ToUpper("Post"),
		"/pki/template",
		CreatePKITemplate,
	},

	Route{
		"CreateSSHTemplate",
		strings.ToUpper("Post"),
		"/ssh/template",
		CreateSSHTemplate,
	},

	Route{
		"DeletePKITemplate",
		strings.ToUpper("Delete"),
		"/pki/template/{templateName}",
		DeletePKITemplate,
	},

	Route{
		"DeleteSSHTemplate",
		strings.ToUpper("Delete"),
		"/ssh/template/{templateName}",
		DeleteSSHTemplate,
	},

	Route{
		"GetPKITemplate",
		strings.ToUpper("Get"),
		"/pki/template/{templateName}",
		GetPKITemplate,
	},

	Route{
		"GetSSHTemplate",
		strings.ToUpper("Get"),
		"/ssh/template/{templateName}",
		GetSSHTemplate,
	},

	Route{
		"ListPKITemplates",
		strings.ToUpper("Get"),
		"/pki/templates",
		ListPKITemplates,
	},

	Route{
		"ListSSHTemplates",
		strings.ToUpper("Get"),
		"/ssh/templates",
		ListSSHTemplates,
	},

	Route{
		"ManagePKITemplate",
		strings.ToUpper("Put"),
		"/pki/template",
		ManagePKITemplate,
	},

	Route{
		"ManageSSHTemplate",
		strings.ToUpper("Put"),
		"/ssh/template",
		ManageSSHTemplate,
	},
}
