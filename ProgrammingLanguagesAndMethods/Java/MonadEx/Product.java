package MonadEx;

import javax.swing.text.html.Option;
import java.util.*;
import java.util.stream.Stream;
class Product {
    int cost, count;
    String name;
    Product (String name, int count, int cost) {
        this.name = name;
        this.count = count;
        this.cost = cost;
    }
}