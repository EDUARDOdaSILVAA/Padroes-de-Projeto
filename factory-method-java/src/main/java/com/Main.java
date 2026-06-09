package com;
import com.creators.DocumentCreator;
import com.creators.PdfCreator;
import com.creators.WordCreator;
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