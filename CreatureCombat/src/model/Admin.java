package model;

public class Admin extends Player {
	
	private int adminPass = 1109;

	public Admin(String n, String p, String e) {
		super(n, p, e);
		
	}

	public int getAdminPass() {
		return adminPass;
	}

	public void setAdminPass(int adminPass) {
		this.adminPass = adminPass;
	}
	
	
	

}
