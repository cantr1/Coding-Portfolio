import java.util.Random;
import java.util.Scanner;
import java.util.concurrent.TimeUnit;

public class Main {
    private static Player player;
    private static Weapon playerWeapon;

    public static void main(String[] args) throws InterruptedException {
        introSequence();
        firstEncounter();
    }

    public static void introSequence() throws InterruptedException {
        //Fields
        String playerName;
        String playerClass;

        //Instantiate resources and scanner for user input
        Resources resources = new Resources();
        Scanner scanner = new Scanner(System.in);

        //Display the intro text
        resources.introDisplay();
        resources.introText();

        // Intro and take user input
        System.out.print("Enter your hero's name: ");
        playerName = scanner.nextLine();

        System.out.println("Select your class");
        System.out.println("Barbarian, Wizard, Rogue");
        playerClass = scanner.nextLine();

        //TODO: Add more classes

        // Conditional to instantiate class
        if(playerClass.equals("Barbarian")){
            String weaponChoice;

            System.out.println("Choose your weapon!");
            System.out.println("Warhammer - Mace - Stick");
            weaponChoice = scanner.nextLine();

            //TODO: Add more weapons
            if(weaponChoice.equals("Warhammer")) {
                playerWeapon = new Warhammer();
            }

            player = new Barbarian(playerName, 120, 2, playerWeapon);

            resources.clearScreen();

            resources.encounterSetup();
        }
    }

    public static void firstEncounter() throws InterruptedException {
        //Instantiate Dragon
        Dragon dragon = new Dragon();

        //Instantiate resources
        Resources resources = new Resources();
        Scanner scanner = new Scanner(System.in);
        Random random = new Random();
        int roll;

        TimeUnit.SECONDS.sleep(7);

        resources.clearScreen();

        resources.dragonDisplay();
        resources.dragonEnters();


        //System.out.println(player.getName() + " encounters a dragon!");

        System.out.println("Current Health: " + player.healthRemaining());

        //Loop for battle
        while(dragon.health() > 0){
            String playerTurn;

            System.out.println("Your Turn: ");
            playerWeapon.getAttacks();
            playerTurn = scanner.nextLine();

            switch (playerTurn) {
                case "Rage" -> ((Barbarian)player).enterRage();
                case "Attack" -> {
                    roll = diceRoll();
                    TimeUnit.SECONDS.sleep(2);
                    System.out.println("You rolled a " + roll + " + Strength = " + ((Barbarian)player).getStrength());
                    if(roll + ((Barbarian)player).getStrength() >= 10){
                        System.out.println("Dice roll = " + (roll + ((Barbarian)player).getStrength()));
                        dragon.setHealth(dragon.health() - playerWeapon.attack());
                        System.out.println("Remaining Enemy Health: " + dragon.health());
                    } else {
                        System.out.println("Dice roll = " + (roll + ((Barbarian)player).getStrength()));
                        System.out.println("Miss!");
                    }
                }
                case "Push" -> {
                    roll = diceRoll();
                    TimeUnit.SECONDS.sleep(2);
                    System.out.println("You rolled a " + roll + " + Strength = " + ((Barbarian)player).getStrength());
                    if(roll + ((Barbarian)player).getStrength() >= 10) {
                        if (playerWeapon instanceof Warhammer) {
                            System.out.println("Dice roll = " + (roll + ((Barbarian)player).getStrength()));
                            dragon.setHealth(dragon.health() - ((Warhammer) playerWeapon).push());
                            System.out.println("Remaining Enemy Health: " + dragon.health());
                        }
                    } else {
                        System.out.println("Dice roll = " + (roll + ((Barbarian)player).getStrength()));
                        System.out.println("Miss!");
                    }
                }
                case "Smash" -> {
                    roll = diceRoll();
                    TimeUnit.SECONDS.sleep(2);
                    System.out.println("You rolled a " + roll + " + Strength = " + ((Barbarian)player).getStrength());
                    if(roll + ((Barbarian)player).getStrength()  >= 10) {
                        if (playerWeapon instanceof Warhammer) {
                            System.out.println("Dice roll = " + (roll + ((Barbarian)player).getStrength()));
                            dragon.setHealth(dragon.health() - ((Warhammer) playerWeapon).smash());
                            System.out.println("Remaining Enemy Health: " + dragon.health());
                        }
                    } else {
                        System.out.println("Dice roll = " + (roll + ((Barbarian)player).getStrength()));
                        System.out.println("Miss!");
                    }
                }
            }

            if (dragon.health() < 0) {
                resources.defeatedDragon();
                resources.victoryText();
                System.exit(0);
            }

            System.out.println("Enemy Turn:");

            // Generate number between 1 and 2 for the dragons turn
            int enemyTurn = random.nextInt(2) + 1;

            switch (enemyTurn) {
                case 1 -> {
                    roll = diceRoll();
                    TimeUnit.SECONDS.sleep(2);
                    if(roll >= 10) {
                        player.loseHealth(dragon.scratch());
                        System.out.println("Remaining Player Health: " + player.healthRemaining());
                    } else {
                        System.out.println("The dragon lunges, but does not connect.");
                    }
                }

                case 2 -> {
                    roll = diceRoll();
                    TimeUnit.SECONDS.sleep(2);
                    if(roll >= 10) {
                        player.loseHealth(dragon.fireBreath());
                        System.out.println("Remaining Player Health: " + player.healthRemaining());
                    } else {
                        System.out.println("The dragon blasts fire but the hero is too quick!");
                    }
                }
            }
        }
    }

    public static int diceRoll(){
        Random random = new Random();
        return random.nextInt(20) + 1;
    }
}
