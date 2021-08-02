package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JOptionPane;

import control.PlayerControl;
import view.PlayerLoginWindow;

public class PlayerLoginListener implements ActionListener {
	private PlayerControl pc;
	private PlayerLoginWindow plw;
	
	public PlayerLoginListener (PlayerControl pc,PlayerLoginWindow plw){
		this.pc = pc;
		this.plw = plw;
	}
	@Override
	public void actionPerformed(ActionEvent e) {
		if (this.pc.isValidPlayer(this.plw.getName(), this.plw.getPasswordField())){
			JOptionPane.showMessageDialog(null, "You are logged " + this.pc.getCurrentPlayer());
			this.plw.setVisible(false);
			this.pc.GameWindow();
		}
		else{
			JOptionPane.showMessageDialog(null, "Invalid user!");
			this.plw.reset();
		}
		

	}

}
