// Generated by ego.
// DO NOT EDIT

package ego
import (
"fmt"
"html"
"io"
"github.com/SlinSo/goTemplateBenchmark/model"
)
var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
//line ego/simple.ego:2
 func EgoSimple(w io.Writer, u *model.User) error  {
//line ego/simple.ego:2
_, _ = io.WriteString(w, "\n")
//line ego/simple.ego:3
_, _ = io.WriteString(w, "\n<html>\n    <body>\n        <h1>")
//line ego/simple.ego:5
_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v",  u.FirstName )))
//line ego/simple.ego:5
_, _ = io.WriteString(w, "</h1>\n\n        <p>Here's a list of your favorite colors:</p>\n        <ul>\n        ")
//line ego/simple.ego:9
 for _, colorName := range u.FavoriteColors { 
//line ego/simple.ego:10
_, _ = io.WriteString(w, "\n            <li>")
//line ego/simple.ego:10
_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v",  colorName )))
//line ego/simple.ego:10
_, _ = io.WriteString(w, "</li>")
//line ego/simple.ego:10
 } 
//line ego/simple.ego:11
_, _ = io.WriteString(w, "\n        </ul>\n    </body>\n</html>")
return nil
}
