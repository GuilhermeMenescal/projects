package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.ManoloDeControle;
import view.LogInWindow;

public class RegisterListener implements ActionListener {
	
	private ManoloDeControle mc;
	private LogInWindow lw;
	
	public RegisterListener(ManoloDeControle mc, LogInWindow lw){
		this.mc = mc;
		this.lw = lw;
	}
	

	@Override
	public void actionPerformed(ActionEvent arg0) {
		this.lw.setVisible(false);
		this.mc.registerPlayer();
		

	}

}
