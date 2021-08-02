package model;

public class Player {
	
	public int getWin() {
		return win;
	}

	public void setWin(int win) {
		this.win = win;
	}

	public int getLosses() {
		return losses;
	}

	public void setLosses(int losses) {
		this.losses = losses;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Creature getC() {
		return c;
	}

	public void setC(Creature c) {
		this.c = c;
	}

	public int getLevel() {
		return level;
	}

	public void setLevel(int level) {
		this.level = level;
	}

	public String getPassword() {
		return password;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	private String name;
	private Creature c;
	private int level;
	private String password;
	private int win;
	private int losses;
	private String email;
	
	public String getEmail() {
		return email;
	}

	public void setEmail(String email) {
		this.email = email;
	}

	public Player (String n, String p, String e){
		this.name = n;
		this.password = p;
		this.email = e;
		
		
		
	}
	
	public String toFileManager(){
		return this.name + ":" + this.password + ":" + this.email; // + ":" + this.c + ":" + this.win + ":" this.losses + ":" + this.level;
	}
	public String toString(){
		return this.name + ": " + this.password + " - " + this.email;
	}
	

}
