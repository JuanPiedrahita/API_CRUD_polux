// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/JuanPiedrahita/Polux/ModeloPoluxPrueba/apiPruebaPolux/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_documento",
			beego.NSInclude(
				&controllers.TipoDocumentoController{},
			),
		),

		beego.NSNamespace("/estado_solicitud",
			beego.NSInclude(
				&controllers.EstadoSolicitudController{},
			),
		),

		beego.NSNamespace("/documento_solicitud",
			beego.NSInclude(
				&controllers.DocumentoSolicitudController{},
			),
		),

		beego.NSNamespace("/respuesta_solicitud",
			beego.NSInclude(
				&controllers.RespuestaSolicitudController{},
			),
		),

		beego.NSNamespace("/usuario_solicitud",
			beego.NSInclude(
				&controllers.UsuarioSolicitudController{},
			),
		),

		beego.NSNamespace("/tipo_solicitud",
			beego.NSInclude(
				&controllers.TipoSolicitudController{},
			),
		),

		beego.NSNamespace("/tipo_detalle",
			beego.NSInclude(
				&controllers.TipoDetalleController{},
			),
		),

		beego.NSNamespace("/detalle",
			beego.NSInclude(
				&controllers.DetalleController{},
			),
		),

		beego.NSNamespace("/detalle_tipo_solicitud",
			beego.NSInclude(
				&controllers.DetalleTipoSolicitudController{},
			),
		),

		beego.NSNamespace("/detalle_solicitud",
			beego.NSInclude(
				&controllers.DetalleSolicitudController{},
			),
		),

		beego.NSNamespace("/documento_entidad",
			beego.NSInclude(
				&controllers.DocumentoEntidadController{},
			),
		),

		beego.NSNamespace("/pregunta",
			beego.NSInclude(
				&controllers.PreguntaController{},
			),
		),

		beego.NSNamespace("/respuesta",
			beego.NSInclude(
				&controllers.RespuestaController{},
			),
		),

		beego.NSNamespace("/respuesta_formato",
			beego.NSInclude(
				&controllers.RespuestaFormatoController{},
			),
		),

		beego.NSNamespace("/respuesta_evaluacion",
			beego.NSInclude(
				&controllers.RespuestaEvaluacionController{},
			),
		),

		beego.NSNamespace("/comentario",
			beego.NSInclude(
				&controllers.ComentarioController{},
			),
		),

		beego.NSNamespace("/correccion",
			beego.NSInclude(
				&controllers.CorreccionController{},
			),
		),

		beego.NSNamespace("/carrera_elegible",
			beego.NSInclude(
				&controllers.CarreraElegibleController{},
			),
		),

		beego.NSNamespace("/espacios_academicos_elegibles",
			beego.NSInclude(
				&controllers.EspaciosAcademicosElegiblesController{},
			),
		),

		beego.NSNamespace("/estructura_investigacion_trabajo_grado",
			beego.NSInclude(
				&controllers.EstructuraInvestigacionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/areas_docente",
			beego.NSInclude(
				&controllers.AreasDocenteController{},
			),
		),

		beego.NSNamespace("/area_conocimiento",
			beego.NSInclude(
				&controllers.AreaConocimientoController{},
			),
		),

		beego.NSNamespace("/areas_trabajo_grado",
			beego.NSInclude(
				&controllers.AreasTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/documento",
			beego.NSInclude(
				&controllers.DocumentoController{},
			),
		),

		beego.NSNamespace("/documento_trabajo_grado",
			beego.NSInclude(
				&controllers.DocumentoTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/modalidad",
			beego.NSInclude(
				&controllers.ModalidadController{},
			),
		),

		beego.NSNamespace("/formato",
			beego.NSInclude(
				&controllers.FormatoController{},
			),
		),

		beego.NSNamespace("/rol_trabajo_grado",
			beego.NSInclude(
				&controllers.RolTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estado_espacio_academico_inscrito",
			beego.NSInclude(
				&controllers.EstadoEspacioAcademicoInscritoController{},
			),
		),

		beego.NSNamespace("/espacio_academico_inscrito",
			beego.NSInclude(
				&controllers.EspacioAcademicoInscritoController{},
			),
		),

		beego.NSNamespace("/vinculacion_trabajo_grado",
			beego.NSInclude(
				&controllers.VinculacionTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/formato_evaluacion_carrera",
			beego.NSInclude(
				&controllers.FormatoEvaluacionCarreraController{},
			),
		),

		beego.NSNamespace("/socializacion",
			beego.NSInclude(
				&controllers.SocializacionController{},
			),
		),

		beego.NSNamespace("/evaluacion",
			beego.NSInclude(
				&controllers.EvaluacionController{},
			),
		),

		beego.NSNamespace("/estado_trabajo_grado",
			beego.NSInclude(
				&controllers.EstadoTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/distincion",
			beego.NSInclude(
				&controllers.DistincionController{},
			),
		),

		beego.NSNamespace("/trabajo_grado",
			beego.NSInclude(
				&controllers.TrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estado_revision",
			beego.NSInclude(
				&controllers.EstadoRevisionController{},
			),
		),

		beego.NSNamespace("/revision",
			beego.NSInclude(
				&controllers.RevisionController{},
			),
		),

		beego.NSNamespace("/tipo_pregunta",
			beego.NSInclude(
				&controllers.TipoPreguntaController{},
			),
		),

		beego.NSNamespace("/pregunta_formato",
			beego.NSInclude(
				&controllers.PreguntaFormatoController{},
			),
		),

		beego.NSNamespace("/estado_estudiante_trabajo_grado",
			beego.NSInclude(
				&controllers.EstadoEstudianteTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/estudiante_trabajo_grado",
			beego.NSInclude(
				&controllers.EstudianteTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/tipo_contacto",
			beego.NSInclude(
				&controllers.TipoContactoController{},
			),
		),

		beego.NSNamespace("/contacto_entidad",
			beego.NSInclude(
				&controllers.ContactoEntidadController{},
			),
		),

		beego.NSNamespace("/tipo_identificacion",
			beego.NSInclude(
				&controllers.TipoIdentificacionController{},
			),
		),

		beego.NSNamespace("/estado_asignatura_trabajo_grado",
			beego.NSInclude(
				&controllers.EstadoAsignaturaTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/asignatura_trabajo_grado",
			beego.NSInclude(
				&controllers.AsignaturaTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/modalidad_tipo_solicitud",
			beego.NSInclude(
				&controllers.ModalidadTipoSolicitudController{},
			),
		),

		beego.NSNamespace("/solicitud_trabajo_grado",
			beego.NSInclude(
				&controllers.SolicitudTrabajoGradoController{},
			),
		),

		beego.NSNamespace("/entidad",
			beego.NSInclude(
				&controllers.EntidadController{},
			),
		),
		beego.NSNamespace("/tr_solicitud",
			beego.NSInclude(
				&controllers.TrSolicitudController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
