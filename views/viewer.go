package views

import (
	"fmt"
	rice "github.com/42wim/go.rice"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:generate rice embed-go

// splitPathURL
func splitPathURL(r *http.Request) []string {
	reqUrl := r.URL.Path
	var splitURL []string
	splitURL = strings.Split(reqUrl, "/")
	splitURL = splitURL[1:]
	return splitURL
}

// Viewer structures the set of app page views
type Viewer struct {
	TemplateMap template.FuncMap
	Templates   *template.Template
	TemplateBox *rice.Box
	Statics     *rice.HTTPBox
}

// InitializeViewer initializes the set of page views
func InitializeViewer() *Viewer {
	templateMap := template.FuncMap{
		"Upper": func(s string) string {
			return strings.ToUpper(s)
		},
	}
	templates := template.New("").Funcs(templateMap)
	var templateBox *rice.Box
	// Load and parse templates (from binary or disk)
	return &Viewer{
		TemplateMap: templateMap,
		Templates:   templates,
		TemplateBox: templateBox,
	}
}

// InitializeTemplates to use in page views
func (p *Viewer) InitializeTemplates() {
	// newTemplate
	var templateBox *rice.Box
	newTemplate := func(path string, fileInfo os.FileInfo, _ error) error {
		if path == "" {
			return nil
		}
		/*
		 * takeRelativeTo function will take the absolute path 'path' which is by default passed to
		 * our 'newTemplate' by Walk function, and will eliminate the intial part of the path up to the end of the
		 * specified directory 'afterDir' ('templates' in this case). Then it will return the rest starting from
		 * the very end of afterDir. If the specified afterDir has more than 1 occurances in the path,
		 * only the first occurance will be considered and the other occurances will be ignored.
		 * eg, If path = "/home/Projects/go/website/templates/html/index.html", then
		 * relativPath := takeRelativeTo(path, "templates") returns "/html/index.html" ;
		 * If path = "/home/Projects/go/website/templates/testing.html", then ;
		 * relativPath := takeRelativeTo(path, "templates") returns "/testing.html" ;
		 * If path = "/home/Projects/go/website/templates/html/templates/components/footer.html", then
		 * relativPath := takeRelativeTo(path, "templates") returns "/html/templates/components/footer.html" .
		 */
		takeRelativeTo := func(givenpath string, afterDir string) string {
			if strings.Contains(givenpath, afterDir+string(filepath.Separator)) {
				wantedpart := strings.SplitAfter(givenpath, afterDir)[1:]
				return filepath.Join(wantedpart...)
			}
			return givenpath
		}
		//if path is a directory, skip Parsing template. Trying to Parse a template from a directory caused an error, now fixed.
		if !fileInfo.IsDir() {
			//get relative path starting from the end of 'templates' .
			relativPath := takeRelativeTo(path, "templates")
			templateString, err := templateBox.String(relativPath)
			if err != nil {
				log.Panicf("Unable to extract: path=%s, err=%s", relativPath, err)
			}
			if _, err = p.Templates.New(filepath.Join("templates", relativPath)).Parse(templateString); err != nil {
				log.Panicf("Unable to parse: path=%s, err=%s", relativPath, err)
			}
		}
		return nil
	}
	templateBox = rice.MustFindBox("templates")
	templateBox.Walk("", newTemplate)
	static := rice.MustFindBox("static").HTTPBox()
	p.Statics = static
	p.TemplateBox = templateBox
}

// RenderTemplate - Render a template given a model
func (p *Viewer) RenderTemplate(w http.ResponseWriter, tmpl string, pInt interface{}) {
	err := p.Templates.ExecuteTemplate(w, tmpl, pInt)
	if err != nil {
		fmt.Println("\n\n\nCHECK ERROR: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
