package MatrixSymmetric;
import java.util.Arrays ;
public class MatrixSymmetric implements Comparable<MatrixSymmetric> {
    private int[][] m;
    private int n;

    public MatrixSymmetric(int[][] m, int n){
        this.m = m;
        this.n = n;
    }

    public int CountAssymetricNumbers(){
        int count = 0;

        for (int i = 0; i < this.n; i++){
            for (int j = 0; j < this.n; j++){
                if (this.m[i][j] != this.m[j][i]) {
                    count++;
                }
            }
        }

        return count;
    }

    public int compareTo(MatrixSymmetric obj){
        int count1 = this.CountAssymetricNumbers();
        int count2 = obj.CountAssymetricNumbers();
        return count1 - count2;
    }

    public String toString() {
        String s = "";
        for (int i = 0; i < this.n; i++){
            s += Arrays.toString(this.m[i]);
            s += "\n";
        }
        return s;
    }
}
