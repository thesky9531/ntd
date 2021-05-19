package form

import "github.com/spf13/pflag"

// Settings contains basic settings for any beat to pass into GenRootCmd
type Settings struct {
	Name          string
	IndexPrefix   string
	Version       string
	HasDashboards bool
	//Monitoring      report.Settings
	RunFlags *pflag.FlagSet
	//ConfigOverrides []cfgfile.ConditionalOverride

	DisableConfigResolver bool

	// load custom index manager. The config object will be the Beats root configuration.
	//IndexManagement idxmgmt.SupportFactory
	//ILM             ilm.SupportFactory

	//Processing processing.SupportFactory

	Umask *int
}
