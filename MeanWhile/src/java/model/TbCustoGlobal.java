package model;
// Generated 06/06/2019 21:55:24 by Hibernate Tools 4.3.1


import java.math.BigDecimal;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

/**
 * TbCustoGlobal generated by hbm2java
 */
@Entity
@Table(name="tb_custo_global"
    ,catalog="meanwhile"
)
public class TbCustoGlobal  implements java.io.Serializable {


     private int idtCustoGlobal;
     private TaEvento taEvento;
     private String nmeCustoGlobal;
     private BigDecimal vlrCustoGlobal;

    public TbCustoGlobal() {
    }

    public TbCustoGlobal(int idtCustoGlobal, TaEvento taEvento, String nmeCustoGlobal, BigDecimal vlrCustoGlobal) {
       this.idtCustoGlobal = idtCustoGlobal;
       this.taEvento = taEvento;
       this.nmeCustoGlobal = nmeCustoGlobal;
       this.vlrCustoGlobal = vlrCustoGlobal;
    }
   
     @Id 

    
    @Column(name="idt_custo_global", unique=true, nullable=false)
    public int getIdtCustoGlobal() {
        return this.idtCustoGlobal;
    }
    
    public void setIdtCustoGlobal(int idtCustoGlobal) {
        this.idtCustoGlobal = idtCustoGlobal;
    }

@ManyToOne(fetch=FetchType.LAZY)
    @JoinColumn(name="cod_evento", nullable=false)
    public TaEvento getTaEvento() {
        return this.taEvento;
    }
    
    public void setTaEvento(TaEvento taEvento) {
        this.taEvento = taEvento;
    }

    
    @Column(name="nme_custo_global", nullable=false, length=50)
    public String getNmeCustoGlobal() {
        return this.nmeCustoGlobal;
    }
    
    public void setNmeCustoGlobal(String nmeCustoGlobal) {
        this.nmeCustoGlobal = nmeCustoGlobal;
    }

    
    @Column(name="vlr_custo_global", nullable=false, precision=8)
    public BigDecimal getVlrCustoGlobal() {
        return this.vlrCustoGlobal;
    }
    
    public void setVlrCustoGlobal(BigDecimal vlrCustoGlobal) {
        this.vlrCustoGlobal = vlrCustoGlobal;
    }




}


