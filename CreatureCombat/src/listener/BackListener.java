package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.ManoloDeControle;
import view.PlayerLoginWindow;

public class BackListener implements ActionListener {
	private ManoloDeControle mc;
	private PlayerLoginWindow plw;
	
	public BackListener (ManoloDeControle mc, PlayerLoginWindow plw){
		this.mc = mc;
		this.plw = plw;
	}
	@Override
	public void actionPerformed(ActionEvent e) {
		this.plw.setVisible(false);
		this.plw.dispose();
		this.mc.start();

	}

}
