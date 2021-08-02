package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.PlayerControl;

public class BattleListener implements ActionListener {
	
	private PlayerControl pc;
	
	public BattleListener(PlayerControl pc){
		this.pc = pc;
		
	}

	@Override
	public void actionPerformed(ActionEvent arg0) {
		
		this.pc.gimmeFight();
	

	}

}
