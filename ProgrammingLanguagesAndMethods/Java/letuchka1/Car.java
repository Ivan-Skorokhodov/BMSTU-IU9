package letuchka1;

public class Car extends Transport{
    public String model;

    public Car(String model){
        super("letuchka1.Car");
        this.model = model;
    }

    public void openDoor(int id){
        System.out.print("open the door with ");
        System.out.print(id);
        System.out.println(" id number");
    }
}

