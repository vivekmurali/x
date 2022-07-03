package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	godotenv.Load()
	r := chi.NewRouter()

	apiBasePath := "/auth"
	websiteBasePath := "/ui"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// try.supertokens.com is for demo purposes. Replace this with the address of your core instance (sign up on supertokens.com), or self host a core.
			ConnectionURI: "http://localhost:3567",
			// APIKey: "IF YOU HAVE AN API KEY FOR THE CORE, ADD IT HERE",
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "x",
			APIDomain:       "http://localhost:3000",
			WebsiteDomain:   "http://localhost:3000",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdparty.Init(&tpmodels.TypeInput{ /*TODO: See next step*/
				SignInAndUpFeature: tpmodels.TypeInputSignInAndUp{
					Providers: []tpmodels.TypeProvider{
						// We have provided you with development keys which you can use for testing.
						// IMPORTANT: Please replace them with your own OAuth keys for production use.
						thirdparty.Google(tpmodels.GoogleConfig{
							ClientSecret: os.Getenv("CLIENT_SECRET"),
							ClientID:     os.Getenv("CLIENT_ID"),
						}),
					},
				},
			}),
			session.Init(nil), // initializes session features
		},
	})

	if err != nil {
		panic(err.Error())
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))

	r.Use(middleware.Logger)
	r.Use(supertokens.Middleware)

	r.Get("/login", login)
	r.Get("/auth/callback/google", callback)

	http.ListenAndServe(":3000", r)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/index.html")
	tmpl.Execute(w, nil)
}

func callback(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("template/callback.html")
	tmpl.Execute(w, nil)

}
