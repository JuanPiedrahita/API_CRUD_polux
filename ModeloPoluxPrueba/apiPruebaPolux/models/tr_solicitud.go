package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrSolicitud struct {
	Solicitud         *SolicitudTrabajoGrado
	DetallesSolicitud *[]DetalleSolicitud
	UsuariosSolicitud *[]UsuarioSolicitud
}

//funcion para la transaccion de solicitudes
func AddTransaccionSolicitud(m *TrSolicitud) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "Success")
	if id, err := o.Insert(m.Solicitud); err == nil {
		fmt.Println(id)
		for _, v := range *m.DetallesSolicitud {
			v.SolicitudTrabajoGrado.Id = int(id)
			if _, err = o.Insert(&v); err != nil {
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "Ocurrio un error al insertar los detalles de la solicitud.")
			}
		}

		for _, j := range *m.UsuariosSolicitud {
			j.SolicitudTrabajoGrado.Id = int(id)
			if _, err = o.Insert(&j); err != nil {
				err = o.Rollback()
				alerta[0] = "Error"
				alerta = append(alerta, "Ocurrio un error al insertar los usuarios de la solicitud.")
			}
		}
		err = o.Commit()
	} else {
		err = o.Rollback()
	}
	return
}
