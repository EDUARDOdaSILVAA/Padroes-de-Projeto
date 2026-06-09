package src.main.java.com.creators;
import src.main.java.com.products.Document;
public abstract class DocumentCreator {

    public abstract Document createDocument();

    public void manageDocument() {
        Document doc = createDocument();
        doc.open();
        doc.close();
    }
}