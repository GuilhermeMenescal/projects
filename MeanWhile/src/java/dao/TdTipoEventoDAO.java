/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package dao;

import java.util.List;
import model.TdTipoEvento;
import org.hibernate.Query;

/**
 *
 * @author letic
 */
public class TdTipoEventoDAO extends BaseDAO<TdTipoEvento>{
    public List<TdTipoEvento> consultarPorTipoEventoNme(Integer idtTipoEvento, String nme) {
        List<TdTipoEvento> lista;
        Query qy = hib.createQuery("SELECT obj FROM TdTipoEvento obj "
                + "WHERE (0=? OR tdTipoEvento.idtTipoEvento=?) AND nmeTipoEvento LIKE ? "
                + "ORDER BY tdTipoEvento.nmeTipoEvento");
        qy.setInteger(0, idtTipoEvento);
        qy.setInteger(1, idtTipoEvento);
        qy.setString(2, "%" + nme + "%");
        lista = qy.list();
        return lista;
    }
}
