package config

type Config struct {
	WebEngine struct {
		ListenAddr string
	}
	SQL struct {
		Engine string
		SQLite struct {
			File string
			DB   string
		}
		MySQL struct {
			Addr     string
			User     string
			DB       string
			Password string
		}
	}
}
