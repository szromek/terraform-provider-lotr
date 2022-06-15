module github.com/hashicorp/terraform-provider-lotr

go 1.17

require (
	github.com/hashicorp/terraform-plugin-docs v0.9.0
	github.com/hashicorp/terraform-plugin-framework v0.8.0
	github.com/hashicorp/terraform-plugin-go v0.9.1
	github.com/hashicorp/terraform-plugin-log v0.4.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.17.0
	github.com/szromek/lotr-client-go v0.0.0-00010101000000-000000000000
)

replace github.com/szromek/lotr-client-go => /Users/krzysztof/Documents/praca/terraform-provider/lotr-client-go
