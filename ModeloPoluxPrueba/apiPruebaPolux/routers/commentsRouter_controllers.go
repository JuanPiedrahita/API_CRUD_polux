package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreaConocimientoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AreasTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:AsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CarreraElegibleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ComentarioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ContactoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:CorreccionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DetalleTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DistincionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoEntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:DocumentoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EntidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EspaciosAcademicosElegiblesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoAsignaturaTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEspacioAcademicoInscritoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoEstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoRevisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstadoTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstructuraInvestigacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EstudianteTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:EvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:FormatoEvaluacionCarreraController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:ModalidadTipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:PreguntaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaEvaluacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaFormatoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RespuestaSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RevisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:RolTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SocializacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:SolicitudTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoContactoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDetalleController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoDocumentoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoIdentificacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoPreguntaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TipoSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:TrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:UsuarioSolicitudController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"] = append(beego.GlobalControllerRouter["github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers:VinculacionTrabajoGradoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
