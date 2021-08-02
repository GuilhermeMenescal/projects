package view;

import java.awt.BorderLayout;
import java.awt.EventQueue;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JTextField;
import javax.swing.border.EmptyBorder;

import control.CreatureControl;
import listener.PlayerCreatureListener;
import listener.WeaponWindowListener;

public class WeaponWindow extends JFrame {

	private JPanel contentPane;
	private JTextField option;
	private CreatureControl cc;


	public WeaponWindow(CreatureControl cc) {
		this.cc = cc;
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setBounds(100, 100, 450, 300);
		contentPane = new JPanel();
		contentPane.setBorder(new EmptyBorder(5, 5, 5, 5));
		setContentPane(contentPane);
		contentPane.setLayout(null);
		
		option = new JTextField();
		option.setBounds(170, 197, 109, 20);
		contentPane.add(option);
		option.setColumns(10);
		
		JButton btnSelectWeapon = new JButton("Select weapon!");
		btnSelectWeapon.setBounds(170, 228, 109, 23);
		btnSelectWeapon.addActionListener(new WeaponWindowListener(this.cc, this));
		contentPane.add(btnSelectWeapon);
		
		JLabel lblTypeTheDesired = new JLabel("Select a weapon for your creature!");
		lblTypeTheDesired.setBounds(140, 11, 254, 14);
		contentPane.add(lblTypeTheDesired);
		
		JButton btnNewButton = new JButton("Carrot");
		btnNewButton.addActionListener(new ActionListener() {
			public void actionPerformed(ActionEvent e) {
			}
		});
		btnNewButton.setBounds(71, 36, 135, 23);
		contentPane.add(btnNewButton);
		
		JButton btnNewButton_1 = new JButton("Child");
		btnNewButton_1.addActionListener(new ActionListener() {
			public void actionPerformed(ActionEvent e) {
			}
		});
		btnNewButton_1.setBounds(71, 70, 135, 23);
		contentPane.add(btnNewButton_1);
		
		JButton btnNewButton_2 = new JButton("Guinsu Knifes");
		btnNewButton_2.setBounds(71, 104, 135, 23);
		contentPane.add(btnNewButton_2);
		
		JButton btnNewButton_3 = new JButton("Abacus");
		btnNewButton_3.setBounds(71, 138, 135, 23);
		contentPane.add(btnNewButton_3);
		
		JButton btnNewButton_4 = new JButton("Hands");
		btnNewButton_4.setBounds(71, 172, 135, 23);
		contentPane.add(btnNewButton_4);
		
		JButton btnNewButton_5 = new JButton("Oddjob's Hat");
		btnNewButton_5.setBounds(279, 36, 145, 23);
		contentPane.add(btnNewButton_5);
		
		JButton btnNewButton_6 = new JButton("Monocyle");
		btnNewButton_6.setBounds(279, 70, 145, 23);
		contentPane.add(btnNewButton_6);
		
		JButton btnNewButton_7 = new JButton("Doritos");
		btnNewButton_7.setBounds(279, 104, 145, 23);
		contentPane.add(btnNewButton_7);
		
		JButton btnNewButton_8 = new JButton("Power of friendship");
		btnNewButton_8.setBounds(279, 138, 145, 23);
		contentPane.add(btnNewButton_8);
		
		JButton btnNewButton_9 = new JButton("Definitely real candy");
		btnNewButton_9.setBounds(279, 172, 145, 23);
		contentPane.add(btnNewButton_9);
		
		JLabel label = new JLabel("1.");
		label.setBounds(50, 40, 46, 14);
		contentPane.add(label);
		
		JLabel label_1 = new JLabel("2.");
		label_1.setBounds(50, 74, 46, 14);
		contentPane.add(label_1);
		
		JLabel label_2 = new JLabel("3.");
		label_2.setBounds(50, 108, 46, 14);
		contentPane.add(label_2);
		
		JLabel label_3 = new JLabel("4.");
		label_3.setBounds(50, 142, 46, 14);
		contentPane.add(label_3);
		
		JLabel label_4 = new JLabel("5.");
		label_4.setBounds(50, 176, 46, 14);
		contentPane.add(label_4);
		
		JLabel label_5 = new JLabel("6.");
		label_5.setBounds(263, 40, 46, 14);
		contentPane.add(label_5);
		
		JLabel label_6 = new JLabel("7.");
		label_6.setBounds(263, 74, 46, 14);
		contentPane.add(label_6);
		
		JLabel label_7 = new JLabel("8.");
		label_7.setBounds(263, 108, 46, 14);
		contentPane.add(label_7);
		
		JLabel label_8 = new JLabel("9.");
		label_8.setBounds(263, 142, 46, 14);
		contentPane.add(label_8);
		
		JLabel label_9 = new JLabel("10.");
		label_9.setBounds(263, 176, 46, 14);
		contentPane.add(label_9);
		
		JLabel lblClickHere = new JLabel("Click here ->");
		lblClickHere.setBounds(102, 232, 73, 14);
		contentPane.add(lblClickHere);
		
		JLabel label_10 = new JLabel("<- Click here");
		label_10.setBounds(287, 232, 73, 14);
		contentPane.add(label_10);
		this.setVisible(true);
	}


	public int getOption() {
		
		return Integer.parseInt(option.getText());
		
	}


	
}
