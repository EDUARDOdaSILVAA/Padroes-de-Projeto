package com.creators;
import com.products.Document;
import com.products.WordDocument;
public class WordCreator extends DocumentCreator {
    @Override
    public Document createDocument() {
        return new WordDocument();
    }
}