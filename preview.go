package main

import (
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailMessages struct {
	FirstName         string
	LastName          string
	TemporaryPassword string
	FBCUrl            string
	FrontAppUrl       string
	// TODO
	Language                    string
	Greeting                    string
	Welcome                     string
	UseAD                       string
	TemporaryPasswordMessage    string
	UpdatePassword              string
	EnterHere                   string
	LinkExpires                 string
	StepsToUpdatePassword       string
	StepsToUpdatePasswordFirst  string
	StepsToUpdatePasswordSecond string
	StepsToUpdatePasswordThird  string
	Team                        string
	Footer                      string
	Internal                    bool
	// internal email
	InternalWelcome     string
	InternalSubtitle    string
	InternalAccess      string
	InternalLink        string
	ToEnterBy           string
	HowToEnterThePortal string
	StepsToEnterFirst   string
	StepsToEnterSecond  string
	FalabellaAzureAD    string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	FBCUrl := os.Getenv("FBC_URL")
	FrontAppUrl := os.Getenv("FRONT_APP_URL")

	t, err := template.ParseFiles("./templates/email-template.html")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("preview.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = t.Execute(file, EmailMessages{
		FirstName:         "Carlos",
		LastName:          "Rojas",
		TemporaryPassword: "xxxxxxxxx",
		FBCUrl:            FBCUrl,
		Team:              "Equipo Falabella",
		FrontAppUrl:       FrontAppUrl,
		// TODO
		Language:                    "es",
		Greeting:                    "Hola",
		Welcome:                     "Te damos la bienvenida a",
		UseAD:                       "y usa Falabella AD",
		TemporaryPasswordMessage:    "Tu contraseña temporal es",
		UpdatePassword:              "Actualiza tu contraseña aquí",
		EnterHere:                   "Ingresa aquí",
		LinkExpires:                 "Este enlace expirará en 48 horas",
		StepsToUpdatePassword:       "Pasos para actualizar tu contraseña",
		StepsToUpdatePasswordFirst:  "Haz clic en el botón",
		StepsToUpdatePasswordSecond: "Accede a fbusinesscenter.com, inicia sesión con tu correo y contraseña temporal, y luego crea una nueva contraseña que contenga al menos 8 caracteres",
		StepsToUpdatePasswordThird:  "¡Listo! Ya has actualizado tu contraseña y puedes ingresar a Falabella Business Center con tu mail y nueva contraseña",
		Footer:                      "Este es un mensaje automático. Por favor, no lo responda",
		Internal:                    true,
		// internal email
		InternalWelcome:     "Te damos la bienvenida",
		InternalSubtitle:    "Te informamos que se te ha otorgado acceso a",
		InternalAccess:      "Accede al portal",
		InternalLink:        FBCUrl,
		ToEnterBy:           "E ingresa por",
		HowToEnterThePortal: "Cómo ingresar al portal",
		StepsToEnterFirst:   "Al ser Colaborador del Grupo Falabella, accede directamente desde",
		StepsToEnterSecond:  "Ahora puedes acceder a los Módulos asignados en el Falabella Business Center",
		FalabellaAzureAD:    "Falabella Azure AD",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Preview generated: preview.html")
}
