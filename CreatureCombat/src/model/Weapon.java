package model;

public class Weapon {
	
	private String name;
	private int strenght;
	private int dps;
	
	public Weapon(String n, int strenght){
		this.name = n;
		this.strenght = strenght;
	}
	
	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getStrenght() {
		return strenght;
	}

	public void setStrenght(int strenght) {
		this.strenght = strenght;
	}

	public int getDps() {
		return dps;
	}

	public void setDps(int dps) {
		this.dps = dps;
	}

	public String toFileManager(){
		return this.name + ";" + this.strenght;
	}
	
	

}
