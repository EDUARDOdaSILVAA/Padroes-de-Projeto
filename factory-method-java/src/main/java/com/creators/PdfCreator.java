package com.creators;
import com.products.Document;
import com.products.PdfDocument;
public class PdfCreator extends DocumentCreator {
    @Override
    public Document createDocument() {
        return new PdfDocument();
    }
}