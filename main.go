
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Project struct {
	Number      string
	Name        string
	Description string
	Technology  string
	Category    string
}

type Portfolio struct {
	Name        string
	Role        string
	Description string
	Class       string
	Hobby       string
	Motto       string
	GitHub      string
	Skills      []string
	Projects    []Project
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := Portfolio{
		Name:        "Arya Wangsa Goparana",
		Role:        "Pelajar TKJ | Linux | Networking | Golang",
		Description: "Saya adalah pelajar Teknik Komputer dan Jaringan yang tertarik pada Linux, infrastruktur jaringan, pengembangan perangkat lunak, dan keamanan siber. Saya terus belajar melalui praktik dan proyek mandiri.",
		Class:       "XII TKJ 2",
		Hobby:       "Menjelajahi alam dan mendaki",
		Motto:       "Terus belajar, terus berkembang.",
		GitHub:      "https://github.com/aryawangsa063-dot",

		Skills: []string{
			"Linux & Debian",
			"Golang",
			"Jaringan Komputer",
			"Docker",
			"HTML & CSS",
			"Server Administration",
		},

		Projects: []Project{
			{
				Number:      "01",
				Name:        "Lab Server Debian",
				Description: "Laboratorium pembelajaran layanan server Linux, meliputi DHCP, DNS, web server, dan FTP.",
				Technology:  "Debian Linux",
				Category:    "Infrastructure",
			},
			{
				Number:      "02",
				Name:        "AI Command Assistant",
				Description: "Eksperimen asisten terminal untuk membantu memahami dan menjalankan perintah Linux.",
				Technology:  "Python · Groq",
				Category:    "AI & CLI",
			},
			{
				Number:      "03",
				Name:        "Smart Home Arduino",
				Description: "Proyek rumah mini dengan konsep RFID, LCD, servo, dan kontrol perangkat.",
				Technology:  "Arduino · IoT",
				Category:    "Internet of Things",
			},
			{
				Number:      "04",
				Name:        "Personal Portfolio",
				Description: "Website CV pribadi yang dibangun sambil mempelajari pemrograman backend Go.",
				Technology:  "Golang · HTML · CSS",
				Category:    "Web Development",
			},
		},
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Gagal memuat halaman", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Render error:", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	staticFiles := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFiles))
	http.HandleFunc("/", homeHandler)

	addr := "0.0.0.0:" + port
	log.Printf("Portfolio server berjalan di %s", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
