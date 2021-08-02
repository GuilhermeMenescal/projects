/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package mb;

import java.util.Date;
import javax.faces.bean.ManagedBean;
import model.TaEvento;
import dao.TaEventoDAO;
import dao.TbUsuarioDAO;
import dao.TdTipoEventoDAO;
import java.io.IOException;
import java.text.ParseException;
import java.util.List;
import model.TbUsuario;
import model.TdTipoEvento;
import mb.UsuarioMB;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.Objects;
import java.util.logging.Level;
import java.util.logging.Logger;
import javax.faces.bean.SessionScoped;
import javax.faces.context.ExternalContext;
import javax.faces.context.FacesContext;
/**
 *
 * @author guilh
 */
@ManagedBean
@SessionScoped
public class EventoMB {

    /**
     * @return the editEvento
     */
    public TaEvento getEditEvento() {
        return editEvento;
    }

    /**
     * @param editEvento the editEvento to set
     */
    public void setEditEvento(TaEvento editEvento) {
        this.editEvento = editEvento;
    }

    private TaEvento evento;
    private TaEvento editEvento;
    private String endEvento;
    private Date dtiIniEvento;
    private Date dtiFimEvento;
    private int minPessoasEvento;
    private int maxPessoasEvento;
    private boolean stsVisibilidade;
    private boolean stsConclusao;
    private TaEventoDAO dao;
    private TbUsuario user;
    private TdTipoEvento tipo;
    private UsuarioMB userMB;
    private String nmeTipoEvento;
    private TdTipoEventoDAO daoTipo;
    private List<TdTipoEvento> tipoLista;
    private int idtTipo;
    private String dataIni;
    private String dataFim;
    private String visibilidade;
    private java.sql.Date dIni;
    private java.sql.Date dFim;       
    private List<TaEvento> taEventos;   
    private TbUsuarioDAO userDAO;
    private List<TaEvento> ativoEventos; 
    private List<TaEvento> histEventos; 
    /**
     * Creates a new instance of EventoMB
     */
    public EventoMB() {
        evento = new TaEvento();
        dao = new TaEventoDAO();
        //this.userMB = new UsuarioMB();
        this.userMB = (UsuarioMB) FacesContext.getCurrentInstance().
				getExternalContext().getSessionMap().get("usuarioMB");
        tipo = new TdTipoEvento();
        daoTipo = new TdTipoEventoDAO();
        tipoLista = daoTipo.consultarTodos();
        userDAO = new TbUsuarioDAO();
        this.taEventos = todosEventos();
        ativoEventos = ativosEventosUsuario();
        histEventos = histEventosUsuario();
        
    }
    public void criarEvento(){
      //System.out.println(this.usuarioLogado.getNmeUsuario());
      //userDAO.consultarPorNme(this.usuarioLogado.getNmeUsuario());
      this.tipo = getDaoTipo().consultarPorIdt(idtTipo);
      this.dateConv();
      this.evento.setEndEvento(getEndEvento());
      this.evento.setDtiIniEvento(dIni);
      this.evento.setDtiFimEvento(dFim);
      this.evento.setMinPessoasEvento(getMinPessoasEvento());
      this.evento.setMaxPessoasEvento(getMaxPessoasEvento());
      this.evento.setStsVisibilidade(isStsVisibilidade());
      this.evento.setStsConclusao(false);
      this.evento.setTbUsuario(this.userMB.getUsuarioLogado());
      this.evento.setTdTipoEvento(tipo);
      dao.incluir(this.evento);
      ExternalContext context = FacesContext.getCurrentInstance().getExternalContext();
        try {
            context.redirect("PaginaInicial.xhtml");
        } catch (IOException ex) {
            Logger.getLogger(EventoMB.class.getName()).log(Level.SEVERE, null, ex);
        }
    }
    public void visiPub(){
    
        setStsVisibilidade(true);
    
    }
    
    public void visiPri(){
    
        setStsVisibilidade(false); 
        
    }
  
    
  public void dateConv(){
        SimpleDateFormat format = new SimpleDateFormat("dd/MM/yyyy");
        try {
            dtiIniEvento = format.parse(dataIni);
            this.dIni = new java.sql.Date(dtiIniEvento.getTime());
            SimpleDateFormat format1 = new SimpleDateFormat("dd/MM/yyyy");
            dtiFimEvento = format1.parse(dataFim);
            this.dFim = new java.sql.Date(dtiFimEvento.getTime());
        } catch (ParseException ex) {
            Logger.getLogger(EventoMB.class.getName()).log(Level.SEVERE, null, ex);
        }        
    }
  
  public List<TaEvento> histEventosUsuario(){
      this.setHistEventos(new ArrayList<>());
       List<TaEvento> eventos = new ArrayList<>();
       this.setHistEventos(dao.consultarTodos());
       
       for (TaEvento evento : this.getHistEventos()) {
          if(Objects.equals(evento.getTbUsuario().getIdtUsuario(), userMB.getUsuarioLogado().getIdtUsuario())){
              eventos.add(evento);
          }
      }
       System.out.println(eventos.toString());
       
       return eventos;
        
  }
  
  public List<TaEvento> ativosEventosUsuario(){
      this.setAtivoEventos(new ArrayList<>());
       List<TaEvento> eventos = new ArrayList<>();
       this.setAtivoEventos(dao.consultarTodos());
       
       for (TaEvento evento : this.getAtivoEventos()) {
          if(Objects.equals(evento.getTbUsuario().getIdtUsuario(), userMB.getUsuarioLogado().getIdtUsuario()) && !evento.isStsConclusao()){
              eventos.add(evento);
          }
      }
       System.out.println(eventos.toString());
      return eventos; 
  }
  
  public List<TaEvento> todosEventos(){
      this.setTaEventos(new ArrayList<>());
       List<TaEvento> eventos = new ArrayList<>();
       this.setTaEventos(dao.consultarTodos());
       
       for (TaEvento evento : this.getTaEventos()) {
          if(!evento.isStsConclusao() && evento.isStsVisibilidade()){
              eventos.add(evento);
          }
      }
       System.out.println(eventos.toString());
      return eventos; 
  }
  
  public void finalizarEvento (TaEvento evento){
      evento.setStsConclusao(true);
      dao.alterar(evento);
  }
  
   public void pgEditarEvento(TaEvento evento){
      this.editEvento = new TaEvento();
      evento = dao.consultarPorIdt(evento.getIdtEvento());
      this.editEvento.setEndEvento(evento.getEndEvento());
      this.setDataIni(evento.getDtiIniEvento().toLocaleString().substring(0, 10));
      this.setDataFim(evento.getDtiFimEvento().toLocaleString().substring(0, 10));      
      this.editEvento.setMinPessoasEvento(evento.getMinPessoasEvento());
      this.editEvento.setMaxPessoasEvento(evento.getMaxPessoasEvento());
      this.editEvento.setStsVisibilidade(evento.isStsVisibilidade());
      this.editEvento.setStsConclusao(false);
      this.editEvento.setTbUsuario(this.userMB.getUsuarioLogado());
      this.editEvento.setTdTipoEvento(evento.getTdTipoEvento());
      this.editEvento = evento;
       
      ExternalContext context = FacesContext.getCurrentInstance().getExternalContext();
      try {
            context.redirect("EditarEvento.xhtml");
        } catch (IOException ex) {
            Logger.getLogger(EventoMB.class.getName()).log(Level.SEVERE, null, ex);
        }
      
  }
    public void EditarEvento (){
      this.dateConv();
      this.editEvento.setDtiIniEvento(dIni);
      this.editEvento.setDtiFimEvento(dFim);
      this.editEvento.setStsVisibilidade(isStsVisibilidade());
      this.editEvento.setStsConclusao(false);
      this.editEvento.setTbUsuario(this.userMB.getUsuarioLogado());
      
      dao.alterar(this.editEvento);
      ExternalContext context = FacesContext.getCurrentInstance().getExternalContext();
      try {
            context.redirect("Eventos_Ativos.xhtml");
        } catch (IOException ex) {
            Logger.getLogger(EventoMB.class.getName()).log(Level.SEVERE, null, ex);
        }
      
  }
    /**
     * @return the evento
     */
    public TaEvento getEvento() {
        return evento;
    }

    /**
     * @param evento the evento to set
     */
    public void setEvento(TaEvento evento) {
        this.evento = evento;
    }
    /**
     * @return the endEvento
     */
    public String getEndEvento() {
        return endEvento;
    }

    /**
     * @param endEvento the endEvento to set
     */
    public void setEndEvento(String endEvento) {
        this.endEvento = endEvento;
    }

    /**
     * @return the dtiIniEvento
     */
    public Date getDtiIniEvento() {
        return dtiIniEvento;
    }

    /**
     * @param dtiIniEvento the dtiIniEvento to set
     */
    public void setDtiIniEvento(Date dtiIniEvento) {
        this.dtiIniEvento = dtiIniEvento;
    }

    /**
     * @return the dtiFimEvento
     */
    public Date getDtiFimEvento() {
        return dtiFimEvento;
    }

    /**
     * @param dtiFimEvento the dtiFimEvento to set
     */
    public void setDtiFimEvento(Date dtiFimEvento) {
        this.dtiFimEvento = dtiFimEvento;
    }

    /**
     * @return the minPessoasEvento
     */
    public int getMinPessoasEvento() {
        return minPessoasEvento;
    }

    /**
     * @param minPessoasEvento the minPessoasEvento to set
     */
    public void setMinPessoasEvento(int minPessoasEvento) {
        this.minPessoasEvento = minPessoasEvento;
    }

    /**
     * @return the maxPessoasEvento
     */
    public int getMaxPessoasEvento() {
        return maxPessoasEvento;
    }

    /**
     * @param maxPessoasEvento the maxPessoasEvento to set
     */
    public void setMaxPessoasEvento(int maxPessoasEvento) {
        this.maxPessoasEvento = maxPessoasEvento;
    }

    /**
     * @return the stsVisibilidade
     */
    public boolean isStsVisibilidade() {
        return stsVisibilidade;
    }

    /**
     * @param stsVisibilidade the stsVisibilidade to set
     */
    public void setStsVisibilidade(boolean stsVisibilidade) {
        this.stsVisibilidade = stsVisibilidade;
    }

    /**
     * @return the stsConclusao
     */
    public boolean isStsConclusao() {
        return stsConclusao;
    }

    /**
     * @param stsConclusao the stsConclusao to set
     */
    public void setStsConclusao(boolean stsConclusao) {
        this.stsConclusao = stsConclusao;
    }

    /**
     * @return the usuario
     */
//    public TbUsuario getUsuario() {
//        return usuario;
//    }
//
//    /**
//     * @param usuario the usuario to set
//     */
//    public void setUsuario(TbUsuario usuario) {
//        this.usuario = usuario;
//    }

    /**
     * @return the tipo
     */
    public TdTipoEvento getTipo() {
        return tipo;
    }

    /**
     * @param tipo the tipo to set
     */
    public void setTipo(TdTipoEvento tipo) {
        this.tipo = tipo;
    }

    /**
     * @return the tipoLista
     */
    public List<TdTipoEvento> getTipoLista() {
        return tipoLista;
    }

    /**
     * @param tipoLista the tipoLista to set
     */
    public void setTipoLista(List<TdTipoEvento> tipoLista) {
        this.tipoLista = tipoLista;
    }

    /**
     * @return the nmeTipoEvento
     */
    public String getNmeTipoEvento() {
        return nmeTipoEvento;
    }

    /**
     * @param nmeTipoEvento the nmeTipoEvento to set
     */
    public void setNmeTipoEvento(String nmeTipoEvento) {
        this.nmeTipoEvento = nmeTipoEvento;
    }

    /**
     * @return the idtTipo
     */
    public int getIdtTipo() {
        return idtTipo;
    }

    /**
     * @param idtTipo the idtTipo to set
     */
    public void setIdtTipo(int idtTipo) {
        this.idtTipo = idtTipo;
    }

    /**
     * @return the dataIni
     */
    public String getDataIni() {
        return dataIni;
    }

    /**
     * @param dataIni the dataIni to set
     */
    public void setDataIni(String dataIni) {
        this.dataIni = dataIni;
    }

    /**
     * @return the dataFim
     */
    public String getDataFim() {
        return dataFim;
    }

    /**
     * @param dataFim the dataFim to set
     */
    public void setDataFim(String dataFim) {
        this.dataFim = dataFim;
    }

    /**
     * @return the visibilidade
     */
    public String getVisibilidade() {
        return visibilidade;
    }

    /**
     * @param visibilidade the visibilidade to set
     */
    public void setVisibilidade(String visibilidade) {
        this.visibilidade = visibilidade;
    }

    /**
     * @return the dIni
     */
    public java.sql.Date getdIni() {
        return dIni;
    }

    /**
     * @param dIni the dIni to set
     */
    public void setdIni(java.sql.Date dIni) {
        this.dIni = dIni;
    }

    /**
     * @return the dFim
     */
    public java.sql.Date getdFim() {
        return dFim;
    }

    /**
     * @param dFim the dFim to set
     */
    public void setdFim(java.sql.Date dFim) {
        this.dFim = dFim;
    }

    /**
     * @return the user
     */
    public TbUsuario getUser() {
        return user;
    }

    /**
     * @param user the user to set
     */
    public void setUser(TbUsuario user) {
        this.user = user;
    }

    /**
     * @return the daoTipo
     */
    public TdTipoEventoDAO getDaoTipo() {
        return daoTipo;
    }

    /**
     * @param daoTipo the daoTipo to set
     */
    public void setDaoTipo(TdTipoEventoDAO daoTipo) {
        this.daoTipo = daoTipo;
    }

    /**
     * @return the eventoLista
     */


    /**
     * @return the taEventos
     */
    public List<TaEvento> getTaEventos() {
        return taEventos;
    }

    /**
     * @param taEventos the taEventos to set
     */
    public void setTaEventos(List<TaEvento> taEventos) {
        this.taEventos = taEventos;
    }

    /**
     * @return the userDAO
     */
    public TbUsuarioDAO getUserDAO() {
        return userDAO;
    }

    /**
     * @param userDAO the userDAO to set
     */
    public void setUserDAO(TbUsuarioDAO userDAO) {
        this.userDAO = userDAO;
    }

    /**
     * @return the ativoEventos
     */
    public List<TaEvento> getAtivoEventos() {
        return ativoEventos;
    }

    /**
     * @param ativoEventos the ativoEventos to set
     */
    public void setAtivoEventos(List<TaEvento> ativoEventos) {
        this.ativoEventos = ativoEventos;
    }

    /**
     * @return the histEventos
     */
    public List<TaEvento> getHistEventos() {
        return histEventos;
    }

    /**
     * @param histEventos the histEventos to set
     */
    public void setHistEventos(List<TaEvento> histEventos) {
        this.histEventos = histEventos;
    }

}
