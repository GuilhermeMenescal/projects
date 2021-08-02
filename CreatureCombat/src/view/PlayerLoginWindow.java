package view;

import java.awt.BorderLayout;

import javax.swing.JFrame;
import javax.swing.JPanel;
import javax.swing.border.EmptyBorder;

import control.ManoloDeControle;
import control.PlayerControl;
import listener.BackListener;
import listener.PlayerLoginListener;

import javax.swing.JTextField;
import javax.swing.JLabel;
import javax.swing.JButton;
import javax.swing.JPasswordField;

public class PlayerLoginWindow extends JFrame {

	private JPanel contentPane;
	private PlayerControl pc;
	private JTextField name;
	private JPasswordField passwordField;
	private ManoloDeControle mc;
	
	public PlayerLoginWindow(PlayerControl pc, ManoloDeControle mc) {
		this.pc = pc;
		this.mc = mc;
		//setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setBounds(100, 100, 450, 300);
		contentPane = new JPanel();
		contentPane.setBorder(new EmptyBorder(5, 5, 5, 5));
		contentPane.setLayout(new BorderLayout(0, 0));
		setContentPane(contentPane);
		
		JPanel panel = new JPanel();
		contentPane.add(panel, BorderLayout.CENTER);
		panel.setLayout(null);
		
		name = new JTextField();
		name.setBounds(193, 71, 116, 22);
		panel.add(name);
		name.setColumns(10);
		
		JLabel lblNewLabel = new JLabel("Username:");
		lblNewLabel.setBounds(53, 74, 63, 16);
		panel.add(lblNewLabel);
		
		JLabel lblPassword = new JLabel("Password:");
		lblPassword.setBounds(53, 142, 81, 16);
		panel.add(lblPassword);
		
		JButton btnLogin = new JButton("Login");
		btnLogin.setBounds(82, 191, 97, 25);
		btnLogin.addActionListener(new PlayerLoginListener (this.pc, this));
		panel.add(btnLogin);
		
		passwordField = new JPasswordField();
		passwordField.setBounds(193, 136, 116, 22);
		panel.add(passwordField);
		
		JButton btnBack = new JButton("Back");
		btnBack.setBounds(242, 191, 97, 25);
		btnBack.addActionListener(new BackListener (this.mc, this));
		panel.add(btnBack);
		
		this.setVisible(true);
		this.setResizable(false);
		
		
	}

	public String getName() {
		return name.getText();
	}

	public String getPasswordField() {
		return passwordField.getText();
	}

	public void reset(){
		this.name.setText("");
		this.passwordField.setText("");
		
	}
}
