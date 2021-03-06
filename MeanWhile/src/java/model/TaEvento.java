package model;
// Generated 06/06/2019 21:55:24 by Hibernate Tools 4.3.1


import java.util.Date;
import java.util.HashSet;
import java.util.Set;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.GeneratedValue;
import static javax.persistence.GenerationType.IDENTITY;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.OneToMany;
import javax.persistence.Table;
import javax.persistence.Temporal;
import javax.persistence.TemporalType;

/**
 * TaEvento generated by hbm2java
 */
@Entity
@Table(name="ta_evento"
    ,catalog="meanwhile"
)
public class TaEvento  implements java.io.Serializable {


     private Integer idtEvento;
     private TbUsuario tbUsuario;
     private TdTipoEvento tdTipoEvento;
     private String dscEvento;
     private String endEvento;
     private Date dtiIniEvento;
     private Date dtiFimEvento;
     private int minPessoasEvento;
     private int maxPessoasEvento;
     private boolean stsVisibilidade;
     private boolean stsConclusao;
     private Set<TaParticipacao> taParticipacaos = new HashSet<TaParticipacao>(0);
     private Set<TbOcorrencia> tbOcorrencias = new HashSet<TbOcorrencia>(0);
     private Set<TbMensagem> tbMensagems = new HashSet<TbMensagem>(0);
     private Set<TbCustoGlobal> tbCustoGlobals = new HashSet<TbCustoGlobal>(0);
     private Set<TbCustoPessoal> tbCustoPessoals = new HashSet<TbCustoPessoal>(0);

    public TaEvento() {
    }

	
    public TaEvento(TbUsuario tbUsuario, TdTipoEvento tdTipoEvento, String endEvento, Date dtiIniEvento, int minPessoasEvento, int maxPessoasEvento, boolean stsVisibilidade, boolean stsConclusao) {
        this.tbUsuario = tbUsuario;
        this.tdTipoEvento = tdTipoEvento;
        this.endEvento = endEvento;
        this.dtiIniEvento = dtiIniEvento;
        this.minPessoasEvento = minPessoasEvento;
        this.maxPessoasEvento = maxPessoasEvento;
        this.stsVisibilidade = stsVisibilidade;
        this.stsConclusao = stsConclusao;
    }
    public TaEvento(TbUsuario tbUsuario, TdTipoEvento tdTipoEvento, String dscEvento, String endEvento, Date dtiIniEvento, Date dtiFimEvento, int minPessoasEvento, int maxPessoasEvento, boolean stsVisibilidade, boolean stsConclusao, Set<TaParticipacao> taParticipacaos, Set<TbOcorrencia> tbOcorrencias, Set<TbMensagem> tbMensagems, Set<TbCustoGlobal> tbCustoGlobals, Set<TbCustoPessoal> tbCustoPessoals) {
       this.tbUsuario = tbUsuario;
       this.tdTipoEvento = tdTipoEvento;
       this.dscEvento = dscEvento;
       this.endEvento = endEvento;
       this.dtiIniEvento = dtiIniEvento;
       this.dtiFimEvento = dtiFimEvento;
       this.minPessoasEvento = minPessoasEvento;
       this.maxPessoasEvento = maxPessoasEvento;
       this.stsVisibilidade = stsVisibilidade;
       this.stsConclusao = stsConclusao;
       this.taParticipacaos = taParticipacaos;
       this.tbOcorrencias = tbOcorrencias;
       this.tbMensagems = tbMensagems;
       this.tbCustoGlobals = tbCustoGlobals;
       this.tbCustoPessoals = tbCustoPessoals;
    }
   
     @Id @GeneratedValue(strategy=IDENTITY)

    
    @Column(name="idt_evento", unique=true, nullable=false)
    public Integer getIdtEvento() {
        return this.idtEvento;
    }
    
    public void setIdtEvento(Integer idtEvento) {
        this.idtEvento = idtEvento;
    }

@ManyToOne(fetch=FetchType.LAZY)
    @JoinColumn(name="cod_usuario", nullable=false)
    public TbUsuario getTbUsuario() {
        return this.tbUsuario;
    }
    
    public void setTbUsuario(TbUsuario tbUsuario) {
        this.tbUsuario = tbUsuario;
    }

@ManyToOne(fetch=FetchType.LAZY)
    @JoinColumn(name="cod_tipo_evento", nullable=false)
    public TdTipoEvento getTdTipoEvento() {
        return this.tdTipoEvento;
    }
    
    public void setTdTipoEvento(TdTipoEvento tdTipoEvento) {
        this.tdTipoEvento = tdTipoEvento;
    }

    
    @Column(name="dsc_evento", length=65535)
    public String getDscEvento() {
        return this.dscEvento;
    }
    
    public void setDscEvento(String dscEvento) {
        this.dscEvento = dscEvento;
    }

    
    @Column(name="end_evento", nullable=false, length=65535)
    public String getEndEvento() {
        return this.endEvento;
    }
    
    public void setEndEvento(String endEvento) {
        this.endEvento = endEvento;
    }

    @Temporal(TemporalType.TIMESTAMP)
    @Column(name="dti_ini_evento", nullable=false, length=19)
    public Date getDtiIniEvento() {
        return this.dtiIniEvento;
    }
    
    public void setDtiIniEvento(Date dtiIniEvento) {
        this.dtiIniEvento = dtiIniEvento;
    }

    @Temporal(TemporalType.TIMESTAMP)
    @Column(name="dti_fim_evento", length=19)
    public Date getDtiFimEvento() {
        return this.dtiFimEvento;
    }
    
    public void setDtiFimEvento(Date dtiFimEvento) {
        this.dtiFimEvento = dtiFimEvento;
    }

    
    @Column(name="min_pessoas_evento", nullable=false)
    public int getMinPessoasEvento() {
        return this.minPessoasEvento;
    }
    
    public void setMinPessoasEvento(int minPessoasEvento) {
        this.minPessoasEvento = minPessoasEvento;
    }

    
    @Column(name="max_pessoas_evento", nullable=false)
    public int getMaxPessoasEvento() {
        return this.maxPessoasEvento;
    }
    
    public void setMaxPessoasEvento(int maxPessoasEvento) {
        this.maxPessoasEvento = maxPessoasEvento;
    }

    
    @Column(name="sts_visibilidade", nullable=false)
    public boolean isStsVisibilidade() {
        return this.stsVisibilidade;
    }
    
    public void setStsVisibilidade(boolean stsVisibilidade) {
        this.stsVisibilidade = stsVisibilidade;
    }

    
    @Column(name="sts_conclusao", nullable=false)
    public boolean isStsConclusao() {
        return this.stsConclusao;
    }
    
    public void setStsConclusao(boolean stsConclusao) {
        this.stsConclusao = stsConclusao;
    }

@OneToMany(fetch=FetchType.LAZY, mappedBy="taEvento")
    public Set<TaParticipacao> getTaParticipacaos() {
        return this.taParticipacaos;
    }
    
    public void setTaParticipacaos(Set<TaParticipacao> taParticipacaos) {
        this.taParticipacaos = taParticipacaos;
    }

@OneToMany(fetch=FetchType.LAZY, mappedBy="taEvento")
    public Set<TbOcorrencia> getTbOcorrencias() {
        return this.tbOcorrencias;
    }
    
    public void setTbOcorrencias(Set<TbOcorrencia> tbOcorrencias) {
        this.tbOcorrencias = tbOcorrencias;
    }

@OneToMany(fetch=FetchType.LAZY, mappedBy="taEvento")
    public Set<TbMensagem> getTbMensagems() {
        return this.tbMensagems;
    }
    
    public void setTbMensagems(Set<TbMensagem> tbMensagems) {
        this.tbMensagems = tbMensagems;
    }

@OneToMany(fetch=FetchType.LAZY, mappedBy="taEvento")
    public Set<TbCustoGlobal> getTbCustoGlobals() {
        return this.tbCustoGlobals;
    }
    
    public void setTbCustoGlobals(Set<TbCustoGlobal> tbCustoGlobals) {
        this.tbCustoGlobals = tbCustoGlobals;
    }

@OneToMany(fetch=FetchType.LAZY, mappedBy="taEvento")
    public Set<TbCustoPessoal> getTbCustoPessoals() {
        return this.tbCustoPessoals;
    }
    
    public void setTbCustoPessoals(Set<TbCustoPessoal> tbCustoPessoals) {
        this.tbCustoPessoals = tbCustoPessoals;
    }




}


