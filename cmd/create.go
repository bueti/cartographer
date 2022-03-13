/*
Copyright © 2022 Benjamin Buetikofer <bbu@ik.me>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/bueti/cartographer/create"
	"github.com/spf13/cobra"
)

var ChartName string
var ChartRepository string
var ChartVersion string
var Project string
var Namespace string
var ValueFiles string
var Secrets bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new App of Apps Chart",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,

	Run: func(cmd *cobra.Command, args []string) {
		c, _ := cmd.Flags().GetString("chartname")
		r, _ := cmd.Flags().GetString("repository")
		v, _ := cmd.Flags().GetString("version")
		p, _ := cmd.Flags().GetString("project")
		n, _ := cmd.Flags().GetString("namespace")
		f, _ := cmd.Flags().GetString("valuefiles")
		s, _ := cmd.Flags().GetBool("secrets")

		create.CreateApplicationCrd(c, r, v, p, n, f, s)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&ChartName, "chartname", "c", "", "Name of the source chart")
	createCmd.MarkFlagRequired("chartname")
	createCmd.Flags().StringVarP(&ChartRepository, "repository", "r", "", "URL of source repository")
	createCmd.MarkFlagRequired("repository")
	createCmd.Flags().StringVarP(&ChartVersion, "version", "v", "", "Chart version")
	createCmd.MarkFlagRequired("version")
	createCmd.Flags().StringVarP(&Project, "project", "p", "default", "Name of the ArgoCD project")
	createCmd.Flags().StringVarP(&Namespace, "namespace", "n", "", "Namespace to install chart into")
	createCmd.MarkFlagRequired("namespace")
	createCmd.Flags().StringVarP(&ValueFiles, "valuefiles", "f", "", "Valuefiles, eg.: file1,file2")
	createCmd.Flags().BoolVarP(&Secrets, "secrets", "s", false, "generate a secrets file?")
}
