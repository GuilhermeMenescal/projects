package util;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.PrintWriter;
import java.io.UnsupportedEncodingException;
import java.util.ArrayList;

import javax.swing.JOptionPane;

public class FileManager {
	
	public void write(String text, String path) {
		
		PrintWriter writer;
		try {
			writer = new PrintWriter(new FileWriter(path, true));
			writer.println(text);
			writer.close();
		} catch (FileNotFoundException e) {
			
		} catch (UnsupportedEncodingException e) {
			
		} catch (IOException e) {
			
		}
	
	}
	
	public ArrayList<String> read(String path) {
		ArrayList<String> lines = new ArrayList<String>();
		
		BufferedReader br = null;

		try {

			String sCurrentLine;

			br = new BufferedReader(new FileReader(path));

			while ((sCurrentLine = br.readLine()) != null) {
				lines.add(sCurrentLine);
			}

		} catch (IOException e) {
			JOptionPane.showMessageDialog(null, "File not found, it will be created when you add something!");
		}
		
		return lines;
		
	}
}
	