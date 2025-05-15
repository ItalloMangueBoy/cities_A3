package views

import (
	"text/template"
)

func ListCities() *template.Template {
	return template.Must(template.New("list_cities").Parse(`
		<!DOCTYPE html>
		<html>
		<body>
				<h1>Lista de cidades</h1>
				<ul>
						{{range .Cities}}
								<li>{{.Name}}</li>
						{{else}}
								<li>Nenhum item encontrado.</li>
						{{end}}
				</ul>

				<p>Token utilizado: {{.Token}}</p>
		</body>
		<footer>
			<p><a href="https://github.com/ItalloMangueBoy">confira o codigo do projeto</a></p>
		</footer>
		</html>
	`))
}
