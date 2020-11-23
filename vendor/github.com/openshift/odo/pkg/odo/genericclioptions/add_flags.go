package genericclioptions

import "github.com/spf13/cobra"

// AddContextFlag adds `context` flag to given cobra command
func AddContextFlag(cmd *cobra.Command, setValueTo *string) {
	helpMessage := "Use given context directory as a source for component settings"
	if setValueTo != nil {
		cmd.Flags().StringVar(setValueTo, ContextFlagName, "", helpMessage)
	} else {
		cmd.Flags().String(ContextFlagName, "", helpMessage)
	}
}

// AddNowFlag adds `now` flag to given cobra command
func AddNowFlag(cmd *cobra.Command, setValueTo *bool) {
	helpMessage := "Push changes to the cluster immediately"
	if setValueTo != nil {
		cmd.Flags().BoolVar(setValueTo, "now", false, helpMessage)
	} else {
		cmd.Flags().Bool("now", false, helpMessage)
	}
}
