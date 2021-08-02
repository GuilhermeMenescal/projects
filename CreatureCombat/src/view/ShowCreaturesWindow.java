package view;

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

public class ShowCreaturesWindow extends JFrame {

	private JPanel contentPane;
	private JTextField textField;
	private CreatureControl cc;


	public ShowCreaturesWindow(CreatureControl cc) {
		this.cc = cc;
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setBounds(100, 100, 450, 300);
		contentPane = new JPanel();
		contentPane.setBorder(new EmptyBorder(5, 5, 5, 5));
		setContentPane(contentPane);
		contentPane.setLayout(null);
		
		textField = new JTextField();
		textField.setBounds(170, 197, 109, 20);
		contentPane.add(textField);
		textField.setColumns(10);
		
		JButton btnSelectCreature = new JButton("Select creature!");
		btnSelectCreature.setBounds(170, 228, 109, 23);
		btnSelectCreature.addActionListener(new PlayerCreatureListener(this.cc, this));
		contentPane.add(btnSelectCreature);
		
		JLabel lblTypeTheDesired = new JLabel("Type the desired creature number");
		lblTypeTheDesired.setBounds(140, 11, 254, 14);
		contentPane.add(lblTypeTheDesired);
		
		JButton btnNewButton = new JButton("Bugs Bunny");
		btnNewButton.addActionListener(new ActionListener() {
			public void actionPerformed(ActionEvent e) {
			}
		});
		btnNewButton.setBounds(81, 36, 125, 23);
		contentPane.add(btnNewButton);
		
		JButton btnNewButton_1 = new JButton("Harambe");
		btnNewButton_1.setBounds(81, 70, 125, 23);
		contentPane.add(btnNewButton_1);
		
		JButton btnNewButton_2 = new JButton("Euphoric Tyrone");
		btnNewButton_2.setBounds(81, 104, 125, 23);
		contentPane.add(btnNewButton_2);
		
		JButton btnNewButton_3 = new JButton("DatBoi");
		btnNewButton_3.setBounds(81, 138, 125, 23);
		contentPane.add(btnNewButton_3);
		
		JButton btnNewButton_4 = new JButton("Putin");
		btnNewButton_4.setBounds(81, 172, 125, 23);
		contentPane.add(btnNewButton_4);
		
		JButton btnNewButton_5 = new JButton("Polishop Man");
		btnNewButton_5.setBounds(279, 36, 125, 23);
		contentPane.add(btnNewButton_5);
		
		JButton btnNewButton_6 = new JButton("Sesame Count");
		btnNewButton_6.setBounds(279, 70, 125, 23);
		contentPane.add(btnNewButton_6);
		
		JButton btnNewButton_7 = new JButton("Guilherme");
		btnNewButton_7.setBounds(279, 104, 125, 23);
		contentPane.add(btnNewButton_7);
		
		JButton btnNewButton_8 = new JButton("Leticia");
		btnNewButton_8.setBounds(279, 138, 125, 23);
		contentPane.add(btnNewButton_8);
		
		JButton btnNewButton_9 = new JButton("Doritos Pope");
		btnNewButton_9.setBounds(279, 172, 125, 23);
		contentPane.add(btnNewButton_9);
		
		JLabel label = new JLabel("1.");
		label.setBounds(65, 40, 46, 14);
		contentPane.add(label);
		
		JLabel label_1 = new JLabel("2.");
		label_1.setBounds(65, 74, 46, 14);
		contentPane.add(label_1);
		
		JLabel label_2 = new JLabel("3.");
		label_2.setBounds(65, 108, 46, 14);
		contentPane.add(label_2);
		
		JLabel label_3 = new JLabel("4.");
		label_3.setBounds(65, 142, 46, 14);
		contentPane.add(label_3);
		
		JLabel label_4 = new JLabel("5.");
		label_4.setBounds(65, 176, 46, 14);
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


	public int getTextField() {
		
		return Integer.parseInt(textField.getText());
		
	}


	
}
