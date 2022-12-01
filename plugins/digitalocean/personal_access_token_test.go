package digitalocean

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestPersonalAccessTokenImporter(t *testing.T) {
	plugintest.TestImporter(t, PersonalAccessToken().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"DIGITALOCEAN_ACCESS_TOKEN": "dop_v1_dk98ysntlv1045mdhneztbd4o1r3q8p0tndohkpfii5m6049a8lacaq4iexample",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[string]string{
						fieldname.Token: "dop_v1_dk98ysntlv1045mdhneztbd4o1r3q8p0tndohkpfii5m6049a8lacaq4iexample",
					},
				},
			},
		},
	})
}

func TestPersonalAccessTokenProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, PersonalAccessToken().Provisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[string]string{
				fieldname.Token: "dop_v1_dk98ysntlv1045mdhneztbd4o1r3q8p0tndohkpfii5m6049a8lacaq4iexample",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"DIGITALOCEAN_ACCESS_TOKEN": "dop_v1_dk98ysntlv1045mdhneztbd4o1r3q8p0tndohkpfii5m6049a8lacaq4iexample",
				},
			},
		},
	})
}