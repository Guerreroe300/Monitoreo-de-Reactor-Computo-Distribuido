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
func (h *Handler) MainHtml(w http.ResponseWriter, req *http.Request) {
	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Nuclar Reactor Momen</title>
		</head>
		<body>
			<h1>Nuclear Reactor Temps</h1>

			<table border="1" id="myTable">
				<thead>
					<tr><th>Time</th><th>Temps</th></tr>
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
							tr.innerHTML = '<td>' + r.date + '</td><td>' + r.temperature + '</td>';
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

func (h *Handler) TableGet(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	temps, err := h.ctrl.GetAllDB(ctx)

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(temps)
}

func (h *Handler) ButtonHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Use POST", http.StatusMethodNotAllowed)
		return
	}
	// Your Go logic here
	fmt.Println("Button was clicked! Doing something in Go...")
	fmt.Fprint(w, "Go logic executed successfully!")
}
