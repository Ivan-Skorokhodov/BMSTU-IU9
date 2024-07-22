package Exception;

public class myCustomException extends Exception{
    public myCustomException(){
        super("Point not in circle");
        System.out.println("hello from exceptions, Point not in circle");
    }
    public void getMsg(String msg){
        System.out.println(msg);
    }
}