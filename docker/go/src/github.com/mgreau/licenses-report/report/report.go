package report

import (
	"fmt"
	types "github.com/mgreau/licenses-report/types"
	"github.com/mitchellh/colorstring"

	"github.com/ryanuber/go-license"

	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// DisplayReport show all licenses for the dependencies found for the project
func DisplayReport(summary types.Summary, format string, outputDir string, reportFile string) {

	fmt.Printf(colorstring.Color("[red] \nProject: %s "), summary.ProjectName)
	fmt.Printf(colorstring.Color("[red] \nNb Dependencies: %d\n"), len(summary.Dependencies))

	if format == "json" {
		summaryJson, _ := json.MarshalIndent(summary, "", " ")
		var PathSeparator = fmt.Sprintf("%c", os.PathSeparator)
		ioutil.WriteFile(outputDir+PathSeparator+reportFile+".json", summaryJson, 0644)
	} else {
		fmt.Printf(colorstring.Color("[red] Dependencies list: \n"))
		for _, d := range summary.Dependencies {
			fmt.Printf(colorstring.Color("[blue] * %s - %s - \n"), d.Name, d.License.ID)
			fmt.Printf("JSON %+v", d)
			fmt.Println("\n\n-----")
		}
	}

}

// GenerateReport Generate report based on path to dependencies
func GenerateReport(params *types.Params) (report types.Summary, err error) {

	summary := types.Summary{
		ProjectName: params.Project,
	}
	var top = params.Path

	fmt.Printf("Project path %s", top)

	fileList := []string{}
	filepath.Walk(top, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})

	for _, file := range fileList {
		l, err := license.NewFromDir(file)
		if err == nil {
			_, f := path.Split(file)

			summary.Dependencies = append(summary.Dependencies,
				types.Dependency{
					Name: f,
					File: file,
					License: types.License{
						ID:   l.Type,
						Text: l.Text,
					},
				})
		}
	}

	DisplayReport(summary, params.Format, params.Output, params.Name)

	return report, nil
}
