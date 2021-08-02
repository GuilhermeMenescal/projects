package view;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextField;
import javax.swing.border.EmptyBorder;

import control.ManoloDeControle;
import control.PlayerControl;
import listener.BackRegisterListener;
import listener.PlayerRegisterListener;

public class RegisterWindow extends JFrame {

	private JPanel contentPane;
	private JTextField name;
	private JTextField password;
	private JTextField email;
	private PlayerControl pc;
	private ManoloDeControle mc;


	public RegisterWindow(PlayerControl pc, ManoloDeControle mc) {
		this.pc = pc;
		this.mc = mc;
		//setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setBounds(100, 100, 450, 300);
		contentPane = new JPanel();
		contentPane.setBorder(new EmptyBorder(5, 5, 5, 5));
		setContentPane(contentPane);
		contentPane.setLayout(null);
		
		name = new JTextField();
		name.setBounds(241, 45, 86, 20);
		contentPane.add(name);
		name.setColumns(10);
		
		password = new JTextField();
		password.setBounds(241, 91, 86, 20);
		contentPane.add(password);
		password.setColumns(10);
		
		email = new JTextField();
		email.setBounds(241, 137, 86, 20);
		contentPane.add(email);
		email.setColumns(10);
		
		JLabel lblNewLabel = new JLabel("Name:");
		lblNewLabel.setBounds(74, 48, 46, 14);
		contentPane.add(lblNewLabel);
		
		JLabel lblNewLabel_1 = new JLabel("Password:");
		lblNewLabel_1.setBounds(74, 94, 68, 14);
		contentPane.add(lblNewLabel_1);
		
		JLabel lblNewLabel_2 = new JLabel("E-mail:");
		lblNewLabel_2.setBounds(74, 140, 46, 14);
		contentPane.add(lblNewLabel_2);
		
		JButton btnRegister = new JButton("Register");
		btnRegister.setBounds(101, 182, 89, 23);
		btnRegister.addActionListener(new PlayerRegisterListener(this.pc, this));
		contentPane.add(btnRegister);
		
		JButton btnBack = new JButton("Back");
		btnBack.setBounds(230, 181, 97, 25);
		btnBack.addActionListener(new BackRegisterListener (this.mc, this));
		contentPane.add(btnBack);
		this.setVisible(true);
	}


	public String getName() {
		return name.getText();
	}


	

	public String getPassword() {
		return password.getText();
	}





	public String getEmail() {
		return email.getText();
	}
	
	public void reset(){
		this.name.setText("");
		this.password.setText("");
		this.email.setText("");
		
	}


}
