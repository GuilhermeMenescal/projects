package control;

import java.util.ArrayList;
import java.util.Random;

import javax.swing.JOptionPane;

import database.SimulatedDatabase;
import model.Creature;
import model.Player;
import model.Weapon;
import util.FileManager;
import view.MainGameWindow;
import view.PlayerLoginWindow;
import view.RegisterWindow;

public class PlayerControl extends GenericControl {
	
	private SimulatedDatabase sdb;
	private FileManager fl;
	private Player currentPlayer;
	private ManoloDeControle mc;
	private Player computerPlayer;
	private Player finalBoss;
	
	public PlayerControl(SimulatedDatabase sdb, ManoloDeControle mc){
		this.sdb = sdb;
		this.mc = mc;
		this.fl = new FileManager();
	
	}

	@Override
	public void registerItem() {
		
		new RegisterWindow(this, this.mc);
		
		
		
		
	}

	@Override
	public void removeItem(String name) {
		
		
	}

	@Override
	public void showList() {
		
		
	}

	public void receivePlayer(String n, String p, String e){
		Player pl = new Player(n,p,e);
		
		this.fl.write(pl.toFileManager(), "players.txt");
		this.sdb.addPlayer(pl);
		
		
	}
	
	public boolean isValidPlayer (String name, String pass){
		ArrayList <Player> players = this.sdb.getPlayers();
		for (Player s : players){
			if (s.getName().equals(name) && s.getPassword().equals(pass)){
				this.currentPlayer = s;
				return true;
			}
		}
		return false;
	}
	
	public void loginWindow(){
		new PlayerLoginWindow(this, this.mc);
	}
	
	public String getCurrentPlayer() {
		return this.currentPlayer.getName();
	}

	public void GameWindow() {
		
		new MainGameWindow(this);
		
	}

	public void showCreatures() {
		
		new CreatureControl(this.sdb).showCreatures(this.currentPlayer);
		
	}
	
	public void calculateDps(){
		int dp1 = this.currentPlayer.getC().getDamage() + this.currentPlayer.getC().getW().getStrenght();
		int dp2 = this.computerPlayer.getC().getDamage() + this.computerPlayer.getC().getW().getStrenght();
		this.currentPlayer.getC().getW().setDps(dp1);
		this.computerPlayer.getC().getW().setDps(dp2) ;
	}
	
	public void playerFight(){
		Random rn = new Random();
		int r1 = rn.nextInt(10) + 1;
		int r2 = rn.nextInt(10) + 1;
		this.computerPlayer.setC(this.sdb.getCreatures().get(r1));
		this.computerPlayer.getC().setW(this.sdb.getWeapons().get(r2));
		
		JOptionPane.showMessageDialog(null, this.currentPlayer.getC().getName() + " VS " + this.computerPlayer.getC().getName());
		
		;
		
		
	}
	public void fight(){
		if(this.currentPlayer.getC().getW().getDps() > this.computerPlayer.getC().getW().getDps()){
			JOptionPane.showMessageDialog(null, "You won!!!");
			this.currentPlayer.setWin(this.currentPlayer.getWin() + 1);
			if(this.currentPlayer.getWin() == 3){
				this.currentPlayer.setLevel(this.currentPlayer.getLevel() + 1);
				JOptionPane.showMessageDialog(null, "You are now lv: " + this.currentPlayer.getLevel());
			}
			else if(this.currentPlayer.getWin() == 6){
				this.currentPlayer.setLevel(this.currentPlayer.getLevel() + 1);
				JOptionPane.showMessageDialog(null, "You are now lv: " + this.currentPlayer.getLevel());
			}
			else if(this.currentPlayer.getWin() == 9){
				this.currentPlayer.setLevel(this.currentPlayer.getLevel() + 1);
				JOptionPane.showMessageDialog(null, "You are now lv: " + this.currentPlayer.getLevel());
			}
			else if(this.currentPlayer.getWin() == 12){
				this.currentPlayer.setLevel(this.currentPlayer.getLevel() + 1);
				JOptionPane.showMessageDialog(null, "You are now lv: " + this.currentPlayer.getLevel());
			}
			else if(this.currentPlayer.getWin() == 18){
				this.currentPlayer.setLevel(this.currentPlayer.getLevel() + 1);
				JOptionPane.showMessageDialog(null, "You are now lv: " + this.currentPlayer.getLevel());
			}
		}
		else{
			JOptionPane.showMessageDialog(null, "HUMILIATING DEFEAT");
			this.currentPlayer.setLosses(this.currentPlayer.getLosses() + 1);
		}
	}
	public void gimmeFight(){
		this.computerPlayer = new Player("Computer", "computer", "computer@");
		this.playerFight();
		this.calculateDps();
		this.fight();
	}
	
	public void faceFinalBoss(){
		this.finalBoss = new Player("Boss","boss","boss@boss");
		this.finalBoss.setC(this.sdb.getCreatures().get(0));
		this.finalBoss.getC().setW(this.sdb.getWeapons().get(0));
		
	}
	
	public void finalFight(){
		
		if(this.currentPlayer.getLevel() < 5){
			JOptionPane.showMessageDialog(null, "You don't have enough power to defeat the boss");
		}
		else{
			JOptionPane.showMessageDialog(null, this.currentPlayer.getC().getName() + "VS" + this.finalBoss.getC().getName());
			JOptionPane.showMessageDialog(null, "You'll never have enough power to defeat the boss");
		}
		
	}
	public void bossFight(){
		this.faceFinalBoss();
		this.finalFight();
	}
	

}
