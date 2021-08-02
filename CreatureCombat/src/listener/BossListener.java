package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.PlayerControl;

public class BossListener implements ActionListener {
	
	private PlayerControl pc;
	
	public BossListener(PlayerControl pc){
		this.pc = pc;
	}

	@Override
	public void actionPerformed(ActionEvent e) {
		
		this.pc.bossFight();
	

	}

}
