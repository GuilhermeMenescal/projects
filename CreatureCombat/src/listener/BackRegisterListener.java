package listener;

import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import control.ManoloDeControle;
import view.RegisterWindow;

public class BackRegisterListener implements ActionListener {
		private ManoloDeControle mc;
		private RegisterWindow rw;
		
		public BackRegisterListener (ManoloDeControle mc, RegisterWindow rw){
			this.mc = mc;
			this.rw = rw;
		}
		@Override
		public void actionPerformed(ActionEvent e) {
			this.rw.setVisible(false);
			this.rw.dispose();
			this.mc.start();

		}

	}
