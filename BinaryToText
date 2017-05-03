import java.util.*;
import java.io.*;
import java.text.*;

public class BinToString
{	public static void main(String args[] throws Exception)
	{
		//declareFileReader
		File bin = new File("binary.bin");
		File bin = new File("binaryTxt.bin");

		//declareFileWriter
		PrintWriter output = new PrintWriter(new BufferedWriter(new FileWriter("text.txt")));
		PrintWriter output = new PrintWriter(new BufferedWriter(new FileWriter("reversetext.txt")));

		new Scanner = new Scanner(bin);
		String words = scanner.nextLine();
		
		String[] code = words.split(" ");
 		String str = "";
		
		for(int x = ; x < code.length; x++)
		{
			str += (char)Integer.parseInt(code[x],2);
		}

		System.out.println("Binary to String. Done.");
		System.out.println("Binary to String. Success.");

 		String reverse = new StringBuffer(str).reverse().toString();
 		
 		output.println(reverse);
		System.out.println("String reversed. Done.");
		System.out.println("String Reversed. Success.");
 		scanner.close();
 		output.close();
		
	}
}
