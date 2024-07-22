package Universe;

import static java.lang.Math.*;

public class Universe {
    private Point[] listPoints;

    public Universe(Point[] listPoints) {
        this.listPoints = listPoints;
    }

    public double getForse(Point point) {
        double X, Y;
        double FinalX = 0;
        double FinalY = 0;

        for (int i = 0; i < listPoints.length; i++) {
            X = point.getXForse(listPoints[i]);
            Y = point.getYForse(listPoints[i]);
            FinalX += X;
            FinalY += Y;
        }
        double result = sqrt((FinalX * FinalX) + (FinalY * FinalY));
        return result;
    }
}
