package routers

import (
	"github.com/astaxie/beego"

	"github.com/dockercn/wharf/controllers"
)

func init() {
	//Web Interface
	beego.Router("/", &controllers.WebController{}, "get:GetIndex")
	beego.Router("/auth", &controllers.WebController{}, "get:GetAuth")
	beego.Router("/setting", &controllers.WebController{}, "get:GetSetting")
	beego.Router("/dashboard", &controllers.WebController{}, "get:GetDashboard")

	//Web API
	web := beego.NewNamespace("/w1",
		beego.NSRouter("/signin", &controllers.AuthWebController{}, "post:Signin"),
		beego.NSRouter("/signup", &controllers.AuthWebController{}, "post:Signup"),
		beego.NSRouter("/profile", &controllers.UsersWebController{}, "get:GetProfile"),

		//team routers
		beego.NSRouter("/users/:username", &controllers.UsersWebController{}, "get:GetUserExist"),
		beego.NSRouter("/team", &controllers.TeamWebController{}, "post:PostTeam"),

		//organization routers
		beego.NSRouter("/organizations", &controllers.OrganizationWebController{}, "get:GetOrganizations"),
		beego.NSRouter("/organization", &controllers.OrganizationWebController{}, "post:PostOrganization"),
		beego.NSRouter("/organization", &controllers.OrganizationWebController{}, "put:PutOrganization"),
		beego.NSRouter("/organizations/:orgName", &controllers.OrganizationWebController{}, "get:GetOrganizationDetail"),
	)

	//Docker Registry API V1 remain
	beego.Router("/_ping", &controllers.PingAPIController{}, "get:GetPing")

	//Docker Registry API V1
	apiv1 := beego.NewNamespace("/v1",
		beego.NSRouter("/_ping", &controllers.PingAPIController{}, "get:GetPing"),
		beego.NSRouter("/users", &controllers.UsersAPIController{}, "get:GetUsers"),
		beego.NSRouter("/users", &controllers.UsersAPIController{}, "post:PostUsers"),

		beego.NSNamespace("/repositories",
			beego.NSRouter("/:namespace/:repo_name/tags/:tag", &controllers.RepositoryAPIController{}, "put:PutTag"),
			beego.NSRouter("/:namespace/:repo_name/images", &controllers.RepositoryAPIController{}, "put:PutRepositoryImages"),
			beego.NSRouter("/:namespace/:repo_name/images", &controllers.RepositoryAPIController{}, "get:GetRepositoryImages"),
			beego.NSRouter("/:namespace/:repo_name/tags", &controllers.RepositoryAPIController{}, "get:GetRepositoryTags"),
			beego.NSRouter("/:namespace/:repo_name", &controllers.RepositoryAPIController{}, "put:PutRepository"),
		),

		beego.NSNamespace("/images",
			beego.NSRouter("/:image_id/ancestry", &controllers.ImageAPIController{}, "get:GetImageAncestry"),
			beego.NSRouter("/:image_id/json", &controllers.ImageAPIController{}, "get:GetImageJSON"),
			beego.NSRouter("/:image_id/layer", &controllers.ImageAPIController{}, "get:GetImageLayer"),
			beego.NSRouter("/:image_id/json", &controllers.ImageAPIController{}, "put:PutImageJSON"),
			beego.NSRouter("/:image_id/layer", &controllers.ImageAPIController{}, "put:PutImageLayer"),
			beego.NSRouter("/:image_id/checksum", &controllers.ImageAPIController{}, "put:PutChecksum"),
		),
	)

	beego.AddNamespace(web)
	beego.AddNamespace(apiv1)
}
