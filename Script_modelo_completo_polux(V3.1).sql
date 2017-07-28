-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.8.2-beta
-- PostgreSQL version: 9.5
-- Project Site: pgmodeler.com.br
-- Model Author: ---


-- Database creation must be done outside an multicommand file.
-- These commands were put in this file only for convenience.
-- -- object: new_database | type: DATABASE --
-- -- DROP DATABASE IF EXISTS new_database;
-- CREATE DATABASE new_database
-- ;
-- -- ddl-end --
-- 

-- object: polux | type: SCHEMA --
-- DROP SCHEMA IF EXISTS polux CASCADE;
CREATE SCHEMA polux;
-- ddl-end --
COMMENT ON SCHEMA polux IS 'schema de la base de datos de trabajos de grado';
-- ddl-end --

SET search_path TO pg_catalog,public,polux;
-- ddl-end --

-- object: polux.trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.trabajo_grado CASCADE;
CREATE TABLE polux.trabajo_grado(
	id serial NOT NULL,
	titulo text NOT NULL,
	modalidad integer NOT NULL,
	estado_trabajo_grado integer NOT NULL,
	distincion integer,
	CONSTRAINT pk_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.trabajo_grado IS 'Tabla sobre la cual se relacionan los ítems de un trabajo de grado con sus respectivas características';
-- ddl-end --
COMMENT ON COLUMN polux.trabajo_grado.id IS 'identificador de la tabla trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.trabajo_grado.titulo IS 'titulo del trabajo de grado a realizar';
-- ddl-end --
COMMENT ON CONSTRAINT pk_trabajo_grado ON polux.trabajo_grado  IS 'constraint de la primary key para la tabla trabajo_grado';
-- ddl-end --

-- object: polux.modalidad | type: TABLE --
-- DROP TABLE IF EXISTS polux.modalidad CASCADE;
CREATE TABLE polux.modalidad(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activa boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_modalidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.modalidad IS 'tabla que almacena las modalidades de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.modalidad.id IS 'identificador de la modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.modalidad.nombre IS 'nombre de la modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.modalidad.descripcion IS 'descripcion corta de la modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.modalidad.codigo_abreviacion IS 'Abreviacion del nombre de la modalidad';
-- ddl-end --
COMMENT ON COLUMN polux.modalidad.activa IS 'Permite saber si una modalidad de sencuentra activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_modalidad ON polux.modalidad  IS 'constraint de la primary key para modalidad';
-- ddl-end --

-- object: polux.documento_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.documento_trabajo_grado CASCADE;
CREATE TABLE polux.documento_trabajo_grado(
	id serial NOT NULL,
	trabajo_grado integer NOT NULL,
	documento integer NOT NULL,
	CONSTRAINT pk_documento_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.documento_trabajo_grado IS 'almacena la informacion de un documento digital pdf relacionado con un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.documento_trabajo_grado.id IS 'identificador documento_tg';
-- ddl-end --

-- object: polux.tipo_documento | type: TABLE --
-- DROP TABLE IF EXISTS polux.tipo_documento CASCADE;
CREATE TABLE polux.tipo_documento(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_tipo_documento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.tipo_documento IS 'tabla que almacena los diferentes tipos de documentos';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_documento.id IS 'identificador de la tabla tipo_documento';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_documento.nombre IS 'indica el nombre del tipo de un documento';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_documento.descripcion IS 'descripcion del tipo de documento';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_documento.codigo_abreviacion IS 'Abreviacion del nombre del documento';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_documento.activo IS 'Campo que permite identificar si un documento se encuentra activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_documento ON polux.tipo_documento  IS 'constraint de primary key para tipo documento';
-- ddl-end --

-- object: polux.revision | type: TABLE --
-- DROP TABLE IF EXISTS polux.revision CASCADE;
CREATE TABLE polux.revision(
	id serial NOT NULL,
	numero_revision numeric(2,0) NOT NULL,
	fecha_recepcion timestamp NOT NULL,
	fecha_revision timestamp,
	estado_revision integer NOT NULL,
	documento_trabajo_grado integer NOT NULL,
	vinculacion_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_revision PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.revision IS 'entidad que aloja el numero de revisiones hechas a un documento';
-- ddl-end --
COMMENT ON COLUMN polux.revision.id IS 'identificador para la tabla revision';
-- ddl-end --
COMMENT ON COLUMN polux.revision.numero_revision IS 'indica el numero de la revision realizada sobre un documento';
-- ddl-end --
COMMENT ON COLUMN polux.revision.fecha_recepcion IS 'campo que indica la fecha en que  la revisión es solicitada a realizarse';
-- ddl-end --
COMMENT ON COLUMN polux.revision.fecha_revision IS 'fecha de realizacion de una revision';
-- ddl-end --
COMMENT ON CONSTRAINT pk_revision ON polux.revision  IS 'constraint primary key revision';
-- ddl-end --

-- object: polux.correccion | type: TABLE --
-- DROP TABLE IF EXISTS polux.correccion CASCADE;
CREATE TABLE polux.correccion(
	id serial NOT NULL,
	observacion text NOT NULL,
	justificacion text,
	pagina numeric(4,0),
	revision integer NOT NULL,
	CONSTRAINT pk_correcion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.correccion IS 'entidad para almacenar las observaciones, correcciones o modificaciones sobre un documento';
-- ddl-end --
COMMENT ON COLUMN polux.correccion.id IS 'identificador de la tabla correccion';
-- ddl-end --
COMMENT ON COLUMN polux.correccion.observacion IS 'correccion que se realiza en una revision a un documento';
-- ddl-end --
COMMENT ON COLUMN polux.correccion.justificacion IS 'Campo que almacena la justificacion que se realiza de la correcion';
-- ddl-end --
COMMENT ON COLUMN polux.correccion.pagina IS 'Pagina del docuento en la que se debe hacer la correcion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_correcion ON polux.correccion  IS 'constraint primary key correccion';
-- ddl-end --

-- object: polux.comentario | type: TABLE --
-- DROP TABLE IF EXISTS polux.comentario CASCADE;
CREATE TABLE polux.comentario(
	id serial NOT NULL,
	comentario text NOT NULL,
	fecha timestamp NOT NULL,
	autor character varying(80) NOT NULL,
	correccion integer NOT NULL,
	CONSTRAINT pk_comentario PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.comentario IS 'comentarios que se realizan sobre una correccion de una revision de un documento';
-- ddl-end --
COMMENT ON COLUMN polux.comentario.id IS 'identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN polux.comentario.comentario IS 'campo en el que se registra el comentario de una persona sobre una corrección';
-- ddl-end --
COMMENT ON COLUMN polux.comentario.fecha IS 'fecha del comentario';
-- ddl-end --
COMMENT ON COLUMN polux.comentario.autor IS 'persona que realiza el comentario';
-- ddl-end --
COMMENT ON CONSTRAINT pk_comentario ON polux.comentario  IS 'constraint primary key comentario';
-- ddl-end --

-- object: polux.espacio_academico_inscrito | type: TABLE --
-- DROP TABLE IF EXISTS polux.espacio_academico_inscrito CASCADE;
CREATE TABLE polux.espacio_academico_inscrito(
	id serial NOT NULL,
	nota numeric(3,2),
	espacios_academicos_elegibles integer NOT NULL,
	estado_espacio_academico_inscrito integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_espacio_academico_inscrito PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.espacio_academico_inscrito IS 'entidad que guarda las asignaturas solicitadas por un estudiante y las asignadas';
-- ddl-end --
COMMENT ON COLUMN polux.espacio_academico_inscrito.id IS 'Identificador del espacio academico inscrito';
-- ddl-end --
COMMENT ON COLUMN polux.espacio_academico_inscrito.nota IS 'Nota de la asignatura vista';
-- ddl-end --
COMMENT ON CONSTRAINT pk_espacio_academico_inscrito ON polux.espacio_academico_inscrito  IS 'constraint primary key espacio_academico_inscrito';
-- ddl-end --

-- object: polux.area_conocimiento | type: TABLE --
-- DROP TABLE IF EXISTS polux.area_conocimiento CASCADE;
CREATE TABLE polux.area_conocimiento(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_area_conocimiento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.area_conocimiento IS 'entidad que almacena la informacion de areas de conocimientos ';
-- ddl-end --
COMMENT ON COLUMN polux.area_conocimiento.id IS 'identificador de area de conocimiento';
-- ddl-end --
COMMENT ON COLUMN polux.area_conocimiento.nombre IS 'nombre que identifica a un area de conocimiento';
-- ddl-end --
COMMENT ON COLUMN polux.area_conocimiento.descripcion IS 'descripcion de un area de conocimiento';
-- ddl-end --
COMMENT ON COLUMN polux.area_conocimiento.codigo_abreviacion IS 'Abreviacion del nombre del area de conocimiento';
-- ddl-end --
COMMENT ON COLUMN polux.area_conocimiento.activo IS 'Indica si el area de conocimiento se encuentra activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_area_conocimiento ON polux.area_conocimiento  IS 'constraint primary key area';
-- ddl-end --

-- object: polux.areas_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.areas_trabajo_grado CASCADE;
CREATE TABLE polux.areas_trabajo_grado(
	id serial NOT NULL,
	area_conocimiento integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_areas_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.areas_trabajo_grado IS 'tabla que relaciona las areas de conocimiento con un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.areas_trabajo_grado.id IS 'Identificador de la relacion entre un area de conocimiento y un area de trabajo';
-- ddl-end --

-- object: polux.areas_docente | type: TABLE --
-- DROP TABLE IF EXISTS polux.areas_docente CASCADE;
CREATE TABLE polux.areas_docente(
	id serial NOT NULL,
	identificacion_docente integer NOT NULL,
	area_conocimiento integer NOT NULL,
	CONSTRAINT pk_areas_docente PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.areas_docente IS 'entidad que almacena areas del conocimiento que maneja el docente';
-- ddl-end --
COMMENT ON COLUMN polux.areas_docente.id IS 'Identificador de la relacion entre el area y el docente';
-- ddl-end --
COMMENT ON COLUMN polux.areas_docente.identificacion_docente IS 'Identificacion del docente (academica: ACDOCENTE)';
-- ddl-end --
COMMENT ON CONSTRAINT pk_areas_docente ON polux.areas_docente  IS 'constraint primary key areas_docente';
-- ddl-end --

-- object: polux.vinculacion_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.vinculacion_trabajo_grado CASCADE;
CREATE TABLE polux.vinculacion_trabajo_grado(
	id serial NOT NULL,
	usuario integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	fecha_inicio date NOT NULL,
	fecha_fin date,
	rol_trabajo_grado integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_vinculacion_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.vinculacion_trabajo_grado IS 'Entidad que almacena los tipos de vinculación del docente con el trabajo de grado ';
-- ddl-end --
COMMENT ON COLUMN polux.vinculacion_trabajo_grado.id IS 'identificacion de la vinculacion docente';
-- ddl-end --
COMMENT ON COLUMN polux.vinculacion_trabajo_grado.usuario IS 'identificación del docente relacionado con el trabajo de grado (academica: ACDOCENTE)';
-- ddl-end --
COMMENT ON COLUMN polux.vinculacion_trabajo_grado.activo IS 'en caso de que se presente cambio de director';
-- ddl-end --
COMMENT ON COLUMN polux.vinculacion_trabajo_grado.fecha_inicio IS 'fecha de inicio de la vinculacion';
-- ddl-end --
COMMENT ON COLUMN polux.vinculacion_trabajo_grado.fecha_fin IS 'fecha de inicio de la vinculacion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_vinculacion_trabajo_grado ON polux.vinculacion_trabajo_grado  IS 'constraint de primary key';
-- ddl-end --

-- object: polux.formato | type: TABLE --
-- DROP TABLE IF EXISTS polux.formato CASCADE;
CREATE TABLE polux.formato(
	id serial NOT NULL,
	nombre character varying(200) NOT NULL,
	introduccion character varying(500),
	fecha_realizacion date NOT NULL,
	CONSTRAINT pk_formato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.formato IS 'entidad que almacena el formato de evaluacion';
-- ddl-end --
COMMENT ON COLUMN polux.formato.id IS 'identificador';
-- ddl-end --
COMMENT ON COLUMN polux.formato.nombre IS 'almacena el nombre del formato de evaluacion';
-- ddl-end --
COMMENT ON COLUMN polux.formato.introduccion IS 'introduccion del formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN polux.formato.fecha_realizacion IS 'Fecha de realizacion del formato de evaluacion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_formato ON polux.formato  IS 'constraint primary key formato';
-- ddl-end --

-- object: polux.pregunta | type: TABLE --
-- DROP TABLE IF EXISTS polux.pregunta CASCADE;
CREATE TABLE polux.pregunta(
	id serial NOT NULL,
	enunciado character varying(1000) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_pregunta PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.pregunta IS 'entidad que almacena las preguntas de una evaluación';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta.id IS 'identificador de la pregunta';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta.enunciado IS 'enunciado de la pregunta que tambien sirve como nombre';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta.descripcion IS 'breve descripcion de la funcionalidad de la pregunta';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta.codigo_abreviacion IS 'Abreviacion del nombre';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta.activo IS 'Permite saber si la pregunta se encuentra activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_pregunta ON polux.pregunta  IS 'constraint primary key pregunta';
-- ddl-end --

-- object: polux.respuesta | type: TABLE --
-- DROP TABLE IF EXISTS polux.respuesta CASCADE;
CREATE TABLE polux.respuesta(
	id serial NOT NULL,
	descripcion character varying(250),
	nombre character varying(1000) NOT NULL,
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_respuesta PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.respuesta IS 'entidad que almacena las respuestas para un formato';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta.id IS 'identificador de la respuesta';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta.descripcion IS 'breve descripción de la respuesta';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta.nombre IS 'Permite identificar a la respuesta';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta.codigo_abreviacion IS 'Abreviacion del nombre';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta.activo IS 'Permite identificar cuando una respuesta se encuentra activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_respuesta ON polux.respuesta  IS 'identificador respuesta';
-- ddl-end --

-- object: polux.socializacion | type: TABLE --
-- DROP TABLE IF EXISTS polux.socializacion CASCADE;
CREATE TABLE polux.socializacion(
	id serial NOT NULL,
	fecha timestamp NOT NULL,
	lugar character varying(80) NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_socializacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.socializacion IS 'entidad que almacena la sustentacion del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.socializacion.id IS 'identificador de la socialización';
-- ddl-end --
COMMENT ON COLUMN polux.socializacion.fecha IS 'fecha de realización de la socialización';
-- ddl-end --
COMMENT ON COLUMN polux.socializacion.lugar IS 'lugar del evento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_socializacion ON polux.socializacion  IS 'constraint primary key tabla socialización';
-- ddl-end --

-- object: polux.respuesta_evaluacion | type: TABLE --
-- DROP TABLE IF EXISTS polux.respuesta_evaluacion CASCADE;
CREATE TABLE polux.respuesta_evaluacion(
	id serial NOT NULL,
	justificacion character varying(500),
	respuesta_formato integer NOT NULL,
	evaluacion integer NOT NULL,
	CONSTRAINT pk_respuesta_evaluacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.respuesta_evaluacion IS 'entidad que almacena las relaciones entre la evaluación y el formato de las respuestas';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_evaluacion.id IS 'identificador de la tabla respuesta_evaluacion';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_evaluacion.justificacion IS 'campo que contiene la justificacion a una respuesta de evaluacion';
-- ddl-end --

-- object: polux.evaluacion | type: TABLE --
-- DROP TABLE IF EXISTS polux.evaluacion CASCADE;
CREATE TABLE polux.evaluacion(
	id serial NOT NULL,
	nota float NOT NULL,
	vinculacion_trabajo_grado integer NOT NULL,
	formato_evaluacion_carrera integer NOT NULL,
	socializacion integer,
	CONSTRAINT pk_evaluacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.evaluacion IS 'almacena la evaluación de la sustentación de un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.evaluacion.id IS 'identificador de la evaluación';
-- ddl-end --
COMMENT ON COLUMN polux.evaluacion.nota IS 'almacena la nota final de la socialización';
-- ddl-end --
COMMENT ON CONSTRAINT pk_evaluacion ON polux.evaluacion  IS 'constraint primary key evaluacion';
-- ddl-end --

-- object: polux.respuesta_formato | type: TABLE --
-- DROP TABLE IF EXISTS polux.respuesta_formato CASCADE;
CREATE TABLE polux.respuesta_formato(
	id serial NOT NULL,
	orden numeric(2,0),
	valoracion float,
	respuesta integer NOT NULL,
	pregunta_formato integer NOT NULL,
	CONSTRAINT pk_respuestas_formato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.respuesta_formato IS 'entidad que guarda las respuestas de un formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_formato.id IS 'identificador';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_formato.orden IS 'representa el orden en que se muestran las respuestas en un formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_formato.valoracion IS 'determina el valor de una respuesta dentro de una pregunta';
-- ddl-end --
COMMENT ON CONSTRAINT pk_respuestas_formato ON polux.respuesta_formato  IS 'constraint primary key respuestas_formato';
-- ddl-end --

-- object: polux.pregunta_formato | type: TABLE --
-- DROP TABLE IF EXISTS polux.pregunta_formato CASCADE;
CREATE TABLE polux.pregunta_formato(
	id serial NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	orden numeric(2,0),
	valoracion float,
	pregunta integer NOT NULL,
	formato integer NOT NULL,
	tipo_pregunta integer NOT NULL,
	CONSTRAINT pk_pregunta_formato PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.pregunta_formato IS 'entidad que almacena las preguntas del formato';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta_formato.id IS 'identificador formato de preguntas';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta_formato.activo IS 'indica si la pregunta esta activa o no';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta_formato.orden IS 'orden de la pregunta en el formato de evaluación';
-- ddl-end --
COMMENT ON COLUMN polux.pregunta_formato.valoracion IS 'determina el valor de una pregunta dentro del formato';
-- ddl-end --
COMMENT ON CONSTRAINT pk_pregunta_formato ON polux.pregunta_formato  IS 'constraint primary key pregunta_formato';
-- ddl-end --

-- object: polux.entidad | type: TABLE --
-- DROP TABLE IF EXISTS polux.entidad CASCADE;
CREATE TABLE polux.entidad(
	id serial NOT NULL,
	activo boolean NOT NULL DEFAULT false,
	identificacion integer NOT NULL,
	entidad_relacionada integer,
	tipo_identificacion integer NOT NULL,
	CONSTRAINT pk_entidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.entidad IS 'se encarga de almacenar una entidad relacionada a una vinculacion externa';
-- ddl-end --
COMMENT ON COLUMN polux.entidad.id IS 'identificador de la tabla entidad';
-- ddl-end --
COMMENT ON COLUMN polux.entidad.activo IS 'estado para saber si se estan solicitando pasantes';
-- ddl-end --
COMMENT ON COLUMN polux.entidad.identificacion IS 'Codigo de la entidad ';
-- ddl-end --
COMMENT ON COLUMN polux.entidad.entidad_relacionada IS 'Entidad a la cual se puede relacionar otra entidad';
-- ddl-end --
COMMENT ON CONSTRAINT pk_entidad ON polux.entidad  IS 'constraint primary key entidad';
-- ddl-end --

-- object: polux.estudiante_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.estudiante_trabajo_grado CASCADE;
CREATE TABLE polux.estudiante_trabajo_grado(
	id serial NOT NULL,
	codigo_estudiante integer NOT NULL,
	trabajo_grado integer NOT NULL,
	estado_estudiante_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_estudiante_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.estudiante_trabajo_grado IS 'Tabla en la que se relaciona uno o varios estudiantes a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.estudiante_trabajo_grado.id IS 'id de la tabla estudiante_TG';
-- ddl-end --
COMMENT ON COLUMN polux.estudiante_trabajo_grado.codigo_estudiante IS 'código del estudiante relacionado con el trabajo de grado, (Academica: ACEST)';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estudiante_trabajo_grado ON polux.estudiante_trabajo_grado  IS 'constraint primary key';
-- ddl-end --

-- object: polux.estructura_investigacion_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.estructura_investigacion_trabajo_grado CASCADE;
CREATE TABLE polux.estructura_investigacion_trabajo_grado(
	id serial NOT NULL,
	codigo_estructura_investigacion integer NOT NULL,
	trabajo_grado integer NOT NULL,
	CONSTRAINT pk_estructura_investigacion_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.estructura_investigacion_trabajo_grado IS 'tabla que almacena las estructuras de investigacion (instituto, grupo o semillero) pertenecientes a un tg';
-- ddl-end --
COMMENT ON COLUMN polux.estructura_investigacion_trabajo_grado.id IS 'tabla que almacena las estructuras de investigacion (instituto, grupo o semillero) pertenecientes a un tg';
-- ddl-end --
COMMENT ON COLUMN polux.estructura_investigacion_trabajo_grado.codigo_estructura_investigacion IS 'código asociado a la estructura de investigación';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estructura_investigacion_trabajo_grado ON polux.estructura_investigacion_trabajo_grado  IS 'constraint primary key estructura_investigacion';
-- ddl-end --

-- object: polux.espacios_academicos_elegibles | type: TABLE --
-- DROP TABLE IF EXISTS polux.espacios_academicos_elegibles CASCADE;
CREATE TABLE polux.espacios_academicos_elegibles(
	id serial NOT NULL,
	codigo_asignatura integer NOT NULL,
	activo boolean NOT NULL DEFAULT false,
	carrera_elegible integer NOT NULL,
	CONSTRAINT pk_espacios_academicos_elegibles PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.espacios_academicos_elegibles IS 'Asignaturas elegibles para optar por la modalidad de espacios academicos de posgrado o profundizacion';
-- ddl-end --
COMMENT ON COLUMN polux.espacios_academicos_elegibles.id IS 'identificador de la tabla asignaturas_elegibles';
-- ddl-end --
COMMENT ON COLUMN polux.espacios_academicos_elegibles.codigo_asignatura IS 'código de la asignatura que puede ser elegible por el estudiante que opte por la modalidad de espacios académicos de posgrado o profundización';
-- ddl-end --
COMMENT ON COLUMN polux.espacios_academicos_elegibles.activo IS 'permite saber si la carrera se encuentra elegible aun o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_espacios_academicos_elegibles ON polux.espacios_academicos_elegibles  IS 'constraint primary key';
-- ddl-end --

-- object: polux.asignatura_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.asignatura_trabajo_grado CASCADE;
CREATE TABLE polux.asignatura_trabajo_grado(
	id serial NOT NULL,
	codigo_asignatura integer NOT NULL,
	periodo numeric NOT NULL,
	anio numeric(4) NOT NULL,
	calificacion float,
	trabajo_grado integer NOT NULL,
	estado_asignatura_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_asignatura_inscrita PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.asignatura_trabajo_grado IS 'guarda la asignatura inscrita para la modalidad de TG';
-- ddl-end --
COMMENT ON COLUMN polux.asignatura_trabajo_grado.id IS 'Identificador de la asignatura inscrita en el trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.asignatura_trabajo_grado.codigo_asignatura IS 'codigo de la asignatura (TGI, TGII, TG Tecnologico), (academica: ACASI)';
-- ddl-end --
COMMENT ON COLUMN polux.asignatura_trabajo_grado.periodo IS 'periodo en el cual se cursa la asignatura';
-- ddl-end --
COMMENT ON COLUMN polux.asignatura_trabajo_grado.anio IS 'Permite almaceñar el año en el cual se cursa la asignatura';
-- ddl-end --
COMMENT ON COLUMN polux.asignatura_trabajo_grado.calificacion IS 'Calificacion de la asignatura inscrita';
-- ddl-end --
COMMENT ON CONSTRAINT pk_asignatura_inscrita ON polux.asignatura_trabajo_grado  IS 'constraint primary key';
-- ddl-end --

-- object: polux.documento | type: TABLE --
-- DROP TABLE IF EXISTS polux.documento CASCADE;
CREATE TABLE polux.documento(
	id serial NOT NULL,
	titulo text NOT NULL,
	enlace character varying(100) NOT NULL,
	resumen text,
	tipo_documento integer NOT NULL,
	CONSTRAINT pk_documento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.documento IS 'Tabla que almacena los documentos registrados en el sistema';
-- ddl-end --
COMMENT ON COLUMN polux.documento.id IS 'identificador del documento';
-- ddl-end --
COMMENT ON COLUMN polux.documento.titulo IS 'titulo del documento asociado a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.documento.enlace IS 'enlace del documento en el servidor';
-- ddl-end --
COMMENT ON COLUMN polux.documento.resumen IS 'resumen del trabajo de grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_documento ON polux.documento  IS 'constraint primary key';
-- ddl-end --

-- object: polux.documento_entidad | type: TABLE --
-- DROP TABLE IF EXISTS polux.documento_entidad CASCADE;
CREATE TABLE polux.documento_entidad(
	id serial NOT NULL,
	documento integer NOT NULL,
	entidad integer NOT NULL,
	CONSTRAINT pk_documento_entidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.documento_entidad IS 'tabla que almacena la informacion de la relacion entre una entidad y documento';
-- ddl-end --
COMMENT ON COLUMN polux.documento_entidad.id IS 'identificador de la tabla documento_entidad';
-- ddl-end --

-- object: polux.formato_evaluacion_carrera | type: TABLE --
-- DROP TABLE IF EXISTS polux.formato_evaluacion_carrera CASCADE;
CREATE TABLE polux.formato_evaluacion_carrera(
	id serial NOT NULL,
	activo boolean NOT NULL DEFAULT false,
	codigo_proyecto integer NOT NULL,
	fecha_inicio date NOT NULL,
	fecha_fin date,
	modalidad integer NOT NULL,
	formato integer NOT NULL,
	CONSTRAINT pk_formato_evaluacion_carrera PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.formato_evaluacion_carrera IS 'tabla que relaciona los registros del formato de evaluacion dependiendo del proyecto curricular';
-- ddl-end --
COMMENT ON COLUMN polux.formato_evaluacion_carrera.id IS 'Identificador de la tabla formato_evaluacion_carrera';
-- ddl-end --
COMMENT ON COLUMN polux.formato_evaluacion_carrera.activo IS 'estado(activo,inactivo)';
-- ddl-end --
COMMENT ON COLUMN polux.formato_evaluacion_carrera.codigo_proyecto IS 'codigo proyecto que viene del web service';
-- ddl-end --
COMMENT ON COLUMN polux.formato_evaluacion_carrera.fecha_inicio IS 'Fecha de inicio del formato  ';
-- ddl-end --
COMMENT ON COLUMN polux.formato_evaluacion_carrera.fecha_fin IS 'Fecha de fin delformato';
-- ddl-end --
COMMENT ON CONSTRAINT pk_formato_evaluacion_carrera ON polux.formato_evaluacion_carrera  IS 'constraint primary key';
-- ddl-end --

-- object: polux.rol_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.rol_trabajo_grado CASCADE;
CREATE TABLE polux.rol_trabajo_grado(
	id serial NOT NULL,
	nombre character varying(80) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_rol_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.rol_trabajo_grado IS 'Tabla que permite parametrizar el rol que cumple una persona u entidad dentro del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.rol_trabajo_grado.id IS 'identificadorde la tabla trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.rol_trabajo_grado.nombre IS 'Nombre del rol que describe la actividad que se realiza dentro del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.rol_trabajo_grado.descripcion IS 'descripcion de la labor que se realiza dentro del trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.rol_trabajo_grado.codigo_abreviacion IS 'Abreviacion del nombre';
-- ddl-end --
COMMENT ON COLUMN polux.rol_trabajo_grado.activo IS 'Permite determinar si el rol se encuentra activo o no en el momento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_rol_trabajo_grado ON polux.rol_trabajo_grado  IS 'constraint primary key rol_tranajo_grado';
-- ddl-end --

-- object: polux.estado_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.estado_trabajo_grado CASCADE;
CREATE TABLE polux.estado_trabajo_grado(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_estado_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.estado_trabajo_grado IS 'Tabla que permite parametrizar los diferentes estados o etapas en los que se puede encontrar un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_trabajo_grado.id IS 'Identificador de la tabla estado_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_trabajo_grado.nombre IS 'Permite identificar por medio de un nombre el estado ';
-- ddl-end --
COMMENT ON COLUMN polux.estado_trabajo_grado.descripcion IS 'breve descripcion donde se explica cuando un trabajo de grado se encuentra en este estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_trabajo_grado.codigo_abreviacion IS 'Abreviacion del nombre que sirve como codigo de identificacion';
-- ddl-end --
COMMENT ON COLUMN polux.estado_trabajo_grado.activo IS 'Permite identificar si el estado se encuentra activo o no';
-- ddl-end --

-- object: polux.carrera_elegible | type: TABLE --
-- DROP TABLE IF EXISTS polux.carrera_elegible CASCADE;
CREATE TABLE polux.carrera_elegible(
	id serial NOT NULL,
	codigo_carrera integer NOT NULL,
	cupos_excelencia numeric(2),
	cupos_adicionales numeric(2),
	periodo numeric NOT NULL,
	anio numeric(4,0) NOT NULL,
	codigo_pensum numeric(3,0) NOT NULL,
	CONSTRAINT pk_carrera_elegible PRIMARY KEY (id),
	CONSTRAINT uq_codigo_carrera_periodo_anio_codigo_pensum UNIQUE (codigo_carrera,periodo,anio,codigo_pensum)

);
-- ddl-end --
COMMENT ON TABLE polux.carrera_elegible IS 'carreras que son elegidas para cursar como modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.carrera_elegible.id IS 'identificador de la tabla carrera_elegible';
-- ddl-end --
COMMENT ON COLUMN polux.carrera_elegible.codigo_carrera IS 'campo en el que se indica la carrera que se puede cursar en la modalidad de trabajo de grado (ACPEN: PEN_CRA_COD)';
-- ddl-end --
COMMENT ON COLUMN polux.carrera_elegible.cupos_excelencia IS 'número de admitidos por excelencia académica y exentos de pago';
-- ddl-end --
COMMENT ON COLUMN polux.carrera_elegible.cupos_adicionales IS 'número de admitidos admitidos por condiciones económicas y calidades académicas, con pago';
-- ddl-end --
COMMENT ON COLUMN polux.carrera_elegible.periodo IS 'campo en el que se registra el periodo academico de una asignatura';
-- ddl-end --
COMMENT ON COLUMN polux.carrera_elegible.anio IS 'campo en el que se indica el año del periodo académico (ACASPERI: APE_ANO)';
-- ddl-end --
COMMENT ON COLUMN polux.carrera_elegible.codigo_pensum IS 'campo en el que se registra el código del pensum al cual pertenece la asignatura (ACPEN: PEN_NRO)';
-- ddl-end --
COMMENT ON CONSTRAINT pk_carrera_elegible ON polux.carrera_elegible  IS 'constraint primary key carrera elegible';
-- ddl-end --
COMMENT ON CONSTRAINT uq_codigo_carrera_periodo_anio_codigo_pensum ON polux.carrera_elegible  IS 'Constraint unique codigo_carrera_periodo_anio_codigo_pensum';
-- ddl-end --

-- object: fk_tipo_documento_documento | type: CONSTRAINT --
-- ALTER TABLE polux.documento DROP CONSTRAINT IF EXISTS fk_tipo_documento_documento CASCADE;
ALTER TABLE polux.documento ADD CONSTRAINT fk_tipo_documento_documento FOREIGN KEY (tipo_documento)
REFERENCES polux.tipo_documento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.respuesta_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.respuesta_solicitud CASCADE;
CREATE TABLE polux.respuesta_solicitud(
	id serial NOT NULL,
	fecha timestamp NOT NULL,
	justificacion text,
	ente_responsable integer,
	usuario integer,
	estado_solicitud integer NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_respuesta_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.respuesta_solicitud IS 'Tabla que permite identificar el estado y la respuesta dada a una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_solicitud.id IS 'Identificador de la tabla respuesta_solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_solicitud.fecha IS 'Fecha en la cual cambia el estado de una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_solicitud.justificacion IS 'almacena la justificación con la cual se hace un cambio al estado de una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_solicitud.ente_responsable IS 'Identificador que relaciona a la organización que se encarga de evaluar y dar respuesta a la solicitud, este campo se va a consumir de un servicio.';
-- ddl-end --
COMMENT ON COLUMN polux.respuesta_solicitud.usuario IS 'Usuario que se encarga de cambiar el estado de la solicitud y de informar lo decidido por la organización responsable';
-- ddl-end --
COMMENT ON CONSTRAINT pk_respuesta_solicitud ON polux.respuesta_solicitud  IS 'Constric primary key tabla respuesta_solicitud';
-- ddl-end --

-- object: polux.estado_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.estado_solicitud CASCADE;
CREATE TABLE polux.estado_solicitud(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_estado_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.estado_solicitud IS 'Tabla que permite parametrizar los diferentes estados en los que se puede encontrar una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.estado_solicitud.id IS 'identificador que diferencia cada uno de los estados en los que se puede encontrar una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.estado_solicitud.nombre IS 'Nombre del tipo de estado de la solicitud.';
-- ddl-end --
COMMENT ON COLUMN polux.estado_solicitud.descripcion IS 'Campo que permite describir los estados en los que se puede encontrar una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.estado_solicitud.codigo_abreviacion IS 'Codigo de abreviacion del nombre del tipo de estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_solicitud.activo IS 'Permite identificar si un tipo de estado esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_solicitud ON polux.estado_solicitud  IS 'Constric primary key tabla estado_solicitud';
-- ddl-end --

-- object: fk_estado_solicitud_respuesta_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_solicitud DROP CONSTRAINT IF EXISTS fk_estado_solicitud_respuesta_solicitud CASCADE;
ALTER TABLE polux.respuesta_solicitud ADD CONSTRAINT fk_estado_solicitud_respuesta_solicitud FOREIGN KEY (estado_solicitud)
REFERENCES polux.estado_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.solicitud_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.solicitud_trabajo_grado CASCADE;
CREATE TABLE polux.solicitud_trabajo_grado(
	id serial NOT NULL,
	fecha timestamp NOT NULL,
	modalidad_tipo_solicitud integer NOT NULL,
	trabajo_grado integer,
	CONSTRAINT pk_solicitud_trabajo_grado PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.solicitud_trabajo_grado IS 'Tabla que permite almacenar las diferentes solicitudes de trabajo de grado que se presenten en el sistema';
-- ddl-end --
COMMENT ON COLUMN polux.solicitud_trabajo_grado.id IS 'Identificador de la tabla solicitud_trabajo_grado';
-- ddl-end --
COMMENT ON COLUMN polux.solicitud_trabajo_grado.fecha IS 'Almacena la fecha en la cual se diligencia y presenta la solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_solicitud_trabajo_grado ON polux.solicitud_trabajo_grado  IS 'Constric primary key tabla solicitud_trabajo_grado';
-- ddl-end --

-- object: polux.documento_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.documento_solicitud CASCADE;
CREATE TABLE polux.documento_solicitud(
	id serial NOT NULL,
	documento integer NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_documento_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.documento_solicitud IS 'Tabla que relaciona un documento con la solicitud a la que pertenece';
-- ddl-end --
COMMENT ON COLUMN polux.documento_solicitud.id IS 'Identificador de la tabla documento_solicitud_trabajo_grado';
-- ddl-end --
COMMENT ON CONSTRAINT pk_documento_solicitud ON polux.documento_solicitud  IS 'Constric primary key tabla documento_solicitud';
-- ddl-end --

-- object: fk_documento_documento_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.documento_solicitud DROP CONSTRAINT IF EXISTS fk_documento_documento_solicitud CASCADE;
ALTER TABLE polux.documento_solicitud ADD CONSTRAINT fk_documento_documento_solicitud FOREIGN KEY (documento)
REFERENCES polux.documento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_solicitud_trabajo_grado_documento_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.documento_solicitud DROP CONSTRAINT IF EXISTS fk_solicitud_trabajo_grado_documento_solicitud CASCADE;
ALTER TABLE polux.documento_solicitud ADD CONSTRAINT fk_solicitud_trabajo_grado_documento_solicitud FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES polux.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_solicitud_trabajo_grado_respuesta_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_solicitud DROP CONSTRAINT IF EXISTS fk_solicitud_trabajo_grado_respuesta_solicitud CASCADE;
ALTER TABLE polux.respuesta_solicitud ADD CONSTRAINT fk_solicitud_trabajo_grado_respuesta_solicitud FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES polux.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.usuario_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.usuario_solicitud CASCADE;
CREATE TABLE polux.usuario_solicitud(
	id serial NOT NULL,
	usuario integer NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	CONSTRAINT pk_usuario_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.usuario_solicitud IS 'Tabla que permite relacionar una solicitud con el usuario que emite la solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.usuario_solicitud.id IS 'Identificador de la tabla usuario_solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.usuario_solicitud.usuario IS 'Código que identifica al usuario que realiza la solicitud.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_usuario_solicitud ON polux.usuario_solicitud  IS 'Constric primary key tabla usuario_solicitud';
-- ddl-end --

-- object: fk_solicitud_trabajo_grado_usuario_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.usuario_solicitud DROP CONSTRAINT IF EXISTS fk_solicitud_trabajo_grado_usuario_solicitud CASCADE;
ALTER TABLE polux.usuario_solicitud ADD CONSTRAINT fk_solicitud_trabajo_grado_usuario_solicitud FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES polux.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_solicitud_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.solicitud_trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_solicitud_trabajo_grado CASCADE;
ALTER TABLE polux.solicitud_trabajo_grado ADD CONSTRAINT fk_trabajo_grado_solicitud_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.tipo_solicitud CASCADE;
CREATE TABLE polux.tipo_solicitud(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_tipo_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.tipo_solicitud IS 'Tabla que almacena los tipos de solicitud que se puedan presentar dentro de cada una de las modalidades de trabajo de grado.';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_solicitud.id IS 'Identificador del tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_solicitud.nombre IS 'Nombre que identifica y deferencia una solicitud de otra';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_solicitud.descripcion IS 'Describe la función de la solicitud y cada una de las caracteristicas de esta asi como tambien a quien va dirigida';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_solicitud.codigo_abreviacion IS 'Abreviación del nombre que identifica un tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_solicitud.activo IS 'Permite definir si un tipo de solicitud se encuentra activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_solicitud ON polux.tipo_solicitud  IS 'Constric primary key tipo_solicitud';
-- ddl-end --

-- object: polux.modalidad_tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.modalidad_tipo_solicitud CASCADE;
CREATE TABLE polux.modalidad_tipo_solicitud(
	id serial NOT NULL,
	tipo_solicitud integer NOT NULL,
	modalidad integer NOT NULL,
	CONSTRAINT pk_modalidad_tipo_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.modalidad_tipo_solicitud IS 'Tabla que relaciona las modalidades de trabajo de grado con los tipos de solicitud.';
-- ddl-end --
COMMENT ON COLUMN polux.modalidad_tipo_solicitud.id IS 'Identificador de la tabla modalidad_tipo_solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_modalidad_tipo_solicitud ON polux.modalidad_tipo_solicitud  IS 'Constric primary key modalidad tipo solicitud';
-- ddl-end --

-- object: fk_tipo_solicitud_modalidad_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.modalidad_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_tipo_solicitud_modalidad_tipo_solicitud CASCADE;
ALTER TABLE polux.modalidad_tipo_solicitud ADD CONSTRAINT fk_tipo_solicitud_modalidad_tipo_solicitud FOREIGN KEY (tipo_solicitud)
REFERENCES polux.tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modalidad_modalidad_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.modalidad_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_modalidad_modalidad_tipo_solicitud CASCADE;
ALTER TABLE polux.modalidad_tipo_solicitud ADD CONSTRAINT fk_modalidad_modalidad_tipo_solicitud FOREIGN KEY (modalidad)
REFERENCES polux.modalidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.tipo_detalle | type: TABLE --
-- DROP TABLE IF EXISTS polux.tipo_detalle CASCADE;
CREATE TABLE polux.tipo_detalle(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	"codigo_abreviación" character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_tipo_detalle PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.tipo_detalle IS 'Tabla que permite almacenar los diferentes tipos de detalle que puede tener una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_detalle.id IS 'Identificador del tipo de detalle ';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_detalle.nombre IS 'Nombre que permite diferenciar a que tipo de detalle se esta refiriendo';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_detalle.descripcion IS 'Descripcion del tipo del detalle';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_detalle."codigo_abreviación" IS 'código de abreviación del nombre que identifica el tipo de detalle';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_detalle.activo IS 'campo que permite saber si un tipo de detalle se encuentra vigente o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_detalle ON polux.tipo_detalle  IS 'Constric primary key tabla tipo_detalle';
-- ddl-end --

-- object: polux.detalle | type: TABLE --
-- DROP TABLE IF EXISTS polux.detalle CASCADE;
CREATE TABLE polux.detalle(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	enunciado character varying(250) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	tipo_detalle integer NOT NULL,
	CONSTRAINT pk_detalle PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.detalle IS 'Tabla que permite almacenar los diferentes detalles de una solicitud, es decir cada uno de los campos que deberá diligenciar un usuario al momento de presentar una solicitud.';
-- ddl-end --
COMMENT ON COLUMN polux.detalle.id IS 'Identificador que permite diferenciar los detalles de una solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.detalle.nombre IS 'Nombre del detalle o campo que se debe diligenciar al realizar una solicitud de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.detalle.enunciado IS 'Enunciado con el cual se solicita al usuario diligenciar el detalle.';
-- ddl-end --
COMMENT ON COLUMN polux.detalle.descripcion IS 'Descripción breve que define la funciónalidad del detalle';
-- ddl-end --
COMMENT ON COLUMN polux.detalle.codigo_abreviacion IS 'Abreviacion del nombre del detalle';
-- ddl-end --
COMMENT ON COLUMN polux.detalle.activo IS 'Campo que permite identificar si el detalle se encuentra activo o no ';
-- ddl-end --
COMMENT ON CONSTRAINT pk_detalle ON polux.detalle  IS 'Constric primary key tabla detalle';
-- ddl-end --

-- object: fk_tipo_detalle_detalle | type: CONSTRAINT --
-- ALTER TABLE polux.detalle DROP CONSTRAINT IF EXISTS fk_tipo_detalle_detalle CASCADE;
ALTER TABLE polux.detalle ADD CONSTRAINT fk_tipo_detalle_detalle FOREIGN KEY (tipo_detalle)
REFERENCES polux.tipo_detalle (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.detalle_tipo_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.detalle_tipo_solicitud CASCADE;
CREATE TABLE polux.detalle_tipo_solicitud(
	id serial NOT NULL,
	detalle integer NOT NULL,
	modalidad_tipo_solicitud integer NOT NULL,
	activo boolean NOT NULL DEFAULT true,
	requerido boolean NOT NULL,
	CONSTRAINT pk_detalle_tipo_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.detalle_tipo_solicitud IS 'Tabla que relaciona cada uno de los detalles que puede tener una solicitud que se realiza dentro de una modalidad de trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.detalle_tipo_solicitud.id IS 'Identificador de la relacion entre el detalle la modalidad y el tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.detalle_tipo_solicitud.activo IS 'Campo que identifica si un detalle se encuentra en vigencia o es requerido para un tipo de solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.detalle_tipo_solicitud.requerido IS 'Indica si el campo es obligatorio para la solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_detalle_tipo_solicitud ON polux.detalle_tipo_solicitud  IS 'Constric primary key detalle_tipo_solicitud';
-- ddl-end --

-- object: fk_detalle_detalle_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.detalle_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_detalle_detalle_tipo_solicitud CASCADE;
ALTER TABLE polux.detalle_tipo_solicitud ADD CONSTRAINT fk_detalle_detalle_tipo_solicitud FOREIGN KEY (detalle)
REFERENCES polux.detalle (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modalidad_tipo_solicitud_detalle_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.detalle_tipo_solicitud DROP CONSTRAINT IF EXISTS fk_modalidad_tipo_solicitud_detalle_tipo_solicitud CASCADE;
ALTER TABLE polux.detalle_tipo_solicitud ADD CONSTRAINT fk_modalidad_tipo_solicitud_detalle_tipo_solicitud FOREIGN KEY (modalidad_tipo_solicitud)
REFERENCES polux.modalidad_tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.detalle_solicitud | type: TABLE --
-- DROP TABLE IF EXISTS polux.detalle_solicitud CASCADE;
CREATE TABLE polux.detalle_solicitud(
	id serial NOT NULL,
	descripcion character varying(250) NOT NULL,
	solicitud_trabajo_grado integer NOT NULL,
	detalle_tipo_solicitud integer NOT NULL,
	CONSTRAINT pk_detalle_solicitud PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.detalle_solicitud IS 'Tabla que almacena la respuesta que el usuario dio al detalle de una solicitud.';
-- ddl-end --
COMMENT ON COLUMN polux.detalle_solicitud.id IS 'Identificador de la respuesta que da el usuario al detalle requerido en la solicitud';
-- ddl-end --
COMMENT ON COLUMN polux.detalle_solicitud.descripcion IS 'Respuesta que da el usuario al detalle requerido en la solicitud';
-- ddl-end --
COMMENT ON CONSTRAINT pk_detalle_solicitud ON polux.detalle_solicitud  IS 'Constric primary key tabla detalle_solicitud';
-- ddl-end --

-- object: fk_solicitud_trabajo_grado_detalle_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.detalle_solicitud DROP CONSTRAINT IF EXISTS fk_solicitud_trabajo_grado_detalle_solicitud CASCADE;
ALTER TABLE polux.detalle_solicitud ADD CONSTRAINT fk_solicitud_trabajo_grado_detalle_solicitud FOREIGN KEY (solicitud_trabajo_grado)
REFERENCES polux.solicitud_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_detalle_tipo_solicitud_detalle_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.detalle_solicitud DROP CONSTRAINT IF EXISTS fk_detalle_tipo_solicitud_detalle_solicitud CASCADE;
ALTER TABLE polux.detalle_solicitud ADD CONSTRAINT fk_detalle_tipo_solicitud_detalle_solicitud FOREIGN KEY (detalle_tipo_solicitud)
REFERENCES polux.detalle_tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_documento_documento_entidad | type: CONSTRAINT --
-- ALTER TABLE polux.documento_entidad DROP CONSTRAINT IF EXISTS fk_documento_documento_entidad CASCADE;
ALTER TABLE polux.documento_entidad ADD CONSTRAINT fk_documento_documento_entidad FOREIGN KEY (documento)
REFERENCES polux.documento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_entidad | type: CONSTRAINT --
-- ALTER TABLE polux.documento_entidad DROP CONSTRAINT IF EXISTS fk_entidad CASCADE;
ALTER TABLE polux.documento_entidad ADD CONSTRAINT fk_entidad FOREIGN KEY (entidad)
REFERENCES polux.entidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_pregunta_pregunta_formato | type: CONSTRAINT --
-- ALTER TABLE polux.pregunta_formato DROP CONSTRAINT IF EXISTS fk_pregunta_pregunta_formato CASCADE;
ALTER TABLE polux.pregunta_formato ADD CONSTRAINT fk_pregunta_pregunta_formato FOREIGN KEY (pregunta)
REFERENCES polux.pregunta (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_formato_pregunta_formato | type: CONSTRAINT --
-- ALTER TABLE polux.pregunta_formato DROP CONSTRAINT IF EXISTS fk_formato_pregunta_formato CASCADE;
ALTER TABLE polux.pregunta_formato ADD CONSTRAINT fk_formato_pregunta_formato FOREIGN KEY (formato)
REFERENCES polux.formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_respuesta_respuesta_formato | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_formato DROP CONSTRAINT IF EXISTS fk_respuesta_respuesta_formato CASCADE;
ALTER TABLE polux.respuesta_formato ADD CONSTRAINT fk_respuesta_respuesta_formato FOREIGN KEY (respuesta)
REFERENCES polux.respuesta (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_pregunta_formato_respuesta_formato | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_formato DROP CONSTRAINT IF EXISTS fk_pregunta_formato_respuesta_formato CASCADE;
ALTER TABLE polux.respuesta_formato ADD CONSTRAINT fk_pregunta_formato_respuesta_formato FOREIGN KEY (pregunta_formato)
REFERENCES polux.pregunta_formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_respuesta_formato_respuesta_evaluacion | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_evaluacion DROP CONSTRAINT IF EXISTS fk_respuesta_formato_respuesta_evaluacion CASCADE;
ALTER TABLE polux.respuesta_evaluacion ADD CONSTRAINT fk_respuesta_formato_respuesta_evaluacion FOREIGN KEY (respuesta_formato)
REFERENCES polux.respuesta_formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_evaluacion_respuesta_evaluacion | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_evaluacion DROP CONSTRAINT IF EXISTS fk_evaluacion_respuesta_evaluacion CASCADE;
ALTER TABLE polux.respuesta_evaluacion ADD CONSTRAINT fk_evaluacion_respuesta_evaluacion FOREIGN KEY (evaluacion)
REFERENCES polux.evaluacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_correccion_comentario | type: CONSTRAINT --
-- ALTER TABLE polux.comentario DROP CONSTRAINT IF EXISTS fk_correccion_comentario CASCADE;
ALTER TABLE polux.comentario ADD CONSTRAINT fk_correccion_comentario FOREIGN KEY (correccion)
REFERENCES polux.correccion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_revision_correccion | type: CONSTRAINT --
-- ALTER TABLE polux.correccion DROP CONSTRAINT IF EXISTS fk_revision_correccion CASCADE;
ALTER TABLE polux.correccion ADD CONSTRAINT fk_revision_correccion FOREIGN KEY (revision)
REFERENCES polux.revision (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_vinculacion_trabajo_grado_revision | type: CONSTRAINT --
-- ALTER TABLE polux.revision DROP CONSTRAINT IF EXISTS fk_vinculacion_trabajo_grado_revision CASCADE;
ALTER TABLE polux.revision ADD CONSTRAINT fk_vinculacion_trabajo_grado_revision FOREIGN KEY (vinculacion_trabajo_grado)
REFERENCES polux.vinculacion_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_documento_trabajo_grado_revision | type: CONSTRAINT --
-- ALTER TABLE polux.revision DROP CONSTRAINT IF EXISTS fk_documento_trabajo_grado_revision CASCADE;
ALTER TABLE polux.revision ADD CONSTRAINT fk_documento_trabajo_grado_revision FOREIGN KEY (documento_trabajo_grado)
REFERENCES polux.documento_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_carrera_elegible_espacios_academicos_elegibles | type: CONSTRAINT --
-- ALTER TABLE polux.espacios_academicos_elegibles DROP CONSTRAINT IF EXISTS fk_carrera_elegible_espacios_academicos_elegibles CASCADE;
ALTER TABLE polux.espacios_academicos_elegibles ADD CONSTRAINT fk_carrera_elegible_espacios_academicos_elegibles FOREIGN KEY (carrera_elegible)
REFERENCES polux.carrera_elegible (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_espacios_academicos_elegibles_espacio_academico_inscrito | type: CONSTRAINT --
-- ALTER TABLE polux.espacio_academico_inscrito DROP CONSTRAINT IF EXISTS fk_espacios_academicos_elegibles_espacio_academico_inscrito CASCADE;
ALTER TABLE polux.espacio_academico_inscrito ADD CONSTRAINT fk_espacios_academicos_elegibles_espacio_academico_inscrito FOREIGN KEY (espacios_academicos_elegibles)
REFERENCES polux.espacios_academicos_elegibles (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_estructura_investigacion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.estructura_investigacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_estructura_investigacion_trabajo_grado CASCADE;
ALTER TABLE polux.estructura_investigacion_trabajo_grado ADD CONSTRAINT fk_trabajo_grado_estructura_investigacion_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_area_conocimiento_areas_docente | type: CONSTRAINT --
-- ALTER TABLE polux.areas_docente DROP CONSTRAINT IF EXISTS fk_area_conocimiento_areas_docente CASCADE;
ALTER TABLE polux.areas_docente ADD CONSTRAINT fk_area_conocimiento_areas_docente FOREIGN KEY (area_conocimiento)
REFERENCES polux.area_conocimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_area_conocimiento_areas_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.areas_trabajo_grado DROP CONSTRAINT IF EXISTS fk_area_conocimiento_areas_trabajo_grado CASCADE;
ALTER TABLE polux.areas_trabajo_grado ADD CONSTRAINT fk_area_conocimiento_areas_trabajo_grado FOREIGN KEY (area_conocimiento)
REFERENCES polux.area_conocimiento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_areas_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.areas_trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_areas_trabajo_grado CASCADE;
ALTER TABLE polux.areas_trabajo_grado ADD CONSTRAINT fk_trabajo_grado_areas_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_asignatura_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.asignatura_trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_asignatura_trabajo_grado CASCADE;
ALTER TABLE polux.asignatura_trabajo_grado ADD CONSTRAINT fk_trabajo_grado_asignatura_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modalidad_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.trabajo_grado DROP CONSTRAINT IF EXISTS fk_modalidad_trabajo_grado CASCADE;
ALTER TABLE polux.trabajo_grado ADD CONSTRAINT fk_modalidad_trabajo_grado FOREIGN KEY (modalidad)
REFERENCES polux.modalidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_documento_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.documento_trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_documento_trabajo_grado CASCADE;
ALTER TABLE polux.documento_trabajo_grado ADD CONSTRAINT fk_trabajo_grado_documento_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_documento_documento_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.documento_trabajo_grado DROP CONSTRAINT IF EXISTS fk_documento_documento_trabajo_grado CASCADE;
ALTER TABLE polux.documento_trabajo_grado ADD CONSTRAINT fk_documento_documento_trabajo_grado FOREIGN KEY (documento)
REFERENCES polux.documento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.socializacion DROP CONSTRAINT IF EXISTS fk_trabajo_grado CASCADE;
ALTER TABLE polux.socializacion ADD CONSTRAINT fk_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modalidad_formato_evaluacion_carrera | type: CONSTRAINT --
-- ALTER TABLE polux.formato_evaluacion_carrera DROP CONSTRAINT IF EXISTS fk_modalidad_formato_evaluacion_carrera CASCADE;
ALTER TABLE polux.formato_evaluacion_carrera ADD CONSTRAINT fk_modalidad_formato_evaluacion_carrera FOREIGN KEY (modalidad)
REFERENCES polux.modalidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_formato_formato_evaluacion_carrera | type: CONSTRAINT --
-- ALTER TABLE polux.formato_evaluacion_carrera DROP CONSTRAINT IF EXISTS fk_formato_formato_evaluacion_carrera CASCADE;
ALTER TABLE polux.formato_evaluacion_carrera ADD CONSTRAINT fk_formato_formato_evaluacion_carrera FOREIGN KEY (formato)
REFERENCES polux.formato (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_rol_trabajo_grado_vinculacion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.vinculacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_rol_trabajo_grado_vinculacion_trabajo_grado CASCADE;
ALTER TABLE polux.vinculacion_trabajo_grado ADD CONSTRAINT fk_rol_trabajo_grado_vinculacion_trabajo_grado FOREIGN KEY (rol_trabajo_grado)
REFERENCES polux.rol_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_vinculacion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.vinculacion_trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_vinculacion_trabajo_grado CASCADE;
ALTER TABLE polux.vinculacion_trabajo_grado ADD CONSTRAINT fk_trabajo_grado_vinculacion_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_vinculacion_trabajo_grado_evaluacion | type: CONSTRAINT --
-- ALTER TABLE polux.evaluacion DROP CONSTRAINT IF EXISTS fk_vinculacion_trabajo_grado_evaluacion CASCADE;
ALTER TABLE polux.evaluacion ADD CONSTRAINT fk_vinculacion_trabajo_grado_evaluacion FOREIGN KEY (vinculacion_trabajo_grado)
REFERENCES polux.vinculacion_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_formato_evaluacion_carrera_evaluacion | type: CONSTRAINT --
-- ALTER TABLE polux.evaluacion DROP CONSTRAINT IF EXISTS fk_formato_evaluacion_carrera_evaluacion CASCADE;
ALTER TABLE polux.evaluacion ADD CONSTRAINT fk_formato_evaluacion_carrera_evaluacion FOREIGN KEY (formato_evaluacion_carrera)
REFERENCES polux.formato_evaluacion_carrera (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_socializacion_evaluacion | type: CONSTRAINT --
-- ALTER TABLE polux.evaluacion DROP CONSTRAINT IF EXISTS fk_socializacion_evaluacion CASCADE;
ALTER TABLE polux.evaluacion ADD CONSTRAINT fk_socializacion_evaluacion FOREIGN KEY (socializacion)
REFERENCES polux.socializacion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: trabajo_grado_titulo | type: INDEX --
-- DROP INDEX IF EXISTS polux.trabajo_grado_titulo CASCADE;
CREATE INDEX trabajo_grado_titulo ON polux.trabajo_grado
	USING btree
	(
	  titulo ASC NULLS LAST
	);
-- ddl-end --
COMMENT ON INDEX polux.trabajo_grado_titulo IS 'indice que facilita las busquedas por titulos de trabajos de grado';
-- ddl-end --

-- object: polux.estado_espacio_academico_inscrito | type: TABLE --
-- DROP TABLE IF EXISTS polux.estado_espacio_academico_inscrito CASCADE;
CREATE TABLE polux.estado_espacio_academico_inscrito(
	id integer NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_estado_espacio_academico_inscrito PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.estado_espacio_academico_inscrito IS 'Tabla que parametriza los estados de los espacios academicos inscritos';
-- ddl-end --
COMMENT ON COLUMN polux.estado_espacio_academico_inscrito.id IS 'Identificador del estado del espacio academico inscrito';
-- ddl-end --
COMMENT ON COLUMN polux.estado_espacio_academico_inscrito.nombre IS 'Nombre del estado que permite identificarlo';
-- ddl-end --
COMMENT ON COLUMN polux.estado_espacio_academico_inscrito.descripcion IS 'Permite describir cuando una asignatura se encuentra en este estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_espacio_academico_inscrito.codigo_abreviacion IS 'Abreviacion del nombre del estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_espacio_academico_inscrito.activo IS 'Permite identificar si el estado esta activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_espacio_academico_inscrito ON polux.estado_espacio_academico_inscrito  IS 'constraint primary key estado espacio academico inscrito';
-- ddl-end --
ALTER TABLE polux.estado_espacio_academico_inscrito OWNER TO postgres;
-- ddl-end --

-- object: fk_estado_espacio_academico_inscrito_espacio_academico_i_4947 | type: CONSTRAINT --
-- ALTER TABLE polux.espacio_academico_inscrito DROP CONSTRAINT IF EXISTS fk_estado_espacio_academico_inscrito_espacio_academico_i_4947 CASCADE;
ALTER TABLE polux.espacio_academico_inscrito ADD CONSTRAINT fk_estado_espacio_academico_inscrito_espacio_academico_i_4947 FOREIGN KEY (estado_espacio_academico_inscrito)
REFERENCES polux.estado_espacio_academico_inscrito (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_espacio_academico_inscrito | type: CONSTRAINT --
-- ALTER TABLE polux.espacio_academico_inscrito DROP CONSTRAINT IF EXISTS fk_trabajo_grado_espacio_academico_inscrito CASCADE;
ALTER TABLE polux.espacio_academico_inscrito ADD CONSTRAINT fk_trabajo_grado_espacio_academico_inscrito FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_estado_trabajo_grado_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.trabajo_grado DROP CONSTRAINT IF EXISTS fk_estado_trabajo_grado_trabajo_grado CASCADE;
ALTER TABLE polux.trabajo_grado ADD CONSTRAINT fk_estado_trabajo_grado_trabajo_grado FOREIGN KEY (estado_trabajo_grado)
REFERENCES polux.estado_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_trabajo_grado_estudiante_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.estudiante_trabajo_grado DROP CONSTRAINT IF EXISTS fk_trabajo_grado_estudiante_trabajo_grado CASCADE;
ALTER TABLE polux.estudiante_trabajo_grado ADD CONSTRAINT fk_trabajo_grado_estudiante_trabajo_grado FOREIGN KEY (trabajo_grado)
REFERENCES polux.trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_solicitud_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_solicitud DROP CONSTRAINT IF EXISTS uq_solicitud_trabajo_grado CASCADE;
ALTER TABLE polux.respuesta_solicitud ADD CONSTRAINT uq_solicitud_trabajo_grado UNIQUE (solicitud_trabajo_grado);
-- ddl-end --

-- object: uq_tipo_solicitud_modalidad | type: CONSTRAINT --
-- ALTER TABLE polux.modalidad_tipo_solicitud DROP CONSTRAINT IF EXISTS uq_tipo_solicitud_modalidad CASCADE;
ALTER TABLE polux.modalidad_tipo_solicitud ADD CONSTRAINT uq_tipo_solicitud_modalidad UNIQUE (tipo_solicitud,modalidad);
-- ddl-end --

-- object: uq_detalle_modalidad_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.detalle_tipo_solicitud DROP CONSTRAINT IF EXISTS uq_detalle_modalidad_tipo_solicitud CASCADE;
ALTER TABLE polux.detalle_tipo_solicitud ADD CONSTRAINT uq_detalle_modalidad_tipo_solicitud UNIQUE (detalle,modalidad_tipo_solicitud);
-- ddl-end --

-- object: uq_solicitud_trabajo_grado_detalle_tipo_solicitud | type: CONSTRAINT --
-- ALTER TABLE polux.detalle_solicitud DROP CONSTRAINT IF EXISTS uq_solicitud_trabajo_grado_detalle_tipo_solicitud CASCADE;
ALTER TABLE polux.detalle_solicitud ADD CONSTRAINT uq_solicitud_trabajo_grado_detalle_tipo_solicitud UNIQUE (solicitud_trabajo_grado,detalle_tipo_solicitud);
-- ddl-end --

-- object: uq_documento_solicitud_tranajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.documento_solicitud DROP CONSTRAINT IF EXISTS uq_documento_solicitud_tranajo_grado CASCADE;
ALTER TABLE polux.documento_solicitud ADD CONSTRAINT uq_documento_solicitud_tranajo_grado UNIQUE (documento,solicitud_trabajo_grado);
-- ddl-end --

-- object: uq_trabajo_grado_documento | type: CONSTRAINT --
-- ALTER TABLE polux.documento_trabajo_grado DROP CONSTRAINT IF EXISTS uq_trabajo_grado_documento CASCADE;
ALTER TABLE polux.documento_trabajo_grado ADD CONSTRAINT uq_trabajo_grado_documento UNIQUE (trabajo_grado,documento);
-- ddl-end --

-- object: uq_documento_entidad | type: CONSTRAINT --
-- ALTER TABLE polux.documento_entidad DROP CONSTRAINT IF EXISTS uq_documento_entidad CASCADE;
ALTER TABLE polux.documento_entidad ADD CONSTRAINT uq_documento_entidad UNIQUE (documento,entidad);
-- ddl-end --

-- object: uq_trabajo_grado_estructura_investigacion | type: CONSTRAINT --
-- ALTER TABLE polux.estructura_investigacion_trabajo_grado DROP CONSTRAINT IF EXISTS uq_trabajo_grado_estructura_investigacion CASCADE;
ALTER TABLE polux.estructura_investigacion_trabajo_grado ADD CONSTRAINT uq_trabajo_grado_estructura_investigacion UNIQUE (codigo_estructura_investigacion,trabajo_grado);
-- ddl-end --

-- object: uq_trabajo_grado_area_conocimiento | type: CONSTRAINT --
-- ALTER TABLE polux.areas_trabajo_grado DROP CONSTRAINT IF EXISTS uq_trabajo_grado_area_conocimiento CASCADE;
ALTER TABLE polux.areas_trabajo_grado ADD CONSTRAINT uq_trabajo_grado_area_conocimiento UNIQUE (trabajo_grado,area_conocimiento);
-- ddl-end --

-- object: uq_areas_docente_identificacion_docente | type: CONSTRAINT --
-- ALTER TABLE polux.areas_docente DROP CONSTRAINT IF EXISTS uq_areas_docente_identificacion_docente CASCADE;
ALTER TABLE polux.areas_docente ADD CONSTRAINT uq_areas_docente_identificacion_docente UNIQUE (identificacion_docente,area_conocimiento);
-- ddl-end --

-- object: uq_trabajo_grado_codigo_estudiante | type: CONSTRAINT --
-- ALTER TABLE polux.estudiante_trabajo_grado DROP CONSTRAINT IF EXISTS uq_trabajo_grado_codigo_estudiante CASCADE;
ALTER TABLE polux.estudiante_trabajo_grado ADD CONSTRAINT uq_trabajo_grado_codigo_estudiante UNIQUE (codigo_estudiante,trabajo_grado);
-- ddl-end --

-- object: uq_espacios_academicos_elegibles_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.espacio_academico_inscrito DROP CONSTRAINT IF EXISTS uq_espacios_academicos_elegibles_trabajo_grado CASCADE;
ALTER TABLE polux.espacio_academico_inscrito ADD CONSTRAINT uq_espacios_academicos_elegibles_trabajo_grado UNIQUE (espacios_academicos_elegibles,trabajo_grado);
-- ddl-end --

-- object: uq_codigo_asignatura_carrera_elegible | type: CONSTRAINT --
-- ALTER TABLE polux.espacios_academicos_elegibles DROP CONSTRAINT IF EXISTS uq_codigo_asignatura_carrera_elegible CASCADE;
ALTER TABLE polux.espacios_academicos_elegibles ADD CONSTRAINT uq_codigo_asignatura_carrera_elegible UNIQUE (codigo_asignatura,carrera_elegible);
-- ddl-end --

-- object: uq_identificacion_docente_rol_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.vinculacion_trabajo_grado DROP CONSTRAINT IF EXISTS uq_identificacion_docente_rol_trabajo_grado CASCADE;
ALTER TABLE polux.vinculacion_trabajo_grado ADD CONSTRAINT uq_identificacion_docente_rol_trabajo_grado UNIQUE (usuario,rol_trabajo_grado,trabajo_grado);
-- ddl-end --

-- object: uq_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.socializacion DROP CONSTRAINT IF EXISTS uq_trabajo_grado CASCADE;
ALTER TABLE polux.socializacion ADD CONSTRAINT uq_trabajo_grado UNIQUE (trabajo_grado);
-- ddl-end --

-- object: uq_pregunta_formato | type: CONSTRAINT --
-- ALTER TABLE polux.pregunta_formato DROP CONSTRAINT IF EXISTS uq_pregunta_formato CASCADE;
ALTER TABLE polux.pregunta_formato ADD CONSTRAINT uq_pregunta_formato UNIQUE (pregunta,formato);
-- ddl-end --

-- object: uq_respuesta_formato_evaluacion | type: CONSTRAINT --
-- ALTER TABLE polux.respuesta_evaluacion DROP CONSTRAINT IF EXISTS uq_respuesta_formato_evaluacion CASCADE;
ALTER TABLE polux.respuesta_evaluacion ADD CONSTRAINT uq_respuesta_formato_evaluacion UNIQUE (respuesta_formato,evaluacion);
-- ddl-end --

-- object: polux.distincion | type: TABLE --
-- DROP TABLE IF EXISTS polux.distincion CASCADE;
CREATE TABLE polux.distincion(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion smallint,
	activo boolean NOT NULL,
	CONSTRAINT pk_distincion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.distincion IS 'Tabla que permite almacenar los diferentes tipos de distinciones que puede tener un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.distincion.id IS 'Identificador de la tabla distincion';
-- ddl-end --
COMMENT ON COLUMN polux.distincion.nombre IS 'Nombre de la disticion que recibe el trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.distincion.descripcion IS 'Campo que permite almacenar una pequeña descripcion de la disticion qeu se otorga';
-- ddl-end --
COMMENT ON COLUMN polux.distincion.codigo_abreviacion IS 'Abreviacion del nombrede la distincion';
-- ddl-end --
COMMENT ON COLUMN polux.distincion.activo IS 'Permite saber si la disticion esta activa o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_distincion ON polux.distincion  IS 'Constraint primary key distincion';
-- ddl-end --
ALTER TABLE polux.distincion OWNER TO postgres;
-- ddl-end --

-- object: fk_distincion_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.trabajo_grado DROP CONSTRAINT IF EXISTS fk_distincion_trabajo_grado CASCADE;
ALTER TABLE polux.trabajo_grado ADD CONSTRAINT fk_distincion_trabajo_grado FOREIGN KEY (distincion)
REFERENCES polux.distincion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.estado_revision | type: TABLE --
-- DROP TABLE IF EXISTS polux.estado_revision CASCADE;
CREATE TABLE polux.estado_revision(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_estado_revision PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.estado_revision IS 'Tabla que permite parametrizar los diferentes estados en los que se puede encontrar una revision';
-- ddl-end --
COMMENT ON COLUMN polux.estado_revision.id IS 'identificador que diferencia cada uno de los estados en los que se puede encontrar una revision';
-- ddl-end --
COMMENT ON COLUMN polux.estado_revision.nombre IS 'Nombre del tipo de estado de la revision';
-- ddl-end --
COMMENT ON COLUMN polux.estado_revision.descripcion IS 'Campo que permite describir los estados en los que se puede encontrar una revision';
-- ddl-end --
COMMENT ON COLUMN polux.estado_revision.codigo_abreviacion IS 'Codigo de abreviacion del nombre del tipo de estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_revision.activo IS 'Permite identificar si un tipo de estado esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_revision ON polux.estado_revision  IS 'Constric primary key tabla estado_revision';
-- ddl-end --

-- object: fk_estado_revision_revision | type: CONSTRAINT --
-- ALTER TABLE polux.revision DROP CONSTRAINT IF EXISTS fk_estado_revision_revision CASCADE;
ALTER TABLE polux.revision ADD CONSTRAINT fk_estado_revision_revision FOREIGN KEY (estado_revision)
REFERENCES polux.estado_revision (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.tipo_pregunta | type: TABLE --
-- DROP TABLE IF EXISTS polux.tipo_pregunta CASCADE;
CREATE TABLE polux.tipo_pregunta(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_tipo_regunta PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.tipo_pregunta IS 'Tabla que permite parametrizar los diferentes tipos de pregunta que se pueden encontrar en un formato';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_pregunta.id IS 'identificador que diferencia cada uno de los tipos de preguntas';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_pregunta.nombre IS 'Nombre del tipo de pregunta  ';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_pregunta.descripcion IS 'Campo que permite describir los tipos de preguntas que puede tener un formato';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_pregunta.codigo_abreviacion IS 'Codigo de abreviacion del nombre del tipo de pregunta';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_pregunta.activo IS 'Permite identificar si un tipo de pregunta esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_regunta ON polux.tipo_pregunta  IS 'Constraint primary key tabla tipo_pregunta';
-- ddl-end --

-- object: fk_tipo_pregunta_pregunta_formato | type: CONSTRAINT --
-- ALTER TABLE polux.pregunta_formato DROP CONSTRAINT IF EXISTS fk_tipo_pregunta_pregunta_formato CASCADE;
ALTER TABLE polux.pregunta_formato ADD CONSTRAINT fk_tipo_pregunta_pregunta_formato FOREIGN KEY (tipo_pregunta)
REFERENCES polux.tipo_pregunta (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.estado_estudiante_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.estado_estudiante_trabajo_grado CASCADE;
CREATE TABLE polux.estado_estudiante_trabajo_grado(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_estado_estudiante_trabajo_grado PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.estado_estudiante_trabajo_grado IS 'Tabla que permite parametrizar los diferentes estados en los que se puede encontrar un estudiante respecto a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_estudiante_trabajo_grado.id IS 'identificador que diferencia cada uno de los estados en los que se puede encontrar una revision';
-- ddl-end --
COMMENT ON COLUMN polux.estado_estudiante_trabajo_grado.nombre IS 'Nombre del estado en el que se puede encontrar un estudiante con respecto a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_estudiante_trabajo_grado.descripcion IS 'Campo que permite describir los estados en los que se puede encontrar un estudiante con respecto a un trabajo de grado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_estudiante_trabajo_grado.codigo_abreviacion IS 'Codigo de abreviacion del nombre del tipo de estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_estudiante_trabajo_grado.activo IS 'Permite identificar si un tipo de estado esta activo para ser usado o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_estudiante_trabajo_grado ON polux.estado_estudiante_trabajo_grado  IS 'Constraint primary keb tabla estado_estudiante_trabajo_grado';
-- ddl-end --

-- object: fk_estado_estudiante_trabajo_grado_estudiante_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.estudiante_trabajo_grado DROP CONSTRAINT IF EXISTS fk_estado_estudiante_trabajo_grado_estudiante_trabajo_grado CASCADE;
ALTER TABLE polux.estudiante_trabajo_grado ADD CONSTRAINT fk_estado_estudiante_trabajo_grado_estudiante_trabajo_grado FOREIGN KEY (estado_estudiante_trabajo_grado)
REFERENCES polux.estado_estudiante_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.tipo_contacto | type: TABLE --
-- DROP TABLE IF EXISTS polux.tipo_contacto CASCADE;
CREATE TABLE polux.tipo_contacto(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	"codigo_abreviación" character varying(20),
	activo boolean NOT NULL DEFAULT true,
	CONSTRAINT pk_tipo_contacto PRIMARY KEY (id)
	 WITH (FILLFACTOR = 100)

);
-- ddl-end --
COMMENT ON TABLE polux.tipo_contacto IS 'Tabla que permite almacenar los diferentes tipos de contacto que puede tener una entidad';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_contacto.id IS 'Identificador del tipo de contacto';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_contacto.nombre IS 'Nombre que permite diferenciar a que tipo de contacto al que se esta refiriendo';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_contacto.descripcion IS 'Descripcion del tipo de contacto';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_contacto."codigo_abreviación" IS 'código de abreviación del nombre que identifica el tipo de contacto';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_contacto.activo IS 'campo que permite saber si un tipo de detalle se encuentra vigente o no.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_contacto ON polux.tipo_contacto  IS 'Constric primary key tabla tipo_detalle';
-- ddl-end --

-- object: polux.contacto_entidad | type: TABLE --
-- DROP TABLE IF EXISTS polux.contacto_entidad CASCADE;
CREATE TABLE polux.contacto_entidad(
	id serial NOT NULL,
	detalle_contacto character varying(250) NOT NULL,
	tipo_contacto integer NOT NULL,
	entidad integer NOT NULL,
	CONSTRAINT pk_contacto_entidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.contacto_entidad IS 'Tabla que permite almacenar los diferentes contactos de una entidad';
-- ddl-end --
COMMENT ON COLUMN polux.contacto_entidad.id IS 'Identificador de la tabla contacto_entidad';
-- ddl-end --
COMMENT ON COLUMN polux.contacto_entidad.detalle_contacto IS 'Campo que permite almacenar el dato del contacto de la entidad';
-- ddl-end --
COMMENT ON CONSTRAINT pk_contacto_entidad ON polux.contacto_entidad  IS 'Constraint primary key tabla contacto_entidad';
-- ddl-end --
ALTER TABLE polux.contacto_entidad OWNER TO postgres;
-- ddl-end --

-- object: fk_tipo_contacto_contacto_entidad | type: CONSTRAINT --
-- ALTER TABLE polux.contacto_entidad DROP CONSTRAINT IF EXISTS fk_tipo_contacto_contacto_entidad CASCADE;
ALTER TABLE polux.contacto_entidad ADD CONSTRAINT fk_tipo_contacto_contacto_entidad FOREIGN KEY (tipo_contacto)
REFERENCES polux.tipo_contacto (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_entidad_contacto_entidad | type: CONSTRAINT --
-- ALTER TABLE polux.contacto_entidad DROP CONSTRAINT IF EXISTS fk_entidad_contacto_entidad CASCADE;
ALTER TABLE polux.contacto_entidad ADD CONSTRAINT fk_entidad_contacto_entidad FOREIGN KEY (entidad)
REFERENCES polux.entidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_entidad_tipo_contacto | type: CONSTRAINT --
-- ALTER TABLE polux.contacto_entidad DROP CONSTRAINT IF EXISTS uq_entidad_tipo_contacto CASCADE;
ALTER TABLE polux.contacto_entidad ADD CONSTRAINT uq_entidad_tipo_contacto UNIQUE (tipo_contacto,entidad);
-- ddl-end --

-- object: polux.tipo_identificacion | type: TABLE --
-- DROP TABLE IF EXISTS polux.tipo_identificacion CASCADE;
CREATE TABLE polux.tipo_identificacion(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_tipo_identificacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.tipo_identificacion IS 'Tabla que permite parametrizar los diferentes tipos de identificacion que puede tener una entidad';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_identificacion.id IS 'Identificador de la tabla tipo_identificacion';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_identificacion.nombre IS 'Permite identificar por medio de un nombre el tipo de identificacion';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_identificacion.descripcion IS 'breve descripcion del tipo del documento';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_identificacion.codigo_abreviacion IS 'Abreviacion del nombre que sirve como codigo de identificacion';
-- ddl-end --
COMMENT ON COLUMN polux.tipo_identificacion.activo IS 'Permite identificar si el estado se encuentra activo o no';
-- ddl-end --

-- object: fk_tipo_identificacion_entidad | type: CONSTRAINT --
-- ALTER TABLE polux.entidad DROP CONSTRAINT IF EXISTS fk_tipo_identificacion_entidad CASCADE;
ALTER TABLE polux.entidad ADD CONSTRAINT fk_tipo_identificacion_entidad FOREIGN KEY (tipo_identificacion)
REFERENCES polux.tipo_identificacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: polux.estado_asignatura_trabajo_grado | type: TABLE --
-- DROP TABLE IF EXISTS polux.estado_asignatura_trabajo_grado CASCADE;
CREATE TABLE polux.estado_asignatura_trabajo_grado(
	id integer NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	CONSTRAINT pk_estado_asignatura_trabajo_grado PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE polux.estado_asignatura_trabajo_grado IS 'Tabla que parametriza los estados de las asignaturas de TG I y TG II';
-- ddl-end --
COMMENT ON COLUMN polux.estado_asignatura_trabajo_grado.id IS 'Identificador del estado de estados de las asignaturas de TG I y TG II';
-- ddl-end --
COMMENT ON COLUMN polux.estado_asignatura_trabajo_grado.nombre IS 'Nombre del estado que permite identificarlo';
-- ddl-end --
COMMENT ON COLUMN polux.estado_asignatura_trabajo_grado.descripcion IS 'Permite describir cuando una asignatura se encuentra en este estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_asignatura_trabajo_grado.codigo_abreviacion IS 'Abreviacion del nombre del estado';
-- ddl-end --
COMMENT ON COLUMN polux.estado_asignatura_trabajo_grado.activo IS 'Permite identificar si el estado esta activo o no';
-- ddl-end --
COMMENT ON CONSTRAINT pk_estado_asignatura_trabajo_grado ON polux.estado_asignatura_trabajo_grado  IS 'constraint primary key estado espacio academico inscrito';
-- ddl-end --
ALTER TABLE polux.estado_asignatura_trabajo_grado OWNER TO postgres;
-- ddl-end --

-- object: fk_estado_asignatura_trabajo_grado_asignatura_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.asignatura_trabajo_grado DROP CONSTRAINT IF EXISTS fk_estado_asignatura_trabajo_grado_asignatura_trabajo_grado CASCADE;
ALTER TABLE polux.asignatura_trabajo_grado ADD CONSTRAINT fk_estado_asignatura_trabajo_grado_asignatura_trabajo_grado FOREIGN KEY (estado_asignatura_trabajo_grado)
REFERENCES polux.estado_asignatura_trabajo_grado (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_modalidad_tipo_solicitud_solicitud_trabajo_grado | type: CONSTRAINT --
-- ALTER TABLE polux.solicitud_trabajo_grado DROP CONSTRAINT IF EXISTS fk_modalidad_tipo_solicitud_solicitud_trabajo_grado CASCADE;
ALTER TABLE polux.solicitud_trabajo_grado ADD CONSTRAINT fk_modalidad_tipo_solicitud_solicitud_trabajo_grado FOREIGN KEY (modalidad_tipo_solicitud)
REFERENCES polux.modalidad_tipo_solicitud (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_entidad_entidad | type: CONSTRAINT --
-- ALTER TABLE polux.entidad DROP CONSTRAINT IF EXISTS fk_entidad_entidad CASCADE;
ALTER TABLE polux.entidad ADD CONSTRAINT fk_entidad_entidad FOREIGN KEY (entidad_relacionada)
REFERENCES polux.entidad (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


