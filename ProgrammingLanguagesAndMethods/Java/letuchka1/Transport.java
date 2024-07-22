package letuchka1;

public class Transport{
    public String typeOfTransport;

    public Transport(String typeOfTransport){
        this.typeOfTransport = typeOfTransport;
    }

    public void move(){
        String s = "i am moving, because i am transport";
        System.out.println(s);
    }
}


