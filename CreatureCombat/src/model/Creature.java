package model;

import java.awt.image.BufferedImage;

public class Creature {
	
	private String name;
	private int healthPoints;
	private int damage;
	private int strenght;
	private Weapon w;
	private BufferedImage image;
	
	public Creature(String n, int hp, int s){
		this.name = n;
		this.healthPoints = hp;
		this.strenght = s;
		
	}
	public Creature(BufferedImage i,String n, int hp, int s){
		this.name = n;
		this.healthPoints = hp;
		this.strenght = s;
		this.image = i;
	}
	public Weapon getW() {
		return w;
	}

	public void setW(Weapon w) {
		this.w = w;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public int getHealthPoints() {
		return healthPoints;
	}

	public void setHealthPoints(int healthPoints) {
		this.healthPoints = healthPoints;
	}

	public int getDamage() {
		return damage;
	}

	public void setDamage(int damage) {
		this.damage = damage;
	}

	public int getStrenght() {
		return strenght;
	}

	public void setStrenght(int strenght) {
		this.strenght = strenght;
	}
	
	public String toString(){
		return "Name: " + this.name + "  HP: " + this.healthPoints+ " Strenght: " + this.strenght;
	}
	
	public String toFileManager(){
		return this.name + "#" + this.healthPoints + "#"+ this.strenght;
	}
	
	

}
