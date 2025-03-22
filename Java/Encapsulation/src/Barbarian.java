public class Barbarian extends Player{

    //Fields
    private int strength;
    private Weapon weapon;

    public Barbarian(String name, int health, int strength, Weapon weapon){
        super(name, health);
        this.strength = strength;
        this.weapon = weapon;
    }

    // Getter for weapon
    public Weapon getWeapon() {
        return weapon;
    }


    public int getStrength() {
        return strength;
    }

    public void setStrength(int strength) {
        this.strength = strength;
    }

    public void enterRage(){
        this.strength = strength * 2;
        System.out.println("ENRAGED");
    }
}
