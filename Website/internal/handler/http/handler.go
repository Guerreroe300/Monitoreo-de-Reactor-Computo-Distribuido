package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	website "github.com/Guerreroe300/Monitoreo-de-Reactor-Computo-Distribuido/Website/internal/controller/website"
)

// to be removed for Temperature
type Row struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Handler struct {
	ctrl *website.Controller
}

func New(ctrl *website.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// HTML
func (h *Handler) MainHtml(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Go Table Example</title>
		</head>
		<body>
			<h1>Quick Go Table</h1>

			<table border="1" id="myTable">
				<thead>
					<tr><th>Name</th><th>Value</th></tr>
				</thead>
				<tbody></tbody>
			</table>
			
			<br>
			<button onclick="doSomething()">Do Something</button>

			<script>
				// Fetch table data
				fetch('/api/data')
					.then(res => res.json())
					.then(rows => {
						const tbody = document.querySelector('#myTable tbody');
						rows.forEach(r => {
							const tr = document.createElement('tr');
							tr.innerHTML = '<td>' + r.name + '</td><td>' + r.value + '</td>';
							tbody.appendChild(tr);
						});
					});

				function doSomething() {
					fetch('/api/doSomething', { method: 'POST' })
						.then(res => res.text())
						.then(msg => alert(msg));
				}
			</script>
		</body>
		</html>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func (h *Handler) TableGet(w http.ResponseWriter, r *http.Request) {
	rows := []Row{
		{"Apples", "10"},
		{"Oranges", "20"},
		{"Bananas", "15"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rows)
}

func (h *Handler) ButtonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST", http.StatusMethodNotAllowed)
		return
	}
	// Your Go logic here
	fmt.Println("Button was clicked! Doing something in Go...")
	fmt.Fprint(w, "Go logic executed successfully!")
}
