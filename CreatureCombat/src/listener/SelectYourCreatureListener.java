package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.PlayerControl;

public class SelectYourCreatureListener implements ActionListener {
	
	private PlayerControl pc;
	
	public SelectYourCreatureListener(PlayerControl pc){
		
		this.pc = pc;
		
	}

	@Override
	public void actionPerformed(ActionEvent e) {
		
		this.pc.showCreatures();
		

	}

}
