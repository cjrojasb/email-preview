package main

import (
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EmailMessages struct {
	FirstName                   string
	LastName                    string
	TemporaryPassword           string
	FBCUrl                      string
	FrontAppUrl                 string
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
	// bulk create info
	BulkCreateInfoSubject        string
	BulkCreateInfoSubjectOfUsers string
	BulkCreateInfoTitle          string
	BulkCreateInfoHeader         string
	BulkCreateInfoHeaderID       string
	BulkCreateInfoHeaderDetail   string
	BulkCreateInfoBody           string
	// bulk create success
	BulkCreateSuccessSubject            string
	BulkCreateSuccessTitle              string
	BulkCreateSuccessHeader             string
	BulkCreateSuccessHeaderDetail       string
	BulkCreateSuccessHeaderDetailSystem string
	BulkCreateSuccessBody               string
	BulkCreateSuccessFooter             string
	// bulk create error
	BulkCreateErrorSubject                             string
	BulkCreateErrorTitle                               string
	BulkCreateErrorHeader                              string
	BulkCreateErrorUsersCreatedSuccessfully            string
	BulkCreateErrorUsersCreatedSuccessfullyDescription string
	BulkCreateErrorUsersNotCreated                     string
	BulkCreateErrorUsersNotCreatedDescription          string
	Greetings                                          string
	// bulk variables values
	Filename   string
	TotalUsers int
	RequestID  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	FBCUrl := os.Getenv("FBC_URL")
	FrontAppUrl := os.Getenv("FRONT_APP_URL")

	t, err := template.ParseFiles("./templates/bulk-create-info.html")
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
		// bulk messages
		BulkCreateInfoSubject:                              "Información importante sobre la solicitud",
		BulkCreateInfoSubjectOfUsers:                       "de creación de usuarios! ℹ️",
		BulkCreateInfoTitle:                                "Información sobre tu solicitud de creación de usuarios",
		BulkCreateInfoHeader:                               "Te informamos que el archivo",
		BulkCreateInfoHeaderID:                             "con ID de solicitud",
		BulkCreateInfoHeaderDetail:                         "fue recibido correctamente, procederemos a crear los usuarios.",
		BulkCreateInfoBody:                                 "Una vez finalizada la creación de usuarios, te informaremos por correo.",
		BulkCreateSuccessSubject:                           "Se han creado con éxito los usuarios solicitados! 🚀",
		BulkCreateSuccessTitle:                             "Información sobre tu solicitud de creación de usuarios",
		BulkCreateSuccessHeader:                            "Te informamos que al cargar el archivo",
		BulkCreateSuccessHeaderDetail:                      "se crearon correctamente los",
		BulkCreateSuccessHeaderDetailSystem:                "usuarios en el sistema.",
		BulkCreateSuccessBody:                              "Puedes revisar y gestionar esta información directamente en la plataforma Falabella Business Center, sección Roles y Perfiles.",
		BulkCreateSuccessFooter:                            "Si tienes alguna duda o necesitas apoyo adicional, no dudes en escribirnos.",
		BulkCreateErrorSubject:                             "Estado de creación de usuarios en la plataformas FBC! 🚀",
		BulkCreateErrorTitle:                               "Tenemos los resultados de tu solicitud de creación de usuarios.",
		BulkCreateErrorHeader:                              "A continuación, te compartimos dos archivos con el  detalle del resultado:",
		BulkCreateErrorUsersCreatedSuccessfully:            "✅ Usuarios creados correctamente:",
		BulkCreateErrorUsersCreatedSuccessfullyDescription: "listado de los usuarios que fueron ingresados exitosamente al sistema.",
		BulkCreateErrorUsersNotCreated:                     "⚠️ Usuarios no creados:",
		BulkCreateErrorUsersNotCreatedDescription:          "listado de los usuarios que no  pudieron ser creados, junto con los errores detectados. Te pedimos por favor revisar estos casos, corregirlos y volver a enviar el archivo actualizado.",
		Greetings:  "Saludos cordiales",
		Filename:   "create_external_users_2023.csv",
		TotalUsers: 100,
		RequestID:  "1234567890",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Preview generated: preview.html")
}
