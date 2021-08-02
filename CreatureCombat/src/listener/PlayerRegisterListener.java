package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JOptionPane;
import control.PlayerControl;
import view.RegisterWindow;

public class PlayerRegisterListener implements ActionListener {
	
	private PlayerControl pc;
	private RegisterWindow rw;
	
	public PlayerRegisterListener(PlayerControl pc, RegisterWindow rw){
		this.pc = pc;
		this.rw = rw;
	}
	

	@Override
	public void actionPerformed(ActionEvent arg0) {	
		if(this.rw.getEmail().contains("@")){
		this.pc.receivePlayer(this.rw.getName(),this.rw.getPassword(), this.rw.getEmail());
		this.rw.setVisible(false);
		}
		else{
			JOptionPane.showMessageDialog(null, "Invalid email!");
			this.rw.reset();
		}
		
		
		
		
		



	}

}
