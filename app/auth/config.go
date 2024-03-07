package auth

import "github.com/twilio/twilio-go"

var accountSid string = "AC8a7bab85f6fe95be36f8c2ba7a7759c5"

var (
	authToken        string = "5e57ff865c1ef80a4d56babe83e1cbf5"
	verifyServiceSid string = "VA5d3d43a2395940b36197e5e66cd55d3a"
)

var client = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: accountSid,
	Password: authToken,
})
