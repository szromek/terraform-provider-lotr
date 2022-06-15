package provider

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ tfsdk.DataSourceType = characterDataSourceType{}
var _ tfsdk.DataSource = characterDataSource{}

type characterDataSourceType struct{}

func (t characterDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Character featuring in Lord of The Rings",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "Character identifier",
				Type:                types.StringType,
				Computed:            true,
			},
			"height": {
				MarkdownDescription: "Height of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"race": {
				MarkdownDescription: "race of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"gender": {
				MarkdownDescription: "gender of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"birth": {
				MarkdownDescription: "birth of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"spouse": {
				MarkdownDescription: "spouse of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"death": {
				MarkdownDescription: "death of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"realm": {
				MarkdownDescription: "realm of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"hair": {
				MarkdownDescription: "hair of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"name": {
				MarkdownDescription: "name of a character",
				Type:                types.StringType,
				Optional:            true,
			},
			"wiki_url": {
				MarkdownDescription: "wiki of a character",
				Type:                types.StringType,
				Optional:            true,
			},
		},
	}, nil
}

func (t characterDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return characterDataSource{
		provider: provider,
	}, diags
}

type characterDataSourceData struct {
	Id      types.String `tfsdk:"id"`
	Height  types.String `tfsdk:"height"`
	Race    types.String `tfsdk:"race"`
	Gender  types.String `tfsdk:"gender"`
	Birth   types.String `tfsdk:"birth"`
	Spouse  types.String `tfsdk:"spouse"`
	Death   types.String `tfsdk:"death"`
	Realm   types.String `tfsdk:"realm"`
	Hair    types.String `tfsdk:"hair"`
	Name    types.String `tfsdk:"name"`
	WikiUrl types.String `tfsdk:"wiki_url"`
}

type characterDataSource struct {
	provider provider
}

func (d characterDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data characterDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	log.Printf("got here")

	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("got here")

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	characters, err := d.provider.client.GetCharacters()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
		return
	}

	// characters[rand.Intn(len(characters))]

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	// data.Id = types.String{Value: "example-id"}

	character := characters[rand.Intn(len(characters))]
	characterState := characterDataSourceData{
		Id:      types.String{Value: character.ID},
		Height:  types.String{Value: character.Height},
		Race:    types.String{Value: character.Race},
		Gender:  types.String{Value: character.Gender},
		Birth:   types.String{Value: character.Birth},
		Spouse:  types.String{Value: character.Spouse},
		Death:   types.String{Value: character.Death},
		Realm:   types.String{Value: character.Realm},
		Hair:    types.String{Value: character.Hair},
		Name:    types.String{Value: character.Name},
		WikiUrl: types.String{Value: character.WikiUrl},
	}

	diags = resp.State.Set(ctx, characterState)
	resp.Diagnostics.Append(diags...)
}
