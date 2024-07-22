package MonadEx;

import javax.swing.text.html.Option;
import java.util.*;
import java.util.stream.Stream;
class ProductTable {
    HashMap<String, Product> Table;
    int total;
    ProductTable() {
        Table = new HashMap<>();
        total = 0;
    }
    void add(Product p) {
        Table.put(p.name, p);
        total += p.count;
    }
    void add(String name, int cost, int count) {
        Table.put(name, new Product(name, count, cost));
        total += count;
    }
    public Stream<String> nameStream(int v) {
        ArrayList<String> result = new ArrayList<>();
        Table.entrySet().stream()
                .filter(x -> x.getValue().count * x.getValue().cost > v)
                .forEach(x -> result.add(x.getKey()));
        return result.stream();
    }
    public Optional<Product> getProduct() {
        Optional<Product> result = Optional.empty();
        Optional<Map.Entry<String, Product>> tmp = Table.entrySet()
                .stream()
                .filter(x -> x.getValue().count * 2 > total)
                .findFirst();
        if (tmp.isPresent()) {
            result = Optional.ofNullable(tmp.get().getValue());
        }
        return result;
    }
}
