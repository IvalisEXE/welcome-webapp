package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {

	//instalasi objek struct Welcome dan diberikan informasi/nilai secara acak
	//dan dijadikan sebagai parameter query dari URL
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//melokasikan file html dan meminta go memparsing file html kemudian membungkusnya ke dalam-
	//template.Must() yang menangani kesalahan dan berhenti jika ada kesalahan fatal
	templates := template.Must(template.ParseFiles("templates/welcome-templates.html"))

	//untuk menghandle css di direktori /static sebagai url, html dapat merujuk ketika mencari css dan file lainya
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	//url ditunjukan di http.Handle("/static/") url ini adalah apa yang kita butuhkan saat-
	//merefensikan file css kita, setelah server dimulai kode html akan menjadi-
	// <link rel = "stylesheet" href = "/ static / stylesheet / ...">, Penting untuk dicatat url di http.Handle bisa apa pun yang kita suka, selama kita konsisten.
	// Metode ini mengambil jalur URL "/" dan fungsi yang mengambil penulis respons, dan permintaan http.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Mengambil nama dari kueri URL misalnya? Name = Arip Saputra, akan menyetel welcome.Name = Arip Saputra.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		// Jika kesalahan menunjukkan pesan kesalahan server internal juga meneruskan struk selamat datang ke file welcome-template.html.
		if err := templates.ExecuteTemplate(w, "welcome-templates.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Jalankan server web, atur port untuk mendengarkan 8080. Tanpa jalur yang diasumsikan localhost
	// Cetak semua kesalahan dari memulai server web menggunakan fmt
	fmt.Println("Start on port :8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
