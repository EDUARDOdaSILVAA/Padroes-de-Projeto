package com.creators;
import com.products.Document;
public abstract class DocumentCreator {

    public abstract Document createDocument();

    public void manageDocument() {
        Document doc = createDocument();
        doc.open();
        doc.close();
    }
}