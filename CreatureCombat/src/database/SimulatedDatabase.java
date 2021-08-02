package database;

import java.util.ArrayList;

import model.Admin;
import model.Creature;
import model.Player;
import model.Weapon;
import util.FileManager;

public class SimulatedDatabase {
	

	private ArrayList<Creature> creatures;
	private ArrayList<Weapon> weapons;
	private ArrayList<Player> players;
	private ArrayList<Admin> admins;
	private FileManager fl;
	private final String adminFile = "admins.txt";
	private final String playerFile = "players.txt";
	private final String creatureFile = "creatures.txt";
	private final String weaponFile = "weapons.txt";
	
	public SimulatedDatabase(){
		
		
		this.fl = new FileManager();
		this.creatures = new ArrayList<Creature>();
		this.weapons = new ArrayList<Weapon>();
		this.players = new ArrayList<Player>();
		this.admins = new ArrayList<Admin>();
		this.populatePlayers();
		this.populateAdmins();
		this.populateCreatures();
		this.populateWeapons();
		//this.writeWeapons();
		//this.populateAdmins();
		//this.writeCreatures();
			
	}
	
	public ArrayList<Admin> getAdmins() {
		return admins;
	}

	public void setAdmins(ArrayList<Admin> admins) {
		this.admins = admins;
	}

	public ArrayList<Creature> getCreatures() {
		return creatures;
	}

	public void setCreatures(ArrayList<Creature> creatures) {
		this.creatures = creatures;
	}

	public ArrayList<Weapon> getWeapons() {
		return weapons;
	}

	public void setWeapons(ArrayList<Weapon> weapons) {
		this.weapons = weapons;
	}

	public ArrayList<Player> getPlayers() {
		return players;
	}

	public void setPlayers(ArrayList<Player> players) {
		this.players = players;
	}

	
	public void addPlayer(Player p){
		this.players.add(p);
		
	}

	public void addWeapon(Weapon w){
		this.weapons.add(w);
		
		
	}
	public void addCreature(Creature c){
		this.creatures.add(c);
		
	}
	public void addAdmin(Admin a){
		this.admins.add(a);
	}
	private void populatePlayers(){
		ArrayList<String> strings = this.fl.read(this.playerFile);
		for(String s :strings ){
		String[] pl = s.split(":");
		this.addPlayer(new Player(pl[0],pl[1],pl[2]));}
		
	}
		
		private void populateAdmins(){
			ArrayList<String> strings = this.fl.read(this.adminFile);
			for(String s :strings ){
			String[] al = s.split(":");
			this.addAdmin(new Admin(al[0],al[1],al[2]));
			
		}
		}
		
		private void populateCreatures(){
			ArrayList<String> strings = this.fl.read(this.creatureFile);
			for(String s :strings ){
			String[] cl = s.split("#");
			this.addCreature(new Creature(cl[0],Integer.parseInt(cl[1]),Integer.parseInt(cl[2])));}
			
		}
		
		private void populateWeapons(){
			ArrayList<String> strings = this.fl.read(this.weaponFile);
			for(String s :strings ){
			String[] wl = s.split(";");
			this.addWeapon(new Weapon(wl[0],Integer.parseInt(wl[1])));}
			
		}
		
	private void writeCreatures(){
		
		 Creature c1 = new Creature( "Caio", 9999, 9999);
		 this.addCreature(c1);
		 this.fl.write(c1.toFileManager(), this.creatureFile);
		 Creature c2 = new Creature("Bugs Bunny", 100, 15);
		 this.addCreature(c2);
		 this.fl.write(c2.toFileManager(), this.creatureFile);
		 
		 Creature c3 = new Creature("Harambe", 250, 100);
		 this.addCreature(c3);
		 this.fl.write(c3.toFileManager(), this.creatureFile);
		 
		 Creature c4 = new Creature("EuphoricTyrone", 150, 30);
		 this.addCreature(c4);
		 this.fl.write(c4.toFileManager(), this.creatureFile);
		 
		 Creature c5 = new Creature("DatBoi", 120, 40);
		 this.addCreature(c5);
		 this.fl.write(c5.toFileManager(), this.creatureFile);
		 
		 Creature c6 = new Creature("Putin", 200, 150);
		 this.addCreature(c6);
		 this.fl.write(c6.toFileManager(), this.creatureFile);
		 
		 Creature c7 = new Creature("Polishop Man", 100, 15);
		 this.addCreature(c7);
		 this.fl.write(c7.toFileManager(), this.creatureFile);
		 
		 Creature c8 = new Creature("Sesame street count", 75, 60);
		 this.addCreature(c8);
		 this.fl.write(c8.toFileManager(), this.creatureFile);
		 
		 Creature c9 = new Creature("Guilherme", 100, 30);
		 this.addCreature(c9);
		 this.fl.write(c9.toFileManager(), this.creatureFile);
		 
		 Creature c10 = new Creature("Leticia", 100, 30);
		 this.addCreature(c10);
		 this.fl.write(c10.toFileManager(), this.creatureFile);
		 
		 Creature c11 = new Creature("Doritos Pope", 100, 50 );
		 this.addCreature(c11);
		 this.fl.write(c11.toFileManager(), this.creatureFile);
		 
	
				 
	}
	private void writeWeapons(){
		Weapon w1 = new Weapon("Cebola", 9999);
		this.addWeapon(w1);
		this.fl.write(w1.toFileManager(), this.weaponFile);
		Weapon w2 = new Weapon("Carrot", 20);
		this.addWeapon(w2);
		this.fl.write(w2.toFileManager(), this.weaponFile);
		Weapon w3 = new Weapon("Child", 40);
		this.addWeapon(w3);
		this.fl.write(w3.toFileManager(), this.weaponFile);
		Weapon w4 = new Weapon("Guinsu", 200);
		this.addWeapon(w4);
		this.fl.write(w4.toFileManager(), this.weaponFile);
		Weapon w5 = new Weapon("Abaco", 10);
		this.addWeapon(w5);
		this.fl.write(w5.toFileManager(), this.weaponFile);
		Weapon w6 = new Weapon("Hands", 50);
		this.addWeapon(w6);
		this.fl.write(w6.toFileManager(), this.weaponFile);
		Weapon w7 = new Weapon("Oddjob's Hat", 200);
		this.addWeapon(w7);
		this.fl.write(w7.toFileManager(), this.weaponFile);
		Weapon w8 = new Weapon("Monocycle", 50);
		this.addWeapon(w8);
		this.fl.write(w8.toFileManager(), this.weaponFile);
		Weapon w9 = new Weapon("Doritos", 15);
		this.addWeapon(w9);
		this.fl.write(w9.toFileManager(), this.weaponFile);
		Weapon w10 = new Weapon("The power of friendship", 300);
		this.addWeapon(w10);
		this.fl.write(w10.toFileManager(), this.weaponFile);
		Weapon w11 = new Weapon("Definitely real candy", 300);
		this.addWeapon(w11);
		this.fl.write(w11.toFileManager(), this.weaponFile);
	}
		
	
		
	
//	private void populateAdmins(){
//		
//		Admin a1 = new Admin("Guilherme", "banana", "guilherme_menescal@hotmail.com");
//		
//		this.addAdmin(a1);
//		Admin a2 = new Admin("Letícia", "limão", "leticiahermanni@gmail.com");
//			
//		this.addAdmin(a2);
//		
//		for(Admin a : this.admins){
//			this.fl.write(a.toFileManager(), this.adminFile);}
//		}
}

	
	





