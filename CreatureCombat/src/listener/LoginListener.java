package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.ManoloDeControle;
import view.LogInWindow;

public class LoginListener implements ActionListener {

	private ManoloDeControle mc;
	private LogInWindow lw;
	
	public LoginListener (ManoloDeControle mc, LogInWindow lw){
		this.mc = mc;
		this.lw = lw;
	}
	@Override
	public void actionPerformed(ActionEvent e) {
		this.mc.loginPlayer();
		this.lw.setVisible(true);

	}

}
