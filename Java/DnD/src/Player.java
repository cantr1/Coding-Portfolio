public class Player {
    // Fields
    public String name;
    public int health;

    // Constructors
    public Player(){
        this("Player 1", 100);
    }

    public Player(String name, int health){
        this.name = name;
        this.health = (health < 0 || health > 100) ? 100 : health;
    }

    // Getters
    public String getName(){
        return name;
    }

    // Setters
    public void setName(String name){
        this.name = name;
    }

    public void setHealth(int health) {
        this.health = health;
    }

    // Methods

    public void loseHealth(int damage){
        this.health -= damage;
        if (health < 0) {
            System.out.println("YOU DIED");
            System.exit(0);
        }
    }

    public void restoreHealth(int healing){
        this.health += healing;
    }

    public int healthRemaining(){
        return this.health;
    }
}
