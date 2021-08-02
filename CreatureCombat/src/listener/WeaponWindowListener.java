package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.CreatureControl;
import model.Weapon;
import view.WeaponWindow;

public class WeaponWindowListener implements ActionListener {
	
	private CreatureControl cc;
	private WeaponWindow ww;
	
	public WeaponWindowListener(CreatureControl cc, WeaponWindow ww){
		
		this.cc = cc;
		this.ww = ww;
		
	}

	@Override
	public void actionPerformed(ActionEvent e) {
		
		Weapon w = this.cc.searchWeapon(this.ww.getOption());
		this.cc.addWeaponPlayer(w);
		this.ww.setVisible(false);
		

	}

}
