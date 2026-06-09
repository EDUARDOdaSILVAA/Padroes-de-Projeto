package src.main.java.com;
import src.main.java.com.creators.DocumentCreator;
import src.main.java.com.creators.PdfCreator;
import src.main.java.com.creators.WordCreator;
public class Main {
    private static DocumentCreator creator;

    public static void main(String[] args) {
        System.out.println(" Configuração para PDF ");
        creator = new PdfCreator();
        creator.manageDocument();

        System.out.println();

        System.out.println(" Configuração para Word ");
        creator = new WordCreator();
        creator.manageDocument();
    }
}