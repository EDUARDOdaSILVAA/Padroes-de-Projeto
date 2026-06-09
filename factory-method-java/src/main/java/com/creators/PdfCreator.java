package src.main.java.com.creators;
import src.main.java.com.products.Document;
import src.main.java.com.products.PdfDocument;
public class PdfCreator extends DocumentCreator {
    @Override
    public Document createDocument() {
        return new PdfDocument();
    }
}