package src.main.java.com.products;
public class WordDocument implements Document {
    @Override
    public void open() {
        System.out.println("Abrindo o documento Word no modo de edição...");
    }

    @Override
    public void close() {
        System.out.println("Guardando as alterações e fechando o documento Word.");
    }
}