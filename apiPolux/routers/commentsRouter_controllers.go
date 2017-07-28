package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/apiPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
