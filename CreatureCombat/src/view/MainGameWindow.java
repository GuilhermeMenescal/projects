package view;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JPanel;
import javax.swing.border.EmptyBorder;

import control.PlayerControl;
import listener.BattleListener;
import listener.BossListener;
import listener.SelectYourCreatureListener;

public class MainGameWindow extends JFrame {

	private JPanel contentPane;
	private PlayerControl pc;
	


	public MainGameWindow(PlayerControl pc) {
		this.pc = pc;
		setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setBounds(100, 100, 450, 300);
		contentPane = new JPanel();
		contentPane.setBorder(new EmptyBorder(5, 5, 5, 5));
		setContentPane(contentPane);
		contentPane.setLayout(null);
		
		JButton btnSelectYourCreature = new JButton("Select your creature!");
		btnSelectYourCreature.setBounds(138, 55, 174, 23);
		btnSelectYourCreature.addActionListener(new SelectYourCreatureListener(this.pc));
		contentPane.add(btnSelectYourCreature);
		
		JButton btnBattle = new JButton("Battle!");
		btnBattle.setBounds(138, 104, 174, 23);
		btnBattle.addActionListener(new BattleListener(this.pc));
		contentPane.add(btnBattle);
		
		JButton btnBoss = new JButton("Face the final boss?");
		btnBoss.addActionListener(new BossListener(this.pc));
		btnBoss.setBounds(138, 152, 174, 23);
		contentPane.add(btnBoss);
		this.setVisible(true);
	}
}
