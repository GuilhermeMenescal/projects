package view;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.border.EmptyBorder;

import control.ManoloDeControle;
import listener.LoginListener;
import listener.RegisterListener;
import javax.swing.ImageIcon;

public class LogInWindow extends JFrame {

	private JPanel contentPane;
	private ManoloDeControle mc;
	

	public LogInWindow(ManoloDeControle mc) {
		this.mc = mc;

		
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setBounds(100, 100, 450, 300);
		contentPane = new JPanel();
		contentPane.setBorder(new EmptyBorder(5, 5, 5, 5));
		setContentPane(contentPane);
		contentPane.setLayout(null);
		
		JPanel panel = new JPanel();
		panel.setBounds(0, -13, 450, 276);
		contentPane.add(panel);
		panel.setLayout(null);
		
		JButton btnRegister = new JButton("Register");
		btnRegister.setBounds(164, 147, 117, 25);
		btnRegister.addActionListener(new RegisterListener(this.mc, this));
		panel.add(btnRegister);
		
		JButton btnLogin = new JButton("Login");
		btnLogin.addActionListener(new LoginListener(this.mc, this));
		btnLogin.setBounds(164, 183, 117, 25);
		panel.add(btnLogin);
		
		JButton btnAdmin = new JButton("Admin ");
		btnAdmin.setBounds(164, 219, 117, 25);
		panel.add(btnAdmin);
		
		JLabel lblNewLabel = new JLabel("");
		lblNewLabel.setIcon(new ImageIcon("C:\\Users\\guilh_000\\Downloads\\CreatureCombat\\assets\\creturecombat.jpg"));
		lblNewLabel.setBounds(132, 25, 181, 120);
		panel.add(lblNewLabel);
		
		
		this.setVisible(true);
		this.setResizable(false);
	}
}
