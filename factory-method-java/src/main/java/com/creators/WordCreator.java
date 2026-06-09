package src.main.java.com.creators;
import src.main.java.com.products.Document;
import src.main.java.com.products.WordDocument;
public class WordCreator extends DocumentCreator {
    @Override
    public Document createDocument() {
        return new WordDocument();
    }
}