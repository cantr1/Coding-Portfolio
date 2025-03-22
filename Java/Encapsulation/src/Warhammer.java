public class Warhammer implements Weapon {
    @Override
    public int attack() {
        System.out.println("ATTACK - 8 Damage");
        return 8;
    }

    @Override
    public void getAttacks() {
        System.out.println("Available Attacks: Rage - Attack - Push - Smash");
    }

    @Override
    public String getName() {
        return "Warhammer";
    }

    public int smash() {
        System.out.println("SMASH - 8 Damage");
        return 8;
    }

    public int push() {
        System.out.println("PUSH - 5 Damage");
        return 5;
    }
}