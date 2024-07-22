package IntDel;

public class IntDel implements Comparable<IntDel>{
    private int number;

    public IntDel(int number){
        this.number = number;
    }

    private int countDel(){
        int i = 1;
        int count = -1;
        while (i * i <= this.number){
            if (this.number % i == 0){
                if (i * i == this.number && IntDel.isPrime(i)){
                    count++;

                } else {
                    if (IntDel.isPrime(i)) {
                        count++;
                    }

                    if (IntDel.isPrime(this.number / i)) {
                        count++;
                    }
                }
            }
            i++;
        }
        return count;
    }

    private static boolean isPrime(int n){
        int i = 2;
        while (i * i <= n){
            if (n % i == 0){
                return false;
            }
            i++;
        }
        return true;
    }

    public int compareTo(IntDel obj){
        return this.countDel() - obj.countDel();
    }

    public String toString(){
        return Integer.toString(this.number);
    }
}
