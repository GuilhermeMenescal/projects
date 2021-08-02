package control;

import javax.swing.JOptionPane;

import database.SimulatedDatabase;
import model.Creature;
import model.Player;
import model.Weapon;
import view.ShowCreaturesWindow;
import view.WeaponWindow;

public class CreatureControl extends GenericControl {
	
	private SimulatedDatabase sdb;
	private Player pl;
	
	public CreatureControl(SimulatedDatabase sdb){
		this.sdb = sdb;
	}
	
	

	@Override
	public void registerItem() {
		
		
		
		
	}

	@Override
	public void removeItem(String name) {
		
		for(int i = 0; i < this.sdb.getCreatures().size(); i++){
		String n = this.sdb.getCreatures().get(i).getName();
		if(name == n){
			this.sdb.getCreatures().remove(i);
		}
		else{
			JOptionPane.showMessageDialog(null, "Creature not found");
		}
		}
		
		
		
	}

	@Override
	public void showList() {
		
		
	}
	
	public Creature searchCreature(int i){
		Creature c;
		return c = this.sdb.getCreatures().get(i);
		
	}
	
	public void showCreatures(Player p){
		this.pl = p;
		
		new ShowCreaturesWindow(this);
	}
	
	
	
	
	
	



	public void addCreaturePlayer(Creature c) {
		
		
		this.pl.setC(c);
		
		
		
		JOptionPane.showMessageDialog(null, pl.getC().getName() + " added!");
		
		
		
		
	}



	public void addWeapon() {
		
		new WeaponWindow(this);
		
	}
	
	public Weapon searchWeapon(int i){
		
		Weapon w;
		
	 return	w = this.sdb.getWeapons().get(i);
		
	}
	
	public void addWeaponPlayer(Weapon w){
		
		this.pl.getC().setW(w);
		
		JOptionPane.showMessageDialog(null, this.pl.getC().getW().getName() + " added!");
		
	}

	
	
	
	
	

}
