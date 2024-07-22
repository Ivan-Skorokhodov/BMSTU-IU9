package MonadEx;

import java.util.*;
class NameComparator implements Comparator<String> {
    public int compare(String a, String b) {
        char a0, b0;
        a0 = a.charAt(0);
        b0 = b.charAt(0);
        if (a0 > b0) { return 1; }
        if (a0 == b0) { return 0; }
        return -1;
    }
}
