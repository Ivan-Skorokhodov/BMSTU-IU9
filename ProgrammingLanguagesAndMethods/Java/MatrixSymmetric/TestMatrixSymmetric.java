package MatrixSymmetric;
import java.util.Arrays ;

public class TestMatrixSymmetric {
    public static void main(String[] args) {
        int n = 3;
        int[][] m1 = new int[n][n];
        int[][] m2 = new int[n][n];
        int[][] m3 = new int[n][n];

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                m1[i][j] = i + j;
                m2[i][j] = i + j;
                m3[i][j] = i + j;
            }
        }

        m2[0][1] = 4;

        m3[0][1] = 10;
        m3[1][2] = 15;

        MatrixSymmetric MatrixSymmetric1 = new MatrixSymmetric(m1, n);
        MatrixSymmetric MatrixSymmetric2 = new MatrixSymmetric(m2, n);
        MatrixSymmetric MatrixSymmetric3 = new MatrixSymmetric(m3, n);

        MatrixSymmetric[] array = new MatrixSymmetric[3];
        array[0] = MatrixSymmetric1;
        array[1] = MatrixSymmetric2;
        array[2] = MatrixSymmetric3;

        Arrays.sort(array);
        for (MatrixSymmetric matrix :array){
            System.out.println(matrix);
        }
    }
}
