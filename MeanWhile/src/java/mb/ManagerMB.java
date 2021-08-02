/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package mb;

import javax.faces.bean.ManagedBean;
import javax.faces.bean.ApplicationScoped;
import model.TbUsuario;

/**
 *
 * @author guilh
 */
@ManagedBean
@ApplicationScoped

public class ManagerMB {
    protected TbUsuario usuarioLogado;
    /**
     * Creates a new instance of ManagerMB
     */
    public ManagerMB() {
        usuarioLogado = new TbUsuario();
    }

    /**
     * @return the usuarioLogado
     */
    public TbUsuario getUsuarioLogado() {
        return usuarioLogado;
    }

    /**
     * @param usuarioLogado the usuarioLogado to set
     */
    public void setUsuarioLogado(TbUsuario usuarioLogado) {
        this.usuarioLogado = usuarioLogado;
    }
    
}
