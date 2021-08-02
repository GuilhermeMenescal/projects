package control;

import javax.swing.JOptionPane;

import database.SimulatedDatabase;
import view.LogInWindow;
import view.RegisterWindow;

public class ManoloDeControle  {
	
	private SimulatedDatabase sdb;
	
	public ManoloDeControle(){
		this.sdb = new SimulatedDatabase();
		this.start();
	}
	
	public void start(){
		new LogInWindow(this);
		
	}

	

	public void registerPlayer() {
		new PlayerControl(this.sdb, this).registerItem();
		
	}

	public void loginPlayer() {
		new PlayerControl(this.sdb, this).loginWindow();
		
	}
	
	

}
