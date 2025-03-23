public class Dragon implements Enemies{

    private int health = 20;

    @Override
    public int attack(){
        return 10;
    }

    @Override
    public int health() {
        return health;
    }

    public void setHealth(int newHealth) {
        this.health = newHealth;
    }

    public int scratch(){
        System.out.println("The dragon slashes with its claws! 8 Damage");
        return 8;
    }

    public int fireBreath(){
        System.out.println("The dragon breathes fire! 15 Damage");
        return 15;
    }

}
