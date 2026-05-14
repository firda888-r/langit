func main() {
	// 1. Koneksi DB (Pastikan Docker Postgres nyala)
	db, _ := sql.Open("postgres", "user=postgres password=mysecretpassword dbname=postgres sslmode=disable")

	// 2. Inisialisasi Layer (Urut dari bawah ke atas)
	repo := repository.NewUserRepository(db)
	svc := service.NewAuthService(repo)
	hdl := handler.NewAuthHandler(svc)

	// 3. Routing
	http.HandleFunc("/api/auth/login", hdl.LoginHandler)
	
	fmt.Println("Account Service running on :8080")
	http.ListenAndServe(":8080", nil)
}