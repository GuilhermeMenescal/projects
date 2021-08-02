package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.CreatureControl;
import model.Creature;
import view.ShowCreaturesWindow;

public class PlayerCreatureListener implements ActionListener {
	
	private CreatureControl cc;
	private ShowCreaturesWindow scw;
	
	public PlayerCreatureListener(CreatureControl cc, ShowCreaturesWindow scw){
		this.cc = cc;
		this.scw = scw;
	}

	@Override
	public void actionPerformed(ActionEvent e) {
		
		Creature c = this.cc.searchCreature(this.scw.getTextField());
		this.cc.addCreaturePlayer(c);
		this.cc.addWeapon();
		this.scw.setVisible(false);
		

	}

}
