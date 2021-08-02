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
 * TaParticipacao generated by hbm2java
 */
@Entity
@Table(name="ta_participacao"
    ,catalog="meanwhile"
)
public class TaParticipacao  implements java.io.Serializable {


     private Integer idtParticipacao;
     private TaEvento taEvento;
     private TbUsuario tbUsuario;
     private Date dtiParticipacao;
     private boolean stsConfirmaParticipacao;
     private boolean stsLiberaContatoParticipacao;
     private boolean stsAvisarParticipacao;
     private boolean stsOcorrenciaParticipacao;
     private boolean stsAceitoParticipacao;
     private Set<TbMensagem> tbMensagems = new HashSet<TbMensagem>(0);
     private Set<TbOcorrencia> tbOcorrencias = new HashSet<TbOcorrencia>(0);

    public TaParticipacao() {
    }

	
    public TaParticipacao(TaEvento taEvento, TbUsuario tbUsuario, Date dtiParticipacao, boolean stsConfirmaParticipacao, boolean stsLiberaContatoParticipacao, boolean stsAvisarParticipacao, boolean stsOcorrenciaParticipacao, boolean stsAceitoParticipacao) {
        this.taEvento = taEvento;
        this.tbUsuario = tbUsuario;
        this.dtiParticipacao = dtiParticipacao;
        this.stsConfirmaParticipacao = stsConfirmaParticipacao;
        this.stsLiberaContatoParticipacao = stsLiberaContatoParticipacao;
        this.stsAvisarParticipacao = stsAvisarParticipacao;
        this.stsOcorrenciaParticipacao = stsOcorrenciaParticipacao;
        this.stsAceitoParticipacao = stsAceitoParticipacao;
    }
    public TaParticipacao(TaEvento taEvento, TbUsuario tbUsuario, Date dtiParticipacao, boolean stsConfirmaParticipacao, boolean stsLiberaContatoParticipacao, boolean stsAvisarParticipacao, boolean stsOcorrenciaParticipacao, boolean stsAceitoParticipacao, Set<TbMensagem> tbMensagems, Set<TbOcorrencia> tbOcorrencias) {
       this.taEvento = taEvento;
       this.tbUsuario = tbUsuario;
       this.dtiParticipacao = dtiParticipacao;
       this.stsConfirmaParticipacao = stsConfirmaParticipacao;
       this.stsLiberaContatoParticipacao = stsLiberaContatoParticipacao;
       this.stsAvisarParticipacao = stsAvisarParticipacao;
       this.stsOcorrenciaParticipacao = stsOcorrenciaParticipacao;
       this.stsAceitoParticipacao = stsAceitoParticipacao;
       this.tbMensagems = tbMensagems;
       this.tbOcorrencias = tbOcorrencias;
    }
   
     @Id @GeneratedValue(strategy=IDENTITY)

    
    @Column(name="idt_participacao", unique=true, nullable=false)
    public Integer getIdtParticipacao() {
        return this.idtParticipacao;
    }
    
    public void setIdtParticipacao(Integer idtParticipacao) {
        this.idtParticipacao = idtParticipacao;
    }

@ManyToOne(fetch=FetchType.LAZY)
    @JoinColumn(name="cod_evento", nullable=false)
    public TaEvento getTaEvento() {
        return this.taEvento;
    }
    
    public void setTaEvento(TaEvento taEvento) {
        this.taEvento = taEvento;
    }

@ManyToOne(fetch=FetchType.LAZY)
    @JoinColumn(name="cod_usuario", nullable=false)
    public TbUsuario getTbUsuario() {
        return this.tbUsuario;
    }
    
    public void setTbUsuario(TbUsuario tbUsuario) {
        this.tbUsuario = tbUsuario;
    }

    @Temporal(TemporalType.TIMESTAMP)
    @Column(name="dti_participacao", nullable=false, length=19)
    public Date getDtiParticipacao() {
        return this.dtiParticipacao;
    }
    
    public void setDtiParticipacao(Date dtiParticipacao) {
        this.dtiParticipacao = dtiParticipacao;
    }

    
    @Column(name="sts_confirma_participacao", nullable=false)
    public boolean isStsConfirmaParticipacao() {
        return this.stsConfirmaParticipacao;
    }
    
    public void setStsConfirmaParticipacao(boolean stsConfirmaParticipacao) {
        this.stsConfirmaParticipacao = stsConfirmaParticipacao;
    }

    
    @Column(name="sts_libera_contato_participacao", nullable=false)
    public boolean isStsLiberaContatoParticipacao() {
        return this.stsLiberaContatoParticipacao;
    }
    
    public void setStsLiberaContatoParticipacao(boolean stsLiberaContatoParticipacao) {
        this.stsLiberaContatoParticipacao = stsLiberaContatoParticipacao;
    }

    
    @Column(name="sts_avisar_participacao", nullable=false)
    public boolean isStsAvisarParticipacao() {
        return this.stsAvisarParticipacao;
    }
    
    public void setStsAvisarParticipacao(boolean stsAvisarParticipacao) {
        this.stsAvisarParticipacao = stsAvisarParticipacao;
    }

    
    @Column(name="sts_ocorrencia_participacao", nullable=false)
    public boolean isStsOcorrenciaParticipacao() {
        return this.stsOcorrenciaParticipacao;
    }
    
    public void setStsOcorrenciaParticipacao(boolean stsOcorrenciaParticipacao) {
        this.stsOcorrenciaParticipacao = stsOcorrenciaParticipacao;
    }

    
    @Column(name="sts_aceito_participacao", nullable=false)
    public boolean isStsAceitoParticipacao() {
        return this.stsAceitoParticipacao;
    }
    
    public void setStsAceitoParticipacao(boolean stsAceitoParticipacao) {
        this.stsAceitoParticipacao = stsAceitoParticipacao;
    }

@OneToMany(fetch=FetchType.LAZY, mappedBy="taParticipacao")
    public Set<TbMensagem> getTbMensagems() {
        return this.tbMensagems;
    }
    
    public void setTbMensagems(Set<TbMensagem> tbMensagems) {
        this.tbMensagems = tbMensagems;
    }

@OneToMany(fetch=FetchType.LAZY, mappedBy="taParticipacao")
    public Set<TbOcorrencia> getTbOcorrencias() {
        return this.tbOcorrencias;
    }
    
    public void setTbOcorrencias(Set<TbOcorrencia> tbOcorrencias) {
        this.tbOcorrencias = tbOcorrencias;
    }




}


