/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package mb;

import dao.TbUsuarioDAO;
import java.io.IOException;
import java.math.BigDecimal;
import java.util.List;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.faces.application.FacesMessage;
import javax.faces.bean.ManagedBean;
import javax.faces.bean.SessionScoped;
import javax.faces.context.ExternalContext;
import javax.faces.context.FacesContext;
import model.TbUsuario;

/**
 *
 * @author guilh
 */
@ManagedBean
@SessionScoped
public class UsuarioMB  {

    private TbUsuario usuario;
    private String nmeUsuario;
    private String email;
    private String celular;
    private String senha;
    private TbUsuario usuarioLogado; 
    private String erro;
    
    public UsuarioMB() {
        usuario = new TbUsuario();
        usuarioLogado = new TbUsuario();
        this.erro = null;
    }
    @SuppressWarnings("uncheked")
    public void cadastrarUsuario(){
        try {           
            
            this.usuario.setNmeUsuario(getNmeUsuario());
            this.usuario.setEmlUsuario(getEmail());
            this.usuario.setCelUsuario(getCelular());
            this.usuario.setPwdUsuario(getSenha());
            
            TbUsuarioDAO dao = new TbUsuarioDAO();
            dao.incluir(this.usuario);
            System.out.println("Usuário Incluido");
            
//        FacesMessage msg = new FacesMessage(FacesMessage.SEVERITY_INFO, "Resultado da Gravação", "Inclusão efetivada na base de dados.");
//        FacesContext.getCurrentInstance().addMessage(null, msg);
            ExternalContext context = FacesContext.getCurrentInstance().getExternalContext();
            context.redirect(context.getRequestContextPath());
            clean();
            
        } catch (IOException ex) {
            Logger.getLogger(UsuarioMB.class.getName()).log(Level.SEVERE, null, ex);
        }
    }
    
    public void login(){
        TbUsuarioDAO dao = new TbUsuarioDAO();
        List<TbUsuario> users = dao.consultarTodos();
        
        for (TbUsuario user : users) {
            if(user.getEmlUsuario().equals(this.email) && user.getPwdUsuario().equals(this.senha)){
                try {
                    this.usuarioLogado = user;
                    this.erro = null;
                    clean();
                    ExternalContext context = FacesContext.getCurrentInstance().getExternalContext();
                    context.redirect("PaginaInicial.xhtml");
                } catch (IOException ex) {
                    Logger.getLogger(UsuarioMB.class.getName()).log(Level.SEVERE, null, ex);
                }
            }
        }                  
                if(!this.email.equals("") || !this.senha.equals("")){
                    this.setErro("Dados Incorretos ou usuário inexistente");
                
                clean();
                ExternalContext context = FacesContext.getCurrentInstance().getExternalContext();
                    try {
                        context.redirect("Login.xhtml");
                    } catch (IOException ex) {
                        Logger.getLogger(UsuarioMB.class.getName()).log(Level.SEVERE, null, ex);
                    }
                }
    }
    
        private void clean(){
            this.setCelular("");
            this.setEmail("");
            this.setNmeUsuario("");
            this.setSenha("");
        }

    /**
     * @return the usuario
     */
    public TbUsuario getUsuario() {
        return usuario;
    }

    /**
     * @param usuario the usuario to set
     */
    public void setUsuario(TbUsuario usuario) {
        this.usuario = usuario;
    }

    /**
     * @return the nmeUsuario
     */
    public String getNmeUsuario() {
        return nmeUsuario;
    }

    /**
     * @param nmeUsuario the nmeUsuario to set
     */
    public void setNmeUsuario(String nmeUsuario) {
        this.nmeUsuario = nmeUsuario;
    }

    /**
     * @return the email
     */
    public String getEmail() {
        return email;
    }

    /**
     * @param email the email to set
     */
    public void setEmail(String email) {
        this.email = email;
    }

    /**
     * @return the celular
     */
    public String getCelular() {
        return celular;
    }

    /**
     * @param celular the celular to set
     */
    public void setCelular(String celular) {
        this.celular = celular;
    }

    /**
     * @return the senha
     */
    public String getSenha() {
        return senha;
    }

    /**
     * @param senha the senha to set
     */
    public void setSenha(String senha) {
        this.senha = senha;
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

    /**
     * @return the erro
     */
    public String getErro() {
        return erro;
    }

    /**
     * @param erro the erro to set
     */
    public void setErro(String erro) {
        this.erro = erro;
    }

    /**
     * @return the usr
     */
}
