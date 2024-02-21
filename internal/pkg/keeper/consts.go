//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package keeper

const (
	ApiVersion                      = "v3"
	ApiBase                         = "/api/" + ApiVersion
	ContentType                     = "Content-Type"
	ContentTypeJSON                 = "application/json"
	ContentTypeText                 = "text/plain"
	ApiPingRoute                    = ApiBase + "/ping"
	ApiRegisterRoute                = ApiBase + "/registry"
	ApiAllRegistrationRoute         = ApiRegisterRoute + "/all"
	ApiRegistrationByServiceIdRoute = ApiRegisterRoute + "/serviceId/"
)
