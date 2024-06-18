/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
package magicapp

import (
	"github.com/magicbutton/magic-zones/services"
	"github.com/nats-io/nats.go/micro"
)

func RegisterServiceEndpoints(root micro.Group) {
    root.AddEndpoint("zonetype", micro.HandlerFunc(services.HandleZoneTypeRequests))
        root.AddEndpoint("zone", micro.HandlerFunc(services.HandleZoneRequests))
        root.AddEndpoint("accessrole", micro.HandlerFunc(services.HandleAccessRoleRequests))
        root.AddEndpoint("accesspass", micro.HandlerFunc(services.HandleAccessPassRequests))
        root.AddEndpoint("person", micro.HandlerFunc(services.HandlePersonRequests))
        root.AddEndpoint("consumerapp", micro.HandlerFunc(services.HandleConsumerAppRequests))
        root.AddEndpoint("appSecret", micro.HandlerFunc(services.HandleAppSecretRequests))
        root.AddEndpoint("auditlog", micro.HandlerFunc(services.HandleAuditLogRequests))
        root.AddEndpoint("user", micro.HandlerFunc(services.HandleUserRequests))
        root.AddEndpoint("service", micro.HandlerFunc(services.HandleServicexrRequests))
        root.AddEndpoint("route", micro.HandlerFunc(services.HandleRouteRequests))
    }
