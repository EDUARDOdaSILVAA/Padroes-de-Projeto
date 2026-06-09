package com.products;
public class PdfDocument implements Document {
    @Override
    public void open() {
        System.out.println("Abrindo o documento PDF");
    }

    @Override
    public void close() {
        System.out.println("Fechando o documento PDF com segurança.");
    }
}